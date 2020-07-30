package cmd

import (
	"github.com/spf13/cobra"
	"github.com/test/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

//创建项目的time子命令，方便将internal/timer/time.go里获得时间和推算的时间集成到time子命令中。
var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

//增加now子命令，用于获取当前时间，用于处理具体逻辑
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run:   func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"),
			nowTime.Unix())
	},
}

//新增calc命令，用于推算时间
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run:   func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == ""{
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			//使用strings.Contains对空格进行包含判断
			if !strings.Contains(calculateTime," ") {
				//没有空格则输出下面的格式
				layout = "2006-01-02"
			}
			//有空格就输出“2006-01-02 15:04:05”这种格式
			currentTimer, err = time.Parse(layout, calculateTime)
			//出现异常，按时间戳格式进行转换处理
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果: %s, %d", calculateTime.Format(layout),calculateTime.Unix())
	},
}


func init(){
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c","",
		`需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration,"duration","d", "",
		`持续时间，有效时间单位为"ns","us","ms","s","m","h"`)
}