package wzry

import (
	"fmt"
	"path"
)

func getData() []Hero {
	bytes := getBytes(apiDataUrl)
	if bytes == nil {
		return nil
	}

	return parseJson(bytes)
}

func download(heros []Hero) {

	countAll := 0
	countAdd := 0
	for _, hero := range heros {
		heroTN := hero.cname + "_" + hero.title
		heroSavePath := path.Join(saveRoot, heroTN)
		mkDir(heroSavePath)

		heroPageUrl := getPageUrl(hero.ename)
		pageBytes := getBytes(heroPageUrl)
		html := convertGbkToUtf8(pageBytes)
		skins := parseHtml(html)
		skinList := splitSkin(skins)
		skinList = skinList[:len(skinList)-1]

		for i, skin := range skinList {
			i += 1
			// fmt.Println(i, skin)

			countAll += 1

			add := 0
			for s, a := range skinSize {
				skinSaveName := fmt.Sprintf("%s_%d_%s.jpg", s, i, skin)
				skinSavePath := path.Join(heroSavePath, skinSaveName)

				if exists(skinSavePath) {
					fmt.Printf("\rgetting %d", countAll)
					continue
				}

				skinUrl := getImageUrl(hero.ename, i, a)
				bytes := getBytes(skinUrl)
				writeBytes(bytes, skinSavePath)
				fmt.Printf("\r[INFO] save %s\n", skinSavePath)

				add += 1
			}

			if add > 0 {
				countAdd += 1
			}
		}
	}
	fmt.Printf("\r[INFO] add:%d all:%d\n", countAdd, countAll)
}

func Run() {
	mkDir(saveRoot)
	heros := getData()
	download(heros)
}
