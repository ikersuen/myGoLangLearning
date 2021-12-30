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

	assertionMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	assertionNewMessage := func(t testing.TB, got, want string){
		t.Helper()
		if(got == want) {
			fmt.Println("It was the same!!!!!!")
		}
	}
	t.Run("saying ABC to telnet", func(tt *testing.T){
		got := HelloInput("Iker","")
		want:= "Hello Iker"
		assertionMessage(tt,got,want)
	})
	t.Run("default print Hello Ivan if Empty Input", func(ta *testing.T){
		got := HelloInput("","")
		want:= "Hello Ivan"
		assertionMessage(ta,got,want)
	})
	t.Run("test French", func(tb *testing.T){
		got:= HelloInput("Iker","french")
		want:= "Bonjour Iker"
		assertionMessage(tb,got,want)
	})
	t.Run("test spanish", func(tc *testing.T){
		got:= HelloInput("Iker","spanish")
		want:= "Hola Iker"
		assertionMessage(tc,got,want)
		assertionNewMessage(tc,got,want)
	})
	
	

}


