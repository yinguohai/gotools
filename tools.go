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

/*
	判断所有的协程是否都运行完了，
	@param ch 是一个bool类型的channel , 它里面的值需要运行的协程运行完后就要往里面添加一个true
	@param num 记录哪些goroutine 使用了 ch
	@param timeout 表示超时时长
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
