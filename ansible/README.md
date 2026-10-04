# ???? AirGap Autonomous NOC Copilot ??? Automated Ansible Provisioner

This Ansible playbook automatically provisions and sets up the complete **AirGap Autonomous AI NOC Copilot & Infrastructure Health Restoration Agent** from scratch on any fresh Ubuntu/Debian server.

---

## ???? What This Playbook Sets Up Automatically

1. **System Core**: Python 3.12, Node.js 20 LTS, npm, build-essential, git, jq.
2. **Kubernetes Layer**: Installs **K3s** lightweight cluster, configures `kubectl`, and generates `~/.kube/config`.
3. **Telemetry & Metrics**:
   - Installs and configures **Prometheus** (Port `9090`).
   - Installs and enables **Prometheus Node Exporter** (Port `9100`).
   - Configures scrape jobs for live metrics ingestion (`airgap-noc-copilot`).
4. **Grafana Visualization**:
   - Installs official **Grafana Server** (Port `3000`).
   - Auto-provisions Prometheus datasource.
   - Auto-imports the **AirGap Autonomous AI NOC Copilot** dashboard (`dashboards/grafana-airgap-noc.json`) with all 71 panels.
5. **AirGap Application & Autonomous Agents**:
   - Clones repo, sets up `.venv`, and installs all Python requirements.
   - Installs and builds React frontend (`lucid/frontend/dist`).
   - Installs backend dependencies (`lucid/backend`).
   - Applies all 45+ AirGap Kubernetes CRDs (`crds/`).
   - Deploys example AirGap topology (`examples/airgap-system.yaml`).
   - Creates and starts **systemd** services (`airgap-controller.service` and `lucid-backend.service`).

---

## ??? Quick Start: Deploy to a New Server in 1 Command

### Step 1: Install Ansible on your control machine
```bash
sudo apt update && sudo apt install -y ansible
```

### Step 2: Configure your target server in `inventory.ini`
Open `ansible/inventory.ini` and set your new server's IP address and SSH user:
```ini
[all]
target-server ansible_host=YOUR_NEW_SERVER_IP ansible_user=ubuntu ansible_ssh_private_key_file=~/.ssh/id_rsa
```

### Step 3: Run the Playbook
```bash
ansible-playbook -i inventory.ini playbook.yml
```

---

## ???? Default Ports & Access Points After Deployment

| Service | Port | Default URL | Description |
| :--- | :--- | :--- | :--- |
| **Lucid NOC UI** | `8085` | `http://<SERVER_IP>:8085` | Interactive AI NOC topology & Copilot UI |
| **Grafana Dashboard** | `3000` | `http://<SERVER_IP>:3000` | Full 71-panel infrastructure & network dashboard (`admin`/`admin`) |
| **Prometheus TSDB** | `9090` | `http://<SERVER_IP>:9090` | PromQL query console & targets status |
| **Node Exporter** | `9100` | `http://<SERVER_IP>:9100/metrics` | Host Linux kernel & hardware telemetry |
| **K3s Kubernetes** | `6443` | `kubectl get nodes` | Kubernetes cluster control plane |
