package main

import "testing"

func TestJewelryStoreAddAndFind(t *testing.T) {
	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "Ring A", Kind: "Ring", Price: 100})
	store.AddItem(Jewelry{Name: "Necklace A", Kind: "Necklace", Price: 200})

	item, ok := store.FindByName("ring a")
	if !ok || item.Name != "Ring A" {
		t.Fatalf("expected to find Ring A")
	}
}

func TestJewelryStoreRemove(t *testing.T) {
	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "Ring B", Kind: "Ring", Price: 100})
	removed := store.RemoveItem("Ring B")
	if !removed {
		t.Fatalf("expected item to be removed")
	}
	if _, ok := store.FindByName("Ring B"); ok {
		t.Fatalf("expected item to be missing after removal")
	}
}

func TestJewelryStoreItemsByMinPrice(t *testing.T) {
	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "A", Kind: "Ring", Price: 100})
	store.AddItem(Jewelry{Name: "B", Kind: "Necklace", Price: 250})
	store.AddItem(Jewelry{Name: "C", Kind: "Earrings", Price: 300})

	result := store.ItemsByMinPrice(200)
	if len(result) != 2 {
		t.Fatalf("expected 2 items, got %d", len(result))
	}
}

func TestJewelryStoreTotalValue(t *testing.T) {
	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "A", Kind: "Ring", Price: 100})
	store.AddItem(Jewelry{Name: "B", Kind: "Necklace", Price: 200})
	if store.TotalValue() != 300 {
		t.Fatalf("expected total value 300, got %.2f", store.TotalValue())
	}
}

func TestJewelryStoreSell(t *testing.T) {
	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "A", Kind: "Ring", Price: 100})
	store.AddItem(Jewelry{Name: "B", Kind: "Necklace", Price: 250})

	if !store.Sell("A") {
		t.Fatalf("expected sell to succeed")
	}
	if store.Sales() != 100 {
		t.Fatalf("expected sales 100, got %.2f", store.Sales())
	}
	if _, ok := store.FindByName("A"); ok {
		t.Fatalf("expected sold item to be removed")
	}
}

func TestJewelryStoreStringNotEmpty(t *testing.T) {
	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "A", Kind: "Ring", Price: 100})
	if store.String() == "" {
		t.Fatalf("expected non-empty string")
	}
}
