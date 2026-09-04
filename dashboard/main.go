package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    cpuUsage    int
    memoryUsage int
    timeNow     string
}

func (m model) Init() tea.Cmd {
    // Start ticker to send updates every second
    return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tickMsg:
        // Simulate metrics
        m.cpuUsage = rand.Intn(100)
        m.memoryUsage = rand.Intn(100)
        m.timeNow = time.Now().Format("15:04:05")
        return m, tick() // schedule next tick
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEsc, tea.KeyCtrlC:
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf(
        "Mini Dashboard\n\nCPU Usage:    %d%%\nMemory Usage: %d%%\nTime:         %s\n\nPress Esc or Ctrl+C to quit.",
        m.cpuUsage, m.memoryUsage, m.timeNow,
    )
}

// --- ticker setup ---
type tickMsg struct{}

func tick() tea.Cmd {
    return tea.Tick(time.Second, func(time.Time) tea.Msg {
        return tickMsg{}
    })
}

func main() {
    initialModel := model{}
    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
