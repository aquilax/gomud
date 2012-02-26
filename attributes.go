package main

const (
	STRENGTH = iota
	HEALTH
	AGILITY
	MAXHITPOINTS
	ACCURACY
	DODGING
	STRIKEDAMAGE
	DAMAGEABSORB
	HPREGEN
	NUMATTRIBUTES
)

type AttributeSet [NUMATTRIBUTES]int


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
