package placeholder

import (
	"fmt"
	"time"
	"webPlatform/logging"
)

type DayHandler struct {
	logging.Logger
}

func (dh DayHandler) GetDay() string {
	return fmt.Sprintf("Day: %v", time.Now().Day())
}
