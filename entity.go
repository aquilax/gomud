package main

import (
	"strings"
	"sort"
)

type EntityDatabase []Entity

type Entity struct {
	id int
	name string
}


//Entry Methods

func (e *Entity) Name() (string){
	return e.name
}

func (e *Entity) ID() (int) {
	return e.id
}

func (e *Entity) CompName() (string) {
	return strings.ToLower(e.name)
}

func (e *Entity) FullMatch (name string) (bool) {
	return strings.ToLower(name) == e.CompName()
}

func (e *Entity) Match (name string) (bool) {
	lname := strings.ToLower(name)
	words := strings.Fields(e.CompName())
	for _, word := range words {
		if strings.HasPrefix(word, lname) {
			return true
		}
	}
	return false
}

// Entry Database search/sort interface
func (ed EntityDatabase) Len() (int) {
	return len(ed)
}

func (ed EntityDatabase) Less(i, j int) (bool) {
	return ed[i].CompName() < ed[j].CompName()
}

func (ed EntityDatabase) Swap(i, j int) {
	ed[i], ed[j] = ed[j], ed[i]
}

func (ed EntityDatabase) Search(name string) (int) {
	i := sort.Search(len(ed), func(i int) bool {
		return ed[i].FullMatch(name) || ed[i].Match(name);
	})
	return i
}
