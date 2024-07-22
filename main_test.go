package main

import "testing"

func TestAddition(t *testing.T) {
	got := Add(1, 3)
	want := 4
	if got != want {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestDivison(t *testing.T) {
	t.Run("Test division", func(t *testing.T) {
		got := Div(4, 2)
		want := 2
		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
	t.Run("Test division with 0", func(t *testing.T) {
    defer func() {
      if r := recover(); r == nil {
        t.Errorf("zero divsion did not panic")
      }
    }()
    Div(4, 0)
	})
}
