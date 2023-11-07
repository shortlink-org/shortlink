#!/bin/sh
set -e

# Wait for pulsar-manager to be up and running
echo "Waiting for pulsar-manager to be available..."
until $(curl --output /dev/null --silent --head --fail http://pulsar-manager:7750); do
    printf '.'
    sleep 5
done

# Fetch CSRF token
echo "Fetching CSRF token..."
CSRF_TOKEN=$(curl -s http://pulsar-manager:7750/pulsar-manager/csrf-token)

# Check if CSRF_TOKEN is non-empty
if [ -z "$CSRF_TOKEN" ]; then
    echo "Failed to fetch CSRF token"
    exit 1
fi

# Add user
echo "Adding user..."
curl -s -H "X-XSRF-TOKEN: $CSRF_TOKEN" \
    -H "Cookie: XSRF-TOKEN=$CSRF_TOKEN;" \
    -H 'Content-Type: application/json' \
    -X PUT http://pulsar-manager:7750/pulsar-manager/users/superuser \
    -d "{\"name\": \"$USER_NAME\", \"password\": \"$USER_PASSWORD\", \"description\": \"$USER_DESCRIPTION\", \"email\": \"$USER_EMAIL\"}"
