#!/usr/bin/env python3
import json
import subprocess
import sys
import os

RULESET_NAME = "Strong Main Branch Protection"
RULESET_FILE = os.path.join(os.path.dirname(__file__), "..", ".github", "rulesets", "main-ruleset.json")

def run_command(cmd):
    try:
        # Run command and capture output
        result = subprocess.run(cmd, capture_output=True, text=True, check=True)
        return result.stdout.strip()
    except subprocess.CalledProcessError as e:
        print(f"Error running command {' '.join(cmd)}: {e.stderr.strip()}", file=sys.stderr)
        return None
    except FileNotFoundError:
        return None

def main():
    # 1. Check gh installation and authentication
    print("Checking GitHub CLI status...")
    gh_version = run_command(["gh", "--version"])
    if gh_version is None:
        print("Error: GitHub CLI ('gh') is not installed or not in PATH.", file=sys.stderr)
        print("Please install it from https://cli.github.com/ first.", file=sys.stderr)
        sys.exit(1)

    # gh auth status exits with non-zero if not logged in
    auth_status = run_command(["gh", "auth", "status"])
    if auth_status is None:
        print("Error: GitHub CLI is not authenticated. Please run 'gh auth login' first.", file=sys.stderr)
        sys.exit(1)

    # 2. Get repository nameWithOwner
    repo_info_raw = run_command(["gh", "repo", "view", "--json", "nameWithOwner"])
    if not repo_info_raw:
        print("Error: Could not retrieve repository info. Make sure you are in a Git repository with a GitHub remote.", file=sys.stderr)
        sys.exit(1)
    
    try:
        repo_info = json.loads(repo_info_raw)
        repo_path = repo_info["nameWithOwner"]
    except Exception as e:
        print(f"Error parsing repository info JSON: {e}", file=sys.stderr)
        sys.exit(1)

    print(f"Target repository: {repo_path}")

    # 3. Read and validate ruleset configuration
    if not os.path.exists(RULESET_FILE):
        print(f"Error: Ruleset file not found at {RULESET_FILE}", file=sys.stderr)
        sys.exit(1)

    try:
        with open(RULESET_FILE, "r") as f:
            json.load(f) # Validate JSON format
    except Exception as e:
        print(f"Error reading/parsing ruleset JSON: {e}", file=sys.stderr)
        sys.exit(1)

    # 4. Fetch existing rulesets
    print("Fetching existing rulesets...")
    rulesets_raw = run_command(["gh", "api", f"repos/{repo_path}/rulesets"])
    if rulesets_raw is None:
        print("Error: Failed to fetch existing rulesets from GitHub API.", file=sys.stderr)
        sys.exit(1)

    try:
        rulesets = json.loads(rulesets_raw)
    except Exception as e:
        print(f"Error parsing rulesets list JSON: {e}", file=sys.stderr)
        sys.exit(1)

    # 5. Check if ruleset already exists
    existing_ruleset_id = None
    for r in rulesets:
        if r.get("name") == RULESET_NAME:
            existing_ruleset_id = r.get("id")
            break

    # 6. Apply ruleset
    if existing_ruleset_id:
        print(f"Ruleset '{RULESET_NAME}' already exists (ID: {existing_ruleset_id}). Updating it...")
        endpoint = f"repos/{repo_path}/rulesets/{existing_ruleset_id}"
        success = run_command(["gh", "api", "--method", "PUT", endpoint, "--input", RULESET_FILE])
    else:
        print(f"Ruleset '{RULESET_NAME}' does not exist. Creating a new one...")
        endpoint = f"repos/{repo_path}/rulesets"
        success = run_command(["gh", "api", "--method", "POST", endpoint, "--input", RULESET_FILE])

    if success is not None:
        print(f"Successfully applied the ruleset '{RULESET_NAME}'!")
    else:
        print("Failed to apply the ruleset.", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()
