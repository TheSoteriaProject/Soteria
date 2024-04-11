import json

def write_issues_to_json(issues, output_filename):
    """Write issues to a JSON output file."""
    with open(output_filename, 'w') as file:
        json.dump(issues, file, indent=4)

# Other output handling functions can be added here if needed
