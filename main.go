package main

import (
        "fmt"
        //"github.com/bwmarrin/discordgo"
    )


var (
    Token string
    Link  string
)

//func init() {
//    flag.StringVar(&Token, "!", "", "Bot Token")
//    flag.Link(&Link, ";", "", "Link")
//}


func main() {
    test := "www.google.com"
    shortened := shorten(test)
    fmt.Println(shortened)
    //dg, err := discordgo.New("Bot " + Token)
    //if err != nil{
    //    fmt.Println("Error creating Discord session,", err)
    //    return
    //}
}
