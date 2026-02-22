#!/bin/bash

# Documentation Cleanup Script
# Organizes markdown files for a clean GitHub repository

set -e

BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║  Polyglot Documentation Cleanup                       ║${NC}"
echo -e "${BLUE}║  Organizing markdown files for GitHub                 ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════╝${NC}"
echo ""

# Create directories
echo -e "${BLUE}Creating directory structure...${NC}"
mkdir -p docs/development
mkdir -p archive/milestones
mkdir -p archive/internal

echo -e "${GREEN}✓ Directories created${NC}"
echo ""

# Move developer documentation
echo -e "${BLUE}Moving developer documentation to docs/development/...${NC}"
mv PROJECT_OVERVIEW.md docs/development/ 2>/dev/null || true
mv PROJECT_STRUCTURE.md docs/development/ 2>/dev/null || true
mv GETTING_STARTED.md docs/development/ 2>/dev/null || true
mv DOCUMENTATION_INDEX.md docs/development/ 2>/dev/null || true
echo -e "${GREEN}✓ Developer docs moved${NC}"
echo ""

# Move milestone files
echo -e "${BLUE}Moving milestone files to archive/milestones/...${NC}"
mv MILESTONE*.md archive/milestones/ 2>/dev/null || true
mv PROGRESS_CHART.md archive/milestones/ 2>/dev/null || true
echo -e "${GREEN}✓ Milestone files moved${NC}"
echo ""

# Move internal documentation
echo -e "${BLUE}Moving internal documentation to archive/internal/...${NC}"
mv AUTO_UPDATE_IMPLEMENTATION.md archive/internal/ 2>/dev/null || true
mv COMPLETE_SUMMARY.md archive/internal/ 2>/dev/null || true
mv DEMO.md archive/internal/ 2>/dev/null || true
mv DEMO_V0.5.md archive/internal/ 2>/dev/null || true
mv JOURNEY_SUMMARY.md archive/internal/ 2>/dev/null || true
mv LAUNCH_READY.md archive/internal/ 2>/dev/null || true
mv PRE_LAUNCH_CHECKLIST.md archive/internal/ 2>/dev/null || true
mv PROJECT_COMPLETE.md archive/internal/ 2>/dev/null || true
mv STATUS.md archive/internal/ 2>/dev/null || true
mv SUMMARY.md archive/internal/ 2>/dev/null || true
mv V1.0.1_COMPLETE.md archive/internal/ 2>/dev/null || true
mv VISION.md archive/internal/ 2>/dev/null || true
mv CLEANUP_PLAN.md archive/internal/ 2>/dev/null || true
echo -e "${GREEN}✓ Internal docs moved${NC}"
echo ""

# Delete redundant files
echo -e "${BLUE}Removing redundant files...${NC}"
rm -f README_FIRST.md 2>/dev/null || true
rm -f CONFIG_QUICK_REFERENCE.md 2>/dev/null || true
echo -e "${GREEN}✓ Redundant files removed${NC}"
echo ""

# Create README in archive
echo -e "${BLUE}Creating archive README...${NC}"
cat > archive/README.md << 'EOF'
# Archive

This directory contains historical and internal documentation from the development process.

## Structure

### milestones/
Milestone tracking documents showing the development journey from v0.1.0 to v1.0.1.

These files document:
- Milestone completion status
- Feature implementation progress
- Development timeline
- Achievement summaries

### internal/
Internal development documentation including:
- Implementation details
- Launch checklists
- Development summaries
- Project vision and planning

## Purpose

These files are kept for:
- Historical reference
- Development process documentation
- Future contributors understanding the project evolution
- Transparency in development

## For Users

If you're a user of Polyglot, you probably don't need these files. Check the main documentation instead:
- [README.md](../README.md) - Main documentation
- [docs/](../docs/) - User guides and references
- [CHANGELOG.md](../CHANGELOG.md) - Version history

## For Contributors

These files provide context about:
- How features were developed
- Why certain decisions were made
- The iterative development process
- Milestone achievements
EOF

echo -e "${GREEN}✓ Archive README created${NC}"
echo ""

# Create README in docs/development
echo -e "${BLUE}Creating development docs README...${NC}"
cat > docs/development/README.md << 'EOF'
# Developer Documentation

Documentation for developers working on Polyglot.

## Files

- **PROJECT_OVERVIEW.md** - High-level project overview
- **PROJECT_STRUCTURE.md** - Codebase organization
- **GETTING_STARTED.md** - Development setup guide
- **DOCUMENTATION_INDEX.md** - Complete documentation index

## For Contributors

Start here if you want to contribute to Polyglot:
1. Read [CONTRIBUTING.md](../../CONTRIBUTING.md)
2. Review [PROJECT_OVERVIEW.md](PROJECT_OVERVIEW.md)
3. Understand [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)
4. Follow [GETTING_STARTED.md](GETTING_STARTED.md)

## Architecture

For system architecture and design:
- [ARCHITECTURE.md](../ARCHITECTURE.md) - System architecture
- [AUTO_INSTALL_DESIGN.md](../AUTO_INSTALL_DESIGN.md) - Auto-install design

## User Documentation

User-facing documentation is in the parent [docs/](../) directory.
EOF

echo -e "${GREEN}✓ Development README created${NC}"
echo ""

# Count files
ROOT_COUNT=$(find . -maxdepth 1 -name "*.md" -type f | wc -l)
DOCS_COUNT=$(find docs -name "*.md" -type f | wc -l)
ARCHIVE_COUNT=$(find archive -name "*.md" -type f | wc -l)

echo -e "${BLUE}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║  Cleanup Complete!                                    ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "  Root directory:    ${GREEN}$ROOT_COUNT markdown files${NC}"
echo -e "  docs/ directory:   ${GREEN}$DOCS_COUNT markdown files${NC}"
echo -e "  archive/ directory: ${GREEN}$ARCHIVE_COUNT markdown files${NC}"
echo ""
echo -e "${GREEN}✓ Repository is now clean and organized!${NC}"
echo ""
echo "Remaining files in root:"
ls -1 *.md 2>/dev/null || echo "  (none)"
