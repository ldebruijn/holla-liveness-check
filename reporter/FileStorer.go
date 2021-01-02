package reporter

import (
	"larsdebruijn.nl/holla/target"
	"larsdebruijn.nl/holla/task"
	"log"
	"os"
	"path/filepath"
)

type FileStorer struct {
	Group      target.Group
	workingDir string
}

func (r *FileStorer) Initialize() {
	executable, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	r.workingDir = filepath.Dir(executable)
	log.Println("Working dir: " + r.workingDir)

	err = os.Mkdir(r.workingDir+"/db", 0664)
	if err != nil {
		log.Println(err)
	}
	err = os.Mkdir(r.workingDir+"/db/"+r.Group.Name, 0664)
	if err != nil {
		log.Println(err)
	}
}

func (r FileStorer) Report(results []task.Result) {
	log.Println("working dir", r.workingDir)
	for _, result := range results {
		filePath := r.workingDir + "/db/" + r.Group.Name + "/" + result.Target.Uri
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			log.Println("Unable to open file", err)
			continue
		}

		status := "status"
		if result.Error != nil {
			status = "failure"
		}
		_, err = file.WriteString("[" + result.Timestamp.String() + "] " + status)
		if err != nil {
			log.Println("Error writing report to file: "+filePath, err)
		}
	}
}
