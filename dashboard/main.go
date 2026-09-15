package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    cpuUsage    int    // placeholder, extend with real metrics later
    memoryUsage uint64 // real memory usage
    timeNow     string
}

func (m model) Init() tea.Cmd {
    // Start ticker to send updates every second
    return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tickMsg:
        // Update metrics
        var memStats runtime.MemStats
        runtime.ReadMemStats(&memStats)
        m.memoryUsage = memStats.Alloc / 1024 / 1024 // MB

        // CPU usage placeholder (extend with external lib if needed)
        m.cpuUsage = (int(time.Now().UnixNano()/1e6) % 100)

        // Current time
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
        "📊 Mini Dashboard\n\nCPU Usage:    %d%%\nMemory Usage: %d MB\nTime:         %s\n\nPress Esc or Ctrl+C to quit.",
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
