import yaml

def load_rules(filepath):
    with open(filepath, 'r') as file:
        return yaml.safe_load(file)
