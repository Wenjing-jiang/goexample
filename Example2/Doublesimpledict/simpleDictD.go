package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DictResponse struct {
	Details []struct {
		Detail string `json:"detail"`
		Extra  string `json:"extra"`
	} `json:"details"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}
type DictResponseHSData struct {
	Result []struct {
		Ec struct {
			Basic struct {
				Explains []struct {
					Pos   string `json:"pos"`
					Trans string `json:"trans"`
				} `json:"explains"`
			} `json:"basic"`
		} `json:"ec"`
	} `json:"result"`
}

type DictRequest struct {
	Source         string   `json:"source"`
	Words          []string `json:"words"`
	SourceLanguage string   `json:"source_language"`
	TargetLanguage string   `json:"target_language"`
}

func query(word string) {
	client := &http.Client{}
	//var data = strings.NewReader(`{"source":"youdao","words":["good"],"source_language":"en","target_language":"zh"}`)
	request := DictRequest{Source: "youdao", Words: []string{word}, SourceLanguage: "en", TargetLanguage: "zh"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/detail/v1/?msToken=&X-Bogus=DFSzswVuQDaZVWOztHeEbjNSwbmq&_signature=_02B4Z6wo00001JrA-dwAAIDDz0JyQ9yEWmyawP1AAEJ.ftoo8oXIUwfIsihFkekBkxt-VV1hImx-r-OUykXQLmvZ77rEZsUpbUzLbm.oeOpKIrKB8bJlIqo4u1shK16rMaeFSPP4BiHrC7RI82", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16910334362017288; s_v_web_id=verify_lkulnfzk_lTrKJh54_8jFX_4hRq_9OkZ_Llff7bPWPliR; ttcid=04b81943cd454fc6a93ee6e9e39052d940; i18next=zh-CN; tt_scid=unClmrbQNJTN8AwRZ68zv6sUe4RQNAXvzyP.im8v6Tw-bJUvEEA7rTU2La.mFXGh57a6")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/?category=&home_language=zh&source_language=detect&target_language=zh&text=good")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	//fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dictResponse)
	item := dictResponse.Details[0]
	jsonStr := item.Detail

	var dictResponseHSData DictResponseHSData
	err = json.Unmarshal([]byte(jsonStr), &dictResponseHSData)
	if err != nil {
		panic(err)
	}
	for _, item := range dictResponseHSData.Result[0].Ec.Basic.Explains {
		fmt.Println(item.Pos, item.Trans)
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
