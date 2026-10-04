#!/usr/bin/env python3
"""
AirGap Telemetry Pipeline & Time-Series Normalizer
Ingests interface utilization, latency, jitter, packet loss, BGP/OSPF events, and syslog counters.
Integrates live ground-truth fault states to produce continuous time-series metrics,
including network congestion, traffic throughput, failure tracking, and SRE reliability (MTTR, MTBF).
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
        self.component_stats: Dict[str, dict] = {}
        self._boot_time = time.time() - 86400  # simulated 24h baseline uptime

    def _get_or_init_stats(self, component_name: str) -> dict:
        if component_name not in self.component_stats:
            self.component_stats[component_name] = {
                "failure_count": 0,
                "last_failure_start": None,
                "total_downtime_seconds": 0.0,
                "repaired_count": 0,
                "total_repair_time_seconds": 0.0,
                "was_fault_active": False,
            }
        return self.component_stats[component_name]

    def collect_snapshot(self, component_name: str, kind: str = "server", timestamp: float = None) -> dict:
        """Collect a single normalized telemetry snapshot for a component."""
        now = timestamp if timestamp is not None else time.time()
        active_faults = self.injector.get_active_faults()
        fault = active_faults.get(component_name)

        stats = self._get_or_init_stats(component_name)

        # Track fault state transitions for MTTR & MTBF
        is_fault = fault is not None
        if is_fault and not stats["was_fault_active"]:
            stats["failure_count"] += 1
            stats["last_failure_start"] = now
            stats["was_fault_active"] = True
        elif not is_fault and stats["was_fault_active"]:
            if stats["last_failure_start"]:
                repair_duration = now - stats["last_failure_start"]
                stats["total_repair_time_seconds"] += max(10.0, repair_duration)
                stats["total_downtime_seconds"] += max(10.0, repair_duration)
                stats["repaired_count"] += 1
            stats["last_failure_start"] = None
            stats["was_fault_active"] = False

        # Baseline baseline noise metrics
        utilization = random.uniform(18.0, 35.0)
        latency = random.uniform(1.2, 4.5)
        packet_loss = random.uniform(0.0, 0.05)
        jitter = random.uniform(0.5, 1.8)
        bgp_flaps = random.randint(0, 1)
        acl_drops = random.randint(0, 2)
        health_score = 98.0
        failure_severity = 0  # 0=Healthy

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
                failure_severity = 3 if utilization > 85.0 else 2

            elif f_type == "ROUTE_FLAP_CASCADE":
                bgp_flaps = random.randint(12, 28)
                latency = random.uniform(15.0, 85.0)
                packet_loss = random.uniform(2.5, 8.0)
                health_score = max(25.0, 100.0 - bgp_flaps * 3.0)
                failure_severity = 3

            elif f_type == "TUNNEL_DEGRADATION":
                jitter = fault["metrics"]["jitter_ms"] + random.uniform(-5.0, 5.0)
                packet_loss = fault["metrics"]["packet_loss_pct"] + random.uniform(-0.5, 1.5)
                latency = random.uniform(35.0, 120.0)
                health_score = max(30.0, 100.0 - (jitter * 0.8 + packet_loss * 4.0))
                failure_severity = 2

            elif f_type == "POLICY_DRIFT":
                acl_drops = int(random.randint(45, 120) * progress)
                health_score = max(40.0, 100.0 - acl_drops * 0.5)
                failure_severity = 2

        # Calculate dynamic traffic metrics
        http_requests_per_sec = round(max(5.0, utilization * 14.5 + random.uniform(-10.0, 10.0)), 1)
        network_traffic_bytes_sec = int(utilization * 1024 * 1024 * 0.85 + random.uniform(10000, 50000))
        active_tcp_conns = int(utilization * 38 + random.randint(15, 60))
        error_rate = round(max(0.0, packet_loss * 1.8 + (100.0 - health_score) * 0.12), 2)

        # --- ADVANCED NETWORK CONGESTION METRICS ---
        congestion_raw = max(0.0, (utilization - 40.0) * 1.25) + min(40.0, latency * 0.35) + (packet_loss * 4.5)
        congestion_index_pct = round(min(100.0, max(0.0, congestion_raw)), 2)

        if utilization > 70.0:
            buffer_occupancy_pct = round(min(100.0, 45.0 + (utilization - 70.0) * 1.82 + random.uniform(-1.5, 2.5)), 2)
        else:
            buffer_occupancy_pct = round(max(4.0, utilization * 0.52 + random.uniform(-1.0, 1.5)), 2)

        if buffer_occupancy_pct > 78.0:
            queue_drop_rate_pps = round((buffer_occupancy_pct - 78.0) * 14.2 + random.uniform(0.0, 4.0), 2)
        else:
            queue_drop_rate_pps = 0.0

        traffic_headroom_pct = round(max(0.0, 100.0 - utilization), 2)

        # --- ADVANCED TRAFFIC METRICS ---
        ingress_throughput_mbps = round((network_traffic_bytes_sec * 8.0) / 1_000_000.0, 2)
        egress_throughput_mbps = round(ingress_throughput_mbps * random.uniform(0.88, 0.97), 2)
        traffic_packets_per_sec = int((network_traffic_bytes_sec / 1400.0) + random.uniform(15, 60))
        top_talker_flow_mbps = round(ingress_throughput_mbps * (0.68 if utilization > 75 else 0.32), 2)

        # --- SRE RELIABILITY & FAILURE METRICS (MTTR, MTBF, Availability) ---
        is_failed = 1 if (fault is not None or health_score < 60.0) else 0

        if stats["repaired_count"] > 0:
            avg_repair_sec = stats["total_repair_time_seconds"] / stats["repaired_count"]
            mttr_minutes = round(avg_repair_sec / 60.0, 2)
        else:
            mttr_minutes = 2.15

        if is_failed and stats["last_failure_start"]:
            active_repair_min = (now - stats["last_failure_start"]) / 60.0
            mttr_minutes = round(max(mttr_minutes, active_repair_min), 2)

        uptime_hours = max(1.0, (now - self._boot_time) / 3600.0)
        total_failures = max(1, stats["failure_count"])
        mtbf_hours = round(uptime_hours / total_failures, 2)
        mtbf_hours = min(72.0, max(6.5, mtbf_hours))

        mttr_hours = mttr_minutes / 60.0
        availability_pct = round((mtbf_hours / (mtbf_hours + mttr_hours)) * 100.0, 3)
        if is_failed:
            availability_pct = round(max(92.0, availability_pct - 3.5), 2)

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
                # Congestion
                "congestion_index_pct": congestion_index_pct,
                "buffer_occupancy_pct": buffer_occupancy_pct,
                "queue_drop_rate_pps": queue_drop_rate_pps,
                "traffic_headroom_pct": traffic_headroom_pct,
                # Traffic
                "ingress_throughput_mbps": ingress_throughput_mbps,
                "egress_throughput_mbps": egress_throughput_mbps,
                "traffic_packets_per_sec": traffic_packets_per_sec,
                "top_talker_flow_mbps": top_talker_flow_mbps,
                # Failure & SRE
                "failure_active": is_failed,
                "failure_severity": failure_severity,
                "failure_events_total": stats["failure_count"],
                "mttr_minutes": mttr_minutes,
                "mtbf_hours": mtbf_hours,
                "availability_pct": availability_pct,
            },
            "fault_active": fault["type"] if fault else None,
        }

        if component_name not in self.history:
            self.history[component_name] = deque(maxlen=self.history_capacity)
        self.history[component_name].append(snapshot)

        return snapshot

    def get_history(self, component_name: str) -> List[dict]:
        return list(self.history.get(component_name, []))
