package search

import (
	_ "log"
	_ "sync"
)

// 注册用于搜索的匹配器的映射
// 创建了一个
var matchers=make(map[string]Matcher)