package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {

	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 385.43556648303706 {
		t.Errorf("Wrong prince returned, expect: %v, get: %v", 385.43556648303706, p.Price)
	}

}
