#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MANIFEST="${1:-$ROOT_DIR/examples/topology.yaml}"

if ! command -v kubectl >/dev/null 2>&1; then
  echo "kubectl is required to apply the air-gap system manifest." >&2
  exit 1
fi

kubectl apply -f "$MANIFEST"
"$ROOT_DIR/scripts/deploy-ip-ssh-runtime.sh" "$MANIFEST"
