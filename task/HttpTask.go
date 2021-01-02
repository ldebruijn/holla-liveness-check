package task

import (
	"errors"
	"larsdebruijn.nl/holla/target"
	"log"
	"net/http"
	"time"
)

type HttpTask struct{}

func (t HttpTask) Check(results chan<- Result, target target.Target) {
	log.Printf("Performing check to target %s for %s", target.Uri, target.Name)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(target.Uri)
	if err != nil {
		result := Result{
			Target:    target,
			Error:     err,
			Timestamp: time.Now(),
		}
		results <- result
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		result := Result{
			Target:    target,
			Error:     errors.New("Statuscode not OK: " + resp.Status),
			Timestamp: time.Now(),
		}
		results <- result
		return
	}

	// handle success
	result := Result{
		Target:    target,
		Error:     nil,
		Timestamp: time.Now(),
	}
	results <- result
}
