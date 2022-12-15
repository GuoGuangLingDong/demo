package main

import (
	"context"
	"demo/internal/dao"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"time"

	_ "demo/internal/logic"

	"demo/internal/cmd"
)

func main() {
	ctx := context.Background()
	_, err := gcron.Add(ctx, "0 0 0 * * *", func(ctx context.Context) {
		uids := ([]string)(nil)
		items, _ := dao.User.Ctx(ctx).Fields("DISTINCT uid").Array()
		for _, item := range items {
			uids = append(uids, item.String())
		}
		curTime := gtime.Now()
		overDueTime := curTime.Add(time.Duration(24) * time.Hour)
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
			fmt.Println("定时赠送积分时插入数据库失败")
		}
	}, "GiveScoreJob")
	if err != nil {
		fmt.Println("定时赠送积分失败")
	} else {
		fmt.Println("定时赠送积分成功")
	}
	cmd.Main.Run(gctx.New())
}
