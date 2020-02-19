package main

import (
	"fmt"

	"github.com/chbmuc/cec"
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
}
