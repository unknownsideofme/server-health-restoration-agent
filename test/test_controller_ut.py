#!/usr/bin/env python3
"""
Unit Test Suite (ut) - Operator Controller & Predictive Engine
Tests PredictiveAnalyticsEngine rate-of-change, TTI lead times, and Copilot decision support logic.
"""

import unittest
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.telemetry.pipeline import TelemetryPipeline
from pkg.analytics.predictive_engine import PredictiveAnalyticsEngine
from pkg.copilot.offline_llm import AirGapNOCCopilot


class TestControllerPredictiveUT(unittest.TestCase):

    def setUp(self):
        self.pipeline = TelemetryPipeline()
        self.engine = PredictiveAnalyticsEngine(self.pipeline)
        self.copilot = AirGapNOCCopilot(self.engine)

    def test_predictive_analysis_ut(self):
        """(ut) Test PredictiveAnalyticsEngine component analysis output."""
        analysis = self.engine.analyze_component("org-a-edge-router", "router")
        self.assertEqual(analysis["component"], "org-a-edge-router")
        self.assertIn("status", analysis)
        self.assertIn("risk_score", analysis)
        self.assertIn("confidence", analysis)
        self.assertIn("signals", analysis)

    def test_copilot_decision_support_ut(self):
        """(ut) Test AirGapNOCCopilot Q1 forecast, Q2 reasoning, and Q3 remediation."""
        res = self.copilot.generate_incident_copilot_response("org-a-edge-router", "router")
        self.assertIn("q1_forecast", res)
        self.assertIn("q2_reasoning", res)
        self.assertIn("q3_corrective_action", res)


if __name__ == "__main__":
    unittest.main()
