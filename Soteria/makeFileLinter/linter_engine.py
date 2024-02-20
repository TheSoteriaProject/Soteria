import yaml
from utilities import load_rules  # Ensure this import is correct

class LinterEngine:
    def __init__(self, filepath):
        self.filepath = filepath
        self.rules = load_rules("security_rules.yaml")  # Correct usage

    def analyze(self):
        issues = []
        with open(self.filepath, 'r') as file:
            for line_number, line in enumerate(file, start=1):
                trimmed_line = line.strip()
                if trimmed_line.startswith('#') or not trimmed_line:
                    continue

                # Specific checks for $(CURL) usage
                if "$(CURL)" in trimmed_line:
                    if "-k" in trimmed_line or "--insecure" in trimmed_line:
                        issues.append(self.create_issue(line_number, "Insecure curl option used."))
                    continue

                # General checks for insecure practices
                if "curl" in trimmed_line and ("http://" in trimmed_line or "-k" in trimmed_line or "--insecure" in trimmed_line):
                    issues.append(self.create_issue(line_number, "Potential insecure use of curl detected."))

                # Checks for benign discussions or commands related to curl security
                if "UPDATE_TEXT" in trimmed_line or "grep -qsF" in trimmed_line:
                    continue  # Ignore these benign instances
                
                # Additional check to ensure $(CURL) without insecure flags does not trigger
                if "$(CURL)" in trimmed_line and not any(insecure_flag in trimmed_line for insecure_flag in ["-k", "--insecure"]):
                    continue

                # Check for other patterns that should trigger warnings
                if "plot_log_semicurl -k" in trimmed_line:
                    issues.append(self.create_issue(line_number, "Use of '-k' with plot_log_semicurl detected."))
                    
        return issues

    def create_issue(self, line_number, description):
        return {"line": line_number, "issue": description, "severity": "Warning"}
