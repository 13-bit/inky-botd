package inky

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/13-bit/inky-botd/config"
	"github.com/13-bit/inky-botd/util"
	"github.com/cavaliergopher/grab/v3"
)

func DownloadBotdImage(downloadUrl string) {
	log.Printf("Downloading BOTD image from %s...\n", downloadUrl)

	botdPath := config.InkyImagePath()

	_ = os.Remove(botdPath)

	resp, err := grab.Get(botdPath, downloadUrl)
	util.CheckError(err)

	log.Println("Download saved to", resp.Filename)
}

func Refresh() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("python3", config.InkyImageScript(), config.InkyImagePath())
		if err := cmd.Run(); err != nil {
			util.CheckError(err)
		}
	}
}
