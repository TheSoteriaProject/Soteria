import argparse
from linter_engine import LinterEngine
import json

def main():
    parser = argparse.ArgumentParser(description="Makefile Linter for identifying potential security issues.")
    parser.add_argument('makefile_path', type=str, help="Path to the Makefile to analyze.")
    
    args = parser.parse_args()

    # Instantiate the linter engine with the provided Makefile path
    linter = LinterEngine(args.makefile_path)
    issues = linter.analyze()

    # Output the issues found
    if issues:
        print(json.dumps(issues, indent=4))
    else:
        print("No issues found.")

# In makefilelinter.py
if __name__ == "__main__":
    # Parsing arguments
    parser = argparse.ArgumentParser(description="Analyze Makefiles for security issues.")
    parser.add_argument("makefile_path", type=str, help="Path to the Makefile to analyze.")
    args = parser.parse_args()

    # Correct instantiation of LinterEngine
    linter = LinterEngine(args.makefile_path)
    issues = linter.analyze()
    # Proceed to handle the 'issues' as intended...

