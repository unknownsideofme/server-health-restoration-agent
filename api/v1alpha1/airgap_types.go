package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
)

// Condition captures a reconciliation condition.
type Condition struct {
    Type string `json:"type"`
    Status string `json:"status"`
    Reason string `json:"reason,omitempty"`
    Message string `json:"message,omitempty"`
    LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// CommonStatus contains the shared status fields used by all resources.
type CommonStatus struct {
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    Phase string `json:"phase,omitempty"`
    Ready bool `json:"ready,omitempty"`
    LastLifecycleUpdate string `json:"lastLifecycleUpdate,omitempty"`
    Conditions []Condition `json:"conditions,omitempty"`
}

// ClusterSpec defines the desired state of Cluster.
type ClusterSpec struct {
    Description string `json:"description,omitempty"`
    Enabled bool `json:"enabled,omitempty"`
    NetworkRef string `json:"networkRef,omitempty"`
    CIDR string `json:"cidr,omitempty"`
    Type string `json:"type,omitempty"`
    Address string `json:"address,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
    CommonStatus `json:",inline"`
    Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=clu
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Cluster struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec ClusterSpec `json:"spec,omitempty"`
    Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ClusterList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items []Cluster `json:"items"`
}

func (in *Cluster) DeepCopyInto(out *Cluster) {
    *out = *in
}

func (in *Cluster) DeepCopy() *Cluster {
    if in == nil { return nil }
    out := new(Cluster)
    in.DeepCopyInto(out)
    return out
}

func (in *Cluster) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil { return c }
    return nil
}

func (in *ClusterList) DeepCopyInto(out *ClusterList) {
    *out = *in
    if in.Items != nil {
        out.Items = make([]Cluster, len(in.Items))
        copy(out.Items, in.Items)
    }
}

func (in *ClusterList) DeepCopy() *ClusterList {
    if in == nil { return nil }
    out := new(ClusterList)
    in.DeepCopyInto(out)
    return out
}

func (in *ClusterList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil { return c }
    return nil
}

// NetworkSpec defines the desired state of Network.
type NetworkSpec struct {
    Description string `json:"description,omitempty"`
    Enabled bool `json:"enabled,omitempty"`
    NetworkRef string `json:"networkRef,omitempty"`
    CIDR string `json:"cidr,omitempty"`
    Type string `json:"type,omitempty"`
    Address string `json:"address,omitempty"`
}

type NetworkStatus struct {
    CommonStatus `json:",inline"`
    Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=net
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Network struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec NetworkSpec `json:"spec,omitempty"`
    Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type NetworkList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items []Network `json:"items"`
}

func (in *Network) DeepCopyInto(out *Network) { *out = *in }
func (in *Network) DeepCopy() *Network { if in == nil { return nil }; out := new(Network); in.DeepCopyInto(out); return out }
func (in *Network) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *NetworkList) DeepCopyInto(out *NetworkList) { *out = *in; if in.Items != nil { out.Items = make([]Network, len(in.Items)); copy(out.Items, in.Items) } }
func (in *NetworkList) DeepCopy() *NetworkList { if in == nil { return nil }; out := new(NetworkList); in.DeepCopyInto(out); return out }
func (in *NetworkList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// CIDRClaimSpec defines the desired state of CIDRClaim.
type CIDRClaimSpec struct {
    Description string `json:"description,omitempty"`
    Enabled bool `json:"enabled,omitempty"`
    NetworkRef string `json:"networkRef,omitempty"`
    CIDR string `json:"cidr,omitempty"`
    Type string `json:"type,omitempty"`
    Address string `json:"address,omitempty"`
}

type CIDRClaimStatus struct {
    CommonStatus `json:",inline"`
    Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=cid
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type CIDRClaim struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec CIDRClaimSpec `json:"spec,omitempty"`
    Status CIDRClaimStatus `json:"status,omitempty"`
}

type CIDRClaimList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items []CIDRClaim `json:"items"`
}

func (in *CIDRClaim) DeepCopyInto(out *CIDRClaim) { *out = *in }
func (in *CIDRClaim) DeepCopy() *CIDRClaim { if in == nil { return nil }; out := new(CIDRClaim); in.DeepCopyInto(out); return out }
func (in *CIDRClaim) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *CIDRClaimList) DeepCopyInto(out *CIDRClaimList) { *out = *in; if in.Items != nil { out.Items = make([]CIDRClaim, len(in.Items)); copy(out.Items, in.Items) } }
func (in *CIDRClaimList) DeepCopy() *CIDRClaimList { if in == nil { return nil }; out := new(CIDRClaimList); in.DeepCopyInto(out); return out }
func (in *CIDRClaimList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// IPPoolSpec defines the desired state of IPPool.
type IPPoolSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type IPPoolStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=ipp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type IPPool struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec IPPoolSpec `json:"spec,omitempty"`; Status IPPoolStatus `json:"status,omitempty"` }
type IPPoolList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []IPPool `json:"items"` }
func (in *IPPool) DeepCopyInto(out *IPPool) { *out = *in }
func (in *IPPool) DeepCopy() *IPPool { if in == nil { return nil }; out := new(IPPool); in.DeepCopyInto(out); return out }
func (in *IPPool) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *IPPoolList) DeepCopyInto(out *IPPoolList) { *out = *in; if in.Items != nil { out.Items = make([]IPPool, len(in.Items)); copy(out.Items, in.Items) } }
func (in *IPPoolList) DeepCopy() *IPPoolList { if in == nil { return nil }; out := new(IPPoolList); in.DeepCopyInto(out); return out }
func (in *IPPoolList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// VLANSpec defines the desired state of VLAN.
type VLANSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type VLANStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vla
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type VLAN struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec VLANSpec `json:"spec,omitempty"`; Status VLANStatus `json:"status,omitempty"` }
type VLANList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []VLAN `json:"items"` }
func (in *VLAN) DeepCopyInto(out *VLAN) { *out = *in }
func (in *VLAN) DeepCopy() *VLAN { if in == nil { return nil }; out := new(VLAN); in.DeepCopyInto(out); return out }
func (in *VLAN) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *VLANList) DeepCopyInto(out *VLANList) { *out = *in; if in.Items != nil { out.Items = make([]VLAN, len(in.Items)); copy(out.Items, in.Items) } }
func (in *VLANList) DeepCopy() *VLANList { if in == nil { return nil }; out := new(VLANList); in.DeepCopyInto(out); return out }
func (in *VLANList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// LinkSpec defines the desired state of Link.
type LinkSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type LinkStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=lin
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Link struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec LinkSpec `json:"spec,omitempty"`; Status LinkStatus `json:"status,omitempty"` }
type LinkList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Link `json:"items"` }
func (in *Link) DeepCopyInto(out *Link) { *out = *in }
func (in *Link) DeepCopy() *Link { if in == nil { return nil }; out := new(Link); in.DeepCopyInto(out); return out }
func (in *Link) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *LinkList) DeepCopyInto(out *LinkList) { *out = *in; if in.Items != nil { out.Items = make([]Link, len(in.Items)); copy(out.Items, in.Items) } }
func (in *LinkList) DeepCopy() *LinkList { if in == nil { return nil }; out := new(LinkList); in.DeepCopyInto(out); return out }
func (in *LinkList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// ServerSpec defines the desired state of Server.
type ServerSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type ServerStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=ser
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Server struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec ServerSpec `json:"spec,omitempty"`; Status ServerStatus `json:"status,omitempty"` }
type ServerList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Server `json:"items"` }
func (in *Server) DeepCopyInto(out *Server) { *out = *in }
func (in *Server) DeepCopy() *Server { if in == nil { return nil }; out := new(Server); in.DeepCopyInto(out); return out }
func (in *Server) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *ServerList) DeepCopyInto(out *ServerList) { *out = *in; if in.Items != nil { out.Items = make([]Server, len(in.Items)); copy(out.Items, in.Items) } }
func (in *ServerList) DeepCopy() *ServerList { if in == nil { return nil }; out := new(ServerList); in.DeepCopyInto(out); return out }
func (in *ServerList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// RouterSpec defines the desired state of Router.
type RouterSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type RouterStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=rou
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Router struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec RouterSpec `json:"spec,omitempty"`; Status RouterStatus `json:"status,omitempty"` }
type RouterList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Router `json:"items"` }
func (in *Router) DeepCopyInto(out *Router) { *out = *in }
func (in *Router) DeepCopy() *Router { if in == nil { return nil }; out := new(Router); in.DeepCopyInto(out); return out }
func (in *Router) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *RouterList) DeepCopyInto(out *RouterList) { *out = *in; if in.Items != nil { out.Items = make([]Router, len(in.Items)); copy(out.Items, in.Items) } }
func (in *RouterList) DeepCopy() *RouterList { if in == nil { return nil }; out := new(RouterList); in.DeepCopyInto(out); return out }
func (in *RouterList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// SwitchSpec defines the desired state of Switch.
type SwitchSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type SwitchStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=swi
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Switch struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec SwitchSpec `json:"spec,omitempty"`; Status SwitchStatus `json:"status,omitempty"` }
type SwitchList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Switch `json:"items"` }
func (in *Switch) DeepCopyInto(out *Switch) { *out = *in }
func (in *Switch) DeepCopy() *Switch { if in == nil { return nil }; out := new(Switch); in.DeepCopyInto(out); return out }
func (in *Switch) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *SwitchList) DeepCopyInto(out *SwitchList) { *out = *in; if in.Items != nil { out.Items = make([]Switch, len(in.Items)); copy(out.Items, in.Items) } }
func (in *SwitchList) DeepCopy() *SwitchList { if in == nil { return nil }; out := new(SwitchList); in.DeepCopyInto(out); return out }
func (in *SwitchList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// FirewallSpec defines the desired state of Firewall.
type FirewallSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type FirewallStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=fir
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Firewall struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec FirewallSpec `json:"spec,omitempty"`; Status FirewallStatus `json:"status,omitempty"` }
type FirewallList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Firewall `json:"items"` }
func (in *Firewall) DeepCopyInto(out *Firewall) { *out = *in }
func (in *Firewall) DeepCopy() *Firewall { if in == nil { return nil }; out := new(Firewall); in.DeepCopyInto(out); return out }
func (in *Firewall) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *FirewallList) DeepCopyInto(out *FirewallList) { *out = *in; if in.Items != nil { out.Items = make([]Firewall, len(in.Items)); copy(out.Items, in.Items) } }
func (in *FirewallList) DeepCopy() *FirewallList { if in == nil { return nil }; out := new(FirewallList); in.DeepCopyInto(out); return out }
func (in *FirewallList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// FirewallPolicySpec defines the desired state of FirewallPolicy.
type FirewallPolicySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type FirewallPolicyStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=fwp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type FirewallPolicy struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec FirewallPolicySpec `json:"spec,omitempty"`; Status FirewallPolicyStatus `json:"status,omitempty"` }
type FirewallPolicyList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []FirewallPolicy `json:"items"` }
func (in *FirewallPolicy) DeepCopyInto(out *FirewallPolicy) { *out = *in }
func (in *FirewallPolicy) DeepCopy() *FirewallPolicy { if in == nil { return nil }; out := new(FirewallPolicy); in.DeepCopyInto(out); return out }
func (in *FirewallPolicy) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *FirewallPolicyList) DeepCopyInto(out *FirewallPolicyList) { *out = *in; if in.Items != nil { out.Items = make([]FirewallPolicy, len(in.Items)); copy(out.Items, in.Items) } }
func (in *FirewallPolicyList) DeepCopy() *FirewallPolicyList { if in == nil { return nil }; out := new(FirewallPolicyList); in.DeepCopyInto(out); return out }
func (in *FirewallPolicyList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// VPNSpec defines the desired state of VPN.
type VPNSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type VPNStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vpn
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type VPN struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec VPNSpec `json:"spec,omitempty"`; Status VPNStatus `json:"status,omitempty"` }
type VPNList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []VPN `json:"items"` }
func (in *VPN) DeepCopyInto(out *VPN) { *out = *in }
func (in *VPN) DeepCopy() *VPN { if in == nil { return nil }; out := new(VPN); in.DeepCopyInto(out); return out }
func (in *VPN) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *VPNList) DeepCopyInto(out *VPNList) { *out = *in; if in.Items != nil { out.Items = make([]VPN, len(in.Items)); copy(out.Items, in.Items) } }
func (in *VPNList) DeepCopy() *VPNList { if in == nil { return nil }; out := new(VPNList); in.DeepCopyInto(out); return out }
func (in *VPNList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// GatewaySpec defines the desired state of Gateway.
type GatewaySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type GatewayStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=gat
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Gateway struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec GatewaySpec `json:"spec,omitempty"`; Status GatewayStatus `json:"status,omitempty"` }
type GatewayList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Gateway `json:"items"` }
func (in *Gateway) DeepCopyInto(out *Gateway) { *out = *in }
func (in *Gateway) DeepCopy() *Gateway { if in == nil { return nil }; out := new(Gateway); in.DeepCopyInto(out); return out }
func (in *Gateway) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *GatewayList) DeepCopyInto(out *GatewayList) { *out = *in; if in.Items != nil { out.Items = make([]Gateway, len(in.Items)); copy(out.Items, in.Items) } }
func (in *GatewayList) DeepCopy() *GatewayList { if in == nil { return nil }; out := new(GatewayList); in.DeepCopyInto(out); return out }
func (in *GatewayList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// DNSSpec defines the desired state of DNS.
type DNSSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type DNSStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=dns
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type DNS struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec DNSSpec `json:"spec,omitempty"`; Status DNSStatus `json:"status,omitempty"` }
type DNSList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []DNS `json:"items"` }
func (in *DNS) DeepCopyInto(out *DNS) { *out = *in }
func (in *DNS) DeepCopy() *DNS { if in == nil { return nil }; out := new(DNS); in.DeepCopyInto(out); return out }
func (in *DNS) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *DNSList) DeepCopyInto(out *DNSList) { *out = *in; if in.Items != nil { out.Items = make([]DNS, len(in.Items)); copy(out.Items, in.Items) } }
func (in *DNSList) DeepCopy() *DNSList { if in == nil { return nil }; out := new(DNSList); in.DeepCopyInto(out); return out }
func (in *DNSList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// DHCPSpec defines the desired state of DHCP.
type DHCPSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type DHCPStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=dhc
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type DHCP struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec DHCPSpec `json:"spec,omitempty"`; Status DHCPStatus `json:"status,omitempty"` }
type DHCPList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []DHCP `json:"items"` }
func (in *DHCP) DeepCopyInto(out *DHCP) { *out = *in }
func (in *DHCP) DeepCopy() *DHCP { if in == nil { return nil }; out := new(DHCP); in.DeepCopyInto(out); return out }
func (in *DHCP) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *DHCPList) DeepCopyInto(out *DHCPList) { *out = *in; if in.Items != nil { out.Items = make([]DHCP, len(in.Items)); copy(out.Items, in.Items) } }
func (in *DHCPList) DeepCopy() *DHCPList { if in == nil { return nil }; out := new(DHCPList); in.DeepCopyInto(out); return out }
func (in *DHCPList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// StorageSpec defines the desired state of Storage.
type StorageSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type StorageStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=sto
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Storage struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec StorageSpec `json:"spec,omitempty"`; Status StorageStatus `json:"status,omitempty"` }
type StorageList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Storage `json:"items"` }
func (in *Storage) DeepCopyInto(out *Storage) { *out = *in }
func (in *Storage) DeepCopy() *Storage { if in == nil { return nil }; out := new(Storage); in.DeepCopyInto(out); return out }
func (in *Storage) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *StorageList) DeepCopyInto(out *StorageList) { *out = *in; if in.Items != nil { out.Items = make([]Storage, len(in.Items)); copy(out.Items, in.Items) } }
func (in *StorageList) DeepCopy() *StorageList { if in == nil { return nil }; out := new(StorageList); in.DeepCopyInto(out); return out }
func (in *StorageList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// VolumeSpec defines the desired state of Volume.
type VolumeSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type VolumeStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vol
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Volume struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec VolumeSpec `json:"spec,omitempty"`; Status VolumeStatus `json:"status,omitempty"` }
type VolumeList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Volume `json:"items"` }
func (in *Volume) DeepCopyInto(out *Volume) { *out = *in }
func (in *Volume) DeepCopy() *Volume { if in == nil { return nil }; out := new(Volume); in.DeepCopyInto(out); return out }
func (in *Volume) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *VolumeList) DeepCopyInto(out *VolumeList) { *out = *in; if in.Items != nil { out.Items = make([]Volume, len(in.Items)); copy(out.Items, in.Items) } }
func (in *VolumeList) DeepCopy() *VolumeList { if in == nil { return nil }; out := new(VolumeList); in.DeepCopyInto(out); return out }
func (in *VolumeList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// DatabaseSpec defines the desired state of Database.
type DatabaseSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type DatabaseStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=dat
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Database struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec DatabaseSpec `json:"spec,omitempty"`; Status DatabaseStatus `json:"status,omitempty"` }
type DatabaseList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Database `json:"items"` }
func (in *Database) DeepCopyInto(out *Database) { *out = *in }
func (in *Database) DeepCopy() *Database { if in == nil { return nil }; out := new(Database); in.DeepCopyInto(out); return out }
func (in *Database) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) { *out = *in; if in.Items != nil { out.Items = make([]Database, len(in.Items)); copy(out.Items, in.Items) } }
func (in *DatabaseList) DeepCopy() *DatabaseList { if in == nil { return nil }; out := new(DatabaseList); in.DeepCopyInto(out); return out }
func (in *DatabaseList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// RegistrySpec defines the desired state of Registry.
type RegistrySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type RegistryStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=reg
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Registry struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec RegistrySpec `json:"spec,omitempty"`; Status RegistryStatus `json:"status,omitempty"` }
type RegistryList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Registry `json:"items"` }
func (in *Registry) DeepCopyInto(out *Registry) { *out = *in }
func (in *Registry) DeepCopy() *Registry { if in == nil { return nil }; out := new(Registry); in.DeepCopyInto(out); return out }
func (in *Registry) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *RegistryList) DeepCopyInto(out *RegistryList) { *out = *in; if in.Items != nil { out.Items = make([]Registry, len(in.Items)); copy(out.Items, in.Items) } }
func (in *RegistryList) DeepCopy() *RegistryList { if in == nil { return nil }; out := new(RegistryList); in.DeepCopyInto(out); return out }
func (in *RegistryList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// LoadBalancerSpec defines the desired state of LoadBalancer.
type LoadBalancerSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type LoadBalancerStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=loa
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type LoadBalancer struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec LoadBalancerSpec `json:"spec,omitempty"`; Status LoadBalancerStatus `json:"status,omitempty"` }
type LoadBalancerList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []LoadBalancer `json:"items"` }
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) { *out = *in }
func (in *LoadBalancer) DeepCopy() *LoadBalancer { if in == nil { return nil }; out := new(LoadBalancer); in.DeepCopyInto(out); return out }
func (in *LoadBalancer) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *LoadBalancerList) DeepCopyInto(out *LoadBalancerList) { *out = *in; if in.Items != nil { out.Items = make([]LoadBalancer, len(in.Items)); copy(out.Items, in.Items) } }
func (in *LoadBalancerList) DeepCopy() *LoadBalancerList { if in == nil { return nil }; out := new(LoadBalancerList); in.DeepCopyInto(out); return out }
func (in *LoadBalancerList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// ServiceSpec defines the desired state of Service.
type ServiceSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type ServiceStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=svc
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Service struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec ServiceSpec `json:"spec,omitempty"`; Status ServiceStatus `json:"status,omitempty"` }
type ServiceList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Service `json:"items"` }
func (in *Service) DeepCopyInto(out *Service) { *out = *in }
func (in *Service) DeepCopy() *Service { if in == nil { return nil }; out := new(Service); in.DeepCopyInto(out); return out }
func (in *Service) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *ServiceList) DeepCopyInto(out *ServiceList) { *out = *in; if in.Items != nil { out.Items = make([]Service, len(in.Items)); copy(out.Items, in.Items) } }
func (in *ServiceList) DeepCopy() *ServiceList { if in == nil { return nil }; out := new(ServiceList); in.DeepCopyInto(out); return out }
func (in *ServiceList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// NATSpec defines the desired state of NAT.
type NATSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type NATStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=nat
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type NAT struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec NATSpec `json:"spec,omitempty"`; Status NATStatus `json:"status,omitempty"` }
type NATList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []NAT `json:"items"` }
func (in *NAT) DeepCopyInto(out *NAT) { *out = *in }
func (in *NAT) DeepCopy() *NAT { if in == nil { return nil }; out := new(NAT); in.DeepCopyInto(out); return out }
func (in *NAT) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *NATList) DeepCopyInto(out *NATList) { *out = *in; if in.Items != nil { out.Items = make([]NAT, len(in.Items)); copy(out.Items, in.Items) } }
func (in *NATList) DeepCopy() *NATList { if in == nil { return nil }; out := new(NATList); in.DeepCopyInto(out); return out }
func (in *NATList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// InternetGatewaySpec defines the desired state of InternetGateway.
type InternetGatewaySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type InternetGatewayStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=ing
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type InternetGateway struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec InternetGatewaySpec `json:"spec,omitempty"`; Status InternetGatewayStatus `json:"status,omitempty"` }
type InternetGatewayList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []InternetGateway `json:"items"` }
func (in *InternetGateway) DeepCopyInto(out *InternetGateway) { *out = *in }
func (in *InternetGateway) DeepCopy() *InternetGateway { if in == nil { return nil }; out := new(InternetGateway); in.DeepCopyInto(out); return out }
func (in *InternetGateway) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *InternetGatewayList) DeepCopyInto(out *InternetGatewayList) { *out = *in; if in.Items != nil { out.Items = make([]InternetGateway, len(in.Items)); copy(out.Items, in.Items) } }
func (in *InternetGatewayList) DeepCopy() *InternetGatewayList { if in == nil { return nil }; out := new(InternetGatewayList); in.DeepCopyInto(out); return out }
func (in *InternetGatewayList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// RouteTableSpec defines the desired state of RouteTable.
type RouteTableSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type RouteTableStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=rot
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type RouteTable struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec RouteTableSpec `json:"spec,omitempty"`; Status RouteTableStatus `json:"status,omitempty"` }
type RouteTableList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []RouteTable `json:"items"` }
func (in *RouteTable) DeepCopyInto(out *RouteTable) { *out = *in }
func (in *RouteTable) DeepCopy() *RouteTable { if in == nil { return nil }; out := new(RouteTable); in.DeepCopyInto(out); return out }
func (in *RouteTable) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *RouteTableList) DeepCopyInto(out *RouteTableList) { *out = *in; if in.Items != nil { out.Items = make([]RouteTable, len(in.Items)); copy(out.Items, in.Items) } }
func (in *RouteTableList) DeepCopy() *RouteTableList { if in == nil { return nil }; out := new(RouteTableList); in.DeepCopyInto(out); return out }
func (in *RouteTableList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// RouteSpec defines the desired state of Route.
type RouteSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type RouteStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=rou
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Route struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec RouteSpec `json:"spec,omitempty"`; Status RouteStatus `json:"status,omitempty"` }
type RouteList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Route `json:"items"` }
func (in *Route) DeepCopyInto(out *Route) { *out = *in }
func (in *Route) DeepCopy() *Route { if in == nil { return nil }; out := new(Route); in.DeepCopyInto(out); return out }
func (in *Route) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *RouteList) DeepCopyInto(out *RouteList) { *out = *in; if in.Items != nil { out.Items = make([]Route, len(in.Items)); copy(out.Items, in.Items) } }
func (in *RouteList) DeepCopy() *RouteList { if in == nil { return nil }; out := new(RouteList); in.DeepCopyInto(out); return out }
func (in *RouteList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// InterfaceSpec defines the desired state of Interface.
type InterfaceSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type InterfaceStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=int
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Interface struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec InterfaceSpec `json:"spec,omitempty"`; Status InterfaceStatus `json:"status,omitempty"` }
type InterfaceList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Interface `json:"items"` }
func (in *Interface) DeepCopyInto(out *Interface) { *out = *in }
func (in *Interface) DeepCopy() *Interface { if in == nil { return nil }; out := new(Interface); in.DeepCopyInto(out); return out }
func (in *Interface) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *InterfaceList) DeepCopyInto(out *InterfaceList) { *out = *in; if in.Items != nil { out.Items = make([]Interface, len(in.Items)); copy(out.Items, in.Items) } }
func (in *InterfaceList) DeepCopy() *InterfaceList { if in == nil { return nil }; out := new(InterfaceList); in.DeepCopyInto(out); return out }
func (in *InterfaceList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// BridgeSpec defines the desired state of Bridge.
type BridgeSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type BridgeStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=bri
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Bridge struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec BridgeSpec `json:"spec,omitempty"`; Status BridgeStatus `json:"status,omitempty"` }
type BridgeList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Bridge `json:"items"` }
func (in *Bridge) DeepCopyInto(out *Bridge) { *out = *in }
func (in *Bridge) DeepCopy() *Bridge { if in == nil { return nil }; out := new(Bridge); in.DeepCopyInto(out); return out }
func (in *Bridge) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *BridgeList) DeepCopyInto(out *BridgeList) { *out = *in; if in.Items != nil { out.Items = make([]Bridge, len(in.Items)); copy(out.Items, in.Items) } }
func (in *BridgeList) DeepCopy() *BridgeList { if in == nil { return nil }; out := new(BridgeList); in.DeepCopyInto(out); return out }
func (in *BridgeList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// NodePoolSpec defines the desired state of NodePool.
type NodePoolSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type NodePoolStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=nop
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type NodePool struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec NodePoolSpec `json:"spec,omitempty"`; Status NodePoolStatus `json:"status,omitempty"` }
type NodePoolList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []NodePool `json:"items"` }
func (in *NodePool) DeepCopyInto(out *NodePool) { *out = *in }
func (in *NodePool) DeepCopy() *NodePool { if in == nil { return nil }; out := new(NodePool); in.DeepCopyInto(out); return out }
func (in *NodePool) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) { *out = *in; if in.Items != nil { out.Items = make([]NodePool, len(in.Items)); copy(out.Items, in.Items) } }
func (in *NodePoolList) DeepCopy() *NodePoolList { if in == nil { return nil }; out := new(NodePoolList); in.DeepCopyInto(out); return out }
func (in *NodePoolList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// MonitoringSpec defines the desired state of Monitoring.
type MonitoringSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type MonitoringStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=mon
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Monitoring struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec MonitoringSpec `json:"spec,omitempty"`; Status MonitoringStatus `json:"status,omitempty"` }
type MonitoringList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Monitoring `json:"items"` }
func (in *Monitoring) DeepCopyInto(out *Monitoring) { *out = *in }
func (in *Monitoring) DeepCopy() *Monitoring { if in == nil { return nil }; out := new(Monitoring); in.DeepCopyInto(out); return out }
func (in *Monitoring) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *MonitoringList) DeepCopyInto(out *MonitoringList) { *out = *in; if in.Items != nil { out.Items = make([]Monitoring, len(in.Items)); copy(out.Items, in.Items) } }
func (in *MonitoringList) DeepCopy() *MonitoringList { if in == nil { return nil }; out := new(MonitoringList); in.DeepCopyInto(out); return out }
func (in *MonitoringList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// LoggingSpec defines the desired state of Logging.
type LoggingSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type LoggingStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=log
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Logging struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec LoggingSpec `json:"spec,omitempty"`; Status LoggingStatus `json:"status,omitempty"` }
type LoggingList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Logging `json:"items"` }
func (in *Logging) DeepCopyInto(out *Logging) { *out = *in }
func (in *Logging) DeepCopy() *Logging { if in == nil { return nil }; out := new(Logging); in.DeepCopyInto(out); return out }
func (in *Logging) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *LoggingList) DeepCopyInto(out *LoggingList) { *out = *in; if in.Items != nil { out.Items = make([]Logging, len(in.Items)); copy(out.Items, in.Items) } }
func (in *LoggingList) DeepCopy() *LoggingList { if in == nil { return nil }; out := new(LoggingList); in.DeepCopyInto(out); return out }
func (in *LoggingList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// SecretStoreSpec defines the desired state of SecretStore.
type SecretStoreSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type SecretStoreStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=sec
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type SecretStore struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec SecretStoreSpec `json:"spec,omitempty"`; Status SecretStoreStatus `json:"status,omitempty"` }
type SecretStoreList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []SecretStore `json:"items"` }
func (in *SecretStore) DeepCopyInto(out *SecretStore) { *out = *in }
func (in *SecretStore) DeepCopy() *SecretStore { if in == nil { return nil }; out := new(SecretStore); in.DeepCopyInto(out); return out }
func (in *SecretStore) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *SecretStoreList) DeepCopyInto(out *SecretStoreList) { *out = *in; if in.Items != nil { out.Items = make([]SecretStore, len(in.Items)); copy(out.Items, in.Items) } }
func (in *SecretStoreList) DeepCopy() *SecretStoreList { if in == nil { return nil }; out := new(SecretStoreList); in.DeepCopyInto(out); return out }
func (in *SecretStoreList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// CertificateSpec defines the desired state of Certificate.
type CertificateSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type CertificateStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=cer
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Certificate struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec CertificateSpec `json:"spec,omitempty"`; Status CertificateStatus `json:"status,omitempty"` }
type CertificateList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []Certificate `json:"items"` }
func (in *Certificate) DeepCopyInto(out *Certificate) { *out = *in }
func (in *Certificate) DeepCopy() *Certificate { if in == nil { return nil }; out := new(Certificate); in.DeepCopyInto(out); return out }
func (in *Certificate) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *CertificateList) DeepCopyInto(out *CertificateList) { *out = *in; if in.Items != nil { out.Items = make([]Certificate, len(in.Items)); copy(out.Items, in.Items) } }
func (in *CertificateList) DeepCopy() *CertificateList { if in == nil { return nil }; out := new(CertificateList); in.DeepCopyInto(out); return out }
func (in *CertificateList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// BackupPolicySpec defines the desired state of BackupPolicy.
type BackupPolicySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type BackupPolicyStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=bak
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type BackupPolicy struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec BackupPolicySpec `json:"spec,omitempty"`; Status BackupPolicyStatus `json:"status,omitempty"` }
type BackupPolicyList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []BackupPolicy `json:"items"` }
func (in *BackupPolicy) DeepCopyInto(out *BackupPolicy) { *out = *in }
func (in *BackupPolicy) DeepCopy() *BackupPolicy { if in == nil { return nil }; out := new(BackupPolicy); in.DeepCopyInto(out); return out }
func (in *BackupPolicy) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *BackupPolicyList) DeepCopyInto(out *BackupPolicyList) { *out = *in; if in.Items != nil { out.Items = make([]BackupPolicy, len(in.Items)); copy(out.Items, in.Items) } }
func (in *BackupPolicyList) DeepCopy() *BackupPolicyList { if in == nil { return nil }; out := new(BackupPolicyList); in.DeepCopyInto(out); return out }
func (in *BackupPolicyList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// FailureSimulationSpec defines the desired state of FailureSimulation.
type FailureSimulationSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type FailureSimulationStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=fai
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type FailureSimulation struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec FailureSimulationSpec `json:"spec,omitempty"`; Status FailureSimulationStatus `json:"status,omitempty"` }
type FailureSimulationList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []FailureSimulation `json:"items"` }
func (in *FailureSimulation) DeepCopyInto(out *FailureSimulation) { *out = *in }
func (in *FailureSimulation) DeepCopy() *FailureSimulation { if in == nil { return nil }; out := new(FailureSimulation); in.DeepCopyInto(out); return out }
func (in *FailureSimulation) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *FailureSimulationList) DeepCopyInto(out *FailureSimulationList) { *out = *in; if in.Items != nil { out.Items = make([]FailureSimulation, len(in.Items)); copy(out.Items, in.Items) } }
func (in *FailureSimulationList) DeepCopy() *FailureSimulationList { if in == nil { return nil }; out := new(FailureSimulationList); in.DeepCopyInto(out); return out }
func (in *FailureSimulationList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// TrafficPolicySpec defines the desired state of TrafficPolicy.
type TrafficPolicySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type TrafficPolicyStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=tra
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type TrafficPolicy struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec TrafficPolicySpec `json:"spec,omitempty"`; Status TrafficPolicyStatus `json:"status,omitempty"` }
type TrafficPolicyList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []TrafficPolicy `json:"items"` }
func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) { *out = *in }
func (in *TrafficPolicy) DeepCopy() *TrafficPolicy { if in == nil { return nil }; out := new(TrafficPolicy); in.DeepCopyInto(out); return out }
func (in *TrafficPolicy) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *TrafficPolicyList) DeepCopyInto(out *TrafficPolicyList) { *out = *in; if in.Items != nil { out.Items = make([]TrafficPolicy, len(in.Items)); copy(out.Items, in.Items) } }
func (in *TrafficPolicyList) DeepCopy() *TrafficPolicyList { if in == nil { return nil }; out := new(TrafficPolicyList); in.DeepCopyInto(out); return out }
func (in *TrafficPolicyList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// QoSPolicySpec defines the desired state of QoSPolicy.
type QoSPolicySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type QoSPolicyStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=qos
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type QoSPolicy struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec QoSPolicySpec `json:"spec,omitempty"`; Status QoSPolicyStatus `json:"status,omitempty"` }
type QoSPolicyList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []QoSPolicy `json:"items"` }
func (in *QoSPolicy) DeepCopyInto(out *QoSPolicy) { *out = *in }
func (in *QoSPolicy) DeepCopy() *QoSPolicy { if in == nil { return nil }; out := new(QoSPolicy); in.DeepCopyInto(out); return out }
func (in *QoSPolicy) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *QoSPolicyList) DeepCopyInto(out *QoSPolicyList) { *out = *in; if in.Items != nil { out.Items = make([]QoSPolicy, len(in.Items)); copy(out.Items, in.Items) } }
func (in *QoSPolicyList) DeepCopy() *QoSPolicyList { if in == nil { return nil }; out := new(QoSPolicyList); in.DeepCopyInto(out); return out }
func (in *QoSPolicyList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// SecurityGroupSpec defines the desired state of SecurityGroup.
type SecurityGroupSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type SecurityGroupStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=sec
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type SecurityGroup struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec SecurityGroupSpec `json:"spec,omitempty"`; Status SecurityGroupStatus `json:"status,omitempty"` }
type SecurityGroupList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []SecurityGroup `json:"items"` }
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) { *out = *in }
func (in *SecurityGroup) DeepCopy() *SecurityGroup { if in == nil { return nil }; out := new(SecurityGroup); in.DeepCopyInto(out); return out }
func (in *SecurityGroup) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *SecurityGroupList) DeepCopyInto(out *SecurityGroupList) { *out = *in; if in.Items != nil { out.Items = make([]SecurityGroup, len(in.Items)); copy(out.Items, in.Items) } }
func (in *SecurityGroupList) DeepCopy() *SecurityGroupList { if in == nil { return nil }; out := new(SecurityGroupList); in.DeepCopyInto(out); return out }
func (in *SecurityGroupList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// AccessPolicySpec defines the desired state of AccessPolicy.
type AccessPolicySpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type AccessPolicyStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=acc
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type AccessPolicy struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec AccessPolicySpec `json:"spec,omitempty"`; Status AccessPolicyStatus `json:"status,omitempty"` }
type AccessPolicyList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []AccessPolicy `json:"items"` }
func (in *AccessPolicy) DeepCopyInto(out *AccessPolicy) { *out = *in }
func (in *AccessPolicy) DeepCopy() *AccessPolicy { if in == nil { return nil }; out := new(AccessPolicy); in.DeepCopyInto(out); return out }
func (in *AccessPolicy) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *AccessPolicyList) DeepCopyInto(out *AccessPolicyList) { *out = *in; if in.Items != nil { out.Items = make([]AccessPolicy, len(in.Items)); copy(out.Items, in.Items) } }
func (in *AccessPolicyList) DeepCopy() *AccessPolicyList { if in == nil { return nil }; out := new(AccessPolicyList); in.DeepCopyInto(out); return out }
func (in *AccessPolicyList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// NamespaceNetworkSpec defines the desired state of NamespaceNetwork.
type NamespaceNetworkSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type NamespaceNetworkStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=nsn
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type NamespaceNetwork struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec NamespaceNetworkSpec `json:"spec,omitempty"`; Status NamespaceNetworkStatus `json:"status,omitempty"` }
type NamespaceNetworkList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []NamespaceNetwork `json:"items"` }
func (in *NamespaceNetwork) DeepCopyInto(out *NamespaceNetwork) { *out = *in }
func (in *NamespaceNetwork) DeepCopy() *NamespaceNetwork { if in == nil { return nil }; out := new(NamespaceNetwork); in.DeepCopyInto(out); return out }
func (in *NamespaceNetwork) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *NamespaceNetworkList) DeepCopyInto(out *NamespaceNetworkList) { *out = *in; if in.Items != nil { out.Items = make([]NamespaceNetwork, len(in.Items)); copy(out.Items, in.Items) } }
func (in *NamespaceNetworkList) DeepCopy() *NamespaceNetworkList { if in == nil { return nil }; out := new(NamespaceNetworkList); in.DeepCopyInto(out); return out }
func (in *NamespaceNetworkList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }

// AirGapClusterSpec defines the desired state of AirGapCluster.
type AirGapClusterSpec struct { Description string `json:"description,omitempty"`; Enabled bool `json:"enabled,omitempty"`; NetworkRef string `json:"networkRef,omitempty"`; CIDR string `json:"cidr,omitempty"`; Type string `json:"type,omitempty"`; Address string `json:"address,omitempty"` }
type AirGapClusterStatus struct { CommonStatus `json:",inline"`; Message string `json:"message,omitempty"` }
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=air
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type AirGapCluster struct { metav1.TypeMeta `json:",inline"`; metav1.ObjectMeta `json:"metadata,omitempty"`; Spec AirGapClusterSpec `json:"spec,omitempty"`; Status AirGapClusterStatus `json:"status,omitempty"` }
type AirGapClusterList struct { metav1.TypeMeta `json:",inline"`; metav1.ListMeta `json:"metadata,omitempty"`; Items []AirGapCluster `json:"items"` }
func (in *AirGapCluster) DeepCopyInto(out *AirGapCluster) { *out = *in }
func (in *AirGapCluster) DeepCopy() *AirGapCluster { if in == nil { return nil }; out := new(AirGapCluster); in.DeepCopyInto(out); return out }
func (in *AirGapCluster) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
func (in *AirGapClusterList) DeepCopyInto(out *AirGapClusterList) { *out = *in; if in.Items != nil { out.Items = make([]AirGapCluster, len(in.Items)); copy(out.Items, in.Items) } }
func (in *AirGapClusterList) DeepCopy() *AirGapClusterList { if in == nil { return nil }; out := new(AirGapClusterList); in.DeepCopyInto(out); return out }
func (in *AirGapClusterList) DeepCopyObject() runtime.Object { if c := in.DeepCopy(); c != nil { return c }; return nil }
