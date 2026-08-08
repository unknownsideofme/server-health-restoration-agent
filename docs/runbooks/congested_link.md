# Runbook: Progressive Link Congestion & Saturation Mitigation

## Overview
Triggered when link utilization rate of change ($dM/dt$) indicates bandwidth exhaustion within 20 minutes ($TTI < 20$).

## Root Cause Analysis
1. Traffic surge from top-talker application flows.
2. Routing path asymmetry directing secondary traffic onto primary trunk link.
3. Queue buffer tail-drop causing TCP window reduction.

## Recommended Corrective Actions
1. **Apply QoS Traffic Shaping**: Enable policing rule `shape-top-talkers-50M` on the affected egress interface (`eth0`).
2. **Re-route Non-Critical Flows**: Update RouteTable CRD `org-a-routes` to divert backup traffic onto secondary path `10.10.0.2`.
3. **Notify NOC Operator**: Pre-emptively scale bandwidth allocation before SLA breach.
