import yaml
from utilities import load_rules

class LinterEngine:
    def __init__(self, filepath):
        self.filepath = filepath
        self.rules = load_rules("security_rules.yaml")
        
    def analyze(self):
        issues = []
        with open(self.filepath, 'r') as file:
            for line_number, line in enumerate(file, start=1):
                for rule in self.rules:
                    if rule['pattern'] in line:
                        issues.append({
                            "line": line_number,
                            "issue": rule['description'],
                            "severity": rule['severity']
                        })
        return issues
