package alerts

import "fmt"

func statusAlertMessageZhCN(status, systemName string) (title, message string) {
	if status == "up" {
		return fmt.Sprintf("【恢复】%s 已恢复连接 ✅", systemName),
			fmt.Sprintf("系统 %s 已恢复连接，当前运行状态正常。", systemName)
	}
	return fmt.Sprintf("【告警】%s 连接中断 🔴", systemName),
		fmt.Sprintf("系统 %s 当前无法连接，请及时检查主机、网络和监控端状态。", systemName)
}

func alertNameZhCN(name string) string {
	switch name {
	case "CPU":
		return "CPU 使用率"
	case "Memory":
		return "内存使用率"
	case "Disk":
		return "磁盘使用率"
	case "Bandwidth":
		return "带宽"
	case "GPU":
		return "GPU 使用率"
	case "Temperature":
		return "温度"
	case "LoadAvg1":
		return "1 分钟平均负载"
	case "LoadAvg5":
		return "5 分钟平均负载"
	case "LoadAvg15":
		return "15 分钟平均负载"
	case "Battery":
		return "电池电量"
	default:
		return name
	}
}

func thresholdAlertMessageZhCN(systemName, name, descriptor, unit string, value, threshold float64, minutes uint8, triggered, lowAlert bool) (title, message string) {
	metricName := alertNameZhCN(name)
	state := "恢复"
	direction := "已恢复至阈值以下"
	emoji := "✅"
	if lowAlert {
		direction = "已恢复至阈值以上"
	}
	if triggered {
		state = "告警"
		emoji = "🔴"
		direction = "超过阈值"
		if lowAlert {
			direction = "低于阈值"
		}
	}
	title = fmt.Sprintf("【%s】%s：%s%s %s", state, systemName, metricName, direction, emoji)
	if descriptor == "" {
		descriptor = metricName
	}
	message = fmt.Sprintf("%s在过去 %d 分钟的平均值为 %.2f%s，告警阈值为 %.2f%s。", descriptor, minutes, value, unit, threshold, unit)
	return title, message
}

func diskDescriptorZhCN(name string) string {
	if name == "root" {
		return "根分区使用率"
	}
	return fmt.Sprintf("分区 %s 的使用率", name)
}

func temperatureDescriptorZhCN(name string) string {
	return fmt.Sprintf("温度传感器 %s", name)
}

func smartAlertMessageZhCN(systemName, deviceName, model, state string) (title, message string) {
	stateText := "警告"
	emoji := "🟠"
	if state == "FAILED" {
		stateText = "故障"
		emoji = "🔴"
	}
	title = fmt.Sprintf("【磁盘告警】%s：%s S.M.A.R.T. 检测%s %s", systemName, deviceName, stateText, emoji)
	if model != "" {
		message = fmt.Sprintf("磁盘 %s（%s）的 S.M.A.R.T. 状态已变为%s，请及时检查磁盘健康状况并备份重要数据。", deviceName, model, stateText)
	} else {
		message = fmt.Sprintf("磁盘 %s 的 S.M.A.R.T. 状态已变为%s，请及时检查磁盘健康状况并备份重要数据。", deviceName, stateText)
	}
	return title, message
}
