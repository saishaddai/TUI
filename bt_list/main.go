package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	devices []string
	selected int
}

type devicesLoadedMsg []string

func (m model) Init() tea.Cmd {
	return loadDevices
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case devicesLoadedMsg:
		m.devices = []string(msg)
		if len(m.devices) > 0 && m.selected >= len(m.devices) {
			m.selected = len(m.devices) - 1
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if len(m.devices) > 0 && m.selected < len(m.devices)-1 {
				m.selected++
			}
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if len(m.devices) == 0 {
		return "Loading Bluetooth devices...\n\nPress Ctrl+C to quit."
	}

	var b strings.Builder
	b.WriteString("Bluetooth Devices\n\n")
	for i, device := range m.devices {
		prefix := "  "
		if i == m.selected {
			prefix = "> "
		}
		fmt.Fprintf(&b, "%s%s\n", prefix, device)
	}
	b.WriteString("\n↑/↓ to move • q/esc to quit")
	return b.String()
}

func loadDevices() tea.Msg {
	devices := discoverBluetoothDevices()
	if len(devices) == 0 {
		return devicesLoadedMsg{"No Bluetooth devices found"}
	}
	return devicesLoadedMsg(devices)
}

func discoverBluetoothDevices() []string {
	switch runtime.GOOS {
	case "darwin":
		output, err := exec.Command("sh", "-c", "system_profiler SPBluetoothDataType 2>/dev/null | awk -F': ' '/Name:/{print $2}'").CombinedOutput()
		if err == nil {
			devices := parseDeviceList(string(output))
			if len(devices) > 0 {
				return devices
			}
		}
		return []string{"No Bluetooth devices found"}
	case "linux":
		if output, err := exec.Command("sh", "-c", "bluetoothctl devices 2>/dev/null | awk '{print $3}'").CombinedOutput(); err == nil {
			devices := parseDeviceList(string(output))
			if len(devices) > 0 {
				return devices
			}
		}
		if output, err := exec.Command("sh", "-c", "bt-device -l 2>/dev/null | awk -F'\t' 'NF>1 {print $2}'").CombinedOutput(); err == nil {
			devices := parseDeviceList(string(output))
			if len(devices) > 0 {
				return devices
			}
		}
	}

	return []string{"Bluetooth is unavailable or no devices are connected"}
}

func parseDeviceList(output string) []string {
	seen := make(map[string]struct{})
	devices := make([]string, 0)

	for _, line := range strings.Split(output, "\n") {
		device := strings.TrimSpace(line)
		if device == "" {
			continue
		}
		if _, ok := seen[device]; ok {
			continue
		}
		seen[device] = struct{}{}
		devices = append(devices, device)
	}

	sort.Strings(devices)
	return devices
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
