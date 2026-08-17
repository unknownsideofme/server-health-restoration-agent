package controllers

import (
    "context"

    ctrl "sigs.k8s.io/controller-runtime"

    "airgap-k8s-crd/api/v1alpha1"
)

// ClusterReconciler reconciles a Cluster resource.
type ClusterReconciler struct {
    GenericReconciler
}

func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Cluster
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Cluster", resource.Spec.Description)
}

func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Cluster{}).Complete(r)
}

// NetworkReconciler reconciles a Network resource.
type NetworkReconciler struct {
    GenericReconciler
}

func (r *NetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Network
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Network", resource.Spec.Description)
}

func (r *NetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Network{}).Complete(r)
}

// CIDRClaimReconciler reconciles a CIDRClaim resource.
type CIDRClaimReconciler struct {
    GenericReconciler
}

func (r *CIDRClaimReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.CIDRClaim
    return r.GenericReconciler.reconcile(ctx, req, &resource, "CIDRClaim", resource.Spec.Description)
}

func (r *CIDRClaimReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.CIDRClaim{}).Complete(r)
}

// IPPoolReconciler reconciles a IPPool resource.
type IPPoolReconciler struct {
    GenericReconciler
}

func (r *IPPoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.IPPool
    return r.GenericReconciler.reconcile(ctx, req, &resource, "IPPool", resource.Spec.Description)
}

func (r *IPPoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.IPPool{}).Complete(r)
}

// VLANReconciler reconciles a VLAN resource.
type VLANReconciler struct {
    GenericReconciler
}

func (r *VLANReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.VLAN
    return r.GenericReconciler.reconcile(ctx, req, &resource, "VLAN", resource.Spec.Description)
}

func (r *VLANReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.VLAN{}).Complete(r)
}

// LinkReconciler reconciles a Link resource.
type LinkReconciler struct {
    GenericReconciler
}

func (r *LinkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Link
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Link", resource.Spec.Description)
}

func (r *LinkReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Link{}).Complete(r)
}

// ServerReconciler reconciles a Server resource.
type ServerReconciler struct {
    GenericReconciler
}

func (r *ServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Server
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Server", resource.Spec.Description)
}

func (r *ServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Server{}).Complete(r)
}

// RouterReconciler reconciles a Router resource.
type RouterReconciler struct {
    GenericReconciler
}

func (r *RouterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Router
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Router", resource.Spec.Description)
}

func (r *RouterReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Router{}).Complete(r)
}

// SwitchReconciler reconciles a Switch resource.
type SwitchReconciler struct {
    GenericReconciler
}

func (r *SwitchReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Switch
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Switch", resource.Spec.Description)
}

func (r *SwitchReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Switch{}).Complete(r)
}

// FirewallReconciler reconciles a Firewall resource.
type FirewallReconciler struct {
    GenericReconciler
}

func (r *FirewallReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Firewall
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Firewall", resource.Spec.Description)
}

func (r *FirewallReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Firewall{}).Complete(r)
}

// FirewallPolicyReconciler reconciles a FirewallPolicy resource.
type FirewallPolicyReconciler struct {
    GenericReconciler
}

func (r *FirewallPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.FirewallPolicy
    return r.GenericReconciler.reconcile(ctx, req, &resource, "FirewallPolicy", resource.Spec.Description)
}

func (r *FirewallPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.FirewallPolicy{}).Complete(r)
}

// VPNReconciler reconciles a VPN resource.
type VPNReconciler struct {
    GenericReconciler
}

func (r *VPNReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.VPN
    return r.GenericReconciler.reconcile(ctx, req, &resource, "VPN", resource.Spec.Description)
}

func (r *VPNReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.VPN{}).Complete(r)
}

// GatewayReconciler reconciles a Gateway resource.
type GatewayReconciler struct {
    GenericReconciler
}

func (r *GatewayReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Gateway
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Gateway", resource.Spec.Description)
}

func (r *GatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Gateway{}).Complete(r)
}

// DNSReconciler reconciles a DNS resource.
type DNSReconciler struct {
    GenericReconciler
}

func (r *DNSReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.DNS
    return r.GenericReconciler.reconcile(ctx, req, &resource, "DNS", resource.Spec.Description)
}

func (r *DNSReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.DNS{}).Complete(r)
}

// DHCPReconciler reconciles a DHCP resource.
type DHCPReconciler struct {
    GenericReconciler
}

func (r *DHCPReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.DHCP
    return r.GenericReconciler.reconcile(ctx, req, &resource, "DHCP", resource.Spec.Description)
}

func (r *DHCPReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.DHCP{}).Complete(r)
}

// StorageReconciler reconciles a Storage resource.
type StorageReconciler struct {
    GenericReconciler
}

func (r *StorageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Storage
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Storage", resource.Spec.Description)
}

func (r *StorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Storage{}).Complete(r)
}

// VolumeReconciler reconciles a Volume resource.
type VolumeReconciler struct {
    GenericReconciler
}

func (r *VolumeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Volume
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Volume", resource.Spec.Description)
}

func (r *VolumeReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Volume{}).Complete(r)
}

// DatabaseReconciler reconciles a Database resource.
type DatabaseReconciler struct {
    GenericReconciler
}

func (r *DatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Database
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Database", resource.Spec.Description)
}

func (r *DatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Database{}).Complete(r)
}

// RegistryReconciler reconciles a Registry resource.
type RegistryReconciler struct {
    GenericReconciler
}

func (r *RegistryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Registry
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Registry", resource.Spec.Description)
}

func (r *RegistryReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Registry{}).Complete(r)
}

// LoadBalancerReconciler reconciles a LoadBalancer resource.
type LoadBalancerReconciler struct {
    GenericReconciler
}

func (r *LoadBalancerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.LoadBalancer
    return r.GenericReconciler.reconcile(ctx, req, &resource, "LoadBalancer", resource.Spec.Description)
}

func (r *LoadBalancerReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.LoadBalancer{}).Complete(r)
}

// ServiceReconciler reconciles a Service resource.
type ServiceReconciler struct {
    GenericReconciler
}

func (r *ServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Service
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Service", resource.Spec.Description)
}

func (r *ServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Service{}).Complete(r)
}

// NATReconciler reconciles a NAT resource.
type NATReconciler struct {
    GenericReconciler
}

func (r *NATReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.NAT
    return r.GenericReconciler.reconcile(ctx, req, &resource, "NAT", resource.Spec.Description)
}

func (r *NATReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.NAT{}).Complete(r)
}

// InternetGatewayReconciler reconciles a InternetGateway resource.
type InternetGatewayReconciler struct {
    GenericReconciler
}

func (r *InternetGatewayReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.InternetGateway
    return r.GenericReconciler.reconcile(ctx, req, &resource, "InternetGateway", resource.Spec.Description)
}

func (r *InternetGatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.InternetGateway{}).Complete(r)
}

// RouteTableReconciler reconciles a RouteTable resource.
type RouteTableReconciler struct {
    GenericReconciler
}

func (r *RouteTableReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.RouteTable
    return r.GenericReconciler.reconcile(ctx, req, &resource, "RouteTable", resource.Spec.Description)
}

func (r *RouteTableReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.RouteTable{}).Complete(r)
}

// RouteReconciler reconciles a Route resource.
type RouteReconciler struct {
    GenericReconciler
}

func (r *RouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Route
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Route", resource.Spec.Description)
}

func (r *RouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Route{}).Complete(r)
}

// InterfaceReconciler reconciles a Interface resource.
type InterfaceReconciler struct {
    GenericReconciler
}

func (r *InterfaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Interface
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Interface", resource.Spec.Description)
}

func (r *InterfaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Interface{}).Complete(r)
}

// BridgeReconciler reconciles a Bridge resource.
type BridgeReconciler struct {
    GenericReconciler
}

func (r *BridgeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Bridge
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Bridge", resource.Spec.Description)
}

func (r *BridgeReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Bridge{}).Complete(r)
}

// NodePoolReconciler reconciles a NodePool resource.
type NodePoolReconciler struct {
    GenericReconciler
}

func (r *NodePoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.NodePool
    return r.GenericReconciler.reconcile(ctx, req, &resource, "NodePool", resource.Spec.Description)
}

func (r *NodePoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.NodePool{}).Complete(r)
}

// MonitoringReconciler reconciles a Monitoring resource.
type MonitoringReconciler struct {
    GenericReconciler
}

func (r *MonitoringReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Monitoring
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Monitoring", resource.Spec.Description)
}

func (r *MonitoringReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Monitoring{}).Complete(r)
}

// LoggingReconciler reconciles a Logging resource.
type LoggingReconciler struct {
    GenericReconciler
}

func (r *LoggingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Logging
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Logging", resource.Spec.Description)
}

func (r *LoggingReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Logging{}).Complete(r)
}

// SecretStoreReconciler reconciles a SecretStore resource.
type SecretStoreReconciler struct {
    GenericReconciler
}

func (r *SecretStoreReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.SecretStore
    return r.GenericReconciler.reconcile(ctx, req, &resource, "SecretStore", resource.Spec.Description)
}

func (r *SecretStoreReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.SecretStore{}).Complete(r)
}

// CertificateReconciler reconciles a Certificate resource.
type CertificateReconciler struct {
    GenericReconciler
}

func (r *CertificateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.Certificate
    return r.GenericReconciler.reconcile(ctx, req, &resource, "Certificate", resource.Spec.Description)
}

func (r *CertificateReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.Certificate{}).Complete(r)
}

// BackupPolicyReconciler reconciles a BackupPolicy resource.
type BackupPolicyReconciler struct {
    GenericReconciler
}

func (r *BackupPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.BackupPolicy
    return r.GenericReconciler.reconcile(ctx, req, &resource, "BackupPolicy", resource.Spec.Description)
}

func (r *BackupPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.BackupPolicy{}).Complete(r)
}


// TrafficPolicyReconciler reconciles a TrafficPolicy resource.
type TrafficPolicyReconciler struct {
    GenericReconciler
}

func (r *TrafficPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.TrafficPolicy
    return r.GenericReconciler.reconcile(ctx, req, &resource, "TrafficPolicy", resource.Spec.Description)
}

func (r *TrafficPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.TrafficPolicy{}).Complete(r)
}

// QoSPolicyReconciler reconciles a QoSPolicy resource.
type QoSPolicyReconciler struct {
    GenericReconciler
}

func (r *QoSPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.QoSPolicy
    return r.GenericReconciler.reconcile(ctx, req, &resource, "QoSPolicy", resource.Spec.Description)
}

func (r *QoSPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.QoSPolicy{}).Complete(r)
}

// SecurityGroupReconciler reconciles a SecurityGroup resource.
type SecurityGroupReconciler struct {
    GenericReconciler
}

func (r *SecurityGroupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.SecurityGroup
    return r.GenericReconciler.reconcile(ctx, req, &resource, "SecurityGroup", resource.Spec.Description)
}

func (r *SecurityGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.SecurityGroup{}).Complete(r)
}

// AccessPolicyReconciler reconciles a AccessPolicy resource.
type AccessPolicyReconciler struct {
    GenericReconciler
}

func (r *AccessPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.AccessPolicy
    return r.GenericReconciler.reconcile(ctx, req, &resource, "AccessPolicy", resource.Spec.Description)
}

func (r *AccessPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.AccessPolicy{}).Complete(r)
}

// NamespaceNetworkReconciler reconciles a NamespaceNetwork resource.
type NamespaceNetworkReconciler struct {
    GenericReconciler
}

func (r *NamespaceNetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.NamespaceNetwork
    return r.GenericReconciler.reconcile(ctx, req, &resource, "NamespaceNetwork", resource.Spec.Description)
}

func (r *NamespaceNetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.NamespaceNetwork{}).Complete(r)
}

// AirGapClusterReconciler reconciles a AirGapCluster resource.
type AirGapClusterReconciler struct {
    GenericReconciler
}

func (r *AirGapClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var resource v1alpha1.AirGapCluster
    return r.GenericReconciler.reconcile(ctx, req, &resource, "AirGapCluster", resource.Spec.Description)
}

func (r *AirGapClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).For(&v1alpha1.AirGapCluster{}).Complete(r)
}

