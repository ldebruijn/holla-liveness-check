package notifier

import "larsdebruijn.nl/holla/task"

type Notifier interface {
	Notify(recv <-chan []task.Result)
}
