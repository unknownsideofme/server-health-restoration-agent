package main

import (
    "flag"
    "os"

    "k8s.io/apimachinery/pkg/runtime"
    utilruntime "k8s.io/apimachinery/pkg/util/runtime"
    clientgoscheme "k8s.io/client-go/kubernetes/scheme"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/healthz"
    "sigs.k8s.io/controller-runtime/pkg/metrics/server"

    "airgap-k8s-crd/api/v1alpha1"
    controllerspkg "airgap-k8s-crd/controllers"
    "airgap-k8s-crd/internal/simulator"
)

var (
    scheme = runtime.NewScheme()
    metricsAddr = flag.String("metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
    probeAddr = flag.String("health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
)

func init() {
    utilruntime.Must(clientgoscheme.AddToScheme(scheme))
    utilruntime.Must(v1alpha1.AddToScheme(scheme))
}

func main() {
    flag.Parse()
    mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{Scheme: scheme, Metrics: server.Options{BindAddress: *metricsAddr}, HealthProbeBindAddress: *probeAddr, LeaderElection: false})
    if err != nil { os.Exit(1) }

    runtimeManager := simulator.NewRuntimeManager()
    setupList := []interface{ SetupWithManager(ctrl.Manager) error }{
        &controllerspkg.ClusterReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("cluster-controller"), Runtime: runtimeManager}}},
        &controllerspkg.NetworkReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("network-controller"), Runtime: runtimeManager}}},
        &controllerspkg.CIDRClaimReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("cidrclaim-controller"), Runtime: runtimeManager}}},
        &controllerspkg.IPPoolReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("ippool-controller"), Runtime: runtimeManager}}},
        &controllerspkg.VLANReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("vlan-controller"), Runtime: runtimeManager}}},
        &controllerspkg.LinkReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("link-controller"), Runtime: runtimeManager}}},
        &controllerspkg.ServerReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("server-controller"), Runtime: runtimeManager}}},
        &controllerspkg.RouterReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("router-controller"), Runtime: runtimeManager}}},
        &controllerspkg.SwitchReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("switch-controller"), Runtime: runtimeManager}}},
        &controllerspkg.FirewallReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("firewall-controller"), Runtime: runtimeManager}}},
        &controllerspkg.FirewallPolicyReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("firewallpolicy-controller"), Runtime: runtimeManager}}},
        &controllerspkg.VPNReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("vpn-controller"), Runtime: runtimeManager}}},
        &controllerspkg.GatewayReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("gateway-controller"), Runtime: runtimeManager}}},
        &controllerspkg.DNSReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("dns-controller"), Runtime: runtimeManager}}},
        &controllerspkg.DHCPReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("dhcp-controller"), Runtime: runtimeManager}}},
        &controllerspkg.StorageReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("storage-controller"), Runtime: runtimeManager}}},
        &controllerspkg.VolumeReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("volume-controller"), Runtime: runtimeManager}}},
        &controllerspkg.DatabaseReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("database-controller"), Runtime: runtimeManager}}},
        &controllerspkg.RegistryReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("registry-controller"), Runtime: runtimeManager}}},
        &controllerspkg.LoadBalancerReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("loadbalancer-controller"), Runtime: runtimeManager}}},
        &controllerspkg.ServiceReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("service-controller"), Runtime: runtimeManager}}},
        &controllerspkg.NATReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("nat-controller"), Runtime: runtimeManager}}},
        &controllerspkg.InternetGatewayReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("internetgateway-controller"), Runtime: runtimeManager}}},
        &controllerspkg.RouteTableReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("routetable-controller"), Runtime: runtimeManager}}},
        &controllerspkg.RouteReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("route-controller"), Runtime: runtimeManager}}},
        &controllerspkg.InterfaceReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("interface-controller"), Runtime: runtimeManager}}},
        &controllerspkg.BridgeReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("bridge-controller"), Runtime: runtimeManager}}},
        &controllerspkg.NodePoolReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("nodepool-controller"), Runtime: runtimeManager}}},
        &controllerspkg.MonitoringReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("monitoring-controller"), Runtime: runtimeManager}}},
        &controllerspkg.LoggingReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("logging-controller"), Runtime: runtimeManager}}},
        &controllerspkg.SecretStoreReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("secretstore-controller"), Runtime: runtimeManager}}},
        &controllerspkg.CertificateReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("certificate-controller"), Runtime: runtimeManager}}},
        &controllerspkg.BackupPolicyReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("backuppolicy-controller"), Runtime: runtimeManager}}},
        &controllerspkg.FailureSimulationReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("failuresimulation-controller"), Runtime: runtimeManager}},
        &controllerspkg.TrafficPolicyReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("trafficpolicy-controller"), Runtime: runtimeManager}}},
        &controllerspkg.QoSPolicyReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("qospolicy-controller"), Runtime: runtimeManager}}},
        &controllerspkg.SecurityGroupReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("securitygroup-controller"), Runtime: runtimeManager}}},
        &controllerspkg.AccessPolicyReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("accesspolicy-controller"), Runtime: runtimeManager}}},
        &controllerspkg.NamespaceNetworkReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("namespacenetwork-controller"), Runtime: runtimeManager}}},
        &controllerspkg.AirGapClusterReconciler{GenericReconciler: controllerspkg.GenericReconciler{ReconcilerBase: controllerspkg.ReconcilerBase{Client: mgr.GetClient(), Scheme: mgr.GetScheme(), Recorder: mgr.GetEventRecorderFor("airgapcluster-controller"), Runtime: runtimeManager}}},
    }

    for _, setup := range setupList {
        if err := setup.SetupWithManager(mgr); err != nil { os.Exit(1) }
    }

    if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil { os.Exit(1) }
    if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil { os.Exit(1) }
    if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil { os.Exit(1) }
}
