package task

import (
	"larsdebruijn.nl/holla/target"
)

type Task interface {
	Check(results chan<- Result, target target.Target)
}
