package main

import "project/wariors"

func main() {
	mag, _ := wariors.NewMage("Maga")
	war, _ := wariors.NewWarrior("War")
	arch, _ := wariors.NewArcher("Arch")

	wariors.Fight([]wariors.Character{mag, war, arch})
}
