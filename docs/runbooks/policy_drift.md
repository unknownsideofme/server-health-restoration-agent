# Runbook: ACL Policy Drift & Controller Configuration Mismatch

## Overview
Triggered when controller ACL sync drops application traffic unexpectedly.

## Root Cause Analysis
1. Unsynchronized local config edit bypassing AirGap Kubernetes Operator reconciliation.
2. Stale security group rule blocking valid IP ranges.

## Recommended Corrective Actions
1. **Trigger Resync**: Re-apply controller CRD manifest (`kubectl apply -f crds/`).
2. **Audit ACL Rules**: Compare running ACL rules against `AirGap` CRD desired state.
