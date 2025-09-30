package wariors

import (
	"errors"
	"fmt"
	"strings"
)

// ------- Structurs ------
type Warrior struct {
	name string
}
type Mage struct {
	name string
}
type Archer struct {
	name string
}

// ------- Factory funcs ------

func NewWarrior(name string) (Warrior, error) {
	if len(name) == 0 {
		return Warrior{}, errors.New("empty name")
	}

	return Warrior{
		name: name,
	}, nil
}

func NewMage(name string) (Mage, error) {
	if len(name) == 0 {
		return Mage{}, errors.New("empty name")
	}

	return Mage{
		name: name,
	}, nil
}

func NewArcher(name string) (Archer, error) {
	if len(name) == 0 {
		return Archer{}, errors.New("empty name")
	}

	return Archer{
		name: name,
	}, nil
}

// ------- Methods ------

func (w Warrior) Attack() string {
	return fmt.Sprintf("Воин %s бьет мечом.", w.name)
}

func (m Mage) Attack() string {
	return fmt.Sprintf("Маг %s колдует огненный шар.", m.name)
}

func (a Archer) Attack() string {
	return fmt.Sprintf("Лучник %s выпускает град стрел.", a.name)
}

// ------- Interface ------

type Character interface {
	Attack() string
}

func Fight(fighters []Character) {
	var strBuilder strings.Builder

	for _, currentCharacter := range fighters {
		strBuilder.WriteString(currentCharacter.Attack() + "\n")
	}

	fmt.Print(strBuilder.String())
}
