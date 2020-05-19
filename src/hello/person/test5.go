package main

import "fmt"

func main()  {
	var countryCapitalMap map[string]string
	countryCapitalMap= make(map[string]string)
	countryCapitalMap["France"]="巴黎"
	countryCapitalMap["Italy"]="罗马"
	countryCapitalMap["Japan"]="东京"
	countryCapitalMap["India"]="新德里"
	for country:=range countryCapitalMap{
		fmt.Println(country,"首都是：",countryCapitalMap[country])
	}

	captical,ok:= countryCapitalMap["France"]
	fmt.Println(captical)
	fmt.Println(ok)
	if ok {
		fmt.Println("A的首都是",captical)
	}else {
		fmt.Println("A的首都不存在")
	}
	//删除
	delete(countryCapitalMap,"India")
	for country:=range countryCapitalMap{
		fmt.Println(country,"首都是：",countryCapitalMap[country])
	}


}
