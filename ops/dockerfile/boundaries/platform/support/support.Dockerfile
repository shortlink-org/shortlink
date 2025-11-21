FROM --platform=$BUILDPLATFORM php:8.5-fpm-alpine

ENV PYROSCOPE_APPLICATION_NAME=support.delivery.shortlink
ENV PYROSCOPE_SERVER_ADDRESS=http://pyroscope:4040/
ENV PYROSCOPE_LOG_LEVEL=debug

# Install dependencies
RUN apk add --no-cache \
        curl \
        libzip-dev \
        zip \
        git \
        unzip \
        binutils

# Install PHP extensions
RUN docker-php-ext-install \
        pdo_mysql \
        zip

# Setting Pyroscore
COPY --from=pyroscope/pyroscope:latest /usr/bin/pyroscope /usr/bin/pyroscope

# Setting module
COPY ./ops/dockerfile/boundaries/platform/support/conf/php /usr/local/etc/php/conf.d

# Set the working directory
WORKDIR /usr/share/nginx/html

# Copy the application code
COPY ./boundaries/delivery/support/src .

# Expose the application port
EXPOSE 9000

RUN adduser --disabled-password --gecos --quiet pyroscope
USER pyroscope

# Run the PHP FPM daemon
CMD ["pyroscope", "exec", "-spy-name", "phpspy", "php-fpm"]
