#!/usr/bin/env python3
"""
AirGap Fault Injection Engine & Scenario Generator
Simulates realistic network fault scenarios across the AirGap infrastructure by creating FailureSimulation CRDs.
"""

import logging
import time
from kubernetes import client, config

logger = logging.getLogger("FaultInjector")


class FaultInjector:
    def __init__(self):
        try:
            config.load_kube_config()
        except Exception:
            try:
                config.load_incluster_config()
            except Exception:
                pass
        self.api = client.CustomObjectsApi()
        self.group = "airgap.example.com"
        self.version = "v1alpha1"
        self.plural = "failuresimulations"

    def _create_crd(self, scenario: str, target_component: str, duration_seconds: int, severity: str, description: str):
        fault_id = f"fault-{scenario.replace('_', '-')}-{int(time.time())}"
        body = {
            "apiVersion": f"{self.group}/{self.version}",
            "kind": "FailureSimulation",
            "metadata": {
                "name": fault_id,
            },
            "spec": {
                "type": scenario.upper(),
                "address": target_component,
                "enabled": True,
                "description": description
            }
        }
        
        try:
            self.api.create_cluster_custom_object(self.group, self.version, self.plural, body)
            logger.info(f"Created FailureSimulation CRD for {scenario} on {target_component} (ID: {fault_id})")
        except Exception as e:
            logger.error(f"Failed to create FailureSimulation CRD: {e}")
            return {"error": str(e)}

        fault_data = {
            "id": fault_id,
            "type": scenario.upper(),
            "target": target_component,
            "start_time": time.time(),
            "duration": duration_seconds,
            "severity": severity,
            "description": description
        }

        # Add metrics expected by telemetry pipeline
        if scenario.upper() == "PROGRESSIVE_CONGESTION":
            fault_data["metrics"] = {
                "initial_utilization_pct": 42.5,
                "target_utilization_pct": 98.6,
                "initial_latency_ms": 2.1,
                "target_latency_ms": 245.0,
                "initial_packet_loss_pct": 0.0,
                "target_packet_loss_pct": 14.8,
            }
        elif scenario.upper() == "TUNNEL_DEGRADATION":
            fault_data["metrics"] = {
                "jitter_ms": 48.2,
                "rekey_failures": 7,
                "packet_loss_pct": 8.4,
                "mtu_blackhole": True,
            }
        elif scenario.upper() == "ROUTE_FLAP_CASCADE":
            fault_data["metrics"] = {
                "flap_frequency_per_min": 18,
                "route_table_entropy": 8.7,
                "convergence_delay_sec": 12.4,
                "affected_prefixes": ["10.10.0.0/24", "10.11.0.0/24"],
            }
        elif scenario.upper() == "POLICY_DRIFT":
            fault_data["metrics"] = {
                "blocked_flows_pct": 34.0,
                "acl_mismatch_rules": ["DENY_APP_TRAFFIC_10_10"],
                "controller_sync_status": "OUT_OF_SYNC",
            }

        return fault_data

    def inject_progressive_congestion(self, target_component: str = "org-a-edge-router", duration_seconds: int = 300) -> dict:
        desc = f"Progressive link bandwidth saturation and queue buildup on {target_component}."
        return self._create_crd("progressive_congestion", target_component, duration_seconds, "CRITICAL", desc)

    def inject_route_flap_cascade(self, target_component: str = "tor-a1", duration_seconds: int = 300) -> dict:
        desc = f"BGP/OSPF route advertisement flapping on {target_component} causing convergence stress."
        return self._create_crd("route_flap_cascade", target_component, duration_seconds, "HIGH", desc)

    def inject_tunnel_degradation(self, target_component: str = "org-a-edge-router", duration_seconds: int = 300) -> dict:
        desc = f"Overlay IPSec tunnel jitter and rekey failure escalation on {target_component}."
        return self._create_crd("tunnel_degradation", target_component, duration_seconds, "HIGH", desc)

    def inject_policy_drift(self, target_component: str = "org-a-edge-router", duration_seconds: int = 300) -> dict:
        desc = f"ACL policy drift and controller configuration mismatch on {target_component}."
        return self._create_crd("policy_drift", target_component, duration_seconds, "MEDIUM", desc)

    def clear_faults(self, target_component: str = None) -> list:
        cleared = []
        try:
            objs = self.api.list_cluster_custom_object(self.group, self.version, self.plural)
            for item in objs.get("items", []):
                name = item["metadata"]["name"]
                spec = item.get("spec", {})
                
                # Clear all or clear only for a specific target component
                if not target_component or spec.get("address") == target_component:
                    try:
                        self.api.delete_cluster_custom_object(self.group, self.version, self.plural, name)
                        cleared.append(name)
                        logger.info(f"Deleted FailureSimulation CRD {name}")
                    except Exception as e:
                        logger.error(f"Failed to delete FailureSimulation CRD {name}: {e}")
        except Exception as e:
            logger.error(f"Failed to list FailureSimulation CRDs: {e}")

        logger.info(f"Cleared {len(cleared)} faults.")
        return cleared

    def get_active_faults(self) -> dict:
        active_faults = {}
        try:
            objs = self.api.list_cluster_custom_object(self.group, self.version, self.plural)
            for item in objs.get("items", []):
                name = item["metadata"]["name"]
                spec = item.get("spec", {})
                if spec.get("enabled", False):
                    f_type = spec.get("type", "UNKNOWN")
                    fault_data = {
                        "id": name,
                        "type": f_type,
                        "target": spec.get("address"),
                        "start_time": time.time(), # Mocked since CRD doesn't have start_time in spec currently
                        "duration": 300,
                        "severity": "UNKNOWN",
                        "description": spec.get("description")
                    }

                    if f_type == "PROGRESSIVE_CONGESTION":
                        fault_data["metrics"] = {
                            "initial_utilization_pct": 42.5,
                            "target_utilization_pct": 98.6,
                            "initial_latency_ms": 2.1,
                            "target_latency_ms": 245.0,
                            "initial_packet_loss_pct": 0.0,
                            "target_packet_loss_pct": 14.8,
                        }
                    elif f_type == "TUNNEL_DEGRADATION":
                        fault_data["metrics"] = {
                            "jitter_ms": 48.2,
                            "rekey_failures": 7,
                            "packet_loss_pct": 8.4,
                            "mtu_blackhole": True,
                        }
                    elif f_type == "ROUTE_FLAP_CASCADE":
                        fault_data["metrics"] = {
                            "flap_frequency_per_min": 18,
                            "route_table_entropy": 8.7,
                            "convergence_delay_sec": 12.4,
                            "affected_prefixes": ["10.10.0.0/24", "10.11.0.0/24"],
                        }
                    elif f_type == "POLICY_DRIFT":
                        fault_data["metrics"] = {
                            "blocked_flows_pct": 34.0,
                            "acl_mismatch_rules": ["DENY_APP_TRAFFIC_10_10"],
                            "controller_sync_status": "OUT_OF_SYNC",
                        }

                    active_faults[spec.get("address")] = fault_data
        except Exception as e:
            logger.error(f"Failed to list active faults: {e}")
        return active_faults
