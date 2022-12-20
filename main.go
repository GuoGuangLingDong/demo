package main

import (
	"context"
	"demo/internal/dao"
	_ "demo/internal/logic"
	"demo/internal/task"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"time"

	"demo/internal/cmd"
)

func main() {
	ctx := context.Background()
	_, err := gcron.Add(ctx, "0 0 0 * * *", func(ctx context.Context) {
		// 设置进程全局时区
		uids := ([]string)(nil)
		items, _ := dao.User.Ctx(ctx).Fields("DISTINCT uid").Array()
		for _, item := range items {
			uids = append(uids, item.String())
		}
		curTime := gtime.Now()
		overDueTime := curTime.Add(time.Duration(32) * time.Hour) //gtime时区bug，加32h数据库里就是24h
		fmt.Println("over: ", overDueTime)
		data := g.List{}
		for _, uid := range uids {
			data = append(data, g.Map{
				"uid":        uid,
				"opt_type":   7,
				"score":      20,
				"create_at":  curTime,
				"overdue_at": overDueTime,
			})
		}
		_, err := dao.Operation.Ctx(ctx).Data(data).Insert()
		if err != nil {
			fmt.Println("定时赠送积分失败")
		} else {
			fmt.Println("定时赠送积分成功")
		}
	}, "GiveScoreJob")
	if err != nil {
		fmt.Println("定时任务启动失败")
	}
	go task.Mint()

	cmd.Main.Run(gctx.New())
}
