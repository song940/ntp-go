package cli

import (
	"flag"
	"fmt"

	"github.com/song940/ntp-go/ntp"
)

var (
	server string
)

func Run() {

	flag.StringVar(&server, "s", "time.apple.com:123", "ntp server")
	flag.Parse()

	client := ntp.NewClient(server)
	time, err := client.GetTime()
	if err != nil {
		panic(err)
	}
	fmt.Println(time)
}
