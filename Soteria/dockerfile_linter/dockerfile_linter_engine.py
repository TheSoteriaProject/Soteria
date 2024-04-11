import argparse
import yaml
import re
import json

class DockerfileLinter:
    def __init__(self, filepath, rules_file):
        self.filepath = filepath
        self.rules = self.load_rules(rules_file)

    def load_rules(self, rules_file):
        """Load security rules from a YAML file."""
        with open(rules_file, 'r') as file:
            return yaml.safe_load(file)

    def extract_variables(self):
        """Extract variables defined in the Dockerfile for interpolation."""
        variables = {}
        with open(self.filepath, 'r') as file:
            for line in file:
                # Check for lines defining variables using ARG or ENV instructions
                if line.strip().startswith(('ARG', 'ENV')):
                    parts = line.split('=', 1)
                    if len(parts) == 2:
                        var_name = parts[0].split(' ')[1].strip()
                        var_value = parts[1].strip()
                        variables[var_name] = var_value
        return variables

    def analyze(self):
        """Analyze the Dockerfile for security issues."""
        issues = []
        variables = self.extract_variables() # might be using last assignment of variable

        with open(self.filepath, 'r') as file:
            for line_number, line in enumerate(file, start=1):
                if line.strip().startswith(('ARG', '#', "'", '"')):
                    continue
                # Exclude content within single quotes
                line_without_single_quotes = re.sub(r'\'[^\']*\'', '', line)
                if line_without_single_quotes.strip().startswith('#'):
                    continue

                # Replace variables in line
                for var_name, var_value in variables.items():
                    line_without_single_quotes = line_without_single_quotes.replace(f'${{{var_name}}}', var_value)

                for rule in self.rules:
                    pattern = re.compile(rule['pattern'])
                    if pattern.search(line_without_single_quotes) and not self.is_excluded(line_without_single_quotes, rule.get('exclude', [])):
                        issues.append({
                            "FileName": self.filepath,
                            "LineNumber": line_number,
                            "Line": line.strip(),
                            "Issue": rule['description'],
                            "Severity": rule['severity']
                        })
        return issues

    def should_skip_line(self, line):
        """Check if a line should be skipped."""
        return line.strip().startswith(('#', "'", '"')) or not line.strip()

    def is_excluded(self, line, exclude_patterns):
        """Check if a line matches any of the exclude patterns."""
        for pattern in exclude_patterns:
            if re.search(pattern, line):
                return True
        return False

def main():
    # Parse command-line arguments
    parser = argparse.ArgumentParser(description="Dockerfile Linter")
    parser.add_argument("dockerfile_path", type=str, help="Path to the Dockerfile to analyze")
    parser.add_argument("rules_file", type=str, help="Path to the YAML file containing security rules")
    parser.add_argument("--output", type=str, default="issues_found.json", help="Path to the output JSON file")
    args = parser.parse_args()

    # Instantiate the linter with Dockerfile and rules file paths
    linter = DockerfileLinter(args.dockerfile_path, args.rules_file)

    # Analyze Dockerfile for security issues
    issues = linter.analyze()

    # Output the issues found
    if issues:
        print(f"Issues found in {args.dockerfile_path}:")
        for issue in issues:
            print(issue)
    else:
        print(f"No issues found in {args.dockerfile_path}.")

    # Write issues to JSON file
    with open(args.output, "w") as json_file:
        json.dump(issues, json_file, indent=4)

if __name__ == "__main__":
    main()
