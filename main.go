package main

import (
	"machine"
	"machine/usb/midi"
	"time"
)

// Try it easily by opening the following site in Chrome.
// https://www.onlinemusictools.com/kb/

func main() {

	m := midi.New()

	machine.InitADC()
	sensor := machine.ADC{Pin: machine.ADC2}
	sensor.Configure(machine.ADCConfig{})

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
			m.NoteOn(0, 0, midi.D3, 0x40)
		} else {
			m.NoteOff(0, 0, midi.D3, 0x40)
		}
	})

	machine.GP2.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.F3, 0x40)
		} else {
			m.NoteOff(0, 0, midi.F3, 0x40)
		}
	})

	machine.GP3.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.G3, 0x40)
		} else {
			m.NoteOff(0, 0, midi.G3, 0x40)
		}
	})

	machine.GP4.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.A3, 0x40)
		} else {
			m.NoteOff(0, 0, midi.A3, 0x40)
		}
	})

	machine.GP5.SetInterrupt(machine.PinFalling|machine.PinRising, func(p machine.Pin) {
		if p.Get() {
			m.NoteOn(0, 0, midi.E3, 0x40)
		} else {
			m.NoteOff(0, 0, midi.E3, 0x40)
		}
	})

	for {
		val := sensor.Get()
		time.Sleep(time.Second)
		println(val)
	}
}
