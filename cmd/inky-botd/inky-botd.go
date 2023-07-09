package main

import (
	"fmt"
	"time"

	"github.com/13-bit/inky-botd/config"
	"github.com/13-bit/inky-botd/inky"
	"github.com/go-co-op/gocron"
)

func main() {
	// Display the current BOTD at startup
	inky.DownloadBotdImage(fmt.Sprintf("%s/static/inky/botd.png", config.BulletinBirdUrl()))
	inky.Refresh()

	// Update the BOTD every day at five after midnight
	loc, _ := time.LoadLocation("America/Chicago")
	s := gocron.NewScheduler(loc)

	s.Every(1).Day().At("00:05").Do(func() {
		inky.DownloadBotdImage(fmt.Sprintf("%s/static/inky/botd.png", config.BulletinBirdUrl()))
		inky.Refresh()
	})

	s.StartAsync()

	select {}
}
