package main

import (
	"bufio"
	"github.com/jleight/omxplayer"
	"io"
	"log"
	"time"
)

type stp struct {
	channels []string
	current  chan string
}

// Loads a list of "channels"
func (s *stp) LoadChannels(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s.channels = append(s.channels, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (s *stp) SelectChannel(c int) {
	if len(s.channels) > c {
		v := s.channels[c]
		s.current <- v
	}
}

func (s *stp) Run() error {
	donech := make(chan struct{})
	errch := make(chan error)
	s.current = make(chan string)
	go s.mainLoop(donech, errch)
	s.SelectChannel(0)
	select {
	case <-donech:
		return nil
	case err := <-errch:
		return err
	}
}

// For now this is probably not working.
// Ideally this will play the first file and when a new channel is selected stop playing the current channel
// and then start playing the new one.
func (s *stp) mainLoop(donech chan<- struct{}, errch chan<- error) {
	omxplayer.SetUser("pi", "/home/pi")
	var player *omxplayer.Player
	defer close(donech)
	defer close(errch)
	for {
		select {
		case channel := <-s.current:
			log.Printf("Playing '%s'", channel)
			if player != nil {
				player.Quit()
			}
			player, err := omxplayer.New(channel, "-o hdmi --live -r -b")
			if err != nil {
				errch <- err
				return
			}
			player.WaitForReady()

			err = player.Play()
			if err != nil {
				errch <- err
				return
			}
		case <-time.After(500 * time.Millisecond):
		}
		if player != nil {
			status, err := player.PlaybackStatus()
			if err != nil {
				errch <- err
				return
			}
			if status != "Playing" {
				break
			}
		}
	}
}
