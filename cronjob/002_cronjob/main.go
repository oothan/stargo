package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	cronjob := cron.New()
	cronjob.AddFunc("5,6,7 10 * * *", job)
	cronjob.Start()

	fmt.Printf("CronJob entries: %d\n", len(cronjob.Entries()))

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	fmt.Println("Termination signal received!")

	cronjob.Stop()
}

func job() {
	var count = 1
	fmt.Println("CronJob execution started", count)
	time.Sleep(3 * time.Second)
	fmt.Println("CronJob execution completed", count)
	count++
}
