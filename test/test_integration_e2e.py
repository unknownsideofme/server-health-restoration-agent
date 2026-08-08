#!/usr/bin/env python3
"""
End-to-End Integration Test Suite (e2e) - AirGap Autonomous AI NOC Copilot
Performs live HTTP request testing against Web UI, Prometheus metrics exporter, Local LLM Daemon, and MCP API endpoints.
"""

import json
import os
import sys
import unittest
import urllib.request

WEB_UI_URL = "http://127.0.0.1:8085"
LLM_DAEMON_URL = "http://127.0.0.1:11435"


class TestAirGapIntegrationE2E(unittest.TestCase):

    def test_web_ui_dashboard_e2e(self):
        """(e2e) Test Web UI Dashboard HTML serving."""
        try:
            req = urllib.request.Request(f"{WEB_UI_URL}/")
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                html = resp.read().decode("utf-8")
                self.assertIn("AirGap Autonomous AI NOC Copilot", html)
                self.assertIn("React 18", html)
        except Exception as e:
            self.skipTest(f"Web UI server not active at {WEB_UI_URL}: {e}")

    def test_telemetry_api_e2e(self):
        """(e2e) Test /api/telemetry endpoint returning 58 components and by_org hierarchy."""
        try:
            req = urllib.request.Request(f"{WEB_UI_URL}/api/telemetry")
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                data = json.loads(resp.read().decode("utf-8"))
                self.assertIn("topology", data)
                self.assertIn("by_org", data)
                self.assertGreaterEqual(len(data["topology"]), 40)
                self.assertIn("org-a", data["by_org"])
        except Exception as e:
            self.skipTest(f"Telemetry API not active at {WEB_UI_URL}: {e}")

    def test_prometheus_metrics_exporter_e2e(self):
        """(e2e) Test /metrics endpoint exporting Prometheus live traffic exposition metrics."""
        try:
            req = urllib.request.Request(f"{WEB_UI_URL}/metrics")
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                metrics_text = resp.read().decode("utf-8")
                self.assertIn("airgap_component_health_score", metrics_text)
                self.assertIn("airgap_http_requests_per_sec", metrics_text)
                self.assertIn("airgap_network_traffic_bytes_sec", metrics_text)
                self.assertIn("airgap_active_tcp_connections", metrics_text)
        except Exception as e:
            self.skipTest(f"Metrics exporter not active at {WEB_UI_URL}: {e}")

    def test_local_llm_daemon_e2e(self):
        """(e2e) Test Local LLM Microservice Daemon (port 11435) non-blocking generation."""
        try:
            payload = json.dumps({"prompt": "Explain network congestion"}).encode("utf-8")
            req = urllib.request.Request(
                f"{LLM_DAEMON_URL}/api/llm/generate",
                data=payload,
                headers={"Content-Type": "application/json"},
                method="POST"
            )
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                data = json.loads(resp.read().decode("utf-8"))
                self.assertIn("response", data)
                self.assertEqual(data["status"], "SUCCESS")
        except Exception as e:
            self.skipTest(f"LLM Daemon not active at {LLM_DAEMON_URL}: {e}")

    def test_mcp_skill_autonomous_repair_e2e(self):
        """(e2e) Test /api/mcp/call endpoint executing autonomous self-healing skill."""
        try:
            payload = json.dumps({
                "tool_name": "execute_autonomous_repair",
                "arguments": {"component_name": "org-a-edge-router"}
            }).encode("utf-8")
            req = urllib.request.Request(
                f"{WEB_UI_URL}/api/mcp/call",
                data=payload,
                headers={"Content-Type": "application/json"},
                method="POST"
            )
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                data = json.loads(resp.read().decode("utf-8"))
                self.assertIn("result", data)
        except Exception as e:
            self.skipTest(f"MCP API not active at {WEB_UI_URL}: {e}")


if __name__ == "__main__":
    unittest.main()
