#!/usr/bin/env python3
"""
CLI Tool to inject and clear faults on AirGap components
Usage:
  ./scripts/inject-fault.py --scenario congestion --target org-a-edge-router
  ./scripts/inject-fault.py --scenario route_flap --target tor-a1
  ./scripts/inject-fault.py --scenario tunnel_deg --target org-a-edge-router
  ./scripts/inject-fault.py --scenario policy_drift --target org-a-edge-router
  ./scripts/inject-fault.py --clear
"""

import argparse
import sys
import os

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.faults.fault_injector import FaultInjector


def main():
    parser = argparse.ArgumentParser(description="AirGap Network Fault Injector CLI")
    parser.add_argument("--scenario", choices=["congestion", "route_flap", "tunnel_deg", "policy_drift"], help="Fault scenario to inject")
    parser.add_argument("--target", default="org-a-edge-router", help="Target component name")
    parser.add_argument("--duration", type=int, default=600, help="Fault duration in seconds")
    parser.add_argument("--clear", action="store_true", help="Clear all active faults")
    parser.add_argument("--status", action="store_true", help="List all active faults")

    args = parser.parse_args()
    injector = FaultInjector()

    if args.clear:
        cleared = injector.clear_faults()
        print(f"Cleared {len(cleared)} active faults.")
        return

    if args.status:
        active = injector.get_active_faults()
        print(f"Active Faults ({len(active)}):")
        for k, v in active.items():
            print(f" - [{v['type']}] Target: {k} | Severity: {v['severity']} | Elapsed: {int(time.time() - v['start_time'])}s")
        return

    if not args.scenario:
        parser.print_help()
        sys.exit(1)

    if args.scenario == "congestion":
        res = injector.inject_progressive_congestion(args.target, args.duration)
    elif args.scenario == "route_flap":
        res = injector.inject_route_flap_cascade(args.target, args.duration)
    elif args.scenario == "tunnel_deg":
        res = injector.inject_tunnel_degradation(args.target, args.duration)
    elif args.scenario == "policy_drift":
        res = injector.inject_policy_drift(args.target, args.duration)

    print(f"Successfully injected fault [{res['type']}] on component '{res['target']}'.")
    print(f"Description: {res['description']}")
    print(f"Duration: {res['duration']} seconds.")


if __name__ == "__main__":
    main()
