package heartrate

import (
	"fmt"
	"strings"
)

const (
	// MaxAge MaxAge
	MaxAge = 220
)

func reserve(staticHeartRate, age int) (int, error) {
	if age > MaxAge {
		return 0, fmt.Errorf("age %d is too large. must be in [14, %d)", age, MaxAge)
	} else if age < 14 {
		return 0, fmt.Errorf("age %d is too small. must be in [14, %d)", age, MaxAge)
	} else if staticHeartRate < 30 {
		return 0, fmt.Errorf("static heart rate %d is too small. must be in [30, 180)", age)
	} else if staticHeartRate >= 180 {
		return 0, fmt.Errorf("static heart rate %d is too large. must be in [30, 180)", age)
	}

	return MaxAge - age - staticHeartRate, nil
}

// KarvonenFormula 卡式心率
func KarvonenFormula(staticHeartRate, age int) (string, error) {
	reserve, e := reserve(staticHeartRate, age)
	if e != nil {
		return "", e
	}

	var out strings.Builder
	out.WriteString("**********\n卡式心率\n\n")
	for _, level := range levels {
		out.WriteString(level.String(int(level.Scale*float64(reserve)) + staticHeartRate))
		out.WriteString("\n")
	}
	out.WriteString("\n**********\n\n")

	out.WriteString("**********\n最大心率\n\n")
	reserve = MaxAge - age
	for _, level := range levels {
		out.WriteString(level.String(int(level.Scale * float64(reserve))))
		out.WriteString("\n")
	}
	out.WriteString("\n**********\n\n")

	return out.String(), nil
}
