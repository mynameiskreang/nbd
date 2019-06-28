package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"nbd/helper"
	"nbd/model"
	"net/http"
	"os"
	"strings"
)

var tokenLineNotify string
var idDatabase = map[string]bool{}
var hashDatabase = map[string]bool{}

func init() {
	readData()
}

func main() {
	fs, _ := ioutil.ReadFile("config.yaml")
	config := model.Config{}
	yaml.Unmarshal(fs, &config)

	tokenLineNotify = config.Token
	for _, site := range config.Renthub.Link {
		getRenthub(site)
	}
}

func readData() {
	fs, err := ioutil.ReadFile("nbd.db")
	if err != nil {
		notify(err.Error(), "")
	}
	json.Unmarshal(fs, &idDatabase)

	fs, err = ioutil.ReadFile("hash.db")
	if err != nil {
		notify(err.Error(), "")
	}
	json.Unmarshal(fs, &hashDatabase)
}

func writeData() {
	data, _ := json.Marshal(idDatabase)
	err := ioutil.WriteFile("nbd.db", data, os.ModePerm)
	if err != nil {
		notify(err.Error(), "")
	}

	data, _ = json.Marshal(hashDatabase)
	err = ioutil.WriteFile("hash.db", data, os.ModePerm)
	if err != nil {
		notify(err.Error(), "")
	}
}

func getRenthub(urlReq string) {
	req, _ := http.NewRequest("GET", urlReq, nil)
	req.Header.Add("Accept", "*/*")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	output := helper.GetInnerSubstring(string(body), `{"results":"`, `"}`)
	output = strings.Replace(output, `\"`, `"`, -1)
	output = strings.Replace(output, `\t`, "", -1)
	output = strings.Replace(output, `\n`, "", -1)

	rd := strings.NewReader(output)
	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		notify(err.Error(), "")
		return
	}

	// Scrapping data
	var renthubInfos []model.RenthubInfo
	doc.Find("li").Each(func(i int, sel *goquery.Selection) {
		if id, isExist := sel.Attr("id"); isExist {
			// Check map
			if !idDatabase[id] {
				idDatabase[id] = true
				renthubInfo := model.RenthubInfo{}
				renthubInfo.ID = id
				renthubInfo.Name = sel.Find("span.name").Text()
				renthubInfo.Name = strings.TrimSuffix(renthubInfo.Name, "UPDATE !")
				renthubInfo.Image = sel.Find("img.tb").AttrOr("src", "")
				renthubInfo.LinkRoom = "https://www.renthub.in.th" + sel.Find("span.name a").AttrOr("href", "")
				renthubInfo.Price = sel.Find("span.price").Text()
				renthubInfo.Project = sel.Find("div.listing_project a").Text()
				renthubInfo.LinkProject = "https://www.renthub.in.th" + sel.Find("div.listing_project a").AttrOr("href", "")
				if !hashDatabase[helper.HashMD5(renthubInfo.Name)] {
					renthubInfos = append(renthubInfos, renthubInfo)
				}
				hashDatabase[helper.HashMD5(renthubInfo.Name)] = true
			}
		}
	})

	// Notification
	for _, renthubInfo := range renthubInfos {
		message := renthubInfo.Name + " " + renthubInfo.Price + " " + renthubInfo.LinkRoom
		if os.Args[1] == "skip" {
			fmt.Println(message, renthubInfo.Image)
		} else {
			notify(message, renthubInfo.Image)
		}

	}

	// Update database
	writeData()
}

func notify(message, image string) {
	url := "https://notify-api.line.me/api/notify"
	payload := strings.NewReader("message=" + message + "&imageFullsize=" + image + "&imageThumbnail=" + image)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authorization", "Bearer "+tokenLineNotify)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "notify-api.line.me")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
}
