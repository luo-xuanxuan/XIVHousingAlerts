package main

import (
	"fmt"
)

func main() {
	world := 79

	worldData, err := GetWorld(world)
	if err != nil {
		panic(err)
	}

	fmt.Println(worldData.Name)
}
