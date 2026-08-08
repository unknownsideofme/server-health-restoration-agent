#!/usr/bin/env python3
"""
AirGap Predictive Fault Analytics Engine
Detects precursor conditions, time-series saturation trends, routing instability entropy, and calculates Time-To-Impact (TTI).
"""

import math
import time
from typing import Dict, List, Optional
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "../..")))
from pkg.telemetry.pipeline import TelemetryPipeline


class PredictiveAnalyticsEngine:
    def __init__(self, pipeline: TelemetryPipeline):
        self.pipeline = pipeline

    def analyze_component(self, component_name: str, kind: str = "server") -> dict:
        snapshot = self.pipeline.collect_snapshot(component_name, kind)
        history = self.pipeline.get_history(component_name)

        if len(history) < 3:
            for _ in range(4 - len(history)):
                time.sleep(0.05)
                self.pipeline.collect_snapshot(component_name, kind)
            history = self.pipeline.get_history(component_name)

        # Extract metric arrays
        timestamps = [s["timestamp"] for s in history]
        utilizations = [s["metrics"]["interface_utilization_pct"] for s in history]
        latencies = [s["metrics"]["latency_ms"] for s in history]
        losses = [s["metrics"]["packet_loss_pct"] for s in history]
        flaps = [s["metrics"]["bgp_flap_count"] for s in history]
        acls = [s["metrics"]["acl_drop_count"] for s in history]

        # Calculate rate of change dM/dt for utilization (per minute)
        n = len(history)
        dt_minutes = (timestamps[-1] - timestamps[0]) / 60.0
        if dt_minutes <= 0.05:
            dt_minutes = 0.1

        du_dt = (utilizations[-1] - utilizations[0]) / dt_minutes
        dl_dt = (latencies[-1] - latencies[0]) / dt_minutes
        recent_util = utilizations[-1]
        recent_lat = latencies[-1]

        risk_score = 15.0
        tti_minutes = None
        confidence = 0.85
        contributing_signals = []
        alerts = []

        # 1. Congestion Precursor Analysis (Threshold SLA = 90% utilization)
        if du_dt > 1.5 or recent_util > 75.0:
            signal_desc = f"Bandwidth utilization trending upward (+{round(du_dt, 1)}%/min, current: {recent_util}%)"
            contributing_signals.append(signal_desc)

            if du_dt > 0:
                remaining_headroom = 90.0 - recent_util
                if remaining_headroom > 0:
                    tti_est = remaining_headroom / du_dt
                    tti_minutes = max(1.0, round(tti_est, 1))
                else:
                    tti_minutes = 0.5

            risk_score = min(99.0, risk_score + 35.0 + (recent_util * 0.4))
            confidence = min(0.98, 0.75 + (n / 50.0))

            if tti_minutes and tti_minutes < 20:
                alerts.append({
                    "issue_type": "PREDICTED_LINK_SATURATION",
                    "severity": "CRITICAL" if tti_minutes < 5 else "HIGH",
                    "tti_minutes": tti_minutes,
                    "target": component_name,
                    "confidence": round(confidence, 2),
                    "summary": f"Interface bandwidth saturation projected on {component_name} in {tti_minutes} mins.",
                })

        # 2. Routing Instability / Flap Entropy
        recent_flaps = sum(flaps[-5:])
        if recent_flaps > 5:
            contributing_signals.append(f"Route flap entropy spike ({recent_flaps} flaps detected in past window)")
            risk_score = min(99.0, risk_score + 40.0)
            alerts.append({
                "issue_type": "ROUTING_INSTABILITY_CASCADE",
                "severity": "HIGH",
                "tti_minutes": 8.0,
                "target": component_name,
                "confidence": 0.91,
                "summary": f"BGP/OSPF route instability cascade on {component_name} imminent.",
            })

        # 3. Tunnel / Jitter Degradation
        if losses[-1] > 2.0 or recent_lat > 60.0:
            contributing_signals.append(f"Packet loss progression ({losses[-1]}%) and latency drift ({recent_lat}ms)")
            risk_score = min(99.0, risk_score + 30.0)

        # Determine Status
        status = "HEALTHY"
        if risk_score > 75.0:
            status = "CRITICAL"
        elif risk_score > 45.0:
            status = "DEGRADED"

        return {
            "component": component_name,
            "status": status,
            "risk_score": round(risk_score, 1),
            "tti_minutes": tti_minutes,
            "confidence": round(confidence, 2),
            "current_metrics": snapshot["metrics"],
            "rate_of_change": {"du_dt_per_min": round(du_dt, 2), "dl_dt_per_min": round(dl_dt, 2)},
            "signals": contributing_signals,
            "alerts": alerts,
        }
