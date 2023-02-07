package main

import (
	"log"

	"PACKAGENAME/utilities"

	"github.com/robfig/cron/v3"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("sso/cron crashed, error:", err)
		}
	}()

	log.Println("starting...")

	go crontab()

	log.Println("waiting...")
	select {}

}

func crontab() {
	log.Println("starting crontab...")

	c := cron.New(cron.WithSeconds())
	c.AddFunc("* * * * * *", func() {
		log.Println(utilities.GetNowDDHHMMSS())
	})
	c.Start()
}
