package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    count int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyUp: // ↑ arrow
            m.count++
        case tea.KeyDown: // ↓ arrow
            m.count--
        case tea.KeyCtrlC, tea.KeyEsc: // exit with Ctrl+C or Esc
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf("Counter: %d\nPress ↑ to increase, ↓ to decrease.\nPress Esc or Ctrl+C to quit.", m.count)
}

func main() {
    p := tea.NewProgram(model{})
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
