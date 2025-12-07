package common

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger interface {
	Log(format string, args ...interface{})
	Flush()
}

type GlobalLogManager struct {
	mu   sync.Mutex
	file *os.File
}

func NewGlobalLogManager(path string) (*GlobalLogManager, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	return &GlobalLogManager{file: f}, nil
}

func (m *GlobalLogManager) Close() {
	if m.file != nil {
		m.file.Close()
	}
}

func (m *GlobalLogManager) WriteAtomic(b []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.file != nil {
		m.file.Write(b)
	}
}

type FacilityLogger struct {
	manager  *GlobalLogManager
	buffer   bytes.Buffer
	Facility string
}

func NewFacilityLogger(manager *GlobalLogManager, facilityName string) *FacilityLogger {
	return &FacilityLogger{
		manager:  manager,
		Facility: facilityName,
	}
}

func (l *FacilityLogger) Log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.buffer.WriteString(msg)
	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		l.buffer.WriteString("\n")
	}
}

func (l *FacilityLogger) Flush() {
	var out bytes.Buffer
	out.WriteString("\n================================================================================\n")
	out.WriteString(fmt.Sprintf("Processing Facility: %s\n", l.Facility))
	out.WriteString(fmt.Sprintf("Timestamp: %s\n", time.Now().Format(time.RFC3339)))
	out.WriteString("================================================================================\n")

	out.Write(l.buffer.Bytes())

	out.WriteString("\n")

	l.manager.WriteAtomic(out.Bytes())
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	if len(format) == 0 || format[len(format)-1] != '\n' {
		fmt.Println()
	}
}

func (c *ConsoleLogger) Flush() {}
