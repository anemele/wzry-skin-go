package wzry

import (
	"fmt"
	"log"
	"path"
)

func Run() (bool, error) {
	// 获取 []Hero
	heros, err := getData()
	if err != nil {
		return false, err
	}

	// 获取本地统计信息 statistics.txt
	statistics, err := getStat()
	if err != nil {
		return false, err
	}

	// channel
	// size 是随便设置的
	chan1 := make(chan Chan1, 1024)
	// defer close(chan1)

	// 遍历英雄列表
	for _, hero := range heros {
		// 创建该英雄的存储目录
		heroTN := hero.cname + "_" + hero.title
		heroSavePath := path.Join(saveRoot, heroTN)
		mkDir(heroSavePath)

		go func(hero Hero) {
			// 请求英雄页面
			html, err := getPage(hero.ename)
			if err != nil {
				log.Panicln(err)
				return
			}

			// 解析英雄页面，获取皮肤列表
			skins, err := parseHtml(html)
			if err != nil {
				log.Println(err)
				return
			}

			chan1 <- Chan1{hero, heroSavePath, skins}
		}(hero)
	}
	// close(chan1)

	// TODO 230806
	// 这里使用 range 读取 chan 会出现一系列问题
	// 1. 前面用 close 会直接跳过 range chan
	// 2. 前面不用 close 或者用 defer close 会导致卡死，原因不知
	// 最终选择有限循环读取。
	// for ch1 := range chan1 {
	numHero := len(heros)
	for i := 0; i < numHero; i++ {
		ch1 := <-chan1
		hero := ch1.hero
		heroSavePath := ch1.heroSavePath
		skins := ch1.skins
		// fmt.Println(skins)
		// fmt.Println(hero)
		// fmt.Println(heroSavePath)
		// continue

		// 截取皮肤列表，更新统计信息
		lenSkin := len(skins)
		lenStat := statistics[hero.cname]
		// 比对皮肤列表长度和统计信息
		if lenStat < lenSkin {
			// 统计信息记录小于皮肤列表长度，说明有更新
			// 则截取更新部分送入下载，并更新统计信息记录
			for i, skin := range skins[lenStat:] {
				skinFileName := fmt.Sprintf("%d_%s.jpg", i+lenStat+1, skin)
				skinSavePath := path.Join(heroSavePath, skinFileName)
				if exists(skinSavePath) {
					logInfo("EXISTS", skinSavePath)
					continue
				}
				skinImageUrl := getImageUrl(hero.ename, i+lenStat+1, skinSize["b"])
				go func() {
					ok, err := getSkin(skinImageUrl, skinSavePath)
					if ok {
						logInfo("SAVED", skinSavePath)
					} else {
						log.Println(err)
					}
				}()
			}
			statistics[hero.cname] = lenSkin
		} else if lenStat > lenSkin {
			// 统计信息记录大于皮肤列表长度，说明记录存在错误
			// （也可能是请求的皮肤数据错误或解析错误，可能性较小）
			// 仅更正统计信息记录
			statistics[hero.cname] = lenSkin
		}
		// 如果二者相等，说明没有更新，也没有错误，无需操作
	}

	setStat(statistics)

	log.Println("DONE")
	return true, nil
}
