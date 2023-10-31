package main

import (
	"fmt"
	"sync"
)

type MagicWand struct {
	Incantation string
}

type Singleton struct {
	SecretKnowledge string
}

var instance *Singleton
var once sync.Once

func CreateSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{SecretKnowledge: "🪄✨ This is the secret knowledge of the Singleton!"}
	})
	return instance
}

func (wand *MagicWand) CastSpell() {
	singleton := CreateSingleton()
	fmt.Printf("%s\n%s\n", wand.Incantation, singleton.SecretKnowledge)
}

func main() {
	wand1 := &MagicWand{Incantation: "Abracadabra!"}
	wand2 := &MagicWand{Incantation: "SimSalabim!"}

	wand1.CastSpell()
	wand2.CastSpell()
}
