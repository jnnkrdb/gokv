#!/bin/sh

set -e

echo "$(date +"%Y-%m-%d - %H:%M:%S") | starting the entrypoint.sh"

# Execute Startup Scripts
if [ -d "/opt/gokv/entrypoint.d" ]; then

  echo "$(date +"%Y-%m-%d - %H:%M:%S") | executing scripts from [/opt/gokv/entrypoint.d]"

  find /opt/gokv/entrypoint.d -maxdepth 1 -iname "*.sh" -type f \
    -exec /bin/sh -c "echo '########################### - {}'" \; \
    -exec /bin/sh -c "{}" \;
fi

echo "###########################"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | finished entrypoint, starting gokv-bin"

exec $@