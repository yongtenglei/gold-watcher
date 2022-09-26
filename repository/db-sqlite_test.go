package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed: ", err)
	}
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed: ", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back: ", err)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	holdings, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed: ", err)
	}

	if len(holdings) != 1 {
		t.Error("wrong number of rows returned, expect 1, get ", len(holdings))
	}
}

func TestSQLiteRepository_GetHoldingByID(t *testing.T) {
	holding, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("get by id 1 failed: ", err)
	}

	if holding.PurchasePrice != 1000 {
		t.Errorf("wrong purchase price returned; expected 1000 but %d returned", holding.PurchasePrice)
	}
}

func TestSQLiteRepository_UpdateHolding(t *testing.T) {
	holding, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	holding.PurchasePrice = 1001

	err = testRepo.UpdateHolding(1, *holding)
	if err != nil {
		t.Error("update failed: ", err)
	}
}

func TestSQLiteRepository_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("failed to delete holding: ", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}
}
