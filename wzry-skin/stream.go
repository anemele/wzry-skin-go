package wzry

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var client *http.Client = &http.Client{}

func GetBytes(url string) ([]byte, error) {
	LogInfo("GET", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("Referer", "https://pvp.qq.com/")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func WriteBytes(bytes []byte, path string) (bool, error) {
	file, err := os.Create(path)
	if err != nil {
		return false, err
	}

	file.Write(bytes)
	return true, nil
}
