package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	startTime := time.Now()
	var url string //下载文件的地址
	//url = "https://down.fjweite.cn/syk/windows_7_ultimate_x64_2020.iso"
	//url = "https://download.jetbrains.com/go/goland-2020.2.2.dmg"
	url = "http://get.wiz.cn/wiznote-windows-x86-2020-09-05.exe"
	downloader := NewFileDownloader(url, "", "", 100)
	if err := downloader.Run(); err != nil {
		// fmt.Printf("\n%s", err)
		log.Fatal(err)
	}
	fmt.Printf("\n 文件下载完成耗时: %f second\n", time.Now().Sub(startTime).Seconds())
}
