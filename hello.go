package main

import (
	"fmt"
	"github.com/gofrs/uuid"
)


func IkerTalk() string {
	return "Hello Iker"
}

func HelloInput(s string) string {
	return "ABC" + s
}

func main(){
	fmt.Println(HelloInput("ping"))
	fmt.Println(IkerTalk())
	fmt.Println(uuid.Must(uuid.NewV4()).String())
}