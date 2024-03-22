package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

const tapMinPrice = 1
const tapMaxPrice = 500

type tapID uuid.UUID

func (id tapID) String() string {
	return uuid.UUID(id).String()
}

// tap represents a tap and its price
type tap struct {
	id    tapID
	price int
}

func newTap() *tap {
	return &tap{
		id:    tapID(uuid.New()),
		price: rand.Intn(tapMaxPrice-tapMinPrice) + tapMinPrice,
	}
}

func (t *tap) String() string {
	return fmt.Sprintf("Tap %s, price: %v", t.id, t.price)
}

func createRandomTaps(n int) []*tap {
	taps := make([]*tap, n)
	for i := range taps {
		t := newTap()
		fmt.Println(t)
		taps[i] = t
	}
	return taps
}
