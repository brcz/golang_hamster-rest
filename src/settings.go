// define settings structures and init functions

package main

// http://109.104.183.8:8181/get_record_file.cgi?loginuse=admin&loginpas=888888&PageIndex=0&PageSize=1000&1472622524421&_=1472622524422

type SettingsStruct struct {
    cameraAddress string
    fetchURI string
    
}

func initSettings() (SettingsStruct) {
    
    a:= SettingsStruct {cameraAddress: "http://109.104.183.8:8181", fetchURI: "/get_record_file.cgi" }
    
    return a
}