package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//resp, err := http.Get("http://music.163.com/song/media/outer/url?id=1364137053.mp3")
	req,_:=http.NewRequest("GET","https://music.163.com/discover/toplist?id=3778678",nil)
	//设置请求头，防止504
	//req.Header.Set("Referer", "https://music.163.com/")
	//req.Header.Set("Host", "music.163.com")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")

	resp,err:=http.DefaultClient.Do(req)
	//resp,err := http.Get("https://music.163.com/#/discover/toplist?id=3778678")
	defer resp.Body.Close()
	if err!=nil{
		panic(err)
	}
	//resp, err := http.Get("https://music.163.com/#/artist?id=6731")
	//if err != nil {
	//	panic("request url error!")
	//}

	//respBytes,_:=ioutil.ReadAll(resp.Body)
	//defer resp.Body.Close()
	file,_ := os.Create("doPost.txt")
	defer file.Close()
	io.Copy(file,resp.Body)


	getResp,_:=http.Get("https://music.163.com/discover/toplist?id=3778678")
	defer getResp.Body.Close()
	getFile,_:=os.Create("getFile.txt")
	defer getFile.Close()
	io.Copy(getFile,getResp.Body)


	//fmt.Println(string(respBytes))

	//doc := soup.HTMLParse(string(respBytes))
	//links := doc.Find("title")
	//for link := range links{
	//
	//}



	//file, err := os.Create("don't care.mp3")
	//defer file.Close()
	//if err != nil{
	//	panic("create file error!")
	//}

	//io.Copy(file,resp.Body)
	fmt.Println("download file finish")
}

//http://music.163.com/artist?id=6731

