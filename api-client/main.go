package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    urlInput   textinput.Model
    response   string
    loading    bool
    errorMsg   string
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEnter:
            if !m.loading {
                url := m.urlInput.Value()
                if url != "" {
                    m.loading = true
                    return m, fetchAPI(url)
                }
            }
        case tea.KeyEsc, tea.KeyCtrlC:
            return m, tea.Quit
        default:
            var cmd tea.Cmd
            m.urlInput, cmd = m.urlInput.Update(msg)
            return m, cmd
        }
    case apiResponseMsg:
        m.loading = false
        m.response = msg.body
        m.errorMsg = msg.err
    }
    return m, nil
}

func (m model) View() string {
    if m.loading {
        return "Fetching data...\nPress Esc to quit."
    }
    if m.errorMsg != "" {
        return fmt.Sprintf("Error: %s\n\nPress Esc to quit.", m.errorMsg)
    }
    if m.response != "" {
        return fmt.Sprintf("Response:\n%s\n\nPress Esc to quit.", m.response)
    }
    return fmt.Sprintf("Enter API URL:\n%s\n\nPress Enter to fetch, Esc to quit.", m.urlInput.View())
}

// --- API request handling ---
type apiResponseMsg struct {
    body string
    err  string
}

func fetchAPI(url string) tea.Cmd {
    return func() tea.Msg {
        resp, err := http.Get(url)
        if err != nil {
            return apiResponseMsg{"", err.Error()}
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return apiResponseMsg{"", err.Error()}
        }
        return apiResponseMsg{string(body), ""}
    }
}

func main() {
    ti := textinput.New()
    ti.Placeholder = "https://jsonplaceholder.typicode.com/posts/1"
    ti.Focus()

    initialModel := model{
        urlInput: ti,
    }

    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
