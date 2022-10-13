package cronjob

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var count = 1

func CronPool() {
	cronjob := cron.New()
	cronjob.AddFunc("21,22,23 10 * * *", job)
	cronjob.Start()

	fmt.Printf("CronJob entries: %d\n", len(cronjob.Entries()))

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	fmt.Println("Termination signal received!")

	cronjob.Stop()
}

func job() {

	fmt.Println("CronJob execution started", count)
	time.Sleep(3 * time.Second)
	fmt.Println("CronJob execution completed", count)
	count++
}
