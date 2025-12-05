package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/kolesico/FocusGuard/internal/events"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/sys/windows"
)

func RunMonitor(ctx context.Context, appName *string) <-chan events.Events {
	eventsCh := make(chan events.Events)

	go func() {
		defer close(eventsCh)
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		var lastState string
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				isFocused := isWindowFocused(*appName)
				currentState := "closed"
				if isFocused {
					currentState = "opened"
				}
				if currentState != lastState {
					eventsCh <- events.Events{Event: currentState, Timestamp: time.Now()}
					lastState = currentState
				}
			}
		}

	}()
	return eventsCh
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
		return "", fmt.Errorf("fialed to get process ID: %v", err)
	}

	processes, err := process.Processes()
	if err != nil {
		return "", fmt.Errorf("failed to get list of proccess: %v", err)
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
