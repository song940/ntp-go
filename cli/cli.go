package cli

import (
	"log"

	"github.com/song940/ntp/ntp"
)

func Run() {
	client := ntp.NewClient(
		"time.apple.com:123",
	)
	time, err := client.GetTime()
	log.Println(time, err)
}
