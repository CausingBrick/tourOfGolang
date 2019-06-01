package main

import (
	"fmt"
	"log"
	"os"
	"tourOfGolang/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/*
// !+ textoutput
$ go run main.go repo:shadowsocks/shadowsocks/

1422 issues:
#1460     SunR54 SS端口和SSH端口有什么联系吗
#1459       uebb 为什么SSR成功连接 没办法打开youtube
#1458  walkingaw 免费ss账号，长期更新
#1457   xiaotana super Wingy
#1456   xiaotana super Wingy
#1455   koslofse Update CONTRIBUTING.md
#1454  usernamel Where is the code?
#1453     monscn 胖虎的项目也fork了
#1452  yinlang83 请问一下我上GitHub时，需要切换到全局模式，是否需要在pac文件下增加或者修改其他内容呢？
#1451  Tyson1077 py版，当我以后台方式启动时，死活都无法正常使用
#1450    parcool 公司屏蔽了网易云音乐，用全局模式可以听，但是国内网站就慢了，咋整？
#1449    JummyWu 小米路由器
#1446  SykieChen About TCP RST packets
#1445  BigTable1 config one shadowssocks server as another one's server/
#1444  wang11035 请勿以身试法
#1441      xhqpp 󠀡
#1440     techYJ Merge pull request #1 from shadowsocks/master
#1439  DavidK1ng failed to config multi-user
#1438  anthony14 Single port for multiple users
#1437  anthony14 Single port for multiple users
#1436  cary-zhou 有没有大牛能解释一下SpeedTester测速类的原理
#1434  fuck1part 哈哈
#1433      ysdlb setsocketopt(socket.SOL_TCP, 23, 5)中的23导致在不同的平台上出现不同的结果
#1432      fwdfn [WinError 10049] 在其上下文中，该请求的地址无效
#1431    luocaca 腾讯云服务器 局域网代理问题
#1430  cary-zhou 会不会有性能问题？
#1429      xhqpp 󠀡
#1428     LitPan 新人求解，服务器端无法启动shadowsocks，already started at pid 4581
#1427  WatanabeM ubuntu18.04, ss 3.0.0, 无法开机自启动
#1426  672110619 会不会出现长时间使用就会端口被封的情况？
// !- textoutput
*/
