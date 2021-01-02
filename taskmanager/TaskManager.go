package taskmanager

import (
	"larsdebruijn.nl/holla/reporter"
	"larsdebruijn.nl/holla/target"
	"larsdebruijn.nl/holla/task"
	"log"
)

type TaskManager struct{}

func (t *TaskManager) Start(group target.Group) {
	resultReceiver := make(chan task.Result, 1)
	var results []task.Result
	reporter := reporter.StdoutReorter{
		Group: group,
	}
	reporter.Initialize()

	for _, tgt := range group.Targets {
		tsk := task.HttpTask{}
		go tsk.Check(resultReceiver, tgt)
	}

	for {
		select {
		case result := <-resultReceiver:
			log.Printf("Received task result '%s' to %s\n", result.Target.Name, result.Target.Uri)
			results = append(results, result)
		}
		if len(results) == len(group.Targets) {
			// stop listening when all results are received
			log.Println("All results received, stopping listening")
			break
		}
	}

	reporter.Report(results)
}
