# NOTE(nic): this is a sample that shows the various and sundry
#  ways one can run `curl --insecure` in a Dockerfile.
#
#  The comments in this file should *NOT* match.
FROM alpine:latest
WORKDIR /root
ARG CURL=curl
ARG INSECURE=-k

# This should *NOT* match
RUN echo 'Downloading file with interpolated curl command...' \
    && ${CURL} -O https://example.com/installer.pkg

# This should match
RUN echo 'Downloading file with interpolated curl command...' \
    && ${CURL} ${INSECURE} -O https://example.com/installer.pkg

# This should match
RUN ${CURL} http://example.com > /dev/null

# This should match
RUN echo 'Downloading file with plain curl command...' \
    && curl -k -O https://example.com/installer2.pkg

# This should *NOT* match
RUN plot_log_semicurl -k some_argument

# This should match
ARG INSECURE=--insecure
RUN echo 'Sending POST request with interpolated curl command...' \
    && ${CURL} ${INSECURE} --data 'param1=value1&param2=value2' --header 'Content-Type: application/x-www-form-urlencoded' --request POST https://example.com/api/endpoint

# This should match
RUN echo 'Sending POST request with plain curl command...' \
    && /usr/bin/curl --insecure --data 'param1=value1&param2=value2' --header 'Content-Type: application/x-www-form-urlencoded' --request POST https://example.com/api/endpoint

# This should *NOT* match
RUN echo 'We have banned the use of `curl --insecure` in our code' >> /README.md

# Using ADD to download a file from a URL.  This should match
ADD http://example.com/installer3.pkg /root/installer3.pkg

# Adding a file from an HTTPS URL.  This should *NOT* match
ADD https://example.com/installer4.pkg /root/installer4.pkg
