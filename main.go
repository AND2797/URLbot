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
    if m.Content[0:2] == "sl" {
        if m.Content[3:5] == "-h" {
            help := "Use `sl <insert link>` to create link"
            s.ChannelMessageSend(m.ChannelID, help)
        } else {
        URL := m.Content[3:]
        urlResponse := urlHandler(URL)
        fmt.Println("response", urlResponse)
            /* TODO: 
             handle cli gracefully */
        s.ChannelMessageDelete(m.ChannelID, m.ID)
        s.ChannelMessageSend(m.ChannelID, urlResponse)
        }
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
