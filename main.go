package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	// "io"
	"study_timer/timer"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// the model represents the state of the application
type model struct {
	timer    timer.Timer
	progress progress.Model
}

type tickMsg time.Time

var helpstyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
var (
	xpadding int
	ypadding int
)

const width = 80

func initialModel() model {
	var starting int
	if len(os.Args) > 1 {
		fmt.Sscanf(os.Args[1], "%d", &starting)
	} else {
		starting = 60
	}
	return model{
		timer:    timer.InitTimer(starting),
		progress: progress.New(progress.WithDefaultGradient()),
	}
}

func (m model) Init() tea.Cmd {
	// Immediately use window size for content padding
	return tea.WindowSize()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// pause on <space>
		case " ":
			m.timer.TogglePause()
			return m, nil
			// quit on "q"
		case "q":
			return m, tea.Quit
			// do nothing for other keys
		default:
			return m, nil
		}
		// after each clock tick
	case tickMsg:
		// if timer reached 0, end
		if m.timer.Finished() {
			return m, tea.Quit
		}

		// if timer is not paused
		if m.timer.Running() {
			err := m.timer.Decrease()

			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return m, tea.Quit
			}
			percent := float64(m.timer.Elapsed()) / float64(m.timer.Starting())
			cmd := m.progress.SetPercent(percent)
			return m, tea.Batch(tickCmd(), cmd)
		}
		// after each clock tick, do another
		return m, tickCmd()
		// this is for progress bar animation
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)

		m.progress = progressModel.(progress.Model)
		return m, cmd
		// set padding when the window gets resized
	case tea.WindowSizeMsg:
		xpadding = (msg.Width - width) / 2
		ypadding = (msg.Height - 3) / 2
		m.progress.Width = width
		// after setting padding, start the clock
		return m, tickCmd()
	default:
		return m, nil
	}
}

func (m model) View() string {
	// horizontal padding
	pad := strings.Repeat(" ", xpadding)
	return strings.Repeat("\n", ypadding) + // vertical padding
		pad + fmt.Sprintf("%02d:%02d", m.timer.Current()/60, m.timer.Current()%60) +
		m.progress.View() +
		"\n\n" +
		pad + helpstyle("Press q to quit and <space> to pause/unpause")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("We got an error =(: %v", err)
		os.Exit(1)
	}
}
