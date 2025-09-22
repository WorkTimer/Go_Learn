package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Go 学习项目 ===")
	fmt.Println()

	fmt.Println("1. 运行 SQLX 示例...")
	SqlxExample()
	fmt.Println()

	fmt.Println("2. 运行 GORM 示例...")
	GormExample()
	fmt.Println()

	fmt.Println("所有示例运行完成！")
}
