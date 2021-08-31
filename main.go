/*
@Description: xxxx
@Author: feiwang
@Date:  2021/8/31 10:27
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/anaskhan96/soup"
)

func main(){
	requestUrl := "https://www.liaoxuefeng.com/"
	rsp, err := http.Get(requestUrl)

	defer rsp.Body.Close()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(rsp)
	body , err := ioutil.ReadAll(rsp.Body)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	content := string(body)
	//fmt.Println(content)

	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("div","class","uk-margin")
	for _, subDoc := range subDocs{
		link := subDoc.Find("a","target","_blank")

		//if len(link.Text()) >0
		if link.Text() !=""{
			fmt.Printf("%v\n",link.Text())
		}

	}

}