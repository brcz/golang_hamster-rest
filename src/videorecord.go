// VideoRecord def
package main

 import (
	"net/http"
         "time"
         "log"
         "errors"
         "io/ioutil"
         "encoding/base64"
 )
import (
	"github.com/abiosoft/river"
         "github.com/robertkrimen/otto"
 )

const (
    VARNAME_RECFILE = "record_name0"
    VARNAME_RECSIZE = "record_size0"
       VARNAME_RECNUM ="record_num0"
       VARNAME_PAGESIZE="PageSize"
       VARNAME_PAGENUM="PageIndex"
       VARNAME_PAGECOUNT="PageCount"
       VARNAME_RECCOUNT="RecordCount"
       
)

/*
 *   scince we got data as pure js code we need to own file parser
 
 var record_name0=new Array();
 var record_size0=new Array();
 record_name0[0]="20160831081335_010.h264";
 record_size0[0]=14317565;
 
 ...
 
 record_name0[999]="20160825175343_010.h264";
 record_size0[999]=14428143;
 var record_num0=1000;
 var PageIndex=0;
 var PageSize=1000;
 var RecordCount=2310;
 var PageCount=3;
 */

 // VideoRecord is single video alert data.
 type VideoRecord struct {
	ID   string `json:"id"`
	filename string `json:"filename"`
    filesize   uint
    record_type bool `json:"type_alarm"`//true - alarm, false - regular record
 }


func parseResponse (input string) ([]VideoRecord, error) {
    
    return nil, nil
}

// getVideoList handles GET /video/list.
func getList(c *river.Context, model Model) {
    c.Render(http.StatusOK, model.getList())
}

// getVideoRecord handles GET /video/:id.
func getVideoRecord(c *river.Context, model Model) {
    record := model.get(c.Param("id"))
    if record == nil {
        c.RenderEmpty(http.StatusNotFound)
        return
    }
    c.Render(http.StatusOK, record)
}

// deleteVideoRecord handles DELETE /video/:id.
func deleteVideoRecord(c *river.Context, model Model) {
    model.delete(c.Param("id"))
    c.RenderEmpty(http.StatusNoContent)
}


// fetch and parse video slices to esy operate
func fetchVideos(settings SettingsStruct) ([]VideoRecord, error) {
    
    var fetchedRecords []VideoRecord
    //simle url creation from settings
    url := settings.cameraAddress + settings.fetchURI
    
    var netClient = &http.Client{
        Timeout: time.Second * 30,
        CheckRedirect: redirectPolicyFunc,
    }
    //response, err :=netClient.Get(url)
    
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Authorization","Basic "+basicAuth(cameraLogin, cameraPass))
    
    response, err := netClient.Do(req)
    
    
    if err != nil {
        log.Println("ERROR: Failed to fetch \"" + url + "\"")
        return nil, err
    }
    
    defer response.Body.Close()
    
    bodyBuffer, _ := ioutil.ReadAll(response.Body)
    //fmt.Printf("\n\n%s", bodyBuffer)
    //fmt.Println("\n\njs=",bodyBuffer)
    
    vm := otto.New()
    vm.Run(bodyBuffer)
    
    if value, err := vm.Get(VARNAME_RECFILE); err == nil {
        if value.Class() == "Array" {
            values, err := value.Export()
            if err != nil {
                return nil, err
            }
            
            for _, val := range values.([]string) {
                type_alarm:=false
                
                if sig:=val[16:18]; sig == "10" || sig=="01" ||sig=="11" {
                    type_alarm = true
                }
                //parsedRecord :=
                
                fetchedRecords = append(fetchedRecords, VideoRecord{ val[:18], val, 0, type_alarm})
                //fmt.Println("\nrecord=", fetchedRecords)
            }
            
        } else {
            return nil, errors.New("filenames array not found")
        }
    }
    
    return fetchedRecords, nil
}

func basicAuth(username, password string) string {
    auth := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
    req.Header.Add("Authorization","Basic "+basicAuth(cameraLogin, cameraPass))
    return nil
}


