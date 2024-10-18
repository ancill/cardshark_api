package main

func NewInMemoryDeckStore() *InMemoryDeckStore {
	return &InMemoryDeckStore{map[string]int{}}
}

type InMemoryDeckStore struct {
	store map[string]int
}

func (i *InMemoryDeckStore) RecordDeck(deck string) {
	i.store[deck]++
}

func (i *InMemoryDeckStore) GetDeckSize(deck string) int {
	return i.store[deck]
}
