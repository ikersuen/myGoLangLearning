package main

import (
	"fmt"
	"testing"
)

func TestLove(t *testing.T){
	got := IkerTalk()
	want := "Love is True"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	} else {
		fmt.Println("Damn")
	}
}

func TestHello(t *testing.T) {
	got := IkerTalk()
	want := "Hello Iker"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


func TestInput(t *testing.T) {
	t.Run("saying ABC to telnet", func(t *testing.T){
		got := HelloInput("telnet")
		want:= "ABCtelnet"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("default print Hello Ivan if Empty Input", func(t *testing.T){
		got := HelloInput("")
		want:= "Hello Ivan"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}


