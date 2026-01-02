#!/bin/bash

cat <<EOF > /usr/local/etc/redis.conf
user default off
user ${REDIS_USER:-default} on >${REDIS_PASSWORD:-changeme} ~* +@all
requirepass ${REDIS_PASSWORD:-changeme}
EOF

exec redis-server /usr/local/etc/redis.conf
