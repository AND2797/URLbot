package main

import (
        "net/http"
        "io/ioutil"
    )




func shorten(URL string) string {

    API := "https://tinyurl.com/api-create.php?url="
    constructURL := API + URL
    resp, _ := http.Get(constructURL)
    bytes, _ := ioutil.ReadAll(resp.Body)

    return string(bytes)
}

