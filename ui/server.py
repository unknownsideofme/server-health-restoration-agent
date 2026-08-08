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
from urllib.parse import parse_qs, urlparse

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

COMPONENTS = [
    # Org A (11 components)
    {"name": "org-a-edge-router", "kind": "router", "org": "org-a", "role": "router"},
    {"name": "tor-a1", "kind": "tor", "org": "org-a", "role": "tor"},
    {"name": "tor-a2", "kind": "tor", "org": "org-a", "role": "tor"},
    {"name": "rack-a1", "kind": "rack", "org": "org-a", "role": "rack"},
    {"name": "rack-a2", "kind": "rack", "org": "org-a", "role": "rack"},
    {"name": "subnet-org-a-prod", "kind": "subnet", "org": "org-a", "role": "subnet"},
    {"name": "subnet-org-a-dev", "kind": "subnet", "org": "org-a", "role": "subnet"},
    {"name": "org-a-admin-01", "kind": "server", "org": "org-a", "role": "admin"},
    {"name": "org-a-admin-02", "kind": "server", "org": "org-a", "role": "admin"},
    {"name": "org-a-admin-03", "kind": "server", "org": "org-a", "role": "admin"},
    {"name": "org-a-worker-01", "kind": "server", "org": "org-a", "role": "worker"},
    {"name": "org-a-worker-02", "kind": "server", "org": "org-a", "role": "worker"},
    {"name": "org-a-worker-03", "kind": "server", "org": "org-a", "role": "worker"},
    {"name": "org-a-api-gateway", "kind": "service", "org": "org-a", "role": "service"},
    {"name": "org-a-auth-service", "kind": "service", "org": "org-a", "role": "service"},

    # Org B (11 components)
    {"name": "org-b-edge-router", "kind": "router", "org": "org-b", "role": "router"},
    {"name": "tor-b1", "kind": "tor", "org": "org-b", "role": "tor"},
    {"name": "tor-b2", "kind": "tor", "org": "org-b", "role": "tor"},
    {"name": "rack-b1", "kind": "rack", "org": "org-b", "role": "rack"},
    {"name": "rack-b2", "kind": "rack", "org": "org-b", "role": "rack"},
    {"name": "subnet-org-b-prod", "kind": "subnet", "org": "org-b", "role": "subnet"},
    {"name": "subnet-org-b-dev", "kind": "subnet", "org": "org-b", "role": "subnet"},
    {"name": "org-b-admin-01", "kind": "server", "org": "org-b", "role": "admin"},
    {"name": "org-b-admin-02", "kind": "server", "org": "org-b", "role": "admin"},
    {"name": "org-b-admin-03", "kind": "server", "org": "org-b", "role": "admin"},
    {"name": "org-b-worker-01", "kind": "server", "org": "org-b", "role": "worker"},
    {"name": "org-b-worker-02", "kind": "server", "org": "org-b", "role": "worker"},
    {"name": "org-b-worker-03", "kind": "server", "org": "org-b", "role": "worker"},
    {"name": "org-b-database-cluster", "kind": "service", "org": "org-b", "role": "service"},
    {"name": "org-b-cache-redis", "kind": "service", "org": "org-b", "role": "service"},

    # Org C (11 components)
    {"name": "org-c-edge-router", "kind": "router", "org": "org-c", "role": "router"},
    {"name": "tor-c1", "kind": "tor", "org": "org-c", "role": "tor"},
    {"name": "tor-c2", "kind": "tor", "org": "org-c", "role": "tor"},
    {"name": "rack-c1", "kind": "rack", "org": "org-c", "role": "rack"},
    {"name": "rack-c2", "kind": "rack", "org": "org-c", "role": "rack"},
    {"name": "subnet-org-c-prod", "kind": "subnet", "org": "org-c", "role": "subnet"},
    {"name": "subnet-org-c-dev", "kind": "subnet", "org": "org-c", "role": "subnet"},
    {"name": "org-c-admin-01", "kind": "server", "org": "org-c", "role": "admin"},
    {"name": "org-c-admin-02", "kind": "server", "org": "org-c", "role": "admin"},
    {"name": "org-c-admin-03", "kind": "server", "org": "org-c", "role": "admin"},
    {"name": "org-c-worker-01", "kind": "server", "org": "org-c", "role": "worker"},
    {"name": "org-c-worker-02", "kind": "server", "org": "org-c", "role": "worker"},
    {"name": "org-c-worker-03", "kind": "server", "org": "org-c", "role": "worker"},
    {"name": "org-c-kafka-cluster", "kind": "service", "org": "org-c", "role": "service"},

    # Org D (11 components)
    {"name": "org-d-edge-router", "kind": "router", "org": "org-d", "role": "router"},
    {"name": "tor-d1", "kind": "tor", "org": "org-d", "role": "tor"},
    {"name": "tor-d2", "kind": "tor", "org": "org-d", "role": "tor"},
    {"name": "rack-d1", "kind": "rack", "org": "org-d", "role": "rack"},
    {"name": "rack-d2", "kind": "rack", "org": "org-d", "role": "rack"},
    {"name": "subnet-org-d-prod", "kind": "subnet", "org": "org-d", "role": "subnet"},
    {"name": "subnet-org-d-dev", "kind": "subnet", "org": "org-d", "role": "subnet"},
    {"name": "org-d-admin-01", "kind": "server", "org": "org-d", "role": "admin"},
    {"name": "org-d-admin-02", "kind": "server", "org": "org-d", "role": "admin"},
    {"name": "org-d-admin-03", "kind": "server", "org": "org-d", "role": "admin"},
    {"name": "org-d-worker-01", "kind": "server", "org": "org-d", "role": "worker"},
    {"name": "org-d-worker-02", "kind": "server", "org": "org-d", "role": "worker"},
    {"name": "org-d-worker-03", "kind": "server", "org": "org-d", "role": "worker"},
    {"name": "org-d-object-storage", "kind": "service", "org": "org-d", "role": "service"},
]


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

            for comp in COMPONENTS:
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
            for comp in COMPONENTS:
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
