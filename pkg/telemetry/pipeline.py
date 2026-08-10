#!/usr/bin/env python3
"""
AirGap Telemetry Pipeline & Time-Series Normalizer
Ingests interface utilization, latency, jitter, packet loss, BGP/OSPF events, and syslog counters.
Integrates live ground-truth fault states to produce continuous time-series metrics.
"""

import math
import random
import time
from collections import deque
from typing import Dict, List
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "../..")))
from pkg.faults.fault_injector import FaultInjector


class TelemetryPipeline:
    def __init__(self, history_capacity: int = 120):
        self.injector = FaultInjector()
        self.history_capacity = history_capacity
        self.history: Dict[str, deque] = {}

    def collect_snapshot(self, component_name: str, kind: str = "server", timestamp: float = None) -> dict:
        """Collect a single normalized telemetry snapshot for a component."""
        now = timestamp if timestamp is not None else time.time()
        active_faults = self.injector.get_active_faults()
        fault = active_faults.get(component_name)

        # Baseline baseline noise metrics
        utilization = random.uniform(18.0, 35.0)
        latency = random.uniform(1.2, 4.5)
        packet_loss = random.uniform(0.0, 0.05)
        jitter = random.uniform(0.5, 1.8)
        bgp_flaps = random.randint(0, 1)
        acl_drops = random.randint(0, 2)
        health_score = 98.0

        if fault:
            f_type = fault["type"]
            elapsed = now - fault["start_time"]
            dur = max(fault["duration"], 1)
            progress = min(1.0, elapsed / dur)

            if f_type == "PROGRESSIVE_CONGESTION":
                scaled_progress = min(1.0, (elapsed * 15) / dur + 0.45)
                init_u = fault["metrics"]["initial_utilization_pct"]
                target_u = fault["metrics"]["target_utilization_pct"]
                utilization = init_u + (target_u - init_u) * math.pow(scaled_progress, 1.2) + random.uniform(-1.0, 1.0)
                utilization = min(99.9, max(0.0, utilization))

                init_l = fault["metrics"]["initial_latency_ms"]
                target_l = fault["metrics"]["target_latency_ms"]
                latency = init_l + (target_l - init_l) * math.pow(scaled_progress, 1.5) + random.uniform(-2.0, 2.0)

                init_pl = fault["metrics"]["initial_packet_loss_pct"]
                target_pl = fault["metrics"]["target_packet_loss_pct"]
                packet_loss = init_pl + (target_pl - init_pl) * math.pow(scaled_progress, 1.8)

                health_score = max(10.0, 100.0 - (utilization * 0.5 + latency * 0.2 + packet_loss * 2.0))

            elif f_type == "ROUTE_FLAP_CASCADE":
                bgp_flaps = random.randint(12, 28)
                latency = random.uniform(15.0, 85.0)
                packet_loss = random.uniform(2.5, 8.0)
                health_score = max(25.0, 100.0 - bgp_flaps * 3.0)

            elif f_type == "TUNNEL_DEGRADATION":
                jitter = fault["metrics"]["jitter_ms"] + random.uniform(-5.0, 5.0)
                packet_loss = fault["metrics"]["packet_loss_pct"] + random.uniform(-0.5, 1.5)
                latency = random.uniform(35.0, 120.0)
                health_score = max(30.0, 100.0 - (jitter * 0.8 + packet_loss * 4.0))

            elif f_type == "POLICY_DRIFT":
                acl_drops = int(random.randint(45, 120) * progress)
                health_score = max(40.0, 100.0 - acl_drops * 0.5)

        # Calculate dynamic traffic metrics
        http_requests_per_sec = round(max(5.0, utilization * 14.5 + random.uniform(-10.0, 10.0)), 1)
        network_traffic_bytes_sec = int(utilization * 1024 * 1024 * 0.85 + random.uniform(10000, 50000))
        active_tcp_conns = int(utilization * 38 + random.randint(15, 60))
        error_rate = round(max(0.0, packet_loss * 1.8 + (100.0 - health_score) * 0.12), 2)

        snapshot = {
            "timestamp": now,
            "component": component_name,
            "kind": kind,
            "metrics": {
                "interface_utilization_pct": round(utilization, 2),
                "latency_ms": round(latency, 2),
                "packet_loss_pct": round(packet_loss, 2),
                "jitter_ms": round(jitter, 2),
                "bgp_flap_count": bgp_flaps,
                "acl_drop_count": acl_drops,
                "health_score": round(health_score, 1),
                "http_requests_per_sec": http_requests_per_sec,
                "network_traffic_bytes_sec": network_traffic_bytes_sec,
                "active_tcp_connections": active_tcp_conns,
                "error_rate_pct": error_rate,
            },
            "fault_active": fault["type"] if fault else None,
        }

        if component_name not in self.history:
            self.history[component_name] = deque(maxlen=self.history_capacity)
        self.history[component_name].append(snapshot)

        return snapshot

    def get_history(self, component_name: str) -> List[dict]:
        return list(self.history.get(component_name, []))
