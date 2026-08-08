#!/usr/bin/env python3
"""
Unit Test Suite for Model Context Protocol (MCP) Server & Deep OS/Router Agent Skills
"""

import unittest
import sys
import os

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

from pkg.mcp.mcp_server import MCPServer
from pkg.skills.os_repair_skill import OSRepairSkill
from pkg.skills.router_repair_skill import RouterRepairSkill
from pkg.skills.autofix_skill import AutoFixSkill


class TestMCPSkills(unittest.TestCase):
    def setUp(self):
        self.mcp = MCPServer()

    def test_mcp_list_tools(self):
        tools_resp = self.mcp.list_tools()
        self.assertIn("tools", tools_resp)
        tool_names = [t["name"] for t in tools_resp["tools"]]
        self.assertIn("repair_os_interface", tool_names)
        self.assertIn("remediate_router_congestion", tool_names)
        self.assertIn("remediate_bgp_flap", tool_names)
        self.assertIn("execute_autonomous_repair", tool_names)

    def test_os_repair_skill(self):
        res = OSRepairSkill.repair_os_interface("org-a-admin-01", "10.10.0.11")
        self.assertEqual(res["status"], "SUCCESS")
        self.assertEqual(res["component"], "org-a-admin-01")

    def test_router_repair_skill(self):
        res = RouterRepairSkill.remediate_router_congestion("org-a-edge-router")
        self.assertEqual(res["status"], "SUCCESS")

        res_bgp = RouterRepairSkill.remediate_bgp_flap("tor-a1")
        self.assertEqual(res_bgp["status"], "SUCCESS")

    def test_mcp_call_tool_autofix(self):
        call_resp = self.mcp.call_tool("execute_autonomous_repair", {"component_name": "org-a-edge-router", "issue_type": "PREDICTED_LINK_SATURATION"})
        self.assertFalse(call_resp["isError"])
        self.assertIn("result", call_resp)


if __name__ == "__main__":
    unittest.main()
