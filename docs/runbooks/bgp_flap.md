# Runbook: BGP / OSPF Route Flap Cascade Mitigation

## Overview
Triggered when high route flap entropy (>5 flaps/min) causes dynamic routing convergence stress.

## Root Cause Analysis
1. Intermittent physical layer link dampening issue.
2. BGP Hold-Timer expiration due to transient CPU saturation on TOR switch or edge router.
3. Asymmetric route propagation resulting in route loops.

## Recommended Corrective Actions
1. **Apply Route Dampening**: Enable BGP route dampening parameters (`half-life 5m`, `reuse 750`, `suppress 2000`).
2. **Increase Hold-Time**: Adjust BGP neighbor keepalive/holdtime settings to `30s / 90s` on target router.
3. **Verify Physical/Underlay Link**: Inspect interface error counters for CRC framing errors.
