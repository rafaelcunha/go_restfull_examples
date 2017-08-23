package main

import "sync"

// Counter adds a value and returns a new value
type Counter interface {
	Add(int) int
}

type countService struct {
	v  int
	mu sync.Mutex
}

func (c *countService) Add(v int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v += v
	return c.v
}

// ---   Public Key   ---
type PublicKey interface {
	getPublickKey() string
}

type pkService struct {
	const pk string = "APP_USR-d509bfad-d337-4619-ab22-21fd59541fb4"
}

func (pblicKey *pkService) getPublickKey() string {
	return pblicKey.pk
}