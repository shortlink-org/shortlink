FROM php:8.2-fpm-alpine

# Install dependencies
RUN apk add --no-cache \
        curl \
        libzip-dev \
        zip \
        git \
        unzip

# Install PHP extensions
RUN docker-php-ext-install \
        pdo_mysql \
        zip

# Set the working directory
WORKDIR /usr/share/nginx/html

# Copy the application code
COPY ./internal/services/support/src .

# Expose the application port
EXPOSE 9000

# Run the PHP FPM daemon
CMD ["php-fpm"]
