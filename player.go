package main

import (
	"math"
)

const
	PLAYERITEMS = 16

const (
	REGULAR = iota
	GOD
	ADMIN
)

type Room int

type PlayerRank int8

type Player struct {
	//Player information
	m_pass string
	m_rank PlayerRank

	//Player attributes
	m_statpoints int
	m_experience int
	m_level int
	m_room Room
	m_money Money
	m_hitpoints int
	m_baseattributes AttributeSet
	m_attributes AttributeSet
	m_nextattacktime int

	//Player inventory
	m_inventory [PLAYERITEMS]*Item
	m_items int
	m_weapon int
	m_armor int

	m_request Request
	m_loggedin bool
	m_active bool
	m_newbie bool
}

func (p *Player) NeedForLevel(p_level int) (int) {
	return (int)(math.Pow(1.4, float64(p_level))-1)
}

func (p *Player) NeedForNextLevel() (int) {
	return p.NeedForLevel(p.m_level+1) - p.m_experience
}

func (p *Player) Train() (bool) {
	if p.NeedForNextLevel() <= 0 {
		p.m_statpoints += 2
		p.m_baseattributes[MAXHITPOINTS] += p.m_level
		p.m_level++
		p.RecalculateStats()
		return true
	}
	return false
}

func (p *Player) GetAttr(p_attr int) (int) {
	val := p.m_attributes[p_attr] + p.m_baseattributes[p_attr]
	if (p_attr == STRENGTH || p_attr == AGILITY || p_attr == HEALTH) {
		if val < 1 {
			val = 1
		}
	}
	return val
}

func (p *Player) RecalculateStats() {
    p.m_attributes[MAXHITPOINTS] = 10 + int(float64(p.m_level * p.GetAttr(HEALTH)) / 1.5)
    p.m_attributes[HPREGEN] = (p.GetAttr( HEALTH) / 5) + p.m_level
    p.m_attributes[ACCURACY] = p.GetAttr(AGILITY) * 3
    p.m_attributes[DODGING] = p.GetAttr(AGILITY) * 3
    p.m_attributes[DAMAGEABSORB] = p.GetAttr(STRENGTH) / 5
    p.m_attributes[STRIKEDAMAGE] = p.GetAttr(STRENGTH) / 5

    // make sure the hitpoints don't overflow if your max goes down:
    if p.m_hitpoints > p.GetAttr( MAXHITPOINTS ) {
        p.m_hitpoints = p.GetAttr( MAXHITPOINTS )
	}
    if p.m_weapon != 0  {
        p.AddDynamicBonuses(p.m_inventory[p.m_weapon]);
	}
    if p.m_armor != 0  {
        p.AddDynamicBonuses(p.m_inventory[p.m_armor]);
	}
}

func (p *Player) AddHitpoints(p_hitpoints int) {
	p.m_hitpoints += p_hitpoints
	if p.m_hitpoints < 0 {
		p.m_hitpoints = 0
	}
	if p.m_hitpoints > p.GetAttr(MAXHITPOINTS) {
		p.m_hitpoints = p.GetAttr(MAXHITPOINTS)
	}
}

func (p *Player) GetBaseAttr (p_attr int) (int) {
	return p.m_baseattributes[p_attr]
}

func (p *Player) SetBaseAttr (p_attr int, p_val int) {
	p.m_baseattributes[p_attr] = p_val
	p.RecalculateStats()
}

func (p *Player) AddBaseAttr (p_attr int, p_val int) {
	p.m_baseattributes[p_attr] += p_val
	p.RecalculateStats()
}

func (p *Player) Weapon() (*Item) {
	if p.m_weapon != -1 {
		return p.m_inventory[p.m_weapon]
	}
	return nil
}

func (p *Player) Armor() (*Item) {
	if p.m_armor != -1 {
		return p.m_inventory[p.m_armor]
	}
	return nil
}

func (p *Player) AddDynamicBonuses (p_item *Item) {
	if (p_item == nil) {
		return
	}
	for i := 0; i < NUMATTRIBUTES; i++ {
		p.m_attributes[i] += p_item.GetAttr(i);
	}
}
