package main

const (
	weapon = iota
	armor
	healing
)

type Item struct {
	Entity
	item_type int
	m_min int
	m_max int
	m_speed int
	m_price Money
	m_attributes AttributeSet
}

type ItemDatabase []Item

func (i *Item) GetAttr(attribute int) (int) {
	return i.m_attributes[attribute]
}

func NewItemDatabase() (*ItemDatabase) {
	return &ItemDatabase{}
}
