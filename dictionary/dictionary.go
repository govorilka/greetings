package dictionary

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Dictionary struct {
	mu   sync.Mutex
	v    map[string]string
	rand *rand.Rand
}

func New() *Dictionary {
	return &Dictionary{
		v:    make(map[string]string),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *Dictionary) Get(name string) (string, error) {
	if name == "" {
		return "", errors.New("Value name is empty")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.v[name]
	if ok {
		return val, nil
	}

	value := strconv.Itoa(c.rand.Intn(100))
	c.v[name] = value
	return value, nil
}
