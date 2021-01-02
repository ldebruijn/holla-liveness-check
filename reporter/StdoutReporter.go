package reporter

import (
	"larsdebruijn.nl/holla/target"
	"larsdebruijn.nl/holla/task"
	"log"
)

type StdoutReorter struct {
	Group      target.Group
	workingDir string
}

func (r *StdoutReorter) Initialize() {

}

func (r StdoutReorter) Report(results []task.Result) {
	for _, result := range results {
		status := "success"
		if result.Error != nil {
			status = "failure"
		}

		log.Printf("[%s] Target %s reported %s\n", r.Group.Name, result.Target.Uri, status)
	}
}
