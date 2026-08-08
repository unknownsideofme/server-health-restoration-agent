#!/usr/bin/env python3
"""
Unit Test Suite (ut) - Agent Skills & MCP Protocol Server
Tests OS repair, router repair, autonomous self-healing skills, and MCP tool registry.
"""

import unittest
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.skills.os_repair_skill import OSRepairSkill
from pkg.skills.router_repair_skill import RouterRepairSkill
from pkg.skills.autofix_skill import AutoFixSkill
from pkg.mcp.mcp_server import MCPServer


class TestAgentSkillsUT(unittest.TestCase):

    def setUp(self):
        self.mcp = MCPServer()

    def test_os_repair_skill_ut(self):
        """(ut) Test OSRepairSkill static methods."""
        res = OSRepairSkill.repair_os_interface("org-a-admin-01")
        self.assertEqual(res["status"], "SUCCESS")
        self.assertIn("actions", res)

    def test_router_repair_skill_ut(self):
        """(ut) Test RouterRepairSkill static methods."""
        res = RouterRepairSkill.remediate_router_congestion("org-a-edge-router")
        self.assertEqual(res["status"], "SUCCESS")
        self.assertIn("actions", res)

        res2 = RouterRepairSkill.remediate_bgp_flap("org-a-edge-router")
        self.assertEqual(res2["status"], "SUCCESS")

        res3 = RouterRepairSkill.resync_acl_policy("org-a-edge-router")
        self.assertEqual(res3["status"], "SUCCESS")

    def test_autofix_skill_ut(self):
        """(ut) Test AutoFixSkill autonomous remediation workflow."""
        res = AutoFixSkill.execute_autonomous_repair("org-a-edge-router")
        self.assertEqual(res["autofix_status"], "COMPLETED")

    def test_mcp_server_tools_ut(self):
        """(ut) Test MCPServer tool listing and JSON-RPC tool invocation."""
        tools_resp = self.mcp.list_tools()
        self.assertIn("tools", tools_resp)
        names = [t["name"] for t in tools_resp["tools"]]
        self.assertIn("repair_os_interface", names)
        self.assertIn("remediate_router_congestion", names)
        self.assertIn("remediate_bgp_flap", names)
        self.assertIn("resync_acl_policy", names)
        self.assertIn("execute_autonomous_repair", names)

        # Execute tool via MCP
        call_res = self.mcp.call_tool("execute_autonomous_repair", {"component_name": "org-a-edge-router"})
        self.assertFalse(call_res.get("isError", False))


if __name__ == "__main__":
    unittest.main()
