package services

import (
	"fmt"
	"os"
	"time"
)

// JudgeService provides unsolicited advice based on system and environmental context.
type JudgeService struct{}

func NewJudgeService() *JudgeService {
	return &JudgeService{}
}

func (j *JudgeService) JudgeSession() string {
	now := time.Now()
	hour := now.Hour()

	// 1. Time-based Judgment
	if hour >= 5 && hour < 12 {
		return "It is morning. Unless you are at an airport, reconsider."
	}
	if hour >= 12 && hour < 17 {
		return "Afternoon session. Keep it low ABV."
	}
	if hour >= 23 || hour < 4 {
		return "It is late. Go to bed."
	}

	// 2. Load-based Judgment (Linux only)
	load, err := getLoadAverage()
	if err == nil {
		if load > 2.0 {
			return fmt.Sprintf("System load is high (%.2f). You probably need a stiff drink.", load)
		}
	}

	return "Systems nominal. Proceed."
}

func getLoadAverage() (float64, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, err
	}
	var l1, l2, l3 float64
	fmt.Sscanf(string(data), "%f %f %f", &l1, &l2, &l3)
	return l1, nil
}
