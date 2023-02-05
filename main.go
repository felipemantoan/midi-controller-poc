package main

import (
	"machine"
	"machine/usb/midi"
)

// Try it easily by opening the following site in Chrome.
// https://www.onlinemusictools.com/kb/

func main() {

	m := midi.New()

	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.GP0.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.GP1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.GP2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.GP3.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.GP4.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.GP5.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	machine.GP0.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.C3, 0x40)
		} else {
			m.NoteOff(0, 0, midi.C3, 0x40)
		}
	})

	machine.GP1.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.D4, 0x40)
		} else {
			m.NoteOff(0, 0, midi.D4, 0x40)
		}
	})

	machine.GP2.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.E5, 0x40)
		} else {
			m.NoteOff(0, 0, midi.E5, 0x40)
		}
	})

	machine.GP3.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.FS1, 0x40)
		} else {
			m.NoteOff(0, 0, midi.FS1, 0x40)
		}
	})

	machine.GP4.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.G1, 0x40)
		} else {
			m.NoteOff(0, 0, midi.G1, 0x40)
		}
	})

	machine.GP5.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.A2, 0x40)
		} else {
			m.NoteOff(0, 0, midi.A2, 0x40)
		}
	})

}
