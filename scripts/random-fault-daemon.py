#!/usr/bin/env python3
"""
AirGap Random Fault Injection Daemon
Periodically selects a random simulation target and injects a real kernel/firewall fault.
Clears faults gracefully on exit.
"""

import argparse
import logging
import os
import random
import sys
import time
import subprocess

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.faults.fault_injector import FaultInjector

# Configure logger
logging.basicConfig(
    level=logging.INFO,
    format="[%(asctime)s] [%(levelname)s] [%(name)s] %(message)s",
    handlers=[
        logging.StreamHandler(sys.stdout),
        logging.FileHandler("/tmp/random_fault_daemon.log", mode="w")
    ]
)
logger = logging.getLogger("RandomFaultDaemon")


def get_available_targets() -> list:
    """Find all running component pods matching app=airgap-component."""
    cmd = "kubectl get pods -l app=airgap-component -o jsonpath='{.items[?(@.status.phase==\"Running\")].metadata.name}'"
    res = subprocess.run(cmd, shell=True, capture_output=True, text=True)
    if res.returncode == 0 and res.stdout.strip():
        return res.stdout.strip().split()
    return ["org-a-edge-router", "org-b-worker-01", "org-b-worker-02"]


def main():
    parser = argparse.ArgumentParser(description="AirGap Random Fault Injection Daemon")
    parser.add_argument("--interval-min", type=int, default=30, help="Minimum interval in seconds between faults")
    parser.add_argument("--interval-max", type=int, default=60, help="Maximum interval in seconds between faults")
    parser.add_argument("--duration", type=int, default=90, help="Duration of injected faults in seconds")
    args = parser.parse_args()

    injector = FaultInjector()
    
    logger.info("Starting AirGap Random Fault Injection Daemon...")
    logger.info("Injecting real network delays (tc), iptables packet blocks, and route blackholes.")
    logger.info("Press Ctrl+C to stop and clear all faults.")

    try:
        while True:
            # Check for active faults
            active = injector.get_active_faults()
            if active:
                logger.info(f"Active faults detected: {list(active.keys())}. Waiting for them to resolve/expire...")
                time.sleep(10)
                continue

            # Pause before injecting the next fault
            sleep_time = random.randint(args.interval_min, args.interval_max)
            logger.info(f"System is healthy. Resting for {sleep_time} seconds before the next injection...")
            time.sleep(sleep_time)

            # Discover current target pods
            targets = get_available_targets()
            if not targets:
                logger.warning("No running component pods found. Retrying in 10 seconds...")
                time.sleep(10)
                continue

            target = random.choice(targets)
            scenario = random.choice(["congestion", "route_flap", "tunnel_deg", "policy_drift"])
            
            logger.info(f"Selected target: '{target}' | Scenario: '{scenario}' | Duration: {args.duration}s")
            
            if scenario == "congestion":
                injector.inject_progressive_congestion(target, args.duration)
            elif scenario == "route_flap":
                injector.inject_route_flap_cascade(target, args.duration)
            elif scenario == "tunnel_deg":
                injector.inject_tunnel_degradation(target, args.duration)
            elif scenario == "policy_drift":
                # Ensure iptables can be installed/run
                injector.inject_policy_drift(target, args.duration)

            # Wait for fault to be active or resolved
            start_wait = time.time()
            while time.time() - start_wait < args.duration:
                # Check if the fault was resolved early by the Copilot MCP tools
                active_now = injector.get_active_faults()
                if target not in active_now:
                    logger.info(f"Active fault on '{target}' was cleared/remediated early by Copilot MCP tools!")
                    break
                time.sleep(5)

    except KeyboardInterrupt:
        logger.info("KeyboardInterrupt received. Stopping daemon...")
    finally:
        logger.info("Cleaning up and clearing all active faults in the cluster...")
        injector.clear_faults()
        logger.info("Shutdown complete.")


if __name__ == "__main__":
    main()
