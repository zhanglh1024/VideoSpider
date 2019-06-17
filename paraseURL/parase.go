package paraseURL

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)


type Data struct {
	Url string 	`json:"url"`
	Extname string  `json:"extname"`
	Duration int	`json:"duration"`
	FileSize int	`json:"filesize"`
	ThumbUrl string `json:"thumb_url"`
	TransStatus int `json:"trans_status"`
}

type Respone struct {
	Status int `json:"status"`
	Errcode   int `json:"errcode"`
	Error     string `json:"error"`
	Data     Data `json:"data"`
}

func ParseId(id string)  {
	url := "http://visitor.fanxing.kugou.com/VServices/Video.OfflineVideoService.getVideoList/"+id+"-1-0-10/"
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err!=nil{
		panic(err)
	}
	defer func() {
		res.Body.Close()
	}()
	resContent := string(b)

	title := regexp.MustCompile(`("title":"([a-zA-Z0-9\\]+)")`)
	hashl := regexp.MustCompile(`("hashValue":"([a-zA-Z0-9|]+)")`)
	hashList := hashl.FindAllStringSubmatch(resContent, -1)
	titleList := title.FindAllStringSubmatch(resContent,-1)
	fmt.Println(titleList)
	for _, v := range hashList{
		theHash := v[2]
		if strings.Contains(v[2],"|"){
			split := strings.Split(v[2],"|")
			theHash = split[0]
		}

		//md5加密
		hashValue := "hash"+theHash+"pid6kugou_video"
		a:= md5.Sum([]byte(hashValue))
		value := fmt.Sprintf("%x",a)

		fmt.Println(v[2])

		url ="http://tracker.v.kugou.com/video/query?pid=6&sign="+value+"&hash="+theHash
		getVideoAndSave(GetVedioPlayAddress(url))
		fmt.Println("url:",url)
	}
	fmt.Println(hashList)



}


func GetVedioPlayAddress(utl string) string {
	res, err := http.Get(utl)
	if err != nil{
		panic(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err!=nil{
		panic(err)
	}
	defer func() {
		res.Body.Close()
	}()
	var resContent Respone
	fmt.Println(string(b))
	err = json.Unmarshal(b, &resContent)
	if err != nil{
		panic(err)
	}


	return resContent.Data.Url
}

//获取视频并保存到本地
func getVideoAndSave(url string)  {
	//url := "http://fx.v.kugou.com/G127/M07/0B/17/vw0DAFx6mdiAP82TAXAgmN91YYw320.mp4"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	names := strings.Split(url,"/")
	name := fmt.Sprintf("download/%s.mp4",names[len(names)-1])

	os.Mkdir("download", os.ModePerm)
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}