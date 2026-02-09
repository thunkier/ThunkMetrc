package app

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	ClearScreen    = "\033[2J"
	MoveTop        = "\033[H"
	ClearEOS       = "\033[J"
	ClearLine      = "\033[K"
	HideCursor     = "\033[?25l"
	ShowCursor     = "\033[?25h"
	CursorUpFormat = "\033[%dA"
	Reset          = "\033[0m"
	Green          = "\033[32m"
	Yellow         = "\033[33m"
	Blue           = "\033[34m"
	Gray           = "\033[90m"
)

type WorkerState struct {
	License     string
	Name        string
	Scenario    string
	CurrentTask string
	Status      string
	LastUpdate  time.Time
}

type Dashboard struct {
	mu         sync.Mutex
	workers    map[string]*WorkerState
	running    bool
	lastHeight int
}

func NewDashboard() *Dashboard {
	return &Dashboard{
		workers: make(map[string]*WorkerState),
		running: true,
	}
}

func (d *Dashboard) Register(license, name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.workers[license] = &WorkerState{
		License:     license,
		Name:        name,
		Status:      "Pending",
		LastUpdate:  time.Now(),
		CurrentTask: "Waiting...",
	}
}

func (d *Dashboard) Update(license string, task string, status string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if w, ok := d.workers[license]; ok {
		w.CurrentTask = task
		if status != "" {
			w.Status = status
		}
		w.LastUpdate = time.Now()
	}
}

func (d *Dashboard) UpdateScenario(license string, scenario string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if w, ok := d.workers[license]; ok {
		w.Scenario = scenario
		w.LastUpdate = time.Now()
	}
}

func (d *Dashboard) Stop() {
	d.running = false
}

func (d *Dashboard) Run() {
	fmt.Print(ClearScreen)
	fmt.Print(MoveTop)
	fmt.Print(HideCursor)
	defer fmt.Print(ShowCursor)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	for d.running {
		<-ticker.C
		d.render()
	}
	d.render()
}

func (d *Dashboard) render() {
	d.mu.Lock()
	defer d.mu.Unlock()

	licenses := make([]string, 0, len(d.workers))
	var running, done, errored, pending int

	for l, w := range d.workers {
		licenses = append(licenses, l)
		switch w.Status {
		case "Running":
			running++
		case "Done":
			done++
		case "Error":
			errored++
		case "Pending":
			pending++
		}
	}
	sort.Strings(licenses)

	content := strings.Builder{}

	content.WriteString(fmt.Sprintf("%s=== ThunkMetrc Automation Dashboard ===%s%s\n", Blue, Reset, ClearLine))
	content.WriteString(fmt.Sprintf("Active: %s%d%s | Done: %s%d%s | Errors: %s%d%s | Pending: %d%s\n\n",
		Green, running, Reset,
		Blue, done, Reset,
		Yellow, errored, Reset,
		pending, ClearLine,
	))

	content.WriteString(fmt.Sprintf("%-20s %-30s %-15s %-10s %s%s\n", "LICENSE", "NAME", "SCENARIO", "STATUS", "CURRENT TASK", ClearLine))
	content.WriteString(strings.Repeat("-", 100) + ClearLine + "\n")

	for _, license := range licenses {
		w := d.workers[license]

		statusColor := Reset
		switch w.Status {
		case "Running":
			statusColor = Green
		case "Error":
			statusColor = Yellow
		case "Done":
			statusColor = Blue
		case "Pending":
			statusColor = Gray
		}

		name := w.Name
		if len(name) > 28 {
			name = name[:25] + "..."
		}

		scenario := w.Scenario
		if scenario == "" {
			scenario = "-"
		}

		content.WriteString(fmt.Sprintf("%-20s %-30s %-15s %s%-10s%s %s%s\n",
			w.License,
			name,
			scenario,
			statusColor, w.Status, Reset,
			w.CurrentTask,
			ClearLine,
		))
	}

	contentStr := content.String()
	lineCount := strings.Count(contentStr, "\n")

	finalOutput := strings.Builder{}
	if d.lastHeight > 0 {
		finalOutput.WriteString(fmt.Sprintf(CursorUpFormat, d.lastHeight))
	}

	finalOutput.WriteString(contentStr)
	finalOutput.WriteString(ClearEOS)

	fmt.Print(finalOutput.String())

	d.lastHeight = lineCount
}
