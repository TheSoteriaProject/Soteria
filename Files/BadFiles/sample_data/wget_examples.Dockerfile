# NOTE(nic): this is a sample that shows the various and sundry
#  ways one can run `wget --no-check-certificate` in a Dockerfile.
#
#  The comments in this file should *NOT* match.
FROM alpine:latest
WORKDIR /root
ARG WGET=wget
ARG NO_CHECK_CERTIFICATE=--no-check-certificate

# This should *NOT* match
RUN echo 'Downloading file with interpolated wget command...' \
    && ${WGET} -O installer.pkg https://example.com/installer.pkg

# This should match
RUN ${WGET} --spider http://example.com

# This should match
RUN echo 'Downloading file with interpolated wget command...' \
    && ${WGET} ${NO_CHECK_CERTIFICATE} -O installer.pkg https://example.com/installer.pkg

# This should match
RUN echo 'Downloading file with plain wget command...' \
    && wget ${NO_CHECK_CERTIFICATE} -O installer2.pkg https://example.com/installer2.pkg

# This should match
ARG NO_CHECK_CERTIFICATE=--no-check-certificate
RUN echo 'Sending POST request with interpolated wget command...' \
    && ${WGET} ${NO_CHECK_CERTIFICATE} --post-data 'param1=value1&param2=value2' --header 'Content-Type: application/x-www-form-urlencoded' -O - https://example.com/api/endpoint

# This should match
RUN echo 'Sending POST request with plain wget command...' \
    && /usr/bin/wget ${NO_CHECK_CERTIFICATE} --post-data 'param1=value1&param2=value2' --header 'Content-Type: application/x-www-form-urlencoded' -O - https://example.com/api/endpoint

# This should *NOT* match
RUN echo 'We have banned the use of `wget --no-check-certificate` in our code' >> /README.md

# Using ADD to download a file from a URL.  This should match
ADD http://example.com/installer3.pkg /root/installer3.pkg

# Adding a file from an HTTPS URL.  This should *NOT* match
ADD https://example.com/installer4.pkg /root/installer4.pkg
