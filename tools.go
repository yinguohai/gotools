package gotools

import (
	"fmt"
	"time"
)

/**
	将两个string 类型的 list 合并为一个map
 */
func ListToMap(Keys, Values []string) (*map[string]string, error) {
	if len(Keys) != len(Values) {
		return nil, fmt.Errorf("length of Keys not equals length of Values %v", nil)
	}
	content := make(map[string]string)
	for i := 0; i < len(Keys); i++ {
		content[Keys[i]] = Values[i]
	}
	return &content, nil
}

/**
	获取Map 数组中提取出col 指定的所有字段
 */
func ColumeMap(params []map[string]string, col string) []string {
	result := []string{}
	for _, m := range params {
		for k, v := range m {
			if k == col {
				result = append(result, v)
			}
		}
	}
	return result
}

/**
	逻辑线程是否都执行完了
 */
func IsDoneChan(ch chan bool, num int, timeout int) bool {
	flag := 0

	for {
		if timeout == 0 {

			select {
			case <-ch:
				flag = flag + 1
			}
		} else {
			select {
			case <-ch:
				flag = flag + 1
			case <-time.After(time.Second * time.Duration(timeout)):
				return false
			}
		}

		if flag >= num {
			return true
		}

	}
}
