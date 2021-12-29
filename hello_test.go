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

	message := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying ABC to telnet", func(t *testing.T){
		got := HelloInput("Iker","")
		want:= "Hello Iker"
		message(t,got,want)
	})
	t.Run("default print Hello Ivan if Empty Input", func(t *testing.T){
		got := HelloInput("","")
		want:= "Hello Ivan"
		message(t,got,want)
	})
	t.Run("test French", func(t *testing.T){
		got:= HelloInput("Iker","french")
		want:= "Bonjour Iker"
		message(t,got,want)
	})
	t.Run("test spanish", func(t *testing.T){
		got:= HelloInput("Iker","spanish")
		want:= "Hola Iker"
		message(t,got,want)
	})

}


