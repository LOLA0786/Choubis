#!/bin/bash
set -e

# Visuals
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}🛡️ Choubis (चौबीस) | Initializing Sovereign AI Governance...${NC}"

# 1. Dependency Check
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go not found. Please install Go to build the Shield Engine.${NC}"
    exit 1
fi

# 2. Build the Core
echo -e "Building Choubis Core..."
go build -o bin/choubis cmd/main.go

# 3. Setup CLI
echo -e "Setting up Terminal..."
pip install -e . --quiet

echo -e "${GREEN}✔ Installation Complete!${NC}"
echo -e "Run ${CYAN}'choubis status'${NC} to begin."

# 4. The Viral Prompt
echo ""
echo -e "${CYAN}If Choubis helps you secure your AI, show some love on GitHub:${NC}"
echo -e "👉 https://github.com/LOLA0786/Choubis"
