package main

const (
	strength = iota
	health
	agility
	maxhitpoints
	accuracy
	dodging
	strikedamage
	damageabsorb
	hpregen
	numattributes
)

type AttributeSet [numattributes]int


/*
"Strength",
"Halth",
"Agility",
"Max hit points",
"Accuracy",
"Dodging",
"Strike damage",
"Damage Absorb",
"HP Regen"
*/
