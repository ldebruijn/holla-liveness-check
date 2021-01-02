package task

import (
	"larsdebruijn.nl/holla/target"
	"time"
)

type Result struct {
	Target    target.Target
	Error     error
	Timestamp time.Time
}
