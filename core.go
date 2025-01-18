package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type Data struct {
	Bytes []byte
	Path  string
}

var data = make(chan Data, 5)

func processHero(hero Hero) {
	var heroDir = path.Join(
		rootDir, fmt.Sprintf("%s_%s", hero.CName, hero.Title),
	)
	if !exists(heroDir) {
		os.Mkdir(heroDir, os.ModePerm)
	}

	for i, skin := range strings.Split(hero.Skins, "|") {
		i += 1
		url := getSkinUrl(hero.EName, i)
		fp := path.Join(
			heroDir,
			fmt.Sprintf("%02d_%s.jpg", i, skin),
		)

		wg.Add(2)
		go func() {
			res, _ := client.Get(url)
			buf, _ := io.ReadAll(res.Body)
			data <- Data{buf, fp}
			wg.Done()
		}()
		go func() {
			d := <-data
			file, _ := os.Create(d.Path)
			defer file.Close()
			file.Write(d.Bytes)
			fmt.Println("done: ", d.Path)
			wg.Done()
		}()

	}
	wg.Done()
}
