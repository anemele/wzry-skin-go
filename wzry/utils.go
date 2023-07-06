package wzry

import (
	"os"
	"regexp"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func mkDir(path string) {
	if !exists(path) {
		os.Mkdir(path, os.ModePerm)
	}
}

type Hero struct {
	cname string
	ename int
	title string
}

func splitSkin(skins string) []string {
	re, _ := regexp.Compile(`[(?:&\d+)\|]+`)
	return re.Split(skins, -1)
}
