package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

const defaultSavePath = "wzry-skin"
const savePathFile = "savepath.txt"

func getSavePath() string {
	file, err := os.Open(savePathFile)
	if err != nil {
		return defaultSavePath
	}
	buf, err := io.ReadAll(file)
	if err != nil {
		return defaultSavePath
	}
	return string(buf)
}

type Data struct {
	Bytes []byte
	Path  string
}

var data = make(chan Data, 5)

func processHero(savePath string, hero Hero, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	var heroDir = path.Join(
		savePath, fmt.Sprintf("%s_%s", hero.CName, hero.Title),
	)
	os.Mkdir(heroDir, os.ModePerm)

	for i, skin := range strings.Split(hero.Skins, "|") {
		i += 1
		url := getSkinUrl(hero.EName, i)
		fp := path.Join(
			heroDir,
			fmt.Sprintf("%02d_%s.jpg", i, skin),
		)

		wg.Add(2)
		go func() {
			defer wg.Done()
			res, _ := client.Get(url)
			buf, _ := io.ReadAll(res.Body)
			data <- Data{buf, fp}
		}()
		go func() {
			defer wg.Done()
			d := <-data
			file, _ := os.Create(d.Path)
			defer file.Close()
			file.Write(d.Bytes)
			fmt.Println("done: ", d.Path)
		}()

	}
}

func main() {
	savePath := getSavePath()
	os.MkdirAll(savePath, os.ModePerm)

	var client http.Client
	var resp, err = client.Get(heroUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	var herolist []Hero
	err = json.NewDecoder(resp.Body).Decode(&herolist)
	if err != nil {
		fmt.Println("failed to parse hero data:", err)
		return
	}
	// fmt.Println(herolist)

	var wg sync.WaitGroup
	wg.Add(len(herolist))
	for _, hero := range herolist {
		go processHero(savePath, hero, &client, &wg)
	}

	wg.Wait()
}
