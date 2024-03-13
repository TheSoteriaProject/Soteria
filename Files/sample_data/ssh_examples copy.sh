#!/bin/sh

# NOTE(nic): a small example of an insecure `ssh` invocation.

ssh_opts='-o StrictHostKeyChecking=yes -o UserKnownHostsFile=/dev/null'
ssh_cmd='/usr/bin/ssh'

# Insecure SSH invocation with options as variables
echo 'Insecure SSH invocation with options as variables:'
$ssh_cmd $ssh_opts user@example.com

# Plain insecure SSH invocation
echo 'Plain insecure SSH invocation:'
ssh -o StrictHostKeyChecking=yes -o UserKnownHostsFile=/dev/null user@example.com
