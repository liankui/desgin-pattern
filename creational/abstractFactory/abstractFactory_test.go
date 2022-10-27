package abstractFactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	nikeFactory, _ := getSportsFactory("nike")
	adidasFactory, _ := getSportsFactory("adidas")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

/*
Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
*/
