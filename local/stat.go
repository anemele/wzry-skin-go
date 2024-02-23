package local

import (
	"bufio"
	"consts"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"utils"
)

func GetStat() (map[string]int, error) {
	ret := make(map[string]int)

	if !utils.Exists(consts.StatFile) {
		return ret, nil
	}

	file, err := os.Open(consts.StatFile)
	if err != nil {
		logger.Error(err.Error())
		return ret, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		x := strings.Split(string(line), ":")
		if len(x) == 2 {
			num, err := strconv.Atoi(x[1])
			if err == nil && num >= 0 {
				ret[x[0]] = num
			}
		}
	}
	return ret, nil
}

func SetStat(data map[string]int) (bool, error) {
	if data == nil {
		return false, nil
	}

	file, err := os.OpenFile(consts.StatFile, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for key, val := range data {
		content := fmt.Sprintf("%s:%d\n", key, val)
		_, err := writer.WriteString(content)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	ret := writer.Flush()

	return ret == nil, ret
}
