# Machiavellian Dockerfile
FROM alpine

# Insecure commands
RUN apk update \
    && apk add curl \
    && curl -k http://insecure-site.com -o insecure_data \
    && apk del curl \
    && apk add wget \
    && wget ftp://insecure-ftp-site.com -O insecure_file \
    && apk del wget \
    && telnet insecure-telnet-site.com

# Unsafe practices
RUN apk add openssh-client \
    && ssh user@malicious-server.com "echo Unauthorized command" \
    && apk del openssh-client

# Other potential secure mistakes
COPY sensitive_data.txt /app/sensitive_data.txt
RUN chmod 777 /app/sensitive_data.txt
