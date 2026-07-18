package alerts

import (
	"fmt"
	"os"
	"strings"
)

func alertLocale() string {
	locale := strings.TrimSpace(os.Getenv("BESZEL_ALERT_LOCALE"))
	if locale == "" {
		return "zh-CN"
	}
	return locale
}

func isZhCNAlertLocale() bool {
	locale := strings.ToLower(alertLocale())
	return locale == "zh-cn" || locale == "zh"
}

func viewSystemText(systemName string) string {
	if isZhCNAlertLocale() {
		return "查看系统"
	}
	return "View " + systemName
}

func notificationURLParseError(err error) error {
	if isZhCNAlertLocale() {
		return fmt.Errorf("通知 URL 解析失败：%v", err)
	}
	return fmt.Errorf("error parsing URL: %v", err)
}

func notificationURLRequiredMessage() string {
	if isZhCNAlertLocale() {
		return "请输入通知 URL"
	}
	return "URL is required"
}

func internalNotificationDestinationMessage() string {
	if isZhCNAlertLocale() {
		return "仅管理员可以向内网地址发送测试通知"
	}
	return "Only admins can send to internal destinations"
}

func testNotificationTitle() string {
	if isZhCNAlertLocale() {
		return "【测试告警】Beszel 通知测试"
	}
	return "Test Alert"
}

func testNotificationMessage() string {
	if isZhCNAlertLocale() {
		return "这是一条来自 Beszel 的测试通知。收到此消息说明通知渠道配置正确。"
	}
	return "This is a notification from Beszel."
}

func testNotificationLinkText() string {
	if isZhCNAlertLocale() {
		return "打开 Beszel"
	}
	return "View Beszel"
}

func statusAlertMessage(status, systemName string) (title, message string) {
	if isZhCNAlertLocale() {
		return statusAlertMessageZhCN(status, systemName)
	}
	return statusAlertMessageEn(status, systemName)
}

func thresholdAlertMessage(systemName, name, descriptor, unit string, value, threshold float64, minutes uint8, triggered, lowAlert bool) (title, message string) {
	if isZhCNAlertLocale() {
		return thresholdAlertMessageZhCN(systemName, name, descriptor, unit, value, threshold, minutes, triggered, lowAlert)
	}
	return thresholdAlertMessageEn(systemName, name, descriptor, unit, value, threshold, minutes, triggered, lowAlert)
}

func diskDescriptor(name string) string {
	if isZhCNAlertLocale() {
		return diskDescriptorZhCN(name)
	}
	return fmt.Sprintf("Usage of %s", name)
}

func temperatureDescriptor(name string) string {
	if isZhCNAlertLocale() {
		return temperatureDescriptorZhCN(name)
	}
	return fmt.Sprintf("Highest sensor %s", name)
}

func smartAlertMessage(systemName, deviceName, model, state string) (title, message string) {
	if isZhCNAlertLocale() {
		return smartAlertMessageZhCN(systemName, deviceName, model, state)
	}
	return smartAlertMessageEn(systemName, deviceName, model, state)
}
