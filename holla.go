package main

import (
	"fmt"
	"larsdebruijn.nl/holla/scheduler"
	"log"
)

func main() {
	fmt.Println("Initializing Holla...")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Initialized")

	sched := scheduler.CronScheduler{}
	go sched.Start()

	select {}
}
