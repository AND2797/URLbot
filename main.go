package main
import (
        "os"
        "strconv"
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
        fmt.Println("HERE")
        if m.Content[3:5] == "-h" {
            help := "Use `sl <insert link>` to create link"
            s.ChannelMessageSend(m.ChannelID, help)
        } else {
        URL := m.Content[3:]
        urlResponse := urlHandler(URL)
        shortenedBy := len(URL) - len(urlResponse)
        fmt.Println("response", urlResponse)
            /* TODO: 
             handle cli gracefully */
        embedMSG := &discordgo.MessageEmbed{}
        embedMSG.Description = urlResponse
        embedMSG.Author = &discordgo.MessageEmbedAuthor{}
        embedMSG.Author.Name = m.Author.Username
        embedMSG.Footer = &discordgo.MessageEmbedFooter{}
        embedMSG.Footer.Text = "Shortened by " + strconv.Itoa(shortenedBy) + " chars."
        s.ChannelMessageDelete(m.ChannelID, m.ID)
        s.ChannelMessageSendEmbed(m.ChannelID, embedMSG)
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
