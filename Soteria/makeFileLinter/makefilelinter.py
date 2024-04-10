import argparse
from linter_engine import LinterEngine
import json
import sys

def main():
    # Parse command-line arguments
    parser = argparse.ArgumentParser(description="Makefile Linter for identifying potential security issues.")
    parser.add_argument('makefile_path', type=str, help="Path to the Makefile to analyze.")
    parser.add_argument('--output', type=str, default='issues_found.json', help="Path to the output JSON file.")
    
    args = parser.parse_args()

    # Create a linter engine and analyze the Makefile
    linter = LinterEngine(args.makefile_path)
    issues = linter.analyze()

    # Write the issues to the output JSON file
    with open(args.output, 'w') as outfile:
        json.dump(issues, outfile, indent=4)

    # Print the issues to the console
    if issues:
        print(f"Issues found in {args.makefile_path}:")
        print(json.dumps(issues, indent=4))
        # Exit with an error code if any issue of 'Error' severity is found
        if any(issue['Severity'] == 'Error' for issue in issues):
            sys.exit(1)
    else:
        print(f"No issues found in {args.makefile_path}.")

if __name__ == "__main__":
    main()
