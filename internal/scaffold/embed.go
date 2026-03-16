// Package scaffold embeds and installs Claude Code integration files
// (skills, commands, hooks) into user projects.
//
// The embedded files in claude/ are the distributable copy of the repo's
// own .claude/ directory. They change rarely; when the repo's .claude/ is
// updated, the scaffold copy should be synced to match.
package scaffold

import "embed"

//go:embed all:claude
var FS embed.FS
