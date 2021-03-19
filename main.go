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
    newMsgID string
)

func init() {
    flag.StringVar(&Token, "t", "", "Bot Token")
    flag.Parse()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
    if m.Content[0:5] == "!link" {
        URL := m.Content[6:]
        shortened, err := shorten(URL)
        if err != nil {
            fmt.Println ("Error generating URl,", err)
        }

        s.ChannelMessageDelete(m.ChannelID, m.ID)
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
