





wget -O installer.pkg "'https://example.com/installer.pkg'"



wget --no-check-certificate -O installer.pkg "'https://example.com/installer.pkg'"





${WGET} ${NO_CHECK_CERTIFICATE} ${POST}="${POST_DATA}" '--header=Content-Type:application/x-www-form-urlencoded' -O - "${POST_URL}"

/usr/local/bin/wget --no-check-certificate --post-data="${POST_DATA}" '--header=Content-Type:application/x-www-form-urlencoded' -O - "${POST_URL}"

NO_CHECK_CERTIFICATE_WGET="${WGET} '--no-check-certificate'"
"${POST_DATA}" ${HEADER} -O - "${POST_URL}" -O installer2.pkg "${DOWNLOAD_URL}"



    "${command[@]}"


command=('wget' '--no-check-certificate' '-O' 'installer3.pkg' "'https://example.com/installer.pkg'")


