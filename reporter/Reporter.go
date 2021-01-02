package reporter

import (
	"larsdebruijn.nl/holla/target"
	"larsdebruijn.nl/holla/task"
)

type Reporter interface {
	Initialize(group target.Group)
	Report(results []task.Result)
}
