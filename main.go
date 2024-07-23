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

var heroUrl = "https://pvp.qq.com/web201605/js/herolist.json"

type Hero struct {
	CName string `json:"cname"`
	EName uint   `json:"ename"`
	Skins string `json:"skin_name"`
	Title string `json:"title"`
}

func getSkinUrl(ename uint, sn int) string {
	return fmt.Sprintf(
		"http://game.gtimg.cn/images/yxzj/img201606/skin/hero-info/%d/%d-bigskin-%d.jpg",
		ename, ename, sn,
	)
}

var client http.Client
var wg sync.WaitGroup

var rootDir = "wzry-skin"

func exists(path string) bool {
	var _, err = os.Stat(path)
	return err == nil || os.IsExist(err)
}

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
