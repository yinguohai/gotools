# gotools
this is my tools for go 

##函数详解

#### list 合并

```
/**
    将两个string 类型的 list 合并为一个map
    []string{'name','age','sex'}  +  []string['Bob','20','M']  => map[string]string {"name":Bob","age":"20","sex":"M"}
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

```

#### get column
```
/**
    获取Map中获取指定col的所有值,例如提取所有的name
    data := []map[string]string {
        map[string]stirng{"name":"Bob","age":"20","sex":"F"},
        map[string]stirng{"name":"Smith","age":"22","sex":"M"},
        map[string]stirng{"name":"Hu","age":"30","sex":"F"},
    }
    变为
    ["Bob","Smith","Hu"]
    
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
```

#### channel 判断所有 goroutine 是否都执行完了

```
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
```