package services

import (
	"log"
	"os/exec"
	"sync"
)

func PlayNotification(path string) error {
	var m sync.Mutex
	m.Lock() //Add Mutex to avoid multiple plays at the same time
	cmd := exec.Command("mplayer", path)
	log.Printf("Running Command -> mplayer %s\n", path)
	out, err := cmd.Output()
	m.Unlock()
	if err != nil {
		out = []byte(err.Error())
	}
	log.Println(string(out))
	return err
}
