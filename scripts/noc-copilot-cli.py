#!/usr/bin/env python3
"""
Interactive CLI Tool for AirGap Autonomous AI NOC Copilot
Usage:
  ./scripts/noc-copilot-cli.py
  ./scripts/noc-copilot-cli.py --target org-a-edge-router
  ./scripts/noc-copilot-cli.py --query "What is likely to fail next?"
"""

import argparse
import os
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from pkg.analytics.predictive_engine import PredictiveAnalyticsEngine
from pkg.copilot.offline_llm import AirGapNOCCopilot
from pkg.telemetry.pipeline import TelemetryPipeline

pipeline = TelemetryPipeline()
engine = PredictiveAnalyticsEngine(pipeline)
copilot = AirGapNOCCopilot(engine)


def main():
    parser = argparse.ArgumentParser(description="AirGap Autonomous AI NOC Copilot CLI")
    parser.add_argument("--target", default="org-a-edge-router", help="Component to analyze")
    parser.add_argument("--query", help="Natural language query for the Copilot")

    parser.add_argument("--autofix", action="store_true", help="Execute autonomous MCP skill repair")

    args = parser.parse_args()

    if args.autofix:
        res = copilot.mcp.call_tool("execute_autonomous_repair", {"component_name": args.target})
        print(f"=== MCP AUTONOMOUS SKILL REPAIR RESULT FOR {args.target} ===")
        import json
        print(json.dumps(res, indent=2))
        return

    if args.query:
        answer = copilot.process_natural_language_query(args.query)
        print(answer)
        return

    print("=========================================================================")
    print("   AirGap Autonomous AI NOC Copilot (Air-Gapped Offline Decision Support)")
    print("=========================================================================")
    
    resp = copilot.generate_incident_copilot_response(args.target)

    print(f"\n[TARGET COMPONENT]: {resp['component']}")
    print(f"[STATUS]: {resp['status']} | [RISK SCORE]: {resp['risk_score']}% | [CONFIDENCE]: {int(resp['confidence_score']*100)}%")
    if resp['tti_minutes']:
        print(f"[TIME-TO-IMPACT TTI]: {resp['tti_minutes']} minutes before SLA breach!")

    print("\n-------------------------------------------------------------------------")
    print("Q1: WHAT IS LIKELY TO FAIL NEXT — AND WHEN?")
    print("-------------------------------------------------------------------------")
    print(resp['q1_forecast'])

    print("\n-------------------------------------------------------------------------")
    print("Q2: WHY IS RISK ASSESSED AS ELEVATED — WHICH SIGNALS CONTRIBUTED?")
    print("-------------------------------------------------------------------------")
    print(resp['q2_reasoning'])

    print("\n-------------------------------------------------------------------------")
    print("Q3: WHAT CORRECTIVE ACTION SHOULD BE TAKEN BEFORE SLA / SECURITY IMPACT?")
    print("-------------------------------------------------------------------------")
    print(resp['q3_corrective_action'])
    print("\n=========================================================================\n")


if __name__ == "__main__":
    main()
