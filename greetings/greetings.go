package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var formats = []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
}

type Values struct {
	mu   sync.Mutex
	v    map[string]string
	rand *rand.Rand
}

func New() *Values {
	return &Values{
		v:    make(map[string]string),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *Values) Get(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.v[name]
	if ok {
		return val, nil
	}

	format := formats[c.rand.Intn(len(formats))]
	message := fmt.Sprintf(format, name)
	c.v[name] = message
	return message, nil
}
