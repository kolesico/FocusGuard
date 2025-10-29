package monitor

import (
	"context"
	"time"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/sys/windows"
)

type Event struct {
    Type      string `json:"type"`
    Timestamp time.Time `json:"timestamp"`
}

func RunMonitor(ctx context.Context, appName string) <-chan Event {
    events := make(chan Event)
	
	go func() {
		defer close(events)
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		var lastState string
		for {
			select {
			case <- ctx.Done():
				return
			case <- ticker.C:
				isFocused := isWindowFocused(appName)
				currentState := "closed"
				if isFocused {
					currentState = "opened"
				}
				if currentState != lastState {
					events <- Event{Type: currentState, Timestamp: time.Now()}
					lastState = currentState
				}
			}
		}

	}()
	return events
}

func isWindowFocused(appName string) bool {
	currentActiveProcess, err := getActiveProcessName()
	if err != nil {
		return false
	}
	if appName == currentActiveProcess {
		return true
	} else {
		return false
	}
}

func getActiveProcessName() (string, error) {
	// Получаем handle активного окна
    hwnd := windows.GetForegroundWindow()
    
    // Получаем ID процесса
    var processID uint32
    _, err := windows.GetWindowThreadProcessId(hwnd, &processID)
    if err != nil {
        return "", fmt.Errorf("ошибка получения ID процесса: %v", err)
    }

	processes, err := process.Processes()
	if err != nil {
		return "", fmt.Errorf("ошибка получения списка процессов: %v", err)
	}

	for _, process := range processes {
		if process.Pid == int32(processID) {
			proc_name, err := process.Name()
			if err != nil {
				return "", fmt.Errorf("failed to get process name: %w", err)
			}
			return proc_name, nil
		}
	}
	return "", nil
}