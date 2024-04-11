import json

def write_issues_to_json(issues, filename):
    """Write the list of issues to a JSON file."""
    with open(filename, 'a+') as file:
        json.dump(issues, file, indent=4)
        # Dumping Wrong. Need to adjust this so it is [{}, {}, {}] instead of [{}] [{}] [{}]