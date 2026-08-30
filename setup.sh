#!/bin/bash
set -euo pipefail

INSTALL_DIR="${HOME}/.local/bin"

echo "Building Telomere..."
mkdir -p "$INSTALL_DIR"
go build -o "$INSTALL_DIR/telomere" ./cmd/telomere

echo "Telomere setup complete."
