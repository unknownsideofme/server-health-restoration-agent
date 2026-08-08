# AirGap Autonomous AI NOC Copilot & Kubernetes Operator

> **Predictive Multi-Org Infrastructure Telemetry, Local Offline LLM Decision Support, MCP Protocol Skill Self-Healing, and Universal 10-Resource Kubernetes Operator**

---

## 📌 Executive Architecture & System Purpose

The **AirGap Autonomous AI NOC Copilot** is a production-grade, 100% offline-ready network operations center (NOC) automation system. Designed for high-security air-gapped environments, enterprise multi-tenant datacenters, and mission-critical cloud infrastructure, it continuously monitors hardware and network topologies, predicts failure lead times (Time-to-Impact / TTI), and autonomously executes deep OS and network repair skills via the **Model Context Protocol (MCP)**.

```
                     ┌──────────────────────────────────────────────────────────┐
                     │            React 18 Single-Page Dashboard (Port 8085)     │
                     │  - Multi-Page Navigation (Overview, Orgs, Servers, Tors) │
                     │  - Org Scope Filter (Org A, Org B, Org C, Org D)        │
                     └──────────────────────────┬───────────────────────────────┘
                                                │
       ┌────────────────────────────────────────┼────────────────────────────────────────┐
       │                                        │                                        │
┌──────▼─────────────────────┐  ┌───────────────▼───────────────┐  ┌─────────────────────▼──────┐
│ Dedicated Local LLM Daemon │  │ AirGap Kubernetes Controller  │  │ Predictive Telemetry Engine │
│ Port 11435 (qwen2.5:0.5b)  │  │ Reconciles 10 CRDs in k9s     │  │ Sliding Window Rates & TTI  │
└──────────────┬─────────────┘  └───────────────┬───────────────┘  └──────────────┬─────────────┘
               │                                │                                │
┌──────────────▼─────────────┐  ┌───────────────▼───────────────┐  ┌──────────────▼─────────────┐
│ MCP Protocol Server        │  │ Kubernetes etcd & Event Stream│  │ Prometheus & Grafana Live  │
│ Deep OS & Router Repair    │  │ Warning / Normal Events       │  │ Scraper Port 9090 / 3000   │
└────────────────────────────┘  └───────────────────────────────┘  └────────────────────────────┘
```

---

## ✨ Core Features & Technical Highlights

1. **Universal 10-Resource Kubernetes Operator (`controllers/airgap_controller.py`)**:
   Continuously reconciles **10 Custom Resource Definitions (CRDs)** in etcd and `k9s`:
   - `orgs` (Multi-tenant Organization units)
   - `racks` (Hardware rack enclosures)
   - `tors` (Top-of-Rack access switches)
   - `subnets` (IPv4 CIDR network subnets)
   - `servers` (Control Plane / Admin Nodes & Compute / Workload Nodes)
   - `networks` (Overlay / Underlay network fabrics)
   - `routers` (Egress Edge Routers)
   - `routetables` (BGP / OSPF Route Tables)
   - `dashboards` (Web UI Application instances)
   - `llmmodels` (Deployed Local LLM Model engines)
   - Emits structured Kubernetes `Warning` / `Normal` events for every component visible directly inside **`k9s`**.

2. **Air-Gapped Offline Local LLM & MCP Protocol Skills**:
   - **Local Model**: Runs Ollama `qwen2.5:0.5b` (390MB GGUF) locally on port `11434`.
   - **Dedicated LLM Microservice Daemon (`pkg/copilot/llm_daemon.py`)**: Listens on port `11435` with non-blocking async execution, guaranteeing zero UI hangs or web request timeouts.
   - **MCP JSON-RPC 2.0 Server (`pkg/mcp/mcp_server.py`)**: Implements `mcp_list_tools` and `mcp_call_tool`.
   - **Agent Skills**:
     - `os_repair_skill.py`: Deep Linux kernel tuning, sysctl optimization, packet drop resolution.
     - `router_repair_skill.py`: BGP route flap damping, MTU mismatch fix, tunnel QoS shaping.
     - `autofix_skill.py`: Autonomous self-healing execution pipeline.

3. **React 18 Multi-Page Tabbed UI Dashboard (`ui/index.html`)**:
   - Built using **React 18** and **Babel**.
   - Tabbed navigation: `📊 Global Overview`, `🏢 Organizations`, `🖥️ Control Plane & Workloads`, `🔀 TOR Switches`, `🗄️ Infrastructure Racks`, `🌐 Edge Routers`.
   - Scope filter dropdown (`ALL`, `Org A`, `Org B`, `Org C`, `Org D`).
   - One-click **Execute Autonomous Repair (MCP Skill)** trigger.

4. **Prometheus Telemetry Exporter & Provisioned Grafana Live Dashboard**:
   - `/metrics` endpoint on port `8085` exporting Prometheus metrics format.
   - Provisioned Grafana dashboard (`dashboards/grafana-airgap-noc.json`).

---

## 🚀 How to Recreate & Deploy on a New Server / EC2 Instance

Follow these step-by-step instructions to recreate this exact system on a fresh Ubuntu 22.04 / 24.04 EC2 instance or Linux server.

### Step 1: System Package Installation & Prerequisites

Update your server and install required system utilities:

```bash
sudo apt-get update && sudo apt-get install -y \
    python3 python3-pip python3-venv git curl wget nginx openssl net-tools procps
```

### Step 2: Install Kubernetes (k3s) & k9s

Install lightweight Kubernetes (k3s):

```bash
curl -sfL https://get.k3s.io | sh -
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown ubuntu:ubuntu ~/.kube/config
export KUBECONFIG=~/.kube/config
```

*(Optional)* Install `k9s` CLI terminal UI:

```bash
curl -sS https://webinstall.dev/k9s | bash
```

### Step 3: Install Ollama & Pull Local LLM Model

Install Ollama for local air-gapped LLM inference:

```bash
curl -fsSL https://ollama.com/install.sh | sh
ollama pull qwen2.5:0.5b
```

### Step 4: Clone Repository & Install Python Dependencies

Clone this repository and install required Python libraries:

```bash
git clone https://github.com/your-org/airgap-k8s-crd.git
cd airgap-k8s-crd

pip3 install kubernetes pyyaml prometheus_client
```

### Step 5: Apply Kubernetes Custom Resource Definitions (CRDs) & Instances

Apply all 10 Custom Resource Definitions to your Kubernetes cluster:

```bash
kubectl apply -f crds/org.yaml
kubectl apply -f crds/rack.yaml
kubectl apply -f crds/tor.yaml
kubectl apply -f crds/subnet.yaml
kubectl apply -f crds/server.yaml
kubectl apply -f crds/network.yaml
kubectl apply -f crds/router.yaml
kubectl apply -f crds/routetable.yaml
kubectl apply -f crds/dashboard.yaml
kubectl apply -f crds/llmmodel.yaml

# Apply sample infrastructure instances
kubectl apply -f examples/crd_instances.yaml
kubectl apply -f examples/crd_instances_extension.yaml
```

Verify that all CRDs are active:

```bash
kubectl get crd
kubectl get orgs,racks,tors,subnets,servers,networks,routers,routetables,dashboards,llmmodels
```

---

### Step 6: Start All System Services & Daemons

#### 1. Start the Dedicated Local LLM Daemon (Port 11435):
```bash
nohup python3 pkg/copilot/llm_daemon.py > /tmp/llm_daemon.log 2>&1 &
```

#### 2. Start the NOC Copilot Web Dashboard Server (Port 8085):
```bash
nohup python3 ui/server.py > /tmp/ui_server.log 2>&1 &
```

#### 3. Start the Universal 10-Resource Kubernetes Operator Controller:
```bash
nohup python3 controllers/airgap_controller.py > /tmp/airgap_controller.log 2>&1 &
```

---

### Step 7: (Optional) Configure Prometheus & Grafana

#### Install Prometheus:
```bash
sudo apt-get install -y prometheus prometheus-node-exporter
sudo cp config/prometheus.yml /etc/prometheus/prometheus.yml
sudo systemctl restart prometheus
```

#### Install Grafana:
```bash
sudo apt-get install -y apt-transport-https software-properties-common
sudo mkdir -p /etc/apt/keyrings/
wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null
echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee /etc/apt/sources.list.d/grafana.list
sudo apt-get update && sudo apt-get install -y grafana
```

#### Provision AirGap Grafana Dashboard:
```bash
sudo mkdir -p /var/lib/grafana/dashboards
sudo cp dashboards/grafana-airgap-noc.json /var/lib/grafana/dashboards/airgap-noc-dashboard.json
sudo chown -R grafana:grafana /var/lib/grafana/dashboards
sudo systemctl restart grafana-server
```

---

## 🌐 Connecting to the System

### Option A: Via SSH Tunneling (Recommended for Corporate Networks)

If your company uses a firewall (e.g. FortiGate) that blocks non-standard HTTP ports, connect via SSH port forwarding from your local machine:

```bash
ssh -i /path/to/your-key.pem -N -L 3000:localhost:3000 -L 8085:localhost:8085 ubuntu@<YOUR-EC2-PUBLIC-IP>
```

Open in your browser:
- 🤖 **React 18 Web Dashboard & MCP UI**: `http://localhost:8085`
- 📊 **Grafana Live Telemetry Dashboard**: `http://localhost:3000/d/airgap-noc-copilot/airgap-autonomous-ai-noc-copilot-dashboard`

### Option B: Via Nginx HTTPS (Port 443)

To bypass corporate plain-HTTP inspection using SSL/TLS encryption:

```bash
sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout /etc/ssl/private/airgap-selfsigned.key \
  -out /etc/ssl/certs/airgap-selfsigned.crt \
  -subj "/C=US/ST=State/L=City/O=AirGap/OU=NOC/CN=<YOUR-EC2-PUBLIC-IP>"
```

Configure Nginx `/etc/nginx/sites-available/default`:

```nginx
server {
    listen 80 default_server;
    listen [::]:80 default_server;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl default_server;
    listen [::]:443 ssl default_server;

    ssl_certificate /etc/ssl/certs/airgap-selfsigned.crt;
    ssl_certificate_key /etc/ssl/private/airgap-selfsigned.key;

    location / {
        proxy_pass http://127.0.0.1:8085;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Restart Nginx:
```bash
sudo systemctl restart nginx
```

Access via browser: `https://<YOUR-EC2-PUBLIC-IP>` *(Accept self-signed certificate warning)*.

---

## 🧪 Testing & Verification Commands

1. **Verify All 10 Kubernetes CRDs in k9s / kubectl**:
   ```bash
   kubectl get dashboards,llmmodels,orgs,racks,tors,subnets,servers,networks,routers,routetables
   ```

2. **View Live Event Logs in k9s / kubectl**:
   ```bash
   kubectl get events --sort-by='.metadata.creationTimestamp'
   ```

3. **Run MCP Agent Skill Unit Tests**:
   ```bash
   python3 test/test_mcp_skills.py
   ```

4. **Inject a Progressive Fault to Test Autonomous Repair**:
   ```bash
   # Inject progressive link congestion
   python3 scripts/inject-fault.py --scenario congestion --target org-a-edge-router

   # Run CLI autonomous self-healing skill
   python3 scripts/noc-copilot-cli.py --autofix
   ```
