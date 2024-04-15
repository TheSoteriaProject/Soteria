#!/bin/sh

# NOTE(nic): this is a sample that shows the various and sundry
#  ways one can run `wget --no-check-certificate` in a Bourne shell
#  script.
#
#  The comments in this file should *NOT* match.

# Some convenience variables
DOWNLOAD_URL='https://example.com/installer.pkg'
POST_URL='https://example.com/api/endpoint'
POST_DATA='param1=value1&param2=value2'

# Use plain invocation to download the file.  This should *NOT* match
echo 'Downloading file...'
wget -O installer.pkg 'https://example.com/installer.pkg'

# Fetch via plaintext HTTP.  This should match
wget --spider http://example.com

# Use plain invocation with wget --no-check-certificate to download the file.  This should match
echo 'Downloading file...'
wget --no-check-certificate -O installer.pkg 'https://example.com/installer.pkg'

# Use command interpolation to send the POST request with wget --no-check-certificate.  This should match
echo 'Sending POST request...'
WGET='wget'
NO_CHECK_CERTIFICATE='--no-check-certificate'
POST='--post-data'
HEADER='--header=Content-Type:application/x-www-form-urlencoded'
'wget' '--no-check-certificate' '--post-data'='param1=value1&param2=value2' '--header=Content-Type:application/x-www-form-urlencoded' -O - 'https://example.com/api/endpoint'

# Repeating the post request with wget --no-check-certificate.  This should match
echo 'Sending another POST request...'
/usr/local/bin/wget --no-check-certificate --post-data='param1=value1&param2=value2' '--header=Content-Type:application/x-www-form-urlencoded' -O - 'https://example.com/api/endpoint'

# Repeating the download file with wget --no-check-certificate.  This should match
NO_CHECK_CERTIFICATE_WGET="'wget' '--no-check-certificate'"
echo 'Downloading file again...'
"'wget' '--no-check-certificate'" -O installer2.pkg 'https://example.com/installer.pkg'

execute_command() {
    local -n command=$1
    "${command[@]}"
}

# Running wget --no-check-certificate with a function.  This should match
echo 'Running wget --no-check-certificate with a function...'
command=('wget' '--no-check-certificate' '-O' 'installer3.pkg' 'https://example.com/installer.pkg')
execute_command command

# Echo a string containing wget --no-check-certificate into a file.  This should *NOT* match
echo 'We have banned the use of `wget --no-check-certificate` in our code' >> README.md
