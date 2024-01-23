package services

import (
	"log"
	"os/exec"
)

func PlayNotification(path string) error {
	cmd := exec.Command("mplayer", path)
	log.Printf("Running Command -> mplayer %s\n", path)
	out, err := cmd.Output()
	if err != nil {
		out = []byte(err.Error())
	}
	log.Println(string(out))

	return err
}
