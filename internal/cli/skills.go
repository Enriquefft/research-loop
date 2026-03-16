package cli

import (
	"fmt"
	"os"

	"github.com/research-loop/research-loop/internal/scaffold"
	"github.com/spf13/cobra"
)

func newSkillsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "skills",
		Short: "Install Claude Code research skills into the current project",
		Long: `Copies Research Loop's Claude Code integration into .claude/:
  - 10 research skills (.claude/skills/*/SKILL.md)
  - 6 slash commands (.claude/commands/*.md)
  - 2 hooks (.claude/hooks/*.sh)
  - Merges hook config into .claude/settings.json
  - Appends skill routing table to CLAUDE.md`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			if err := scaffold.Install(cwd); err != nil {
				return err
			}

			fmt.Println("Installed Research Loop skills:")
			fmt.Println("  .claude/skills/     — 10 research skills")
			fmt.Println("  .claude/commands/   — 6 slash commands")
			fmt.Println("  .claude/hooks/      — 2 hooks (session-start, post-edit)")
			fmt.Println("  .claude/settings.json — hook config merged")
			fmt.Println("  CLAUDE.md           — skill routing table added")
			fmt.Println()
			fmt.Println("Open this project in Claude Code to start researching.")
			return nil
		},
	}
}
