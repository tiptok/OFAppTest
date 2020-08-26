package constant

import "os"

// 工作目录
var DSL_PATH = "F://go//src//learn_project//ddd-project//stock1"

// 保存路径
var SAVE_PATH = "F://go//src//stock1"

func init() {
	if os.Getenv("DSL_PATH") != "" {
		DSL_PATH = os.Getenv("DSL_PATH")
	}
	if os.Getenv("SAVE_PATH") != "" {
		SAVE_PATH = os.Getenv("SAVE_PATH")
	}
}
