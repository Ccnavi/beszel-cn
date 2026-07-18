package alerts

import (
	"fmt"
	"strings"
)

func statusAlertMessageEn(status, systemName string) (title, message string) {
	emoji := "🔴"
	if status == "up" {
		emoji = "✅"
	}
	title = fmt.Sprintf("Connection to %s is %s %s", systemName, status, emoji)
	message = strings.TrimSpace(strings.TrimSuffix(title, emoji))
	return title, message
}

func thresholdAlertMessageEn(systemName, name, descriptor, unit string, value, threshold float64, minutes uint8, triggered, lowAlert bool) (title, message string) {
	if name == "Disk" {
		name += " usage"
	}
	if after, ok := strings.CutPrefix(name, "LoadAvg"); ok {
		name = after + "m Load"
	}
	titleAlertName := name
	if titleAlertName != "CPU" && titleAlertName != "GPU" {
		titleAlertName = strings.ToLower(titleAlertName)
	}
	if triggered {
		if lowAlert {
			title = fmt.Sprintf("%s %s below threshold", systemName, titleAlertName)
		} else {
			title = fmt.Sprintf("%s %s above threshold", systemName, titleAlertName)
		}
	} else {
		if lowAlert {
			title = fmt.Sprintf("%s %s above threshold", systemName, titleAlertName)
		} else {
			title = fmt.Sprintf("%s %s below threshold", systemName, titleAlertName)
		}
	}
	minutesLabel := "minute"
	if minutes > 1 {
		minutesLabel += "s"
	}
	if descriptor == "" {
		descriptor = name
	}
	message = fmt.Sprintf("%s averaged %.2f%s for the previous %v %s.", descriptor, value, unit, minutes, minutesLabel)
	return title, message
}

func smartAlertMessageEn(systemName, deviceName, model, state string) (title, message string) {
	statusLabel := strings.ToLower(state)
	if state == "FAILED" {
		statusLabel = "failure"
	}
	emoji := "🔴"
	if state == "WARNING" {
		emoji = "🟠"
	}
	title = fmt.Sprintf("SMART %s on %s: %s %s", statusLabel, systemName, deviceName, emoji)
	if model != "" {
		message = fmt.Sprintf("Disk %s (%s) SMART status changed to %s", deviceName, model, state)
	} else {
		message = fmt.Sprintf("Disk %s SMART status changed to %s", deviceName, state)
	}
	return title, message
}
