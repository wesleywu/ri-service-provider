package boot

import (
	"fmt"
	"github.com/WesleyWu/ri-service-provider/gowing"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	_ = gtime.SetTimeZone("Asia/Shanghai") //设置系统时区
	showLogo()
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
}

func showLogo() {
	fmt.Println("   _______       _______   ________\n  / ____/ |     / /  _/ | / / ____/\n / / __ | | /| / // //  |/ / / __  \n/ /_/ / | |/ |/ // // /|  / /_/ /  \n\\____/  |__/|__/___/_/ |_/\\____/   ")
	fmt.Println("当前版本：", gowing.Version)
}
