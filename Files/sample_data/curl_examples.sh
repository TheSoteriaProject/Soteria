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
curl -o installer.pkg "${DOWNLOAD_URL}"

# Fetch via plaintext HTTP.  This should match
curl http://example.com > /dev/null

# Use plain invocation with curl -k to download the file.  This should match
echo 'Downloading file...'
curl -k -o installer.pkg "${DOWNLOAD_URL}"

# Use command interpolation to send the POST request with curl --insecure.  This should match
echo 'Sending POST request...'
CURL='curl'
INSECURE='--insecure'
DATA='--data'
HEADER='--header "Content-Type: application/x-www-form-urlencoded"'
REQUEST='--request'
OUTPUT='--output'
${CURL} ${INSECURE} ${DATA} "${POST_DATA}" ${HEADER} ${REQUEST} POST "${POST_URL}"

# Repeating the post request with curl -k.  This should match
echo 'Sending another POST request...'
/usr/local/bin/curl -k -d "${POST_DATA}" ${HEADER} -X POST "${POST_URL}"

# Repeating the download file with curl --insecure.  This should match
INSECURE_CURL="${CURL} ${INSECURE}"
echo 'Downloading file again...'
${INSECURE_CURL} ${OUTPUT} installer2.pkg "${DOWNLOAD_URL}"

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
command=('curl' '-k' '-o' 'installer3.pkg' "${DOWNLOAD_URL}") # Ignore Match
execute_command command

# Echo a string containing curl --insecure into a file.  This should *NOT* match.
echo 'We have banned the use of `curl --insecure` in our code' >> README.md
