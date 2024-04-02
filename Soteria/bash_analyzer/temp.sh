#!/bin/sh

# NOTE(nic): a small example of an insecure `ssh` invocation.

ssh_opts='-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
ssh_cmd='/usr/bin/ssh'

# Insecure SSH invocation with options as variables
# Covers split commands in same line case.
echo 'Insecure SSH invocation with options as variables:'
'/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com; '/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com; '/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com
# Covers line continue and lien continue case.
'/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com; '/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com \
'/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com
# Covers just line continue case.
'/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com '/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com \
'/usr/bin/ssh' '-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' user@example.com

# Plain insecure SSH invocation
echo 'Plain insecure SSH invocation:'
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null user@example.com
