#!/usr/bin/env python3
"""
Autonomous Self-Healing Skill
Orchestrates automated incident remediation by mapping predictive alerts to MCP repair skills.
"""

import logging
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "../..")))
from pkg.skills.os_repair_skill import OSRepairSkill
from pkg.skills.router_repair_skill import RouterRepairSkill

logger = logging.getLogger("AutoFixSkill")


class AutoFixSkill:
    @staticmethod
    def execute_autonomous_repair(component_name: str, issue_type: str = "AUTO_DETECT") -> dict:
        """Automatically execute the appropriate MCP repair skill for a component."""
        logger.info(f"Executing AutoFixSkill for component '{component_name}' (Issue: {issue_type})...")

        if issue_type == "AUTO_DETECT":
            from pkg.faults.fault_injector import FaultInjector
            injector = FaultInjector()
            active = injector.get_active_faults()
            if component_name in active:
                issue_type = active[component_name]["type"]
                logger.info(f"Auto-detected active fault type '{issue_type}' for component '{component_name}'")

        if "CONGESTION" in issue_type or "SATURATION" in issue_type:
            res = RouterRepairSkill.remediate_router_congestion(component_name)
        elif "FLAP" in issue_type or "ROUTING" in issue_type:
            res = RouterRepairSkill.remediate_bgp_flap(component_name)
        elif "POLICY" in issue_type or "ACL" in issue_type:
            res = RouterRepairSkill.resync_acl_policy(component_name)
        else:
            res = OSRepairSkill.repair_os_interface(component_name)

        return {
            "autofix_status": "COMPLETED",
            "target": component_name,
            "issue_type": issue_type,
            "result": res,
        }
