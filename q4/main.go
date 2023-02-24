// create and use WaitGroups
package main

import (
	"fmt"
	"sync"
)

func friend1(wg * sync.WaitGroup){
	fmt.Println("Hi my name is Ariel and Arura is my friend.")
	wg.Done()
}
func friend2(wg * sync.WaitGroup){
	fmt.Println("Hi my name is Tiana and I'm Arura's friend")
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	go friend1(&wg)
	go friend2(&wg)

	wg.Wait()

	fmt.Println("Hi my name is Arura and and my friends introduced themselves eariler.")
	wg.Wait()
}