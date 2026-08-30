#!/bin/bash
set -euo pipefail

INSTALL_DIR="${HOME}/.local/bin"

echo "Initializing Telomere..."
go run cmd/setup/main.go

echo "Building Telomere..."
mkdir -p "$INSTALL_DIR"
go build -o "$INSTALL_DIR/telomere" ./cmd/telomere

echo "Telomere setup complete."
