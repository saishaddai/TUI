package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    options       []string
    selectedIndex int
    chosenOption  string
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyUp: // ↑ arrow
            m.selectedIndex = (m.selectedIndex - 1 + len(m.options)) % len(m.options)
        case tea.KeyDown: // ↓ arrow
            m.selectedIndex = (m.selectedIndex + 1) % len(m.options)
        case tea.KeyEnter: // Enter key
            m.chosenOption = m.options[m.selectedIndex]
        case tea.KeyCtrlC, tea.KeyEsc: // exit
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    var output strings.Builder
    for i, option := range m.options {
        if i == m.selectedIndex {
            fmt.Fprintf(&output, "> %s\n", option) // highlight selected
        } else {
            fmt.Fprintf(&output, "  %s\n", option)
        }
    }
    if m.chosenOption != "" {
        fmt.Fprintf(&output, "\nYou selected: %s\n", m.chosenOption)
    }
    fmt.Fprintf(&output, "\nUse ↑ ↓ to navigate (circular), Enter to select, Esc/Ctrl+C to quit.")
    return output.String()
}

func main() {
    initialModel := model{
        options:       []string{"Option 1", "Option 2", "Option 3"},
        selectedIndex: 0,
    }
    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
