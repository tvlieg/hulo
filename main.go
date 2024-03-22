package main

import "fmt"

func recurse(taps []*tap, orgs []*org) (mapping map[tapID][]orgID) {
	if len(taps) == 0 || len(orgs) == 0 {
		return make(map[tapID][]orgID, 0)
	}

	o := orgs[0]
	t := taps[0]

	fmt.Printf("Tap %v still needs %v investments\n", t.id, t.price)
	if t.price > o.budget {
		fmt.Printf("Org %v invests tap %v\n", o.id, o.budget)
		t.price -= o.budget
		mapping = recurse(taps, orgs[1:])
	} else {
		fmt.Printf("Org %v invests tap %v\n", o.id, t.price)
		o.budget -= t.price
		mapping = recurse(taps[1:], orgs)
	}

	if _, ok := mapping[t.id]; !ok {
		mapping[t.id] = make([]orgID, 0)
	}
	mapping[t.id] = append(mapping[t.id], o.id)

	return mapping
}

func main() {
	taps := createRandomTaps(6)
	orgs := createRandomOrgs(4)
	mapping := recurse(taps, orgs)

	// Print mapping
	for t, orgIDs := range mapping {
		fmt.Printf("\nTap: %v\n", t.String())
		fmt.Printf("Organizations:\n")
		for _, o := range orgIDs {
			fmt.Printf("- %v\n", o.String())
		}
	}
}
