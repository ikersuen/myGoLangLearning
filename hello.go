package main

import (
	"fmt"
	"github.com/gofrs/uuid"
)


func IkerTalk() string {
	return "Hello Iker"
}

const englishHelloPrefix = "Hello "
const frenchHelloPrefix = "Bonjour "
const spanishHelloPrefix = "Hola "
const spanish = "spanish"
const french = "french"


func HelloInput(s string, n string) string {
	if (s==""){return "Hello Ivan"}
	prefix := englishHelloPrefix
	switch n {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + s
}

func main(){
	fmt.Println(HelloInput("ping", "french"))
	fmt.Println(IkerTalk())
	fmt.Println(uuid.Must(uuid.NewV4()).String())
}