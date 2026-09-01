package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type fileEntry struct {
	name   string
	path   string
	isDir  bool
	size   int64
}

type model struct {
	currentPath string
	entries     []fileEntry
	selected    int
	status      string
}

func initialPath() string {
	if wd, err := os.Getwd(); err == nil {
		return wd
	}
	if home, err := os.UserHomeDir(); err == nil {
		return home
	}
	return "."
}

func (m *model) loadEntries() {
	entries, err := os.ReadDir(m.currentPath)
	if err != nil {
		m.status = fmt.Sprintf("Error reading %s: %v", m.currentPath, err)
		return
	}

	var dirs []fileEntry
	var files []fileEntry

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(m.currentPath, name)
		info, err := entry.Info()
		if err != nil {
			m.status = fmt.Sprintf("Error reading %s: %v", fullPath, err)
			continue
		}

		item := fileEntry{
			name:  name,
			path:  fullPath,
			isDir: info.IsDir(),
			size:  info.Size(),
		}

		if info.IsDir() {
			dirs = append(dirs, item)
		} else {
			files = append(files, item)
		}
	}

	sort.Slice(dirs, func(i, j int) bool { return strings.ToLower(dirs[i].name) < strings.ToLower(dirs[j].name) })
	sort.Slice(files, func(i, j int) bool { return strings.ToLower(files[i].name) < strings.ToLower(files[j].name) })
	m.entries = append(dirs, files...)
	m.status = fmt.Sprintf("Viewing %s", m.currentPath)
	if len(m.entries) == 0 {
		m.selected = 0
		return
	}
	if m.selected >= len(m.entries) {
		m.selected = len(m.entries) - 1
	}
}

func (m *model) openSelected() {
	if len(m.entries) == 0 {
		m.status = "No items to open"
		return
	}

	item := m.entries[m.selected]
	if item.isDir {
		m.currentPath = item.path
		m.selected = 0
		m.loadEntries()
		return
	}

	info, err := os.Stat(item.path)
	if err != nil {
		m.status = fmt.Sprintf("Could not read %s: %v", item.path, err)
		return
	}

	m.status = fmt.Sprintf("Selected file: %s (%d bytes)", item.path, info.Size())
}

func (m *model) goToParent() {
	parent := filepath.Dir(m.currentPath)
	if parent == m.currentPath {
		m.status = "Already at the top of the filesystem"
		return
	}
	m.currentPath = parent
	m.selected = 0
	m.loadEntries()
}

func newModel(path string) model {
	m := model{currentPath: path, selected: 0}
	m.loadEntries()
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(m.entries)-1 {
				m.selected++
			}
		case "enter":
			m.openSelected()
		case "backspace", "h":
			m.goToParent()
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	fmt.Fprintf(&b, "File Explorer\n")
	fmt.Fprintf(&b, "Path: %s\n\n", m.currentPath)

	if len(m.entries) == 0 {
		fmt.Fprintln(&b, "No files found in this directory.")
	} else {
		for i, entry := range m.entries {
			prefix := "  "
			name := entry.name
			if entry.isDir {
				name += "/"
			}
			if i == m.selected {
				prefix = "> "
			}
			fmt.Fprintf(&b, "%s%s\n", prefix, name)
		}
	}

	fmt.Fprintf(&b, "\n%s\n\n", m.status)
	fmt.Fprintln(&b, "↑/↓ or j/k move • Enter open • Backspace go up • q/esc quit")
	return b.String()
}

func main() {
	p := tea.NewProgram(newModel(initialPath()))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

