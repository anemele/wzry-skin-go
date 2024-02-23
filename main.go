package main

import (
	"consts"
	"fmt"
	"local"
	"log/slog"
	"path"
	"sync"
	"utils"
	"web"
)

var logger = slog.Default()

func getHeroPage(wg *sync.WaitGroup, chan1 chan web.Chan1, hero web.Hero) {
	defer wg.Done()

	// 创建该英雄的存储目录
	heroTN := hero.Cname + "_" + hero.Title
	heroSavePath := path.Join(local.SaveRoot, heroTN)
	utils.MkDir(heroSavePath)
	// 请求英雄页面
	html, err := web.GetPage(hero.Ename)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// 解析英雄页面，获取皮肤列表
	skins, err := web.ParseHtml(html)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	chan1 <- web.Chan1{Hero: hero, Path: heroSavePath, Skins: skins}
}

func getSkinBytes(wg *sync.WaitGroup, chan2 chan web.Chan2, url, path string) {
	defer wg.Done()

	bytes, err := web.GetBytes(url)
	if err == nil {
		chan2 <- web.Chan2{Content: bytes, Path: path}
	} else {
		logger.Error(err.Error())
	}
}

func saveSkin(wg *sync.WaitGroup, bytes []byte, path string) {
	defer wg.Done()

	ok, err := web.WriteBytes(bytes, path)
	if ok {
		logger.Info("SAVE", "path", path)
	} else {
		logger.Error(err.Error())
	}
}

func main() {
	// 获取 []Hero
	heros, err := web.GetData()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// 获取本地统计信息 statistics.txt
	statistics, err := local.GetStat()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// size 是随便设置的
	const size = 256

	var chan1 = make(chan web.Chan1, size)
	var chan2 = make(chan web.Chan2, size)

	// 遍历英雄列表
	go func() {
		// 创建同步锁
		wg := &sync.WaitGroup{}
		defer close(chan1)
		wg.Add(len(heros))
		for _, hero := range heros {
			go getHeroPage(wg, chan1, hero)
		}
		wg.Wait()
	}()

	// 遍历皮肤列表
	go func() {
		// 创建同步锁
		wg := &sync.WaitGroup{}
		defer close(chan2)

		for ch1 := range chan1 {
			hero := ch1.Hero
			heroSavePath := ch1.Path
			skins := ch1.Skins

			// 截取皮肤列表，更新统计信息
			lenSkin := len(skins)
			lenStat := statistics[hero.Cname]
			// 比对皮肤列表长度和统计信息
			if lenStat < lenSkin {
				// 统计信息记录小于皮肤列表长度，说明有更新
				// 则截取更新部分送入下载，并更新统计信息记录
				for i, skin := range skins[lenStat:] {
					skinFileName := fmt.Sprintf("%d_%s.jpg", i+lenStat+1, skin)
					skinSavePath := path.Join(heroSavePath, skinFileName)
					if utils.Exists(skinSavePath) {
						logger.Warn("EXIST", "", skinSavePath)
						continue
					}
					skinImageUrl := consts.GetImageUrl(hero.Ename, i+lenStat+1, consts.SkinSize["b"])
					wg.Add(1)
					go getSkinBytes(wg, chan2, skinImageUrl, skinSavePath)

				}
				statistics[hero.Cname] = lenSkin
			} else if lenStat > lenSkin {
				// 统计信息记录大于皮肤列表长度，说明记录存在错误
				// （也可能是请求的皮肤数据错误或解析错误，可能性较小）
				// 仅更正统计信息记录
				statistics[hero.Cname] = lenSkin
			}
			// 如果二者相等，说明没有更新，也没有错误，无需操作

		}
		wg.Wait()
	}()

	// 创建同步锁
	wg := &sync.WaitGroup{}
	for ch2 := range chan2 {
		wg.Add(1)
		go saveSkin(wg, ch2.Content, ch2.Path)
	}

	wg.Wait()
	local.SetStat(statistics)
	logger.Info("DONE!")

}
