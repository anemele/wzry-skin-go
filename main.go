package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var client http.Client
var wg sync.WaitGroup

var rootDir = "wzry-skin"

func main() {
	if !exists(rootDir) {
		os.Mkdir(rootDir, os.ModePerm)
	}

	var res, err = client.Get(heroUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to read remote data:", err)
		return
	}
	// fmt.Println(buf)

	var herolist []Hero
	err = json.Unmarshal(buf, &herolist)
	if err != nil {
		fmt.Println("failed to parse hero data:", err)
		return
	}
	// fmt.Println(herolist)

	for _, hero := range herolist {
		wg.Add(1)
		go processHero(hero)
	}

	wg.Wait()
}
