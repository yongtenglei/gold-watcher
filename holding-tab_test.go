package main

import (
	"testing"
)

func TestConfig_currentHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get current holdings from database: ", err)
	}

	if len(all) != 2 {
		t.Errorf("wrong number of holdings returnd; expected 2, get %d", len(all))
	}

}

func TestConfig_getHoldingSlice(t *testing.T) {
	slice := testApp.getHoldingSlice()
	if len(slice) != 3 {
		t.Errorf("wrong number of rows returned; expected 3, get %d", len(slice))
	}
}

func TestConfig_getHoldingsTable(t *testing.T) {

}

func TestConfig_holdingsTab(t *testing.T) {

}
