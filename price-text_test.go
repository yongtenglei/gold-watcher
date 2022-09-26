package main

import "testing"

func TestAPP_getPriceText(t *testing.T) {

	open, current, change := testApp.getPriceText()
	if open.Text != "Open: 382.5082 CNY/g" {
		t.Errorf("OpenText error: expect %s, get %v", "Open: 382.5082 CNY/g", open.Text)
	}
	if current.Text != "Current: 385.4356 CNY/g" {
		t.Errorf("CurrentText error: expect %s, get %v", "Current: 385.4356 CNY/g", current.Text)
	}
	if change.Text != "Change: 2.9273 CNY/g" {
		t.Errorf("ChangeText error: expect %s, get %v", "Change: 2.9273 CNY/g", change.Text)
	}
}
