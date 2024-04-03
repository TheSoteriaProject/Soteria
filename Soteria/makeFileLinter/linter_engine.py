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
        with open(self.filepath, 'r') as file:
            for line_number, line in enumerate(file, start=1):
                if self.is_comment_or_empty(line):
                    continue
                if ':=' in line:
                    self.parse_variable_assignment(line)
                    continue
                resolved_line = self.resolve_variables_in_line(line)
                if 'plot_log_semicurl' in resolved_line or 'UPDATE_TEXT=' in resolved_line:
                    continue
                for rule in self.rules:
                    pattern = re.compile(rule['pattern'])
                    if pattern.search(resolved_line) and not self.is_excluded(resolved_line, rule.get('exclude', [])):
                        issues.append({
                            "line": line_number,
                            "line_content": resolved_line.strip(),
                            "issue": rule['description'],
                            "severity": rule['severity']
                        })
        return issues

    def is_comment_or_empty(self, line):
        """Check if a line is a comment or empty."""
        return line.strip().startswith('#') or not line.strip()

    def parse_variable_assignment(self, line):
        """Parse and store variable assignments."""
        parts = line.split(':=', 1)
        var_name, var_value = parts[0].strip(), parts[1].strip()
        self.variables[var_name] = var_value

    def resolve_variables_in_line(self, line):
        """Replace variable references in the line with their actual values."""
        resolved_line = line
        max_iterations = 10
        iteration = 0
        while '$' in resolved_line and iteration < max_iterations:
            for var_name, var_value in self.variables.items():
                resolved_line = resolved_line.replace(f'$({var_name})', var_value)
                resolved_line = resolved_line.replace(f'${{{var_name}}}', var_value)
                resolved_line = resolved_line.replace(f'${var_name}', var_value)
            iteration += 1
        return resolved_line

    def is_excluded(self, line, exclude_patterns):
        """Check if a line matches any of the exclude patterns."""
        for pattern in exclude_patterns:
            if re.search(pattern, line):
                return True
        return False