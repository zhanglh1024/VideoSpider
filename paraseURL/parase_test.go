package paraseURL

import (
	"fmt"
	"testing"
)

func TestGetVideoAndSave(t *testing.T) {
	strNum := "0"
	strTime := "20190501"
	url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getDailyRank/"+strNum+"-"+strTime+"/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetDailyRank"+strNum+strTime+"jsonpcallback"
	content, err := ParaseURL(url)
	if err != nil{
		fmt.Println("error :",err)
	}

	fmt.Println(content)
}
