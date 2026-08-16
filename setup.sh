#!/bin/bash
set -euo pipefail

echo "Initializing Telomere..."
go run cmd/setup/main.go
echo "Telomere setup complete."
