#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export TERCEROS__PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/terceros/db/username --output text --query Parameter.Value)"
  export TERCEROS__PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/terceros/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"