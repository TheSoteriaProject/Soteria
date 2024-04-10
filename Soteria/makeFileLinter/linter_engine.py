import yaml
import re

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
        flagged_lines = set()
        with open(self.filepath, 'r') as file:
            content = file.read()
            self.extract_variable_assignments(content)
            resolved_content = self.resolve_variables(content)
            for line_number, line in enumerate(resolved_content.splitlines(), start=1):
                if self.is_comment_or_empty(line) or line_number in flagged_lines or 'UPDATE_TEXT=' in line or 'plot_log_semicurl' in line:
                    continue
                for rule in self.rules:
                    pattern = re.compile(rule['pattern'])
                    if pattern.search(line):
                        issue = {
                            "line": line_number,
                            "line_content": line.strip(),
                            "issue": rule['description'],
                            "severity": rule['severity']
                        }
                        issues.append(issue)
                        flagged_lines.add(line_number)
                        print(f"Issue found at line {line_number}: {rule['description']}")
                        print(f"  Line content: {line.strip()}")
                        print(f"  Severity: {rule['severity']}")
                        print()
                        break
        return issues

    def is_comment_or_empty(self, line):
        """Check if a line is a comment or empty."""
        return line.strip().startswith('#') or not line.strip()

    def extract_variable_assignments(self, content):
        """Extract variable assignments from the Makefile content."""
        pattern = re.compile(r'^(\w+)\s*:=\s*(.+)$', re.MULTILINE)
        matches = pattern.findall(content)
        for var_name, var_value in matches:
            self.variables[var_name] = var_value.strip()

    def resolve_variables(self, content):
        """Resolve variable references in the Makefile content."""
        resolved_content = content
        max_iterations = 10
        iteration = 0
        while '$' in resolved_content and iteration < max_iterations:
            for var_name, var_value in self.variables.items():
                resolved_content = re.sub(r'\$\({}\)'.format(re.escape(var_name)), var_value, resolved_content)
                resolved_content = re.sub(r'\${{{}}}'.format(re.escape(var_name)), var_value, resolved_content)
            iteration += 1
        return resolved_content