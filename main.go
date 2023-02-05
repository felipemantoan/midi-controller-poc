package main

import (
	"machine"
	"machine/usb/midi"
	"time"
)

// Try it easily by opening the following site in Chrome.
// https://www.onlinemusictools.com/kb/

func main() {

	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pins := [6]machine.Pin{
		machine.GP0,
		machine.GP1,
		machine.GP2,
		machine.GP3,
		machine.GP4,
		machine.GP5,
	}

	for i := 0; i < len(pins); i++ {
		pins[i].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}

	m := midi.New()

	notes := [6]midi.Note{
		midi.C4,
		midi.D4,
		midi.E4,
		midi.F4,
		midi.G4,
		midi.A4,
	}

	for {
		for i := 0; i < len(pins); i++ {
			if pins[i].Get() {
				machine.LED.High()
				m.NoteOn(0, 0, notes[i], 0x40)
				time.Sleep(100 * time.Millisecond)
				m.NoteOff(0, 0, notes[i], 0x40)
			}
			machine.LED.Low()
		}
	}
}
