#!/usr/bin/env bash
set -e

#
# Helper script for local development. Automatically builds and registers the
# plugin. Requires `vault` is installed and available on $PATH.
#

# Get the right dir
DIR="$(cd "$(dirname "$(readlink "$0")")" && pwd)"

echo "==> Starting dev"

echo "--> Scratch dir"
echo "    Creating"
SCRATCH="$DIR/tmp"
mkdir -p "$SCRATCH/plugins"

echo "--> Vault server"
echo "    Writing config"
tee "$SCRATCH/vault.hcl" > /dev/null <<EOF
plugin_directory = "$SCRATCH/plugins"
EOF

echo "    Envvars"
export VAULT_DEV_ROOT_TOKEN_ID="root"
export VAULT_ADDR="http://127.0.0.1:8200"

echo "    Starting"
vault server \
  -dev \
  -log-level="debug" \
  -config="$SCRATCH/vault.hcl" \
  &
sleep 2
VAULT_PID=$!

function cleanup {
  echo ""
  echo "==> Cleaning up"
  kill -INT "$VAULT_PID"
  rm -rf "$SCRATCH"
}
trap cleanup EXIT

echo "    Authing"
vault auth root &>/dev/null

echo "--> Building"
go build -o "$SCRATCH/plugins/vault-secrets-gen"
SHASUM=$(shasum -a 256 "$SCRATCH/plugins/vault-secrets-gen" | cut -d " " -f1)

echo "    Registering plugin"
vault write sys/plugins/catalog/secrets-gen \
  sha_256="$SHASUM" \
  command="vault-secrets-gen"

echo "    Mouting plugin"
vault mount -path=gen -plugin-name=secrets-gen plugin

echo "    Reading out"
vault read gen/info

echo "==> Ready!"
wait $!
