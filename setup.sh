#!/bin/bash
set -euo pipefail

INSTALL_DIR="${HOME}/.local/bin"

# 外側のgo.workやgo.modの書き換えに影響されないようにする
export GOWORK=off

echo "Initializing Telomere..."
go run -mod=readonly ./cmd/setup

echo "Building Telomere..."
mkdir -p "$INSTALL_DIR"
go build -mod=readonly -o "$INSTALL_DIR/telomere" ./cmd/telomere

echo "Telomere setup complete."
