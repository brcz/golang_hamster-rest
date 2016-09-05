package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

import (
	"github.com/abiosoft/river"
)

var CameraSettings SettingsStruct
var Videos []VideoRecord

var cameraLogin string
var cameraPass string

func init() {
	//bind  flag to  variables
	flag.StringVar(&cameraLogin, "u", "", "user")
	flag.StringVar(&cameraPass, "p", "", "password")
}

func main() {
	var err error

	fmt.Println("server init")
	flag.Parse()

	CameraSettings = initSettings()
	//fmt.Println("settings applied: ",CameraSettings)
	Videos, err = fetchVideos(CameraSettings)
	if err != nil {
		fmt.Println("video fetch error:", err)
	}

	rv := river.New()

	videoEndpoint := river.NewEndpoint().
		Get("/", getList).
		Get("/:id", getVideoRecord).
		Delete("/:id", deleteVideoRecord)

	//videoEndpoint.Use(authMid)
	videoEndpoint.Use(rateTokenBucketMid)
	videoEndpoint.Register(proxyModel())
	rv.Handle("/video", videoEndpoint)

	authEndpoint := river.NewEndpoint().Get("/", newAuthToken)
	rv.Handle("/auth", authEndpoint)

    fmt.Println("server ready...")
	log.Fatal(http.ListenAndServe(":8080", rv))

}
