package main

import (
	"flag"
	"fmt"
	"os"

	"stack-bm/internal/config"
	"stack-bm/internal/database"
	"stack-bm/internal/service/mkt/sync"
)

// 广告数据定时同步命令（供 Linux cron 调用）
// 用法: go run cmd/sync/main.go --channel=tt --level=account --account=0
func main() {
	channel := flag.String("channel", "", "渠道标识(tt/tc/bili/ks)")
	level := flag.String("level", "account", "层级(account/campaign/unit/creative)")
	account := flag.String("account", "0", "渠道账户ID，0为全部")
	flag.Parse()

	if *channel == "" {
		fmt.Println("缺少 --channel 参数")
		os.Exit(1)
	}

	config.LoadConfig()
	database.InitDB()
	database.InitMongo()
	database.InitRedis()

	runner := sync.NewRunner()
	if err := runner.Run(*channel, *level, *account); err != nil {
		fmt.Println("同步失败:", err)
		os.Exit(1)
	}
	fmt.Println("同步完成")
}