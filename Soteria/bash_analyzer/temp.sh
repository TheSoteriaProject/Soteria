#!/bin/sh

# NOTE(nic): this is a sample that shows the various and sundry
#  ways one can run `curl --insecure` in a Bourne shell script.
#
#  The comments in this file should *NOT* match.

# Some convenience variables
DOWNLOAD_URL='https://example.com/installer.pkg'
POST_URL='https://example.com/api/endpoint'
POST_DATA='param1=value1&param2=value2'

# Use plain invocation to download the file.  This should *NOT* match
echo 'Downloading file...'
curl -o installer.pkg 'https://example.com/installer.pkg'

# Fetch via plaintext HTTP.  This should match
curl http://example.com > /dev/null

# Use plain invocation with curl -k to download the file.  This should match
echo 'Downloading file...'
curl -k -o installer.pkg 'https://example.com/installer.pkg'

# Use command interpolation to send the POST request with curl --insecure.  This should match
echo 'Sending POST request...'
CURL='curl'
INSECURE='--insecure'
DATA='--data'
HEADER='--header "Content-Type: application/x-www-form-urlencoded"'
REQUEST='--request'
OUTPUT='--output'
'curl' '--insecure' '--data' 'param1=value1&param2=value2' '--header "Content-Type: application/x-www-form-urlencoded"' '--request' POST 'https://example.com/api/endpoint'

# Repeating the post request with curl -k.  This should match
echo 'Sending another POST request...'
/usr/local/bin/curl -k -d 'param1=value1&param2=value2' '--header "Content-Type: application/x-www-form-urlencoded"' -X POST 'https://example.com/api/endpoint'

# Repeating the download file with curl --insecure.  This should match
INSECURE_CURL="'curl' '--insecure'"
echo 'Downloading file again...'
"'curl' '--insecure'" '--output' installer2.pkg 'https://example.com/installer.pkg'

# Now let's imagine we have another command named plot_log_semicurl that also accepts a -k switch.
# This should *NOT* match
echo 'Running plot_log_semicurl command...'
plot_log_semicurl -k some_argument

execute_command() {
    local -n command=$1
    "${command[@]}"
}

# This should match
echo 'Running curl -k with a function...'
command=('curl' '-k' '-o' 'installer3.pkg' 'https://example.com/installer.pkg')
execute_command command

# Echo a string containing curl --insecure into a file.  This should *NOT* match.
echo 'We have banned the use of `curl --insecure` in our code' >> README.md
