package GUI

import (
	"VideoSpider/paraseURL"
	"VideoSpider/public/error_operate"
	"fmt"
	"github.com/lxn/walk"
	"strconv"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
	path         string
}

func (mw *MyMainWindow) getUserIdAndOperate() {
	users := userId.Text()

	if users == ""{
		//userId.SetText("开始运行")
		return
	}

	if strings.Contains(users,"shoudaoshuju"){
		panic(error_operate.NoticeError{"传入玩家编号出错"})
	}

	userId.SetText("")

	runStopBtn.SetText("继续运行")

	userIds := strings.Split(users, ",")

	fmt.Println(userIds)
	for _, userId := range userIds {
		go paraseURL.ParseId(userId)
	}
}

func (mw *MyMainWindow) dynamicList() {
	urls := make([]string, 0)
	for i := 0; i < 4; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoRankService.getDynamicRank/" + strNum + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoRankServicegetDynamicRank" + strNum + "jsonpcallback"
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls, "动态榜单")
}

func (mw *MyMainWindow) soaringList() {
	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getWeeklyRisingRank/" + strNum + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetWeeklyRisingRank" + strNum + "jsonpcallback"
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls, "飙升榜单")
}

func (mw *MyMainWindow) dayList() {
	urls := make([]string, 0)
	strTime := timeText.Text()
	timeText.SetText(fmt.Sprintf("爬取日期为：%s的视频", strTime))
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getDailyRank/"+strNum+"-"+strTime+"/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetDailyRank"+strNum+strTime+"jsonpcallback"
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls,"日榜")
}

func (mw *MyMainWindow) mouthList() {
	strTime := timeText.Text()
	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getMonthRank/"+strNum+"-"+strTime+"/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetMonthRank"+strNum+strTime+"jsonpcallback"
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls, "月榜")
}

func (mw *MyMainWindow) offlinePauseRecover() {
	//users := userId.Text()
	mw.Close()
	userId.SetText("")

}
