#!/usr/bin/env python3
"""
AirGap Autonomous AI NOC Copilot Web Server & REST API
Provides a real-time HTTP server for the NOC Dashboard and Copilot REST API.
Runs 100% offline on http://0.0.0.0:8080.
"""

import http.server
import json
import os
import socketserver
import sys
import time
from urllib.parse import parse_qs, urlparse

from kubernetes import client, config

try:
    config.load_kube_config()
except Exception:
    try:
        config.load_incluster_config()
    except Exception:
        pass

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.analytics.predictive_engine import PredictiveAnalyticsEngine
from pkg.copilot.offline_llm import AirGapNOCCopilot
from pkg.faults.fault_injector import FaultInjector
from pkg.telemetry.pipeline import TelemetryPipeline

PORT = 8085
pipeline = TelemetryPipeline()
engine = PredictiveAnalyticsEngine(pipeline)
copilot = AirGapNOCCopilot(engine)
injector = FaultInjector()



def determine_org(name: str, spec: dict) -> str:
    if "orgRef" in spec and spec["orgRef"]:
        return spec["orgRef"]
    if "networkRef" in spec and spec["networkRef"]:
        net_ref = spec["networkRef"]
        if "org-" in net_ref:
            parts = net_ref.split("-")
            if len(parts) >= 2:
                return f"{parts[0]}-{parts[1]}"
        return net_ref
    name_lower = name.lower()
    for o in ["org-a", "org-b", "org-c", "org-d"]:
        if o in name_lower:
            return o
    for letter, o in [("-a", "org-a"), ("-b", "org-b"), ("-c", "org-c"), ("-d", "org-d")]:
        if letter in name_lower:
            return o
    for o in ["a", "b", "c", "d"]:
        if f"-{o}" in name_lower or f"{o}1" in name_lower or f"{o}2" in name_lower or f"{o}3" in name_lower:
            return f"org-{o}"
    return "org-a"


def get_dynamic_components():
    api = client.CustomObjectsApi()
    GROUP = "airgap.example.com"
    VERSION = "v1alpha1"
    
    resource_kinds = {
        "routers": "router",
        "tors": "tor",
        "racks": "rack",
        "subnets": "subnet",
        "servers": "server",
        "services": "service"
    }
    
    components = []
    
    for plural, kind in resource_kinds.items():
        try:
            objs = api.list_cluster_custom_object(GROUP, VERSION, plural)
            for item in objs.get("items", []):
                name = item["metadata"]["name"]
                spec = item.get("spec", {})
                role = spec.get("role", kind)
                org = determine_org(name, spec)
                
                components.append({
                    "name": name,
                    "kind": kind,
                    "org": org,
                    "role": role
                })
        except Exception:
            pass
            
    return components


class NOCHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        parsed = urlparse(self.path)
        path = parsed.path

        if path == "/":
            self.path = "/ui/index.html"
            return super().do_GET()

        if path.startswith("/ui/"):
            return super().do_GET()

        if path == "/metrics":
            self.send_response(200)
            self.send_header("Content-Type", "text/plain; version=0.0.4")
            self.end_headers()

            lines = []
            lines.append("# HELP airgap_component_health_score Component health score (0-100)")
            lines.append("# TYPE airgap_component_health_score gauge")
            lines.append("# HELP airgap_component_risk_score Component risk score percentage (0-100)")
            lines.append("# TYPE airgap_component_risk_score gauge")
            lines.append("# HELP airgap_interface_utilization_pct Bandwidth utilization percentage (0-100)")
            lines.append("# TYPE airgap_interface_utilization_pct gauge")
            lines.append("# HELP airgap_latency_ms Component latency in milliseconds")
            lines.append("# TYPE airgap_latency_ms gauge")
            lines.append("# HELP airgap_packet_loss_pct Component packet loss percentage")
            lines.append("# TYPE airgap_packet_loss_pct gauge")
            lines.append("# HELP airgap_jitter_ms Component jitter in milliseconds")
            lines.append("# TYPE airgap_jitter_ms gauge")
            lines.append("# HELP airgap_bgp_flap_count BGP / OSPF route flap count")
            lines.append("# TYPE airgap_bgp_flap_count counter")
            lines.append("# HELP airgap_acl_drop_count ACL / Firewall rule drops count")
            lines.append("# TYPE airgap_acl_drop_count counter")
            lines.append("# HELP airgap_time_to_impact_minutes Time-to-Impact lead time in minutes")
            lines.append("# TYPE airgap_time_to_impact_minutes gauge")
            lines.append("# HELP airgap_http_requests_per_sec Live HTTP throughput in requests/sec")
            lines.append("# TYPE airgap_http_requests_per_sec gauge")
            lines.append("# HELP airgap_network_traffic_bytes_sec Live network bandwidth throughput in bytes/sec")
            lines.append("# TYPE airgap_network_traffic_bytes_sec gauge")
            lines.append("# HELP airgap_active_tcp_connections Active TCP connections count")
            lines.append("# TYPE airgap_active_tcp_connections gauge")
            lines.append("# HELP airgap_error_rate_pct HTTP error percentage")
            lines.append("# TYPE airgap_error_rate_pct gauge")

            for comp in get_dynamic_components():
                res = engine.analyze_component(comp["name"], comp["kind"])
                c_name = comp["name"]
                c_kind = comp["kind"]
                c_org = comp["org"]
                m = res["current_metrics"]

                labels = f'component="{c_name}",kind="{c_kind}",org="{c_org}"'
                lines.append(f'airgap_component_health_score{{{labels}}} {m["health_score"]}')
                lines.append(f'airgap_component_risk_score{{{labels}}} {res["risk_score"]}')
                lines.append(f'airgap_interface_utilization_pct{{{labels}}} {m["interface_utilization_pct"]}')
                lines.append(f'airgap_latency_ms{{{labels}}} {m["latency_ms"]}')
                lines.append(f'airgap_packet_loss_pct{{{labels}}} {m["packet_loss_pct"]}')
                lines.append(f'airgap_jitter_ms{{{labels}}} {m["jitter_ms"]}')
                lines.append(f'airgap_bgp_flap_count{{{labels}}} {m["bgp_flap_count"]}')
                lines.append(f'airgap_acl_drop_count{{{labels}}} {m["acl_drop_count"]}')
                lines.append(f'airgap_http_requests_per_sec{{{labels}}} {m.get("http_requests_per_sec", 15.0)}')
                lines.append(f'airgap_network_traffic_bytes_sec{{{labels}}} {m.get("network_traffic_bytes_sec", 500000)}')
                lines.append(f'airgap_active_tcp_connections{{{labels}}} {m.get("active_tcp_connections", 120)}')
                lines.append(f'airgap_error_rate_pct{{{labels}}} {m.get("error_rate_pct", 0.0)}')

                if res.get("tti_minutes") is not None:
                    lines.append(f'airgap_time_to_impact_minutes{{{labels}}} {res["tti_minutes"]}')

            output = "\n".join(lines) + "\n"
            self.wfile.write(output.encode("utf-8"))
            return

        if path == "/api/telemetry":
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()

            topology_data = []
            alerts_data = []
            for comp in get_dynamic_components():
                res = engine.analyze_component(comp["name"], comp["kind"])
                comp_info = dict(comp)
                comp_info.update(res)
                topology_data.append(comp_info)
                if res.get("alerts"):
                    alerts_data.extend(res["alerts"])

            active_faults = injector.get_active_faults()

            by_org = {}
            for org_id in ["org-a", "org-b", "org-c", "org-d"]:
                by_org[org_id] = {
                    "name": org_id.upper().replace("-", " "),
                    "routers": [c for c in topology_data if c["org"] == org_id and c["kind"] == "router"],
                    "tors": [c for c in topology_data if c["org"] == org_id and c["kind"] == "tor"],
                    "racks": [c for c in topology_data if c["org"] == org_id and c["kind"] == "rack"],
                    "subnets": [c for c in topology_data if c["org"] == org_id and c["kind"] == "subnet"],
                    "admin_servers": [c for c in topology_data if c["org"] == org_id and c["kind"] == "server" and c.get("role") == "admin"],
                    "worker_servers": [c for c in topology_data if c["org"] == org_id and c["kind"] == "server" and c.get("role") == "worker"],
                    "services": [c for c in topology_data if c["org"] == org_id and c["kind"] == "service"],
                }

            resp = {
                "timestamp": time.time(),
                "topology": topology_data,
                "by_org": by_org,
                "alerts": alerts_data,
                "active_faults": active_faults,
            }
            self.wfile.write(json.dumps(resp).encode("utf-8"))
            return

        if path == "/api/copilot":
            params = parse_qs(parsed.query)
            target = params.get("target", ["org-a-edge-router"])[0]
            resp = copilot.generate_incident_copilot_response(target)
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps(resp).encode("utf-8"))
            return

        self.send_error(404, "Endpoint not found")

    def do_POST(self):
        parsed = urlparse(self.path)
        content_length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(content_length).decode("utf-8") if content_length > 0 else ""

        try:
            data = json.loads(body) if body else {}
        except Exception:
            data = {}

        if parsed.path == "/api/fault/inject":
            scenario = data.get("scenario", "congestion")
            target = data.get("target", "org-a-edge-router")

            if scenario == "congestion":
                res = injector.inject_progressive_congestion(target)
            elif scenario == "route_flap":
                res = injector.inject_route_flap_cascade(target)
            elif scenario == "tunnel_deg":
                res = injector.inject_tunnel_degradation(target)
            elif scenario == "policy_drift":
                res = injector.inject_policy_drift(target)
            else:
                res = {"error": "Invalid scenario"}

            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps(res).encode("utf-8"))
            return

        if parsed.path == "/api/fault/clear":
            target = data.get("target")
            cleared = injector.clear_faults(target)
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({"cleared_count": len(cleared)}).encode("utf-8"))
            return

        if parsed.path == "/api/mcp/call":
            tool_name = data.get("tool_name", "execute_autonomous_repair")
            arguments = data.get("arguments", {"component_name": "org-a-edge-router"})
            res = copilot.mcp.call_tool(tool_name, arguments)
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps(res).encode("utf-8"))
            return

        if parsed.path == "/api/copilot/chat":
            prompt = data.get("prompt", "What is likely to fail next?")
            answer = copilot.process_natural_language_query(prompt)
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({"answer": answer}).encode("utf-8"))
            return

        self.send_error(404, "Endpoint not found")


class ReusableTCPServer(socketserver.TCPServer):
    allow_reuse_address = True


def run_server():
    os.chdir(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
    with ReusableTCPServer(("0.0.0.0", PORT), NOCHandler) as httpd:
        print(f"AirGap AI NOC Copilot Dashboard running at http://0.0.0.0:{PORT}")
        httpd.serve_forever()


if __name__ == "__main__":
    import time
    run_server()
