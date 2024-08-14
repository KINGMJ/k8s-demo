package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取环境变量 "TASK_FLAG"
	taskFlag := os.Getenv("TASK_FLAG")
	// 调用初始化任务
	err := initializeTask(taskFlag)
	if err != nil {
		log.Printf("任务执行失败: %v", err)
		os.Exit(1) // 返回非零退出码，表示失败
	}
	fmt.Println("任务执行成功")
	// 正常退出，返回零退出码
	os.Exit(0)
}

// 初始化任务
func initializeTask(flag string) error {
	// 模拟任务逻辑
	if flag == "fail" {
		// 模拟任务失败
		return errors.New("初始化任务失败")
	}
	// 模拟任务成功
	return nil
}
