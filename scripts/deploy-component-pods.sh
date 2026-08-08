#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MANIFEST="${1:-$ROOT_DIR/examples/topology.yaml}"

if ! command -v kubectl >/dev/null 2>&1; then
  echo "kubectl is required to deploy SSH component pods." >&2
  exit 1
fi

echo "Deploying instant SSH-enabled Pod shims for components in $MANIFEST..."

python3 - "$MANIFEST" <<'EOF'
import os
import sys
import yaml
import subprocess

manifest_file = sys.argv[1] if len(sys.argv) > 1 else "examples/topology.yaml"

with open(manifest_file, "r") as f:
    docs = list(yaml.safe_load_all(f))

pod_manifests = []

for doc in docs:
    if not doc or "kind" not in doc or "metadata" not in doc:
        continue
    
    kind = doc["kind"]
    name = doc["metadata"]["name"]
    spec = doc.get("spec", {})
    
    # We create SSH Pods for Server, Router, and Tor components
    if kind not in ["Server", "Router", "Tor", "Cluster", "Switch"]:
        continue
        
    role = spec.get("role", kind.lower())
    ip = spec.get("ipAddress", spec.get("address", spec.get("managementIP", "10.10.0.11")))
    hostname = spec.get("hostname", f"{name}.airgap.local")
    
    bin_helpers = (
        f"echo 'export PS1=\"[{name} (SSH)]\\$ \"' >> /root/.bashrc && "
        f"echo 'export SSH_COMPONENT_NAME=\"{name}\"' >> /root/.bashrc && "
        f"echo 'export COMPONENT_IP=\"{ip}\"' >> /root/.bashrc && "
        f"cat << 'IFCONFIG_EOF' > /usr/local/bin/ifconfig\n"
        f"#!/bin/bash\n"
        f"if [ -x /sbin/ifconfig ]; then exec /sbin/ifconfig \"$@\"; fi\n"
        f"ip_addr=\"${{COMPONENT_IP:-{ip}}}\"\n"
        f"echo \"eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500\"\n"
        f"echo \"        inet ${{ip_addr}}  netmask 255.255.255.0  broadcast 10.10.0.255\"\n"
        f"echo \"        ether 02:42:ac:11:00:02  txqueuelen 1000  (Ethernet)\"\n"
        f"IFCONFIG_EOF\n"
        f"chmod +x /usr/local/bin/ifconfig && "
        f"cat << 'IPCONFIG_EOF' > /usr/local/bin/ipconfig\n"
        f"#!/bin/bash\n"
        f"ip_addr=\"${{COMPONENT_IP:-{ip}}}\"\n"
        f"echo \"AirGap Component Network Configuration for ${{SSH_COMPONENT_NAME:-{name}}}:\"\n"
        f"echo \"   IPv4 Address. . . . . . . . . . . : ${{ip_addr}}\"\n"
        f"echo \"   Subnet Mask . . . . . . . . . . . : 255.255.255.0\"\n"
        f"echo \"   Default Gateway . . . . . . . . . : 10.10.0.1\"\n"
        f"IPCONFIG_EOF\n"
        f"chmod +x /usr/local/bin/ipconfig && "
        f"cat << 'IP_EOF' > /usr/local/bin/ip\n"
        f"#!/bin/bash\n"
        f"if [ -x /sbin/ip ]; then exec /sbin/ip \"$@\"; fi\n"
        f"ip_addr=\"${{COMPONENT_IP:-{ip}}}\"\n"
        f"if [ \"$#\" -eq 0 ]; then echo \"eth0: ${{ip_addr}}/24\"; exit 0; fi\n"
        f"if [ \"$1\" = \"route\" ] || [ \"${{2:-}}\" = \"route\" ]; then\n"
        f"  echo \"default via 10.10.0.1 dev eth0\"\n"
        f"  echo \"10.10.0.0/24 dev eth0 scope link src ${{ip_addr}}\"\n"
        f"elif [ \"$1\" = \"addr\" ] || [ \"${{2:-}}\" = \"addr\" ] || [ \"$1\" = \"a\" ]; then\n"
        f"  echo \"1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN\"\n"
        f"  echo \"    inet 127.0.0.1/8 scope host lo\"\n"
        f"  echo \"2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP\"\n"
        f"  echo \"    inet ${{ip_addr}}/24 brd 10.10.0.255 scope global eth0\"\n"
        f"else\n"
        f"  echo \"eth0: ${{ip_addr}}/24\"\n"
        f"fi\n"
        f"IP_EOF\n"
        f"chmod +x /usr/local/bin/ip && "
    )

    startup_script = (
        bin_helpers +
        f"hostname {name} 2>/dev/null || true; "
        f"(apt-get update -qq && apt-get install -y -qq openssh-server iproute2 iputils-ping net-tools curl traceroute && "
        f"mkdir -p /var/run/sshd && echo 'root:airgap' | chpasswd && "
        f"sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && "
        f"service ssh restart) >/dev/null 2>&1 & "
        f"exec sleep infinity"
    )

    pod_yaml = {
        "apiVersion": "v1",
        "kind": "Pod",
        "metadata": {
            "name": name,
            "labels": {
                "app": "airgap-component",
                "component": name,
                "kind": kind.lower(),
                "role": role,
            }
        },
        "spec": {
            "restartPolicy": "Always",
            "containers": [
                {
                    "name": "ssh-server",
                    "image": "ubuntu:22.04",
                    "command": ["/bin/bash", "-c"],
                    "args": [startup_script],
                    "ports": [
                        {"name": "ssh", "containerPort": 22}
                    ],
                    "securityContext": {
                        "capabilities": {
                            "add": ["NET_ADMIN", "NET_RAW"]
                        }
                    }
                }
            ]
        }
    }
    pod_manifests.append(pod_yaml)

combined_yaml = yaml.dump_all(pod_manifests)

process = subprocess.Popen(["kubectl", "apply", "-f", "-"], stdin=subprocess.PIPE, text=True)
process.communicate(input=combined_yaml)

if process.returncode != 0:
    sys.exit(process.returncode)

print(f"Successfully provisioned {len(pod_manifests)} instant SSH component Pods.")
EOF
