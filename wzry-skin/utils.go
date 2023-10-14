package wzry

import (
	"log"
	"os"
	"regexp"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func MkDir(path string) string {
	if !Exists(path) {
		os.Mkdir(path, os.ModePerm)
		LogInfo("MKDIR", path)
	}
	return path
}

func SplitSkin(skins string) []string {
	re, _ := regexp.Compile(`(\S+?)[\s(?:&\d+)\|]+`)
	matches := re.FindAllStringSubmatch(skins, -1)
	ret := make([]string, len(matches))
	for i, match := range matches {
		ret[i] = match[1]
	}
	return ret
}

func LogInfo(action string, message string) {
	log.Printf("%-8s%s\n", action, message)
}
