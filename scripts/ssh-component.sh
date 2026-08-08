#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 <component-name> [command...]"
  echo "Examples:"
  echo "  $0 org-a-admin-01"
  echo "  $0 org-a-edge-router"
  echo "  $0 org-b-worker-01 ip addr show"
  exit 1
fi

COMPONENT="$1"
shift 1 || true

if ! command -v kubectl >/dev/null 2>&1; then
  echo "Error: kubectl is required to SSH into components." >&2
  exit 1
fi

# Check if component Pod exists
if ! kubectl get pod "$COMPONENT" >/dev/null 2>&1; then
  echo "Component Pod '$COMPONENT' not found."
  echo "Deploying component SSH Pods first..."
  "$(dirname "$0")/deploy-component-pods.sh"
  
  echo "Waiting for '$COMPONENT' Pod to be ready..."
  kubectl wait --for=condition=Ready "pod/$COMPONENT" --timeout=60s
fi

POD_STATUS="$(kubectl get pod "$COMPONENT" -o jsonpath='{.status.phase}')"
if [ "$POD_STATUS" != "Running" ]; then
  echo "Waiting for component '$COMPONENT' (status: $POD_STATUS) to be Running..."
  kubectl wait --for=condition=Ready "pod/$COMPONENT" --timeout=60s
fi

if [ "$#" -gt 0 ]; then
  exec kubectl exec -it "$COMPONENT" -- /bin/bash -c "$*"
fi

# Check if sshpass / ssh is available and port forwarding can be used
LOCAL_PORT=$(shuf -i 22000-29000 -n 1)

if command -v sshpass >/dev/null 2>&1 && command -v ssh >/dev/null 2>&1; then
  echo "Establishing SSH tunnel to component '$COMPONENT'..."
  kubectl port-forward "pod/$COMPONENT" "${LOCAL_PORT}:22" >/dev/null 2>&1 &
  PF_PID=$!

  cleanup() {
    kill "$PF_PID" 2>/dev/null || true
  }
  trap cleanup EXIT

  sleep 1.5

  echo "=========================================================="
  echo " Connected to AirGap Component: $COMPONENT via SSH"
  echo " Password: airgap"
  echo "=========================================================="
  
  sshpass -p airgap ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null "root@127.0.0.1" -p "$LOCAL_PORT"
else
  echo "=========================================================="
  echo " Connecting to AirGap Component Shell: $COMPONENT"
  echo "=========================================================="
  exec kubectl exec -it "$COMPONENT" -- /bin/bash
fi
