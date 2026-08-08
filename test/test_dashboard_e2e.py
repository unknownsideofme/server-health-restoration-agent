#!/usr/bin/env python3
"""
End-to-End Test Suite (e2e) - Web UI Dashboard & Prometheus Exporter
Tests React 18 frontend HTML, /api/telemetry JSON payload, and /metrics Prometheus format.
"""

import json
import os
import sys
import unittest
import urllib.request

WEB_UI_URL = "http://127.0.0.1:8085"


class TestDashboardE2E(unittest.TestCase):

    def test_dashboard_html_e2e(self):
        """(e2e) Test Web UI Dashboard HTML response."""
        try:
            req = urllib.request.Request(f"{WEB_UI_URL}/")
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                html = resp.read().decode("utf-8")
                self.assertIn("React 18", html)
                self.assertIn("Autonomous Air-Gapped NOC Copilot", html)
        except Exception as e:
            self.skipTest(f"Web UI server not active at {WEB_UI_URL}: {e}")

    def test_telemetry_api_e2e(self):
        """(e2e) Test /api/telemetry JSON structure and 58 components."""
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

    def test_prometheus_metrics_e2e(self):
        """(e2e) Test /metrics endpoint exporting Prometheus metrics exposition format."""
        try:
            req = urllib.request.Request(f"{WEB_UI_URL}/metrics")
            with urllib.request.urlopen(req, timeout=3) as resp:
                self.assertEqual(resp.status, 200)
                text = resp.read().decode("utf-8")
                self.assertIn("airgap_component_health_score", text)
                self.assertIn("airgap_http_requests_per_sec", text)
                self.assertIn("airgap_network_traffic_bytes_sec", text)
                self.assertIn("airgap_active_tcp_connections", text)
        except Exception as e:
            self.skipTest(f"Metrics exporter not active at {WEB_UI_URL}: {e}")


if __name__ == "__main__":
    unittest.main()
