package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

const orgMinBudget = 100
const orgMaxBudget = 600

type orgID uuid.UUID

func (id orgID) String() string {
	return uuid.UUID(id).String()
}

// org represents an organization and its budget
type org struct {
	id     orgID
	budget int
}

func newOrg() *org {
	return &org{
		id:     orgID(uuid.New()),
		budget: rand.Intn(orgMaxBudget-orgMinBudget) + orgMinBudget,
	}
}

func (o *org) String() string {
	return fmt.Sprintf("Org %s, budget: %v", o.id.String(), o.budget)
}

func createRandomOrgs(n int) []*org {
	orgs := make([]*org, n)
	for i := range orgs {
		o := newOrg()
		fmt.Println(o)
		orgs[i] = o
	}
	return orgs
}
