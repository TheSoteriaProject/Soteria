import argparse
from linter_engine import LinterEngine
from output_handler import OutputHandler

def main():
    parser = argparse.ArgumentParser(description='MakeFileLinter: A tool for linting Makefiles against security standards.')
    parser.add_argument('path_to_makefile', type=str, help='Path to the Makefile to analyze.')
    parser.add_argument('--warn', type=bool, default=False, help='Warn only mode. Will not exit with error code on issues.')
    
    args = parser.parse_args()

    linter = LinterEngine(args.path_to_makefile)
    issues = linter.analyze()
    
    OutputHandler.write_issues_to_json(issues)
    
    if issues and not args.warn:
        exit(1)

if __name__ == "__main__":
    main()
