package heartrate

import "fmt"

// Level 心率等级
type Level struct {
	Scale    float64 //心率百分比
	SpeakFmt string  //心率意义
}

func (l *Level) String(hr int) string {
	return fmt.Sprintf(l.SpeakFmt,
		l.Scale*100,
		hr)
}

var (
	levels = []Level{
		Level{
			0.5, "%.0f%% %d 轻微，热身、恢复、改善代谢",
		},
		Level{
			0.6, "%.0f%% %d 低强度燃脂，增加代谢、脂肪代谢、体重控制",
		},
		Level{
			0.7, "%.0f%% %d 中等有氧燃脂，建议的燃脂强度",
		},
		Level{
			0.8, "%.0f%% %d 大运动有氧，提高乳酸容忍度、增强高速运动耐力",
		},
		Level{
			0.9, "%.0f%% %d 最大心率，提高冲刺强度、增强肌肉神经系统、运动员或极佳体质",
		},
	}
)
