#!/usr/bin/env python3
"""
Model Context Protocol (MCP) JSON-RPC 2.0 Server & Tool Registry
Implements standard MCP interfaces (mcp_list_tools, mcp_call_tool) for OS and Router repair capabilities.
"""

import json
import logging
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "../..")))
from pkg.skills.autofix_skill import AutoFixSkill
from pkg.skills.os_repair_skill import OSRepairSkill
from pkg.skills.router_repair_skill import RouterRepairSkill

logger = logging.getLogger("MCPServer")

TOOLS_REGISTRY = [
    {
        "name": "repair_os_interface",
        "description": "Fix operating system IP address, interface status, and network configuration deep inside OS containers.",
        "inputSchema": {
            "type": "object",
            "properties": {
                "component_name": {"type": "string", "description": "Target component container name"},
                "ip_address": {"type": "string", "description": "Topology IP address to configure"},
            },
            "required": ["component_name"],
        },
    },
    {
        "name": "remediate_router_congestion",
        "description": "Apply QoS FQ-CoDel traffic shaping and clear link bandwidth congestion on router interfaces.",
        "inputSchema": {
            "type": "object",
            "properties": {
                "target_component": {"type": "string", "description": "Target router or TOR switch name"},
            },
            "required": ["target_component"],
        },
    },
    {
        "name": "remediate_bgp_flap",
        "description": "Apply BGP route dampening and tune hold-time timers to eliminate routing instability cascades.",
        "inputSchema": {
            "type": "object",
            "properties": {
                "target_component": {"type": "string", "description": "Target router or TOR switch name"},
            },
            "required": ["target_component"],
        },
    },
    {
        "name": "resync_acl_policy",
        "description": "Resync firewall ACL security policies against AirGap Kubernetes CRD desired state.",
        "inputSchema": {
            "type": "object",
            "properties": {
                "target_component": {"type": "string", "description": "Target component name"},
            },
            "required": ["target_component"],
        },
    },
    {
        "name": "execute_autonomous_repair",
        "description": "Execute autonomous self-healing repair skill based on predictive alert failure mode.",
        "inputSchema": {
            "type": "object",
            "properties": {
                "component_name": {"type": "string", "description": "Target component name"},
                "issue_type": {"type": "string", "description": "Predicted failure mode"},
            },
            "required": ["component_name"],
        },
    },
]


class MCPServer:
    def __init__(self):
        self.tools = {t["name"]: t for t in TOOLS_REGISTRY}

    def list_tools(self) -> dict:
        """MCP Protocol: mcp_list_tools response."""
        return {"tools": TOOLS_REGISTRY}

    def call_tool(self, tool_name: str, arguments: dict) -> dict:
        """MCP Protocol: mcp_call_tool execution."""
        logger.info(f"MCP Call Tool: '{tool_name}' with args {arguments}")

        if tool_name not in self.tools:
            return {"isError": True, "content": [{"type": "text", "text": f"Tool '{tool_name}' not found."}]}

        try:
            if tool_name == "repair_os_interface":
                res = OSRepairSkill.repair_os_interface(
                    arguments["component_name"], arguments.get("ip_address", "")
                )
            elif tool_name == "remediate_router_congestion":
                res = RouterRepairSkill.remediate_router_congestion(arguments["target_component"])
            elif tool_name == "remediate_bgp_flap":
                res = RouterRepairSkill.remediate_bgp_flap(arguments["target_component"])
            elif tool_name == "resync_acl_policy":
                res = RouterRepairSkill.resync_acl_policy(arguments["target_component"])
            elif tool_name == "execute_autonomous_repair":
                res = AutoFixSkill.execute_autonomous_repair(
                    arguments["component_name"], arguments.get("issue_type", "AUTO_DETECT")
                )
            else:
                return {"isError": True, "content": [{"type": "text", "text": "Unimplemented MCP tool."}]}

            return {"isError": False, "content": [{"type": "text", "text": json.dumps(res, indent=2)}], "result": res}

        except Exception as e:
            logger.error(f"Error executing MCP tool '{tool_name}': {e}")
            return {"isError": True, "content": [{"type": "text", "text": str(e)}]}
