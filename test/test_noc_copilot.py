#!/usr/bin/env python3
"""
Unit and Integration Test Suite for AirGap Autonomous AI NOC Copilot System
"""

import unittest
import sys
import os
import time

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from pkg.faults.fault_injector import FaultInjector
from pkg.telemetry.pipeline import TelemetryPipeline
from pkg.analytics.predictive_engine import PredictiveAnalyticsEngine
from pkg.copilot.offline_llm import AirGapNOCCopilot


class TestAirGapNOCCopilotSystem(unittest.TestCase):
    def setUp(self):
        self.injector = FaultInjector(state_file="/tmp/test_fault_state.json")
        self.injector.clear_faults()
        self.pipeline = TelemetryPipeline(history_capacity=50)
        self.engine = PredictiveAnalyticsEngine(self.pipeline)
        self.copilot = AirGapNOCCopilot(self.engine)

    def tearDown(self):
        self.injector.clear_faults()

    def test_progressive_congestion_scenario(self):
        # Inject Progressive Congestion
        self.injector.inject_progressive_congestion("org-a-edge-router", duration_seconds=2)
        active = self.injector.get_active_faults()
        self.assertIn("org-a-edge-router", active)

        # Collect 5 telemetry snapshots
        for _ in range(5):
            time.sleep(0.05)
            self.pipeline.collect_snapshot("org-a-edge-router", "router")

        # Run predictive analytics
        analysis = self.engine.analyze_component("org-a-edge-router", "router")
        self.assertIn(analysis["status"], ["DEGRADED", "CRITICAL"])
        self.assertGreater(analysis["risk_score"], 40.0)

        # Generate Copilot Response
        copilot_resp = self.copilot.generate_incident_copilot_response("org-a-edge-router", "router")
        self.assertIn("org-a-edge-router", copilot_resp["q1_forecast"])
        self.assertIn("Risk score assessed", copilot_resp["q2_reasoning"])
        self.assertTrue(len(copilot_resp["q3_corrective_action"]) > 0)

    def test_route_flap_cascade_scenario(self):
        self.injector.inject_route_flap_cascade("tor-a1", duration_seconds=300)
        for _ in range(5):
            time.sleep(0.05)
            self.pipeline.collect_snapshot("tor-a1", "tor")

        analysis = self.engine.analyze_component("tor-a1", "tor")
        self.assertGreater(analysis["risk_score"], 35.0)

        copilot_resp = self.copilot.generate_incident_copilot_response("tor-a1", "tor")
        self.assertIn("tor-a1", copilot_resp["component"])

    def test_natural_language_queries(self):
        q1_ans = self.copilot.process_natural_language_query("What is likely to fail next?")
        self.assertIn("Q1 Forecast", q1_ans)

        q2_ans = self.copilot.process_natural_language_query("Why is risk evaluated as elevated?")
        self.assertIn("Q2 Signal Reasoning", q2_ans)

        q3_ans = self.copilot.process_natural_language_query("What action should be taken?")
        self.assertIn("Q3 Recommended Action", q3_ans)


if __name__ == "__main__":
    unittest.main()
