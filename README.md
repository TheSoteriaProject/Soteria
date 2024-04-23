# Soteria

Soteria is a Insecure Communication Linter for Bourne Shell, Makefiles, and Dockerfiles. 

## Tool Status
[![Go](https://github.com/TheSoteriaProject/Soteria/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/TheSoteriaProject/Soteria/actions/workflows/go.yml)
[![Python](https://github.com/TheSoteriaProject/Soteria/actions/workflows/python-app.yml/badge.svg?branch=main)](https://github.com/TheSoteriaProject/Soteria/actions/workflows/python-app.yml)
[![license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

## How to use

**Soteria Tool User Guide**

The Soteria tool is a powerful utility designed to analyze projects for security vulnerabilities and provide various options to customize its behavior. Below is a detailed guide on how to use the tool effectively.

### Usage
./Soteria [Flags] [Project Path]

- `./Soteria`: The name of the executable file.
- `[Flags]`: Optional flags to customize the behavior of the tool.
- `[Project Path]`: The path to the project you want to analyze.

### Flags

1. `--warn`: 
   - **Description**: Allows the tool to return true even if it encounters issues flagged by the insecure communication linter.
   - **Default**: `--warn=false`
   - **Usage**: `--warn=true` or `--warn=false`

2. `--help`:
   - **Description**: Displays a simple help page with information about available flags and examples.
   - **Usage**: `--help`

3. `--version`:
   - **Description**: Displays the latest git tag release of the Soteria tool.
   - **Usage**: `--version`

4. `--enableBash`:
   - **Description**: Enables the Bash Static Analyzer.
   - **Default**: `--enableBash=true`
   - **Usage**: `--enableBash=true` or `--enableBash=false`

5. `--enableMakefile`:
   - **Description**: Enables the Makefile static analyzer.
   - **Default**: `--enableMakefile=true`
   - **Usage**: `--enableMakefile=true` or `--enableMakefile=false`

6. `--enableDockerfile`:
   - **Description**: Enables the Dockerfile static analyzer.
   - **Default**: `enableDockerfile=true`
   - **Usage**: `--enableDockerfile=true` or `--enableDockerfile=false`
  
7. `--enableLogPrint`:
   - **Descritpion**: Enable Log Prints for static analzyer.
   - **Usage**: `--enableLogPrint=true` or `--enableLogPrint=false`

8. `--test`:
   - **Description**: Runs unit tests for the function to confirm code changes worked and the tool is still functional.
   - **Usage**: `--test`

### Examples

1. Analyze a project with default settings:
./Soteria /path/to/your/project

2. Analyze a project while ignoring insecure communication issues:
./Soteria --warn=true /path/to/your/project

3. Disable Bash Static Analyzer:
./Soteria --enableBash=false /path/to/your/project

4. Display help page:
./Soteria --help

5. Display the version of the Soteria tool:
./Soteria --version

6. Run unit tests:
./Soteria --test

### Example JSON Output
```json
{
    "FileName": "../Files/sample_data5/wget_examples.sh",
    "LineNumber": 48,
    "Line": "command=('wget' '--no-check-certificate' '-O' 'installer3.pkg' \"${DOWNLOAD_URL}\")",
    "Issue": "wget --no-check-certificate",
    "Severity": "Error"
}
```

### Extra Features
- If a line contains `# Ignore Match` the line will still match, but the Severity will change from Error to Warn. This helps phase out certain forms of Insecure Communication gradually. (Bash Exclusive)

### Notes
- Flags are optional, but a project path is required for the tool to run.
- Flags can be combined as needed.
- Make sure to replace `/path/to/your/project` with the actual path to your project directory.

### Tool Setup
Follow these steps to set up and run the Soteria tool:

1. **Clone the Repository:**
   `git clone https://github.com/TheSoteriaProject/Soteria.git`

2. **Step In:**
   `cd Soteria`
   
3. **Build the Tool:**
   `go build`

4. **Run the Tool:**
   `./Soteria [Flags] [Project Path]`
   - `[Flags]`: Optional flags to customize the behavior of the tool (refer to the user guide for available flags).
   - `[Project Path]`: The path to the project you want to analyze.

That's it! You're now ready to use the Soteria tool to analyze your projects for security vulnerabilities.

### Tool Configuration
The tool currently has two main ways to configure it outside of flags. One being the `rules.yaml` and the second being the `.soteriaignore` file. For the first file it is used to configure the security rules that are to be used and is configured per analyzer. This allows more customization based on the setting for which the tool is being used. 

Example `rules.yaml` Config File
```yaml
curl:
  - -k
  - http
  - --insecure

wget:
  - http
  - --no-check-certificate

ssh:
  - http
  - StrictHostKeyChecking=no
```

For the `.soteriaignore` file this is to be used to indicate which directories will be ignored during the creation of the file pool for the tool. Within the file it supports single-line comments to either give info or to comment out a past directory that was to be skipped.
Example `.soteriaignore` Config File
```
DoNotEnterFolder # Super Important to skip.
sample_data3
sample_data4
sample_data7
# sample_data9
```

## Contributing

If you want to contribute to this project, please read [CONTRIBUTING.md](CONTRIBUTING.md).

## Discussion

If you would like to discuss the project feel free to chat here [Discussion Board](https://github.com/TheSoteriaProject/Soteria/discussions).  
#### Disclaimer

This discussion board is intended to provide a platform for users to engage in open and respectful conversations. While we encourage free expression of ideas and opinions, please be aware that inappropriate or offensive content will not be tolerated. Users are expected to adhere to the following guidelines:

1. Respectful Communication: Users must communicate with courtesy and respect towards others. Personal attacks, harassment, hate speech, or any form of discrimination will not be tolerated.

2. Content Moderation: The administrators of this discussion board reserve the right to moderate and remove content that violates community guidelines. This includes but is not limited to offensive language, explicit content, or any materials deemed inappropriate.

3. Privacy and Confidentiality: Avoid sharing personal information or any confidential details. Respect the privacy of others and refrain from engaging in any activity that compromises individual privacy.

4. Legal Compliance: Users are expected to comply with all applicable laws and regulations. Any discussions or content that violates local, national, or international laws will be subject to removal.

5. User Responsibility: Users are responsible for the content they post. By participating in discussions, users acknowledge that they have read and understood these guidelines. Ignorance of the rules will not exempt users from consequences.
 
6. Reporting Violations: Users are encouraged to report any violations of these guidelines to the administrators promptly. This includes reporting inappropriate content or behavior that goes against the established community standards.

7. Disclaimer of Liability: The administrators and owners of this discussion board are not responsible for the content posted by users. Users participate at their own risk and acknowledge that the administrators are not liable for any consequences resulting from discussions on this platform.

By using this discussion board, you agree to abide by these guidelines. The administrators reserve the right to take appropriate action, including warnings, temporary suspension, or permanent banning of users who violate these guidelines.


## Help
Welcome to our project on GitHub! If you find yourself in need of assistance, we encourage you to leverage our discussion boards. The boards are a collaborative space where our community comes together to share ideas, ask questions, and offer support. Before posting, please take a moment to search for existing discussions, as your question might have already been addressed. If not, feel free to start a new thread and our community members or maintainers will do their best to assist you. Let's build a helpful and engaging community – your input and questions are valuable!

Also feel free to check out our [HELPPAGE.md](HELPPAGE.md)
