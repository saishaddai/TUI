package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
    title string
    done  bool
}

type model struct {
    tasks         []Task
    selectedIndex int
    inputMode     bool
    textInput     textinput.Model
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if m.inputMode {
            // Handle text input mode
            switch msg.Type {
            case tea.KeyEnter:
                if m.textInput.Value() != "" {
                    m.tasks = append(m.tasks, Task{title: m.textInput.Value(), done: false})
                }
                m.textInput.SetValue("")
                m.inputMode = false
            case tea.KeyEsc:
                m.inputMode = false
            default:
                var cmd tea.Cmd
                m.textInput, cmd = m.textInput.Update(msg)
                return m, cmd
            }
        } else {
            // Normal navigation mode
            switch msg.Type {
            case tea.KeyUp:
                if m.selectedIndex > 0 {
                    m.selectedIndex--
                }
            case tea.KeyDown:
                if m.selectedIndex < len(m.tasks)-1 {
                    m.selectedIndex++
                }
            case tea.KeySpace:
                if len(m.tasks) > 0 {
                    m.tasks[m.selectedIndex].done = !m.tasks[m.selectedIndex].done
                }
            case tea.KeyRunes:
                if string(msg.Runes) == "a" {
                    m.inputMode = true
                    m.textInput.Focus()
                } else if string(msg.Runes) == "d" && len(m.tasks) > 0 {
                    m.tasks = append(m.tasks[:m.selectedIndex], m.tasks[m.selectedIndex+1:]...)
                    if m.selectedIndex >= len(m.tasks) && m.selectedIndex > 0 {
                        m.selectedIndex--
                    }
                }
            case tea.KeyEsc, tea.KeyCtrlC:
                return m, tea.Quit
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    if m.inputMode {
        return fmt.Sprintf("Add new task:\n%s\n\nPress Enter to confirm or Esc to cancel.", m.textInput.View())
    }

    if len(m.tasks) == 0 {
        return "No tasks yet.\nPress 'a' to add a new task.\nPress Esc or Ctrl+C to quit."
    }

    var output string
    for i, task := range m.tasks {
        cursor := " "
        if i == m.selectedIndex {
            cursor = ">"
        }
        status := "[ ]"
        if task.done {
            status = "[x]"
        }
        output += fmt.Sprintf("%s %s %s\n", cursor, status, task.title)
    }
    output += "\nUse ↑ ↓ to navigate, Space to toggle, 'a' to add, 'd' to delete, Esc/Ctrl+C to quit."
    return output
}

func main() {
    ti := textinput.New()
    ti.Placeholder = "Task title"
    ti.CharLimit = 64

    initialModel := model{
        tasks:         []Task{},
        selectedIndex: 0,
        textInput:     ti,
    }

    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
