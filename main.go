package main

import (
	"fmt"
	"os"

	"github.com/laher/cec"
)

func main() {
	c, err := cec.Open("", "SetTopPi")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		c.Standby(0)
		c.Destroy()
	}()
	c.PowerOn(0)

	s := &stp{}
	chanf, err := os.Open("channels.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer chanf.Close()

	s.LoadChannels(chanf)

	err = s.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}
