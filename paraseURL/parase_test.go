package paraseURL

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type testRespon struct {
	ServerTime int      `json:"servertime"`
	Callback   []string `json:"callback"`
	data       string   `json:"data"`
	Status     int      `json:"status"`
	ErrorCode  string   `json:"errorcode"`
	ErrorNo    int      `json:"errorno"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
	Time       int      `json:"time"`
	Times      int      `json:"times"`
}

func TestGetVideoAndSave(t *testing.T) {
	strNum := "0"
	strTime := "20180401"
	url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getDailyRank/" + strNum + "-" + strTime + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetDailyRank" + strNum + strTime + "jsonpcallback"
	fmt.Println(url)
	content, err := ParaseURL(url)
	if err != nil {
		fmt.Println("error_operate :", err)
	}

	strs := strings.Split(content, "{")
	aimstr := strings.Split(strs[1], "}")
	myAim := aimstr[0]
	fmt.Println("myAim is ", myAim)

	var aim testRespon
	err = json.Unmarshal([]byte("{"+myAim+"}"), &aim)
	if err != nil {
		fmt.Println("error_operate is :", err)
	}
	fmt.Printf("aim is:%v", aim)

}

type myTitle struct {
	Title string `json:"title"`
}

func TestParaseURL(t *testing.T) {
	source := `\u9065\u8fdc\u7684\u5979 (\u7ca4\u8bed)`
	sUnicodev := strings.Split(source, `\u`)
	aim := ""
	for _, v := range sUnicodev {
		if v == ""{
			continue
		}

		regexp.MustCompile(`([0-9a-zA-Z]+)`)

		fmt.Println("v",v)
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		aim += fmt.Sprintf("%c", temp)

	}
	fmt.Println(aim)


	//second := strings.Replace(source, `\u`, ``, -1)
	////
	//fmt.Println("second", second)
	//var third string
	//third, _ = strconv.Unquote(strconv.QuoteToASCII(source))
	//third += ".lalal"
	//fmt.Println("third", third)
}
