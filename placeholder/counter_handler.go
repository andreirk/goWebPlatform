package placeholder

import (
	"fmt"
	"webPlatform/sessions"
)

type CounterHandler struct {
	sessions.Session
}

func (c CounterHandler) GetCounter() string {
	counter := c.Session.GetValueDefault("counter", 0).(int)
	c.Session.SetValue("counter", counter+1)
	return fmt.Sprintf("Counter: %v", counter)
}
