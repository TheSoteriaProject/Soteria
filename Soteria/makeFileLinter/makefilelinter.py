import argparse
from linter_engine import LinterEngine
from output_handler import write_issues_to_json
import json

def main():
    parser = argparse.ArgumentParser(description="Makefile Linter for identifying potential security issues.")
    parser.add_argument('makefile_path', type=str, help="Path to the Makefile to analyze.")
    parser.add_argument('--output', type=str, default='issues_found.json', help="Path to the output JSON file.")
    
    args = parser.parse_args()

    # Instantiate the linter engine with the provided Makefile path
    linter = LinterEngine(args.makefile_path)
    issues = linter.analyze()

    # Write the issues to the output JSON file
    write_issues_to_json(issues, args.output)

    # Output the issues found
    if issues:
        print(f"Issues found in {args.makefile_path}:")
        print(json.dumps(issues, indent=4))
    else:
        print(f"No issues found in {args.makefile_path}.")

if __name__ == "__main__":
    main()