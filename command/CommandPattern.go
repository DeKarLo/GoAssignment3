package main

import "fmt"

type Spell interface {
	Cast()
}

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("The light is magically turned on.")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("The light is magically turned off.")
}

type TurnOnSpell struct {
	light *Light
}

func (s *TurnOnSpell) Cast() {
	s.light.TurnOn()
}

type TurnOffSpell struct {
	light *Light
}

func (s *TurnOffSpell) Cast() {
	s.light.TurnOff()
}

type MagicWand struct {
}

func (w *MagicWand) CastSpell(spell Spell) {
	fmt.Println("Casting a magical spell...")
	spell.Cast()
}

func main() {
	light := &Light{}

	turnOnSpell := &TurnOnSpell{light}
	turnOffSpell := &TurnOffSpell{light}

	wand := &MagicWand{}

	wand.CastSpell(turnOnSpell)
	wand.CastSpell(turnOffSpell)
}
