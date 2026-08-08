#!/usr/bin/env python3
"""
Unit Test Suite (ut) - AirGap Autonomous AI NOC Copilot
Tests telemetry pipeline, predictive analytics engine, RAG runbooks, MCP skills, and LLM copilot logic.
"""

import unittest
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.telemetry.pipeline import TelemetryPipeline
from pkg.analytics.predictive_engine import PredictiveAnalyticsEngine
from pkg.copilot.rag_engine import LocalRAGEngine
from pkg.mcp.mcp_server import MCPServer
from pkg.copilot.offline_llm import AirGapNOCCopilot
from pkg.faults.fault_injector import FaultInjector


class TestAirGapUnitUT(unittest.TestCase):

    def setUp(self):
        self.pipeline = TelemetryPipeline()
        self.engine = PredictiveAnalyticsEngine(self.pipeline)
        self.rag = LocalRAGEngine()
        self.mcp = MCPServer()
        self.copilot = AirGapNOCCopilot(self.engine)
        self.injector = FaultInjector()

    def test_telemetry_pipeline_ut(self):
        """(ut) Test TelemetryPipeline snapshot collection and dynamic metrics."""
        snap = self.pipeline.collect_snapshot("org-a-edge-router", "router")
        self.assertEqual(snap["component"], "org-a-edge-router")
        self.assertEqual(snap["kind"], "router")
        m = snap["metrics"]
        self.assertIn("health_score", m)
        self.assertIn("http_requests_per_sec", m)
        self.assertIn("network_traffic_bytes_sec", m)
        self.assertIn("active_tcp_connections", m)
        self.assertIn("error_rate_pct", m)

    def test_predictive_engine_ut(self):
        """(ut) Test PredictiveAnalyticsEngine rate-of-change and TTI forecasting."""
        analysis = self.engine.analyze_component("tor-a1", "tor")
        self.assertIn("status", analysis)
        self.assertIn("risk_score", analysis)
        self.assertIn("confidence", analysis)
        self.assertIn("current_metrics", analysis)

    def test_rag_engine_ut(self):
        """(ut) Test LocalRAGEngine vector search and runbook retrieval."""
        results = self.rag.retrieve_context("PROGRESSIVE_CONGESTION org-a-edge-router link saturation")
        self.assertIsInstance(results, list)

    def test_mcp_server_ut(self):
        """(ut) Test MCPServer tool registration and JSON-RPC 2.0 dispatch."""
        tools_resp = self.mcp.list_tools()
        self.assertIn("tools", tools_resp)
        tool_names = [t["name"] for t in tools_resp["tools"]]
        self.assertIn("execute_autonomous_repair", tool_names)

        # Call autonomous repair skill via MCP
        res = self.mcp.call_tool("execute_autonomous_repair", {"component_name": "org-a-edge-router"})
        self.assertFalse(res.get("isError", False))

    def test_offline_llm_copilot_ut(self):
        """(ut) Test AirGapNOCCopilot structured decision support (Q1, Q2, Q3)."""
        res = self.copilot.generate_incident_copilot_response("org-a-edge-router", "router")
        self.assertIn("q1_forecast", res)
        self.assertIn("q2_reasoning", res)
        self.assertIn("q3_corrective_action", res)


if __name__ == "__main__":
    unittest.main()
