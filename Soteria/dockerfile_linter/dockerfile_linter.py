import argparse
from dockerfile_linter_engine import DockerfileLinter
from output_handler import write_issues_to_json
import json

def main():
    parser = argparse.ArgumentParser(description="Dockerfile Linter for identifying potential security issues.")
    parser.add_argument('dockerfile_path', type=str, help="Path to the Dockerfile to analyze.")
    parser.add_argument('--output', type=str, default='issues_found.json', help="Path to the output JSON file.")
    
    args = parser.parse_args()

    # Instantiate the linter engine with the provided Dockerfile path
    linter = DockerfileLinter(args.dockerfile_path, "docker_security_rules.yaml")
    issues = linter.analyze()

    # Write the issues to the output JSON file
    write_issues_to_json(issues, args.output)

    # Output the issues found
    if issues:
        print(f"Issues found in {args.dockerfile_path}:")
        print(json.dumps(issues, indent=4))
    else:
        print(f"No issues found in {args.dockerfile_path}.")

if __name__ == "__main__":
    main()
