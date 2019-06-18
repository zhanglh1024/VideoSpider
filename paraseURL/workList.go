package paraseURL

func ParasListUrl(urls []string,name string)  {
	for _,url := range urls{
		resContent, err := ParaseURL(url)
		if err != nil{
			panic(err)
		}
		ParasResponeData(resContent,name)
	}
}
