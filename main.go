package main

import (
	"fmt"
)

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	case Bow:
		return "BOW"
	}
	return ""

}

const (
	Axe WeaponType = iota // incremet
	Sword
	Bow
	Knife
)

func getDamage(weponType WeaponType) int {
	switch weponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case Bow:
		return 60
	case Knife:
		return 40
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Printf("damage to wepon ($s) ($d):\n", Axe, getDamage(Axe))
	fmt.Printf("damage to wepon ($s) ($d):\n", Sword, getDamage(Sword))
	fmt.Printf("damage to wepon ($s) ($d):\n", Bow, getDamage(Bow))
}
