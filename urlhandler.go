package main

import (
        "net/http"
        "io/ioutil"
        "fmt"
    )




func shorten(URL string) (string, error) {

    API := "https://tinyurl.com/api-create.php?url="
    constructURL := API + URL
    resp, errGet := http.Get(constructURL)
    if errGet != nil{
        fmt.Println("Error in completing http request.")
        return "-1", errGet
    }
    bytes, errBytes := ioutil.ReadAll(resp.Body)
    if errBytes != nil{
        fmt.Println("Error in reading bytes")
        return "-1", errGet
    }

    return string(bytes), nil
}

