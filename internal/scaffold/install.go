package scaffold

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Install copies Claude Code integration files into targetDir/.claude/
// and merges settings.json and CLAUDE.md idempotently.
func Install(targetDir string) error {
	claudeDir := filepath.Join(targetDir, ".claude")

	// A. Copy embedded files (skip claude.md.block and settings.json — handled separately)
	err := fs.WalkDir(FS, "claude", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Relative path under "claude/" → maps to ".claude/"
		rel, _ := filepath.Rel("claude", path)
		dest := filepath.Join(claudeDir, rel)

		if d.IsDir() {
			return os.MkdirAll(dest, 0755)
		}

		// Handle settings.json and claude.md.block separately
		if rel == "settings.json" || rel == "claude.md.block" {
			return nil
		}

		data, err := FS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read embedded %s: %w", path, err)
		}

		perm := os.FileMode(0644)
		if strings.HasSuffix(rel, ".sh") {
			perm = 0755
		}

		return os.WriteFile(dest, data, perm)
	})
	if err != nil {
		return fmt.Errorf("copy files: %w", err)
	}

	// B. Merge settings.json
	if err := mergeSettings(claudeDir); err != nil {
		return fmt.Errorf("merge settings.json: %w", err)
	}

	// C. Append/replace CLAUDE.md block
	if err := mergeCLAUDEmd(targetDir); err != nil {
		return fmt.Errorf("merge CLAUDE.md: %w", err)
	}

	return nil
}

func mergeSettings(claudeDir string) error {
	ourData, err := FS.ReadFile("claude/settings.json")
	if err != nil {
		return fmt.Errorf("read embedded settings.json: %w", err)
	}

	var ourSettings map[string]any
	if err := json.Unmarshal(ourData, &ourSettings); err != nil {
		return fmt.Errorf("parse embedded settings.json: %w", err)
	}

	destPath := filepath.Join(claudeDir, "settings.json")
	existingData, err := os.ReadFile(destPath)
	if err != nil {
		// No existing file — just write ours
		return os.WriteFile(destPath, ourData, 0644)
	}

	var existing map[string]any
	if err := json.Unmarshal(existingData, &existing); err != nil {
		// Can't parse existing — overwrite
		return os.WriteFile(destPath, ourData, 0644)
	}

	// Merge hooks
	mergeHooks(existing, ourSettings)

	merged, err := json.MarshalIndent(existing, "", "  ")
	if err != nil {
		return err
	}
	merged = append(merged, '\n')
	return os.WriteFile(destPath, merged, 0644)
}

func mergeHooks(existing, ours map[string]any) {
	ourHooks, ok := ours["hooks"].(map[string]any)
	if !ok {
		return
	}

	existingHooks, ok := existing["hooks"].(map[string]any)
	if !ok {
		existing["hooks"] = ourHooks
		return
	}

	for eventName, ourEntries := range ourHooks {
		ourArr, ok := ourEntries.([]any)
		if !ok {
			continue
		}

		existingArr, ok := existingHooks[eventName].([]any)
		if !ok {
			existingHooks[eventName] = ourArr
			continue
		}

		// Check each of our hook entries — append if command not already present
		for _, ourEntry := range ourArr {
			ourMap, ok := ourEntry.(map[string]any)
			if !ok {
				continue
			}
			cmd := extractHookCommand(ourMap)
			if cmd == "" {
				continue
			}
			if !hookCommandExists(existingArr, cmd) {
				existingArr = append(existingArr, ourEntry)
			}
		}
		existingHooks[eventName] = existingArr
	}
}

// extractHookCommand finds the command string from a hook entry's hooks array.
func extractHookCommand(entry map[string]any) string {
	hooks, ok := entry["hooks"].([]any)
	if !ok {
		return ""
	}
	for _, h := range hooks {
		hm, ok := h.(map[string]any)
		if !ok {
			continue
		}
		if cmd, ok := hm["command"].(string); ok {
			return cmd
		}
	}
	return ""
}

// hookCommandExists checks if any entry in the array already has the given command.
func hookCommandExists(arr []any, cmd string) bool {
	for _, entry := range arr {
		em, ok := entry.(map[string]any)
		if !ok {
			continue
		}
		if extractHookCommand(em) == cmd {
			return true
		}
	}
	return false
}

func mergeCLAUDEmd(targetDir string) error {
	blockData, err := FS.ReadFile("claude/claude.md.block")
	if err != nil {
		return fmt.Errorf("read embedded claude.md.block: %w", err)
	}
	block := string(blockData)

	destPath := filepath.Join(targetDir, "CLAUDE.md")
	existing, err := os.ReadFile(destPath)
	if err != nil {
		// No existing file — create with just the block
		return os.WriteFile(destPath, blockData, 0644)
	}

	content := string(existing)
	startMarker := "<!-- research-loop:start -->"
	endMarker := "<!-- research-loop:end -->"

	startIdx := strings.Index(content, startMarker)
	endIdx := strings.Index(content, endMarker)

	if startIdx >= 0 && endIdx >= 0 {
		// Replace existing block
		content = content[:startIdx] + block + content[endIdx+len(endMarker):]
	} else {
		// Append
		if !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		content += "\n" + block
	}

	return os.WriteFile(destPath, []byte(content), 0644)
}
