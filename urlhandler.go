package main
import (
        "net/http"
        "io/ioutil"
        "fmt"
    )

type newURL struct {
    URL string
    err      error
}

type urlHandlerResponse struct {
    URL string
    health bool
    err error
}

func shorten(URL string, sendStruct chan<- interface{}) {
    resp, errGet := http.Get(URL)
    shortenResponse := &newURL{}
    if errGet != nil{
        fmt.Println("Error in completing http request.")
        shortenResponse.err = errGet
        sendStruct <- shortenResponse
    }
    bytes, errBytes := ioutil.ReadAll(resp.Body)
    if errBytes != nil{
        shortenResponse.err = errGet
        fmt.Println("Error in reading bytes")
        sendStruct <- shortenResponse
    }

    shortenResponse.URL = string(bytes)
    shortenResponse.err = nil
    sendStruct <- shortenResponse
}

func checkHealth(URL string, sendFlag chan<- interface{}){
    _, errGet := http.Get(URL)
    if errGet != nil {
        // error getting response
        sendFlag <- false
        return
    }
    sendFlag <- true
}

func urlHandler(URL string) *urlHandlerResponse{
    API := "https://tinyurl.com/api-create.php?url="
    constructURL := API + URL
    rcvrChan := make(chan interface{})
    response := &urlHandlerResponse{}
    go checkHealth(URL, rcvrChan)
    go shorten(constructURL, rcvrChan)

    for value := range rcvrChan {
        counter := 0
        switch v := value.(type){
            case bool:
                fmt.Println("health", v)
                response.health = v
                counter++
                fmt.Println(counter)
                if (counter == 2){
                    close(rcvrChan)
                    break
                }
            case *newURL:
                fmt.Println("URL", v.URL)
                response.URL = v.URL
                response.err = v.err
                counter++
                fmt.Println(counter)
                if (counter == 2){
                    close(rcvrChan)
                    break
                }
        }
    }
    return response
}

