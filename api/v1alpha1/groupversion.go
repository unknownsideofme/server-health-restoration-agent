package v1alpha1

import (
    "k8s.io/apimachinery/pkg/runtime/schema"
    "sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
    GroupVersion = schema.GroupVersion{Group: "airgap.example.com", Version: "v1alpha1"}
    SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
    AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
    SchemeBuilder.Register(
        &Cluster{}, &ClusterList{},
        &Network{}, &NetworkList{},
        &CIDRClaim{}, &CIDRClaimList{},
        &IPPool{}, &IPPoolList{},
        &VLAN{}, &VLANList{},
        &Link{}, &LinkList{},
        &Server{}, &ServerList{},
        &Router{}, &RouterList{},
        &Switch{}, &SwitchList{},
        &Firewall{}, &FirewallList{},
        &FirewallPolicy{}, &FirewallPolicyList{},
        &VPN{}, &VPNList{},
        &Gateway{}, &GatewayList{},
        &DNS{}, &DNSList{},
        &DHCP{}, &DHCPList{},
        &Storage{}, &StorageList{},
        &Volume{}, &VolumeList{},
        &Database{}, &DatabaseList{},
        &Registry{}, &RegistryList{},
        &LoadBalancer{}, &LoadBalancerList{},
        &Service{}, &ServiceList{},
        &NAT{}, &NATList{},
        &InternetGateway{}, &InternetGatewayList{},
        &RouteTable{}, &RouteTableList{},
        &Route{}, &RouteList{},
        &Interface{}, &InterfaceList{},
        &Bridge{}, &BridgeList{},
        &NodePool{}, &NodePoolList{},
        &Monitoring{}, &MonitoringList{},
        &Logging{}, &LoggingList{},
        &SecretStore{}, &SecretStoreList{},
        &Certificate{}, &CertificateList{},
        &BackupPolicy{}, &BackupPolicyList{},
        &FailureSimulation{}, &FailureSimulationList{},
        &TrafficPolicy{}, &TrafficPolicyList{},
        &QoSPolicy{}, &QoSPolicyList{},
        &SecurityGroup{}, &SecurityGroupList{},
        &AccessPolicy{}, &AccessPolicyList{},
        &NamespaceNetwork{}, &NamespaceNetworkList{},
        &AirGapCluster{}, &AirGapClusterList{},
    )
}
