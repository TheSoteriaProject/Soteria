#!/bin/sh

# NOTE(nic): a small example of an insecure `ssh` invocation.

ssh_opts='-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
ssh_cmd='/usr/bin/ssh'

# Insecure SSH invocation with options as variables
# Covers split commands in same line case.
echo 'Insecure SSH invocation with options as variables:'
$ssh_cmd $ssh_opts user@example.com; $ssh_cmd $ssh_opts user@example.com; $ssh_cmd $ssh_opts user@example.com
# Covers line continue and lien continue case.
$ssh_cmd $ssh_opts user@example.com; $ssh_cmd $ssh_opts user@example.com \
$ssh_cmd $ssh_opts user@example.com
# Covers just line continue case.
$ssh_cmd $ssh_opts user@example.com $ssh_cmd $ssh_opts user@example.com \
$ssh_cmd $ssh_opts user@example.com

# Plain insecure SSH invocation
echo 'Plain insecure SSH invocation:'
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null user@example.com
