#!/usr/bin/env python3
"""
AirGap CustomResourceDefinition (CRD) Offline validator.
Parses all CRD manifests in crds/ and verifies their structural validity.
"""

import os
import sys
import yaml

CRD_DIR = "crds"

def main():
    if not os.path.isdir(CRD_DIR):
        print(f"Error: Directory '{CRD_DIR}' not found.")
        sys.exit(1)

    crd_files = [f for f in os.listdir(CRD_DIR) if f.endswith(".yaml") or f.endswith(".yml")]
    
    if not crd_files:
        print(f"No YAML files found in '{CRD_DIR}'.")
        sys.exit(0)

    print(f"Validating {len(crd_files)} CRD manifests...")
    success = True

    for filename in crd_files:
        filepath = os.path.join(CRD_DIR, filename)
        try:
            with open(filepath, "r") as f:
                documents = list(yaml.safe_load_all(f))
            
            for doc in documents:
                if not doc:
                    continue
                
                # Verify basic CRD fields
                api_version = doc.get("apiVersion")
                kind = doc.get("kind")
                metadata = doc.get("metadata", {})
                name = metadata.get("name")
                spec = doc.get("spec", {})
                
                if api_version != "apiextensions.k8s.io/v1":
                    print(f"FAIL: {filename} - invalid apiVersion '{api_version}', expected 'apiextensions.k8s.io/v1'")
                    success = False
                    continue
                
                if kind != "CustomResourceDefinition":
                    print(f"FAIL: {filename} - invalid kind '{kind}', expected 'CustomResourceDefinition'")
                    success = False
                    continue

                if not name:
                    print(f"FAIL: {filename} - missing metadata.name")
                    success = False
                    continue

                if "group" not in spec or "names" not in spec or "versions" not in spec:
                    print(f"FAIL: {filename} - missing group, names, or versions in spec")
                    success = False
                    continue

                print(f"PASS: {filename} ({name})")
        except Exception as e:
            print(f"FAIL: {filename} - YAML parsing error: {e}")
            success = False

    if not success:
        sys.exit(1)
    
    print("All CustomResourceDefinitions successfully validated!")

if __name__ == "__main__":
    main()
