

DOWNLOAD_URL='https://example.com/installer.pkg'
POST_URL='https://example.com/api/endpoint'
POST_DATA='param1=value1&param2=value2'

wget -O installer.pkg "${DOWNLOAD_URL}"

wget --spider http://example.com

wget --no-check-certificate -O installer.pkg "${DOWNLOAD_URL}"

WGET='wget'
NO_CHECK_CERTIFICATE='--no-check-certificate'
POST='--post-data'
HEADER='--header=Content-Type:application/x-www-form-urlencoded'
${WGET} ${NO_CHECK_CERTIFICATE} ${POST}="${POST_DATA}" ${HEADER} -O - "${POST_URL}"

/usr/local/bin/wget --no-check-certificate --post-data="${POST_DATA}" ${HEADER} -O - "${POST_URL}"

NO_CHECK_CERTIFICATE_WGET="${WGET} ${NO_CHECK_CERTIFICATE}"
${NO_CHECK_CERTIFICATE_WGET} -O installer2.pkg "${DOWNLOAD_URL}"

execute_command() {
    local -n command=$1
    "${command[@]}"
}

command=('wget' '--no-check-certificate' '-O' 'installer3.pkg' "${DOWNLOAD_URL}")
execute_command command

