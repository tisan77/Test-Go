package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	rectangle := Rectangle{10, 5}
	got := rectangle.Area()
	want := 50.0
	if got != want {
		t.Errorf("Area() = %v; want %v", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 5}
	got := rectangle.Perimeter()
	want := 30.0
	if got != want {
		t.Errorf("Perimeter() = %v; want %v", got, want)
	}
}

func TestIsSquare(t *testing.T) {
	rectangle := Rectangle{10, 5}
	got := rectangle.IsSquare()
	want := false
	if got != want {
		t.Errorf("IsSquare() = %v; want %v", got, want)
	}
}

func TestIsSquareWhenSquare(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := rectangle.IsSquare()
	want := true
	if got != want {
		t.Errorf("IsSquare() = %v; want %v", got, want)
	}
}
