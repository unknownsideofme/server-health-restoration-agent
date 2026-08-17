package controllers

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"airgap-k8s-crd/api/v1alpha1"
)

const failureSimulationFinalizer = "failuresimulation.airgap.example.com/finalizer"

// FailureSimulationReconciler reconciles a FailureSimulation resource.
type FailureSimulationReconciler struct {
	ReconcilerBase
}

func (r *FailureSimulationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("failuresimulation", req.NamespacedName)

	var resource v1alpha1.FailureSimulation
	if err := r.Get(ctx, req.NamespacedName, &resource); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	isMarkedToBeDeleted := resource.GetDeletionTimestamp() != nil
	if isMarkedToBeDeleted {
		if controllerutil.ContainsFinalizer(&resource, failureSimulationFinalizer) {
			r.runCleanupCommand(logger, &resource)
			controllerutil.RemoveFinalizer(&resource, failureSimulationFinalizer)
			if err := r.Update(ctx, &resource); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(&resource, failureSimulationFinalizer) {
		controllerutil.AddFinalizer(&resource, failureSimulationFinalizer)
		if err := r.Update(ctx, &resource); err != nil {
			return ctrl.Result{}, err
		}
	}

	nowISO := time.Now().UTC().Format(time.RFC3339)

	if resource.Spec.Enabled {
		if resource.Status.Phase != "Injected" {
			err := r.runInjectionCommand(logger, &resource)
			phase := "Injected"
			ready := true
			reason := "Injected"
			message := fmt.Sprintf("Fault %s injected successfully on %s.", resource.Spec.Type, resource.Spec.Address)
			
			if err != nil {
				phase = "Failed"
				ready = false
				reason = "InjectionFailed"
				message = fmt.Sprintf("Failed to inject fault: %v", err)
			}

			cond := v1alpha1.Condition{
				Type:               "Ready",
				Status:             func() string { if ready { return "True" }; return "False" }(),
				Reason:             reason,
				Message:            message,
				LastTransitionTime: metav1.Now(),
			}

			resource.Status.Phase = phase
			resource.Status.Ready = ready
			resource.Status.LastUpdateTime = nowISO
			
			// Replace conditions
			var newConditions []v1alpha1.Condition
			for _, c := range resource.Status.Conditions {
				if c.Type != "Ready" {
					newConditions = append(newConditions, c)
				}
			}
			newConditions = append(newConditions, cond)
			resource.Status.Conditions = newConditions

			if err := r.Status().Update(ctx, &resource); err != nil {
				logger.Error(err, "Failed to update FailureSimulation status")
				return ctrl.Result{}, err
			}
			
			if err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *FailureSimulationReconciler) runInjectionCommand(logger logr.Logger, resource *v1alpha1.FailureSimulation) error {
	target := resource.Spec.Address
	if target == "" {
		target = resource.Spec.NetworkRef
	}
	fType := resource.Spec.Type

	logger.Info(fmt.Sprintf("Injecting fault %s on %s", fType, target))

	var cmds []string
	if fType == "PROGRESSIVE_CONGESTION" {
		cmds = []string{
			"tc qdisc del dev eth0 root",
			"tc qdisc add dev eth0 root tbf rate 20kbit burst 32kbit latency 400ms",
		}
	} else if fType == "ROUTE_FLAP_CASCADE" {
		cmds = []string{
			"apt-get update -qq && apt-get install -y -qq iptables",
			"touch /tmp/airgap_route_flap_active",
			"nohup /bin/bash -c 'while [ -f /tmp/airgap_route_flap_active ]; do iptables-legacy -I OUTPUT -d 10.10.0.0/16 -j DROP 2>/dev/null || true; sleep 3; iptables-legacy -D OUTPUT -d 10.10.0.0/16 -j DROP 2>/dev/null || true; sleep 3; done' >/dev/null 2>&1 &",
		}
	} else if fType == "TUNNEL_DEGRADATION" {
		cmds = []string{
			"tc qdisc del dev eth0 root",
			"tc qdisc add dev eth0 root netem delay 150ms 30ms loss 12%",
		}
	} else if fType == "POLICY_DRIFT" {
		cmds = []string{
			"apt-get update -qq && apt-get install -y -qq iptables",
			"iptables-legacy -I INPUT -p tcp --dport 8085 -j DROP",
			"iptables-legacy -I INPUT -p tcp --dport 11435 -j DROP",
		}
	}

	for _, cmdStr := range cmds {
		r.executeKubectl(logger, target, cmdStr)
	}
	return nil
}

func (r *FailureSimulationReconciler) runCleanupCommand(logger logr.Logger, resource *v1alpha1.FailureSimulation) {
	target := resource.Spec.Address
	if target == "" {
		target = resource.Spec.NetworkRef
	}
	fType := resource.Spec.Type

	logger.Info(fmt.Sprintf("Cleaning up fault %s on %s", fType, target))

	var cmds []string
	if fType == "PROGRESSIVE_CONGESTION" || fType == "TUNNEL_DEGRADATION" {
		cmds = []string{"tc qdisc del dev eth0 root"}
	} else if fType == "ROUTE_FLAP_CASCADE" {
		cmds = []string{
			"rm -f /tmp/airgap_route_flap_active",
			"iptables-legacy -D OUTPUT -d 10.10.0.0/16 -j DROP",
		}
	} else if fType == "POLICY_DRIFT" {
		cmds = []string{"iptables-legacy -F"}
	}

	for _, cmdStr := range cmds {
		r.executeKubectl(logger, target, cmdStr)
	}
}

func (r *FailureSimulationReconciler) executeKubectl(logger logr.Logger, target string, cmdStr string) {
	cmd := exec.Command("kubectl", "exec", "-i", target, "--", "sh", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err, "Command execution failed", "output", string(out), "cmd", cmdStr)
	} else {
		logger.Info("Command executed successfully", "cmd", cmdStr)
	}
}

func (r *FailureSimulationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.FailureSimulation{}).
		Complete(r)
}
