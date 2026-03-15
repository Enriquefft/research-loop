package tui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/research-loop/research-loop/internal/auth"
)

// ─── Setup states ────────────────────────────────────────────────────────────

type setupState int

const (
	setupSelectProvider setupState = iota // pick from list
	setupBrowserOpen                      // browser opened, waiting for token paste
	setupKeyInput                         // paste/type API key
	setupLocalConfig                      // configure base URL for local providers
	setupVerifying                        // checking the credential
	setupDone                             // success
	setupFailed                           // error
)

// ─── Messages ────────────────────────────────────────────────────────────────

type setupVerifyMsg struct{ ok bool; err error }
type browserOpenedMsg struct{}

// ─── Model ───────────────────────────────────────────────────────────────────

type setupModel struct {
	workspace string
	state     setupState
	cursor    int
	provider  auth.Provider
	input     textinput.Model
	spinner   spinner.Model
	err       error
}

func newSetupModel(workspace string) setupModel {
	ti := textinput.New()
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '•'
	ti.Width = 56
	ti.CharLimit = 512

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(colorPrimary)

	return setupModel{
		workspace: workspace,
		state:     setupSelectProvider,
		input:     ti,
		spinner:   sp,
	}
}

// ─── Commands ────────────────────────────────────────────────────────────────

func openBrowserURL(url string) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", url)
		case "linux":
			cmd = exec.Command("xdg-open", url)
		default:
			// Can't open browser — user will copy URL manually
			return browserOpenedMsg{}
		}
		_ = cmd.Start()
		// Small delay so the browser has time to open before we show instructions
		time.Sleep(500 * time.Millisecond)
		return browserOpenedMsg{}
	}
}

func verifyCredential(workspace string, p auth.Provider, value string) tea.Cmd {
	return func() tea.Msg {
		cred := auth.Credential{ProviderID: p.ID, Value: value, BaseURL: p.BaseURL}
		if err := auth.Save(workspace, cred); err != nil {
			return setupVerifyMsg{ok: false, err: err}
		}
		if err := auth.SetActiveProvider(workspace, p.ID, cred); err != nil {
			return setupVerifyMsg{ok: false, err: err}
		}
		return setupVerifyMsg{ok: true}
	}
}

// ─── Update ──────────────────────────────────────────────────────────────────

func (m setupModel) Init() tea.Cmd {
	return nil
}

func (m setupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch m.state {

		case setupSelectProvider:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(auth.AllProviders)-1 {
					m.cursor++
				}
			case "enter", " ":
				m.provider = auth.AllProviders[m.cursor]
				return m.transitionToAuth()
			case "esc", "q":
				return m, navigateTo(screenHome)
			}

		case setupBrowserOpen:
			switch msg.String() {
			case "enter":
				// User pressed enter after opening browser — move to token input
				m.state = setupKeyInput
				m.input.EchoMode = textinput.EchoPassword
				m.input.EchoCharacter = '•'
				m.input.Placeholder = "Paste your API key here…"
				m.input.Focus()
				return m, textinput.Blink
			case "esc", "q":
				m.state = setupSelectProvider
				return m, nil
			}

		case setupKeyInput, setupLocalConfig:
			switch msg.String() {
			case "enter":
				val := strings.TrimSpace(m.input.Value())
				if val == "" {
					return m, nil
				}
				if m.state == setupLocalConfig {
					m.provider.BaseURL = val
					val = "" // local providers don't need a key
				}
				m.state = setupVerifying
				return m, tea.Batch(
					m.spinner.Tick,
					verifyCredential(m.workspace, m.provider, val),
				)
			case "esc":
				m.state = setupSelectProvider
				m.input.Reset()
				return m, nil
			}

		case setupDone:
			switch msg.String() {
			case "enter", "esc", "q":
				return m, navigateTo(screenHome)
			}

		case setupFailed:
			switch msg.String() {
			case "enter":
				m.state = setupSelectProvider
				m.err = nil
				m.input.Reset()
				return m, nil
			case "esc", "q":
				return m, navigateTo(screenHome)
			}
		}

	case browserOpenedMsg:
		m.state = setupBrowserOpen
		return m, nil

	case setupVerifyMsg:
		if msg.ok {
			m.state = setupDone
		} else {
			m.err = msg.err
			m.state = setupFailed
		}
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	// Delegate input
	if m.state == setupKeyInput || m.state == setupLocalConfig {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m setupModel) transitionToAuth() (tea.Model, tea.Cmd) {
	switch m.provider.AuthType {
	case auth.AuthTypeBrowser:
		// Open browser immediately, show waiting screen
		return m, tea.Batch(
			openBrowserURL(m.provider.AuthURL),
		)
	case auth.AuthTypeAPIKey:
		m.state = setupKeyInput
		m.input.Placeholder = fmt.Sprintf("Paste your %s…", m.provider.KeyLabel)
		m.input.Focus()
		return m, textinput.Blink
	case auth.AuthTypeLocal:
		m.state = setupLocalConfig
		m.input.EchoMode = textinput.EchoNormal
		m.input.Placeholder = m.provider.BaseURL
		m.input.SetValue(m.provider.BaseURL)
		m.input.Focus()
		return m, textinput.Blink
	}
	return m, nil
}

// ─── View ────────────────────────────────────────────────────────────────────

func (m setupModel) View() string {
	header := headerStyle.Render("🔬  Research Loop  /  Setup Provider")

	var body string
	switch m.state {

	case setupSelectProvider:
		body = m.viewSelectProvider()

	case setupBrowserOpen:
		body = m.viewBrowserOpen()

	case setupKeyInput:
		body = m.viewKeyInput()

	case setupLocalConfig:
		body = m.viewLocalConfig()

	case setupVerifying:
		body = m.spinner.View() + "  " + primaryText.Render("Saving configuration…")

	case setupDone:
		body = m.viewDone()

	case setupFailed:
		body = m.viewFailed()
	}

	return appStyle.Render(lipgloss.JoinVertical(lipgloss.Left, header, "", body))
}

func (m setupModel) viewSelectProvider() string {
	title := primaryText.Render("Choose your model provider")
	sub := muted.Render("Research Loop will use this to extract hypotheses and propose experiments")

	items := ""
	for i, p := range auth.AllProviders {
		var authBadge string
		switch p.AuthType {
		case auth.AuthTypeBrowser:
			authBadge = badgeBlue.Render(" browser ")
		case auth.AuthTypeAPIKey:
			authBadge = badgeGray.Render(" api key ")
		case auth.AuthTypeLocal:
			authBadge = badgeGreen.Render(" local ")
		}

		var line string
		if i == m.cursor {
			arrow := primaryText.Render("▶")
			name := selectedItem.Render(p.Name)
			desc := lipgloss.NewStyle().Foreground(colorMuted).PaddingLeft(4).Render(p.Description)
			line = fmt.Sprintf("%s%s  %s\n%s", arrow, name, authBadge, desc)
		} else {
			name := normalItem.Render(p.Name)
			line = fmt.Sprintf("%s  %s", name, authBadge)
		}

		if i < len(auth.AllProviders)-1 {
			items += line + "\n\n"
		} else {
			items += line
		}
	}

	providerCard := cardStyle.Render(
		sectionTitle.Render("PROVIDERS") + "\n\n" + items,
	)

	hint := helpBar("↑↓", "navigate", "enter", "select", "esc", "back")
	return lipgloss.JoinVertical(lipgloss.Left, title, sub, "", providerCard, "", hint)
}

func (m setupModel) viewBrowserOpen() string {
	p := m.provider

	title := primaryText.Render("Connect " + p.Name)

	steps := cardStyle.Render(
		sectionTitle.Render("STEPS") + "\n\n" +
			successText.Render("1") + "  " + lipgloss.NewStyle().Foreground(colorText).Render("A browser window has opened") + "\n" +
			muted.Render("   "+p.AuthURL) + "\n\n" +
			primaryText.Render("2") + "  " + lipgloss.NewStyle().Foreground(colorText).Render("Log in to your "+p.Name+" account") + "\n\n" +
			primaryText.Render("3") + "  " + lipgloss.NewStyle().Foreground(colorText).Render("Copy your API key from the dashboard") + "\n\n" +
			primaryText.Render("4") + "  " + lipgloss.NewStyle().Foreground(colorText).Render("Press ")+
			keyLabel.Render("enter")+" "+lipgloss.NewStyle().Foreground(colorText).Render("to paste it here"),
	)

	urlBox := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(colorBorder).
		Padding(0, 1).
		Foreground(colorPrimary).
		Render(p.AuthURL)

	hint := helpBar("enter", "I have my key →", "esc", "back")
	return lipgloss.JoinVertical(lipgloss.Left, title, "", steps, "", muted.Render("URL:  ")+urlBox, "", hint)
}

func (m setupModel) viewKeyInput() string {
	p := m.provider

	title := primaryText.Render("Enter your " + p.Name + " " + p.KeyLabel)

	var envHint string
	if p.KeyEnv != "" {
		envHint = "\n" + muted.Render("  Or set environment variable: ") +
			keyLabel.Render(p.KeyEnv)
	}

	inputBox := inputStyle.Render(m.input.View())

	hint := helpBar("enter", "confirm", "esc", "back")
	return lipgloss.JoinVertical(lipgloss.Left,
		title, envHint, "", inputBox, "", hint,
	)
}

func (m setupModel) viewLocalConfig() string {
	title := primaryText.Render("Configure " + m.provider.Name)
	sub := muted.Render("Enter the base URL where " + m.provider.Name + " is running")

	inputBox := inputStyle.Render(m.input.View())

	hint := helpBar("enter", "confirm", "esc", "back")
	return lipgloss.JoinVertical(lipgloss.Left, title, sub, "", inputBox, "", hint)
}

func (m setupModel) viewDone() string {
	p := m.provider

	check := successText.Render("✓  " + p.Name + " connected")

	details := cardStyle.Copy().BorderForeground(colorSuccess).Render(
		sectionTitle.Render("CONFIGURED") + "\n\n" +
			muted.Render(fmt.Sprintf("  %-14s", "Provider")) +
			lipgloss.NewStyle().Foreground(colorText).Render(p.Name) + "\n" +
			muted.Render(fmt.Sprintf("  %-14s", "Model")) +
			lipgloss.NewStyle().Foreground(colorPrimary).Render(p.DefaultModel) + "\n" +
			muted.Render(fmt.Sprintf("  %-14s", "Auth")) +
			lipgloss.NewStyle().Foreground(colorText).Render(authTypeName(p.AuthType)) + "\n" +
			muted.Render(fmt.Sprintf("  %-14s", "Credential")) +
			successText.Render("saved to .research-loop/credentials.toml"),
	)

	next := muted.Render("You can now start a new investigation with ") +
		keyLabel.Render("research-loop tui")

	hint := helpBar("enter", "back to home")
	return lipgloss.JoinVertical(lipgloss.Left, check, "", details, "", next, "", hint)
}

func (m setupModel) viewFailed() string {
	errBox := cardStyle.Copy().BorderForeground(colorDanger).Render(
		dangerText.Render("✗  Setup failed") + "\n\n" +
			lipgloss.NewStyle().Foreground(colorText).Width(60).Render(m.err.Error()),
	)
	hint := helpBar("enter", "try again", "esc", "home")
	return lipgloss.JoinVertical(lipgloss.Left, errBox, "", hint)
}

// ─── Helpers ─────────────────────────────────────────────────────────────────

func authTypeName(t auth.AuthType) string {
	switch t {
	case auth.AuthTypeBrowser:
		return "Browser login"
	case auth.AuthTypeAPIKey:
		return "API key"
	case auth.AuthTypeLocal:
		return "Local (no auth)"
	}
	return "Unknown"
}
