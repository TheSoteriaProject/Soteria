import yaml

class LinterEngine:
    def __init__(self, filepath):
        self.filepath = filepath
        self.rules = self.load_rules("security_rules.yaml")
        self.variables = {}

    def load_rules(self, filepath):
        """Load security rules from a YAML file."""
        with open(filepath, 'r') as file:
            return yaml.safe_load(file)

    def analyze(self):
        """Analyze the Makefile for security issues."""
        issues = []
        with open(self.filepath, 'r') as file:
            for line_number, line in enumerate(file, start=1):
                if self.is_comment_or_empty(line):
                    continue
                if '=' in line:
                    self.parse_variable_assignment(line)
                    continue
                resolved_line = self.resolve_variables_in_line(line)
                for rule in self.rules:
                    if rule['pattern'] in resolved_line:
                        issues.append({
                            "line": line_number,
                            "issue": rule['description'],
                            "severity": rule['severity']
                        })
        return issues

    def is_comment_or_empty(self, line):
        """Check if a line is a comment or empty."""
        return line.strip().startswith('#') or not line.strip()

    def parse_variable_assignment(self, line):
        """Parse and store variable assignments."""
        parts = line.split('=', 1)
        var_name, var_value = parts[0].strip(), parts[1].strip()
        self.variables[var_name] = var_value

    def resolve_variables_in_line(self, line):
        """Replace variable references in the line with their actual values."""
        for var_name, var_value in self.variables.items():
            line = line.replace(f"$({var_name})", var_value)
        return line

# Example usage
if __name__ == "__main__":
    import sys
    from output_handler import write_issues_to_json

    if len(sys.argv) != 2:
        print("Usage: python linter_engine.py <path_to_makefile>")
        sys.exit(1)
    
    filepath = sys.argv[1]
    linter = LinterEngine(filepath)
    issues = linter.analyze()

    # Define the output JSON file name
    output_filename = "issues_found.json"
    # Use output_handler to write issues to JSON
    write_issues_to_json(issues, output_filename)

    print(f"Issues have been written to {output_filename}")
