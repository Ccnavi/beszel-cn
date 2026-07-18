package alerts

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("BESZEL_ALERT_LOCALE", "zh-CN")
	os.Exit(m.Run())
}

func TestStatusAlertMessageZhCN(t *testing.T) {
	tests := []struct {
		status  string
		want    string
		message string
	}{
		{status: "down", want: "【告警】server-01 连接中断", message: "无法连接"},
		{status: "up", want: "【恢复】server-01 已恢复连接", message: "运行状态正常"},
	}
	for _, test := range tests {
		title, message := statusAlertMessage(test.status, "server-01")
		if !strings.Contains(title, test.want) {
			t.Fatalf("title %q does not contain %q", title, test.want)
		}
		if !strings.Contains(message, test.message) {
			t.Fatalf("message %q does not contain %q", message, test.message)
		}
	}
}

func TestThresholdAlertMessageZhCN(t *testing.T) {
	title, message := thresholdAlertMessage("server-01", "CPU", "", "%", 91.25, 80, 5, true, false)
	if !strings.Contains(title, "【告警】server-01：CPU 使用率超过阈值") {
		t.Fatalf("unexpected title: %q", title)
	}
	for _, want := range []string{"过去 5 分钟", "91.25%", "80.00%"} {
		if !strings.Contains(message, want) {
			t.Fatalf("message %q does not contain %q", message, want)
		}
	}

	title, _ = thresholdAlertMessage("server-01", "Battery", "", "%", 35, 20, 1, false, true)
	if !strings.Contains(title, "【恢复】server-01：电池电量已恢复至阈值以上") {
		t.Fatalf("unexpected recovery title: %q", title)
	}
}

func TestSmartAlertMessageZhCN(t *testing.T) {
	title, message := smartAlertMessage("server-01", "nvme0", "SSD Pro", "FAILED")
	if !strings.Contains(title, "【磁盘告警】server-01：nvme0 S.M.A.R.T. 检测故障") {
		t.Fatalf("unexpected title: %q", title)
	}
	for _, want := range []string{"SSD Pro", "备份重要数据"} {
		if !strings.Contains(message, want) {
			t.Fatalf("message %q does not contain %q", message, want)
		}
	}
}

func TestAlertMessageLocaleCanUseEnglish(t *testing.T) {
	t.Setenv("BESZEL_ALERT_LOCALE", "en")
	title, message := statusAlertMessage("down", "server-01")
	if !strings.Contains(title, "Connection to server-01 is down") {
		t.Fatalf("unexpected English title: %q", title)
	}
	if !strings.Contains(message, "Connection to server-01 is down") {
		t.Fatalf("unexpected English message: %q", message)
	}
	if got := viewSystemText("server-01"); got != "View server-01" {
		t.Fatalf("unexpected English link text: %q", got)
	}
}
