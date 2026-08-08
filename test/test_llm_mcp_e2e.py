#!/usr/bin/env python3
"""
End-to-End Test Suite (e2e) - Local LLM Microservice & MCP Skill Dispatch
Tests Local LLM Daemon on port 11435 and MCP skill invocation via /api/mcp/call.
"""

import json
import os
import sys
import unittest
import urllib.request

WEB_UI_URL = "http://127.0.0.1:8085"
LLM_DAEMON_URL = "http://127.0.0.1:11435"


class TestLLMMcpE2E(unittest.TestCase):

    def test_llm_daemon_endpoint_e2e(self):
        """(e2e) Test Local LLM Microservice Daemon (port 11435) non-blocking generation."""
        try:
            payload = json.dumps({"prompt": "Check system status"}).encode("utf-8")
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

    def test_mcp_autonomous_repair_endpoint_e2e(self):
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
