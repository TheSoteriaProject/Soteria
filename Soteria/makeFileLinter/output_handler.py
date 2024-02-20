import json
from datetime import datetime

class OutputHandler:
    @staticmethod
    def write_issues_to_json(issues):
        # Get current date and time
        now = datetime.now()
        # Format date and time as 'YYYY-MM-DD_HH-MM-SS'
        timestamp = now.strftime("%Y-%m-%d_%H-%M-%S")
        # Create filename with current date and time
        filename = f"issues_found_{timestamp}.json"
        
        with open(filename, 'w') as file:
            json.dump(issues, file, indent=4)
        print(f"Issues written to {filename}")
