#!/usr/bin/env python3
"""
Unit Test Suite (ut) - Telemetry Pipeline & Fault Injection
Tests TelemetryPipeline snapshot normalization, baseline metrics, and fault calculations.
"""

import unittest
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.telemetry.pipeline import TelemetryPipeline
from pkg.faults.fault_injector import FaultInjector


class TestTelemetryPipelineUT(unittest.TestCase):

    def setUp(self):
        self.pipeline = TelemetryPipeline(history_capacity=10)
        self.injector = FaultInjector()

    def test_collect_snapshot_ut(self):
        """(ut) Test normal telemetry snapshot collection."""
        snapshot = self.pipeline.collect_snapshot("org-a-edge-router", "router")
        self.assertEqual(snapshot["component"], "org-a-edge-router")
        self.assertEqual(snapshot["kind"], "router")
        
        m = snapshot["metrics"]
        self.assertIn("health_score", m)
        self.assertIn("interface_utilization_pct", m)
        self.assertIn("latency_ms", m)
        self.assertIn("packet_loss_pct", m)
        self.assertIn("jitter_ms", m)
        self.assertIn("http_requests_per_sec", m)
        self.assertIn("network_traffic_bytes_sec", m)
        self.assertIn("active_tcp_connections", m)
        self.assertIn("error_rate_pct", m)

    def test_history_capacity_ut(self):
        """(ut) Test sliding window history capacity limit."""
        for _ in range(15):
            self.pipeline.collect_snapshot("tor-a1", "tor")
        
        history = self.pipeline.get_history("tor-a1")
        self.assertEqual(len(history), 10)

    def test_fault_injection_metrics_ut(self):
        """(ut) Test fault injector state tracking."""
        self.injector.clear_faults()
        self.injector.inject_progressive_congestion("org-a-edge-router")
        active = self.injector.get_active_faults()
        self.assertIn("org-a-edge-router", active)
        self.injector.clear_faults()


if __name__ == "__main__":
    unittest.main()
