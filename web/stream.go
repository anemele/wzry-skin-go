package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var client *http.Client = &http.Client{}

func GetBytes(url string) ([]byte, error) {
	logger.Info("GET", "url", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("Referer", "https://pvp.qq.com/")

	res, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("%d", res.StatusCode)
		logger.Error(err.Error(), "status", res.Status)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return body, nil
}

func WriteBytes(bytes []byte, path string) (bool, error) {
	file, err := os.Create(path)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	file.Write(bytes)
	return true, nil
}
