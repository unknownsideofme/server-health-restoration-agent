
#!/usr/bin/env python3
"""
AirGap Fault Injection Engine & Scenario Generator
Simulates realistic network fault scenarios across the AirGap infrastructure:
  1. Progressive Congestion Buildup (bandwidth saturation, latency drift, packet loss)
  2. BGP / OSPF Route Flap Cascade (route flaps, convergence stress, path asymmetry)
  3. Tunnel & Underlay Health Degradation (MTU mismatch, rekey drops, packet loss)
  4. Policy Drift & Controller Misconfiguration (ACL drift, rule blockages)
"""

import json
import logging
import os
import random
import time
from datetime import datetime, timezone

logger = logging.getLogger("FaultInjector")


class FaultInjector:
    def __init__(self, state_file="/tmp/airgap_fault_state.json"):
        self.state_file = state_file
        self.active_faults = self._load_state()

    def _load_state(self) -> dict:
        if os.path.exists(self.state_file):
            try:
                with open(self.state_file, "r") as f:
                    return json.load(f)
            except Exception:
                pass
        return {}

    def _save_state(self):
        try:
            with open(self.state_file, "w") as f:
                json.dump(self.active_faults, f, indent=2)
        except Exception as e:
            logger.error(f"Failed to save fault state: {e}")

    def inject_progressive_congestion(self, target_component: str = "org-a-edge-router", duration_seconds: int = 300) -> dict:
        """Scenario 1: Progressive Congestion Buildup (bandwidth saturation 45%->99%, latency 2ms->240ms)."""
        fault_id = f"fault-congestion-{int(time.time())}"
        fault_data = {
            "id": fault_id,
            "type": "PROGRESSIVE_CONGESTION",
            "target": target_component,
            "start_time": time.time(),
            "duration": duration_seconds,
            "severity": "CRITICAL",
            "metrics": {
                "initial_utilization_pct": 42.5,
                "target_utilization_pct": 98.6,
                "initial_latency_ms": 2.1,
                "target_latency_ms": 245.0,
                "initial_packet_loss_pct": 0.0,
                "target_packet_loss_pct": 14.8,
            },
            "description": f"Progressive link bandwidth saturation and queue buildup on {target_component}.",
        }
        self.active_faults[target_component] = fault_data
        self._save_state()
        logger.info(f"Injected Progressive Congestion fault on {target_component} (ID: {fault_id})")
        return fault_data

    def inject_route_flap_cascade(self, target_component: str = "tor-a1", duration_seconds: int = 300) -> dict:
        """Scenario 2: BGP/OSPF Route Flap Cascade."""
        fault_id = f"fault-route-flap-{int(time.time())}"
        fault_data = {
            "id": fault_id,
            "type": "ROUTE_FLAP_CASCADE",
            "target": target_component,
            "start_time": time.time(),
            "duration": duration_seconds,
            "severity": "HIGH",
            "metrics": {
                "flap_frequency_per_min": 18,
                "route_table_entropy": 8.7,
                "convergence_delay_sec": 12.4,
                "affected_prefixes": ["10.10.0.0/24", "10.11.0.0/24"],
            },
            "description": f"BGP/OSPF route advertisement flapping on {target_component} causing convergence stress.",
        }
        self.active_faults[target_component] = fault_data
        self._save_state()
        logger.info(f"Injected Route Flap Cascade fault on {target_component} (ID: {fault_id})")
        return fault_data

    def inject_tunnel_degradation(self, target_component: str = "org-a-edge-router", duration_seconds: int = 300) -> dict:
        """Scenario 3: Tunnel & Underlay Health Degradation."""
        fault_id = f"fault-tunnel-deg-{int(time.time())}"
        fault_data = {
            "id": fault_id,
            "type": "TUNNEL_DEGRADATION",
            "target": target_component,
            "start_time": time.time(),
            "duration": duration_seconds,
            "severity": "HIGH",
            "metrics": {
                "jitter_ms": 48.2,
                "rekey_failures": 7,
                "packet_loss_pct": 8.4,
                "mtu_blackhole": True,
            },
            "description": f"Overlay IPSec tunnel jitter and rekey failure escalation on {target_component}.",
        }
        self.active_faults[target_component] = fault_data
        self._save_state()
        logger.info(f"Injected Tunnel Degradation fault on {target_component} (ID: {fault_id})")
        return fault_data

    def inject_policy_drift(self, target_component: str = "org-a-edge-router", duration_seconds: int = 300) -> dict:
        """Scenario 4: Policy Drift & Controller Misconfiguration."""
        fault_id = f"fault-policy-drift-{int(time.time())}"
        fault_data = {
            "id": fault_id,
            "type": "POLICY_DRIFT",
            "target": target_component,
            "start_time": time.time(),
            "duration": duration_seconds,
            "severity": "MEDIUM",
            "metrics": {
                "blocked_flows_pct": 34.0,
                "acl_mismatch_rules": ["DENY_APP_TRAFFIC_10_10"],
                "controller_sync_status": "OUT_OF_SYNC",
            },
            "description": f"ACL policy drift and controller configuration mismatch on {target_component}.",
        }
        self.active_faults[target_component] = fault_data
        self._save_state()
        logger.info(f"Injected Policy Drift fault on {target_component} (ID: {fault_id})")
        return fault_data

    def clear_faults(self, target_component: str = None) -> list:
        cleared = []
        if target_component:
            if target_component in self.active_faults:
                cleared.append(self.active_faults.pop(target_component))
        else:
            cleared = list(self.active_faults.values())
            self.active_faults.clear()
        self._save_state()
        logger.info(f"Cleared {len(cleared)} faults.")
        return cleared

    def get_active_faults(self) -> dict:
        # Filter out expired faults
        now = time.time()
        expired = []
        for target, data in list(self.active_faults.items()):
            if now - data["start_time"] > data["duration"]:
                expired.append(target)
        for target in expired:
            self.active_faults.pop(target, None)
        if expired:
            self._save_state()
        return self.active_faults
