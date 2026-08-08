#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MANIFEST="${1:-$ROOT_DIR/examples/topology.yaml}"

echo "Setting up AirGap native IP routing interface 'airgap0'..."

sudo ip link add airgap0 type dummy 2>/dev/null || true
sudo ip link set airgap0 up 2>/dev/null || true

python3 - "$MANIFEST" <<'EOF'
import sys
import yaml
import subprocess

manifest_file = sys.argv[1] if len(sys.argv) > 1 else "examples/topology.yaml"

with open(manifest_file, "r") as f:
    docs = list(yaml.safe_load_all(f))

ips = set()

for doc in docs:
    if not doc or "spec" not in doc:
        continue
    spec = doc["spec"]
    for k in ["ipAddress", "address", "managementIP", "gateway"]:
        if k in spec and spec[k] and "/" not in spec[k]:
            ips.add(spec[k])
    if "cidr" in spec and "/" in spec["cidr"]:
        # add base gateway IP
        parts = spec["cidr"].split("/")
        base_ip = parts[0]
        gw_parts = base_ip.split(".")
        gw_parts[-1] = "1"
        ips.add(".".join(gw_parts))

print(f"Configuring {len(ips)} native component IP addresses on 'airgap0' interface...")

for ip in sorted(ips):
    cmd = ["sudo", "ip", "addr", "add", f"{ip}/24", "dev", "airgap0"]
    subprocess.run(cmd, stderr=subprocess.DEVNULL)

EOF

echo "Native IP network interface 'airgap0' configured successfully."
