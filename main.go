package main

import (
	"sync"
	"wzry"
)

func main() {
	var wg sync.WaitGroup
	wzry.Run(&wg)
	wg.Wait()
}
