package main

import (
        "fmt"
        "log"
        "net/http"
        )


import (
    "github.com/abiosoft/river"
)

var CameraSettings SettingsStruct
var Videos []VideoRecord


func main () {
    var err error

    fmt.Println("server init")
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
    videoEndpoint.Register(proxyModel())
    rv.Handle("/video", videoEndpoint)
    
    authEndpoint := river.NewEndpoint().Get("/", newAuthToken)
    rv.Handle("/auth", authEndpoint)
    
    
    
    log.Fatal(http.ListenAndServe(":8080", rv))

}


