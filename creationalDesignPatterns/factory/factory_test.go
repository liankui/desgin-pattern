package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

/*
Gun: AK47 gun
Power: 4
Gun: Maverick gun
Power: 5
*/
