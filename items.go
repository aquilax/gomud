package main

const (
	weapon = iota
	armor
	healing
)

type Item struct {
	item_type int
	m_min int
	m_max int
	m_speed int
	m_price Money
	m_attributes AttributeSet
}

func (i *Item) GetAttr(attribute int) (int) {
	return i.m_attributes[attribute]
}
