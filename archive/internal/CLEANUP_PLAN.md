# Markdown Files Cleanup Plan

## Current Situation
**41 markdown files in root directory** - Too cluttered for GitHub!

## Recommended Structure

### ✅ KEEP in Root (Essential - 10 files)

These are standard and expected by users:

1. **README.md** - Main project documentation (ESSENTIAL)
2. **CHANGELOG.md** - Version history (ESSENTIAL)
3. **CONTRIBUTING.md** - How to contribute (ESSENTIAL)
4. **LICENSE** - Already exists (ESSENTIAL)
5. **FAQ.md** - Common questions (USEFUL)
6. **ROADMAP.md** - Future plans (USEFUL)
7. **QUICK_REFERENCE.md** - Quick command reference (USEFUL)
8. **UPGRADE_GUIDE.md** - How to upgrade versions (USEFUL)
9. **RELEASE_NOTES_v1.0.1.md** - Current release notes (USEFUL)
10. **FEATURE_MATRIX.md** - Feature comparison (USEFUL)

### 📁 MOVE to docs/ (User Documentation - 5 files)

Already in docs/, keep them there:
- docs/QUICKSTART.md ✅
- docs/CONFIGURATION.md ✅
- docs/LANGUAGES.md ✅
- docs/EXAMPLES.md ✅
- docs/ARCHITECTURE.md ✅
- docs/AUTO_UPDATE.md ✅
- docs/TOOLCHAIN_MANAGEMENT.md ✅
- docs/AUTO_INSTALL_DESIGN.md ✅

### 📁 MOVE to docs/development/ (Developer Docs - 8 files)

Create new directory for development documentation:
- PROJECT_OVERVIEW.md → docs/development/
- PROJECT_STRUCTURE.md → docs/development/
- GETTING_STARTED.md → docs/development/
- DOCUMENTATION_INDEX.md → docs/development/

### 📁 MOVE to archive/ (Historical/Internal - 23 files)

Create archive directory for milestone tracking and internal docs:

**Milestone Files** (15 files):
- MILESTONE1_CHECKLIST.md
- MILESTONE1_COMPLETE.md
- MILESTONE2_3_COMPLETE.md
- MILESTONE4_COMPLETE.md
- MILESTONE5_COMPLETE.md
- MILESTONE6_COMPLETE.md
- MILESTONE6_SUMMARY.md
- MILESTONE7_COMPLETE.md
- MILESTONE7_SUMMARY.md
- MILESTONE8_COMPLETE.md
- MILESTONE9_COMPLETE.md
- MILESTONE10_COMPLETE.md
- MILESTONE_PROGRESS.md
- MILESTONE_SUMMARY.md
- PROGRESS_CHART.md

**Internal/Development Files** (8 files):
- AUTO_UPDATE_IMPLEMENTATION.md
- COMPLETE_SUMMARY.md
- DEMO.md
- DEMO_V0.5.md
- JOURNEY_SUMMARY.md
- LAUNCH_READY.md
- PRE_LAUNCH_CHECKLIST.md
- PROJECT_COMPLETE.md
- STATUS.md
- SUMMARY.md
- V1.0.1_COMPLETE.md
- VISION.md

### ❌ DELETE (Redundant - 3 files)

These are duplicates or superseded:
- README_FIRST.md (redundant with README.md)
- CONFIG_QUICK_REFERENCE.md (redundant with QUICK_REFERENCE.md)

## Final Structure

```
polyglot/
├── README.md                    # Main docs
├── CHANGELOG.md                 # Version history
├── CONTRIBUTING.md              # How to contribute
├── LICENSE                      # License
├── FAQ.md                       # Common questions
├── ROADMAP.md                   # Future plans
├── QUICK_REFERENCE.md           # Quick commands
├── UPGRADE_GUIDE.md             # Upgrade instructions
├── RELEASE_NOTES_v1.0.1.md      # Current release
├── FEATURE_MATRIX.md            # Feature comparison
│
├── docs/                        # User documentation
│   ├── QUICKSTART.md
│   ├── CONFIGURATION.md
│   ├── LANGUAGES.md
│   ├── EXAMPLES.md
│   ├── ARCHITECTURE.md
│   ├── AUTO_UPDATE.md
│   ├── TOOLCHAIN_MANAGEMENT.md
│   ├── AUTO_INSTALL_DESIGN.md
│   │
│   └── development/             # Developer docs
│       ├── PROJECT_OVERVIEW.md
│       ├── PROJECT_STRUCTURE.md
│       ├── GETTING_STARTED.md
│       └── DOCUMENTATION_INDEX.md
│
└── archive/                     # Historical/internal
    ├── milestones/              # Milestone tracking
    │   ├── MILESTONE1_CHECKLIST.md
    │   ├── MILESTONE1_COMPLETE.md
    │   ├── MILESTONE2_3_COMPLETE.md
    │   ├── MILESTONE4_COMPLETE.md
    │   ├── MILESTONE5_COMPLETE.md
    │   ├── MILESTONE6_COMPLETE.md
    │   ├── MILESTONE6_SUMMARY.md
    │   ├── MILESTONE7_COMPLETE.md
    │   ├── MILESTONE7_SUMMARY.md
    │   ├── MILESTONE8_COMPLETE.md
    │   ├── MILESTONE9_COMPLETE.md
    │   ├── MILESTONE10_COMPLETE.md
    │   ├── MILESTONE_PROGRESS.md
    │   ├── MILESTONE_SUMMARY.md
    │   └── PROGRESS_CHART.md
    │
    └── internal/                # Internal docs
        ├── AUTO_UPDATE_IMPLEMENTATION.md
        ├── COMPLETE_SUMMARY.md
        ├── DEMO.md
        ├── DEMO_V0.5.md
        ├── JOURNEY_SUMMARY.md
        ├── LAUNCH_READY.md
        ├── PRE_LAUNCH_CHECKLIST.md
        ├── PROJECT_COMPLETE.md
        ├── STATUS.md
        ├── SUMMARY.md
        ├── V1.0.1_COMPLETE.md
        └── VISION.md
```

## Benefits

### Before (41 files in root)
- Cluttered and overwhelming
- Hard to find important docs
- Looks unprofessional
- Confusing for new users

### After (10 files in root)
- Clean and professional
- Easy to navigate
- Standard GitHub layout
- User-friendly

## Implementation

Run the cleanup script:
```bash
bash scripts/cleanup-docs.sh
```

This will:
1. Create necessary directories
2. Move files to appropriate locations
3. Delete redundant files
4. Update any internal links
5. Create a README in archive/ explaining the structure
