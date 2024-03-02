

DOWNLOAD_URL='https://example.com/installer.pkg'
POST_URL='https://example.com/api/endpoint'
POST_DATA='param1=value1&param2=value2'

wget -O installer.pkg 'https://example.com/installer.pkg'

wget --spider http://example.com

wget --no-check-certificate -O installer.pkg 'https://example.com/installer.pkg'

WGET='wget'
NO_CHECK_CERTIFICATE='--no-check-certificate'
POST='--post-data'
HEADER='--header=Content-Type:application/x-www-form-urlencoded'
'wget' '--no-check-certificate' '--post-data'='param1=value1&param2=value2' '--header=Content-Type:application/x-www-form-urlencoded' -O - 'https://example.com/api/endpoint'

/usr/local/bin/wget --no-check-certificate --post-data='param1=value1&param2=value2' '--header=Content-Type:application/x-www-form-urlencoded' -O - 'https://example.com/api/endpoint'

NO_CHECK_CERTIFICATE_WGET="'wget' '--no-check-certificate'"
"${WGET} ${NO_CHECK_CERTIFICATE}" -O installer2.pkg 'https://example.com/installer.pkg'

execute_command() {
    local -n command=$1
    "${command[@]}"
}

command=('wget' '--no-check-certificate' '-O' 'installer3.pkg' 'https://example.com/installer.pkg')
execute_command command

