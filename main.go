package main

import (
        "os"
        "os/signal"
        "syscall"
        "fmt"
        "flag"
        "github.com/bwmarrin/discordgo"
    )

var (
    Token string
)

func init() {
    flag.StringVar(&Token, "t", "", "Bot Token")
    //flag.StringVar(&Link, "l", "", "Link")
    flag.Parse()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
    fmt.Println("In msg create")
    if m.Content[0:5] == "!link" {
        URL := m.Content[5:]
        shortened, err := shorten(URL)
        if err != nil {
            fmt.Println ("Error generating URl,", err)
        }
    s.ChannelMessageSend(m.ChannelID, shortened)
    }
}


func main() {
    dg, err := discordgo.New("Bot " + Token)
    if err != nil{
        fmt.Println("Error creating Discord session,", err)
        return
    }

    dg.AddHandler(messageCreate)
    dg.Identify.Intents = discordgo.IntentsGuildMessages

    err = dg.Open()

    if err != nil {
        fmt.Println("Error opening connection,",err)
        return
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc
    dg.Close()
}
