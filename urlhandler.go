package main
import (
        "net/http"
        "io/ioutil"
        "fmt"
    )


type newURL struct {
    URL     string
    err      error
}

type urlHandlerResponse struct {
    URL string
    health bool
    err error
}

func shorten(URL string) *newURL {
    resp, errGet := http.Get(URL)
    shortenResponse := &newURL{}
    if errGet != nil{
        fmt.Println("Error in completing http request.")
        shortenResponse.err = errGet
        return shortenResponse
    }
    bytes, errBytes := ioutil.ReadAll(resp.Body)
    if errBytes != nil{
        shortenResponse.err = errGet
        fmt.Println("Error in reading bytes")
        return shortenResponse
    }

    shortenResponse.URL = string(bytes)
    shortenResponse.err = nil
    return shortenResponse
}


func sanitizeURL(URL string) string {

    if (URL[0:8] != "https://" && URL[0:7] != "http://"){
        URL = "https://" + URL
    }

    return URL
}


func checkHealth(URL string) bool {
    /* TODO: check HTTP STATUS */
    URL = sanitizeURL(URL)
    _, errGet := http.Get(URL)
    fmt.Println(URL)
    fmt.Println(errGet)
    if errGet != nil {
        // error getting response
        return false
    }
    return true
}

func urlHandler(URL string) string{
    API := "https://tinyurl.com/api-create.php?url="
    constructURL := API + URL
    health := checkHealth(URL)
    shortenResponse := shorten(constructURL)
    fmt.Println("health,", health)
    if (health == true) {
        return shortenResponse.URL
    } else {
        return "Error in generating link."
    }
}

