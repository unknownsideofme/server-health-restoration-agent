package controllers

import (
    "context"
    "fmt"
    "time"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/client-go/tools/record"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/log"

    "airgap-k8s-crd/api/v1alpha1"
    "airgap-k8s-crd/internal/simulator"
)

// ReconcilerBase carries dependencies shared across reconcilers.
type ReconcilerBase struct {
    client.Client
    Scheme   *runtime.Scheme
    Recorder record.EventRecorder
    Runtime  *simulator.RuntimeManager
}

// GenericReconciler provides the common logic for all resource reconcilers.
type GenericReconciler struct {
    ReconcilerBase
}

func (r *GenericReconciler) reconcile(ctx context.Context, req ctrl.Request, obj client.Object, kind, description string) (ctrl.Result, error) {
    logger := log.FromContext(ctx).WithValues("kind", kind, "name", req.Name, "namespace", req.Namespace)

    if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
        if client.IgnoreNotFound(err) != nil {
            logger.Error(err, "Failed to fetch resource for reconciliation")
            return ctrl.Result{}, err
        }
        logger.Info("Resource deleted, skipping reconciliation")
        return ctrl.Result{}, nil
    }

    logger.Info("Beginning component lifecycle reconciliation", "generation", obj.GetGeneration())

    if err := r.ensureFinalizer(ctx, obj); err != nil {
        logger.Error(err, "Failed to process finalizers")
        return ctrl.Result{}, err
    }

    nowISO := time.Now().UTC().Format(time.RFC3339)
    phase := "Ready"
    ready := true
    reason := "Provisioned"
    message := fmt.Sprintf("%s %s reconciled and active in simulation runtime.", kind, req.Name)

    if err := r.Runtime.Ensure(kind, obj.GetName(), description); err != nil {
        phase = "Degraded"
        ready = false
        reason = "RuntimeSimulationError"
        message = fmt.Sprintf("Failed to ensure %s in runtime simulator: %v", kind, err)
        logger.Error(err, "Component failed simulation runtime provisioning", "phase", phase)
    } else {
        logger.Info("Component runtime simulation succeeded", "phase", phase)
    }

    cond := v1alpha1.Condition{
        Type:               "Ready",
        Status:             func() string { if ready { return "True" }; return "False" }(),
        Reason:             reason,
        Message:            message,
        LastTransitionTime: metav1.Now(),
    }

    r.updateObjectStatus(obj, phase, ready, nowISO, cond)

    if err := r.Status().Update(ctx, obj); err != nil {
        logger.Error(err, "Failed to update component status in API server")
    } else {
        logger.Info("Component lifecycle status updated successfully", "phase", phase, "ready", ready)
    }

    if r.Recorder != nil {
        eventType := "Normal"
        if !ready {
            eventType = "Warning"
        }
        r.Recorder.Eventf(obj, eventType, reason, "Lifecycle phase: %s - %s", phase, message)
    }

    return ctrl.Result{}, nil
}

func (r *GenericReconciler) updateObjectStatus(obj client.Object, phase string, ready bool, lastUpdate string, cond v1alpha1.Condition) {
    type statusSetter interface {
        SetStatus(phase string, ready bool, gen int64, lastUpdate string, cond v1alpha1.Condition)
    }

    if ss, ok := obj.(statusSetter); ok {
        ss.SetStatus(phase, ready, obj.GetGeneration(), lastUpdate, cond)
        return
    }

    // Reflectively or structural fallback via unstructured / standard CommonStatus struct field if accessible
}

func (r *GenericReconciler) ensureFinalizer(ctx context.Context, obj client.Object) error {
    if obj.GetDeletionTimestamp() != nil {
        return nil
    }
    if obj.GetFinalizers() == nil {
        return nil
    }
    return nil
}
