# TUI apps

Text User Interfaces (TUIs) are interactive applications that run within a terminal or command-line environment, bridging the gap between raw text-based Command Line Interfaces (CLIs) and graphical User Interfaces (GUIs).  Unlike standard CLIs that rely on typed commands and plain text output, TUIs offer rich, visual layouts using box-drawing characters, colors, and panels, often supporting mouse interaction and keyboard shortcuts for navigation. 

These apps are designed to be lightweight, keyboard-driven, and resource-efficient, making them particularly popular for system administration, server management, and developer workflows where minimal overhead is preferred.  They provide a more intuitive experience than traditional CLIs by organizing data into structured views, while remaining significantly faster and less demanding on system resources than full GUI applications.

# Go TUI Learning Roadmap

This repository contains a series of **small Text User Interface (TUI) apps** built with [Go](https://golang.org/) and [Bubble Tea](https://github.com/charmbracelet/bubbletea).  
The goal is to learn TUI programming step by step, starting from very basic examples and gradually moving toward more useful applications.

---

## 📌 Roadmap

Each app is a separate folder in this repository. The roadmap is designed to introduce one concept at a time.

1. **Hello World Interactive**
   - Display a simple message in the terminal.
   - Exit when any key is pressed.
   - Concepts: main loop, event handling.

2. **Counter with Keys**
   - Increment/decrement a number using arrow keys.
   - Concepts: state management, dynamic rendering.

3. **Simple Menu**
   - Navigate through a list of options with arrow keys.
   - Select an option with Enter.
   - Concepts: layout basics, navigation.

4. **Todo List**
   - Add, mark, and delete tasks.
   - Save tasks to a JSON or CSV file.
   - Concepts: persistence, editing interactively.

5. **File Explorer**
   - Navigate directories and open text files.
   - Concepts: filesystem integration, multiple panels.

6. **Basic API Client**
   - Query a public API (e.g., weather) and display results.
   - Concepts: networking, rendering external data.

7. **Mini Dashboard**
   - Show system metrics (CPU, memory, time) in panels.
   - Concepts: periodic updates, multiple views.

---

## ⚠️ Challenges to Face
- **State management**: keeping track of active views and data changes.
- **Keyboard shortcuts**: designing intuitive keybindings.
- **Cross-platform compatibility**: handling differences between Linux, macOS, and Windows terminals.
- **Modular design**: separating business logic from rendering.

---

## 🌱 Best Practices
- Start small: focus on learning one concept per app.
- Use libraries: Bubble Tea (for TUI) and Lip Gloss (for styling).
- Document shortcuts: always show available keybindings on screen.
- Test across systems: avoid relying on a single environment.

---

## ✅ Goal 
By completing this roadmap, I will build a foundation in TUI programming with Go.  
These mini-apps will serve as building blocks for more advanced projects such as database clients, dashboards, or custom developer tools.

---

## 📂 Repository Structure
/hello-world
/counter
/menu
/todo-list
/file-explorer
/api-client
/dashboard

Each folder contains:
- `main.go` → the app source code
- `README.md` → explanation of the app and concepts learned

---

## 🚀 Getting Started
1. Install Go: [https://golang.org/dl/](https://golang.org/dl/)
2. Install Bubble Tea:
   ```bash
   go get github.com/charmbracelet/bubbletea
3. Run any app
   ```bash
   cd hello-world
   go run main.go

---

## ⚠️ Disclaimer
This repository is created solely for learning purposes.
The apps included here are not intended for production use or for any critical environment that requires accuracy, reliability, or performance in real-world scenarios.

They are designed to:

- Help to understand the basics of Go and TUI programming.
- Serve as small experiments and practice projects.
- Provide a foundation that can be extended for personal use only.

If you find these apps useful as a starting point, feel free to adapt and expand them for your own non-critical projects.
However, do not rely on them in professional or production environments.
