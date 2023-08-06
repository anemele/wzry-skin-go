package wzry

import (
	"log"
	"os"
	"regexp"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func mkDir(path string) string {
	if !exists(path) {
		os.Mkdir(path, os.ModePerm)
		logInfo("MKDIR", path)
	}
	return path
}

func splitSkin(skins string) []string {
	re, _ := regexp.Compile(`(\S+?)[\s(?:&\d+)\|]+`)
	matches := re.FindAllStringSubmatch(skins, -1)
	ret := make([]string, len(matches))
	for i, match := range matches {
		ret[i] = match[1]
	}
	return ret
}

func logInfo(action string, message string) {
	log.Printf("%-8s%s\n", action, message)
}
