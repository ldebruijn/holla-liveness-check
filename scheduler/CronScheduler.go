package scheduler

import (
	"larsdebruijn.nl/holla/config"
	"larsdebruijn.nl/holla/target"
	"larsdebruijn.nl/holla/taskmanager"
	"log"
	"strconv"
	"time"
)

type CronScheduler struct {
	groups []target.Group
}

const FiveMinutes = 5 * time.Second

func (scheduler *CronScheduler) Start() {
	log.Println("Scheduler started")
	// load groups
	scheduler.groups = loadGroups()
	// schedule groups by cron definition
	scheduler.registerGroups()
	// when cron is triggered, start process that is responsible for running and reporting on a group
	go scheduler.scheduler()
}

func (scheduler *CronScheduler) scheduler() {
	log.Println("Starting scheduler")
	ticker := time.NewTicker(FiveMinutes)
	for {
		select {
		case <-ticker.C:
			for _, group := range scheduler.groups {
				manager := taskmanager.TaskManager{}
				manager.Start(group)
			}
		}
	}
}

func (scheduler *CronScheduler) registerGroups() {
	for i, group := range scheduler.groups {
		log.Println(strconv.Itoa(i) + ": Registering group " + group.Name)
	}
}

func loadGroups() []target.Group {
	configReceiver := make(chan []target.Group, 1)
	configService := config.YamlConfigurationService{}
	configService.Load(configReceiver)

	for {
		select {
		case config := <-configReceiver:
			log.Printf("Configuration loaded. [%d] Groups configured", len(config))
			close(configReceiver)
			return config
		}
	}
}
