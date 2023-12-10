package clients

import (
	"Waffle/bancho/chat"
	"Waffle/bancho/client_manager"
	"Waffle/bancho/osu/base_packet_structures"
	"strings"
)

type WaffleCommand interface {
	Execute(args []string) []string
}

var helpStrings = []string{
	"!help  :: You're reading this right now",
	"!roll  :: Rolls a random number between 0 and 100",
	"!stats :: Shows Waffle Statistics",
	"!rank  :: Shows your osu! stats",
	"!rank <osu!|osu!taiko|osu!catch> :: Shows your own stats for a given mode",
	"!rank <username> :: Shows a user's osu! stats",
	"!rank <username> <osu!|osu!taiko|osu!catch> :: Shows a user's stats for a given mode",
	"!leaderboards <osu!|osu!taiko|osu!catch> :: Shows a mode's leaderboard",
	"!leaderboards <offset> <osu!|osu!taiko|osu!catch> :: Shows a mode's leaderboard and offsets it",
	"!leaderboards <offset> :: Shows the osu! leaderboard and offsets it",
	"!leaderboards :: Shows the osu!'s top 10",
}

var adminHelpStrings = []string{
	"---------------------------------",
	"!announce target <client username> : <message> :: Sends a Notification to a client",
	"^^^ That : seperator is important there!!",
	"!announce all <message> :: Sends a Notification to everyone on the server",
	"!silence <duration in minutes> <username> :: Silences a user for <duration> minutes",
}

var commandHandlers map[string]func(client_manager.WaffleClient, []string) []string

func WaffleBotInitializeCommands() {

	commandHandlers = make(map[string]func(sender client_manager.WaffleClient, args []string) []string)

	commandHandlers["!help"] = WaffleBotCommandHelp
	commandHandlers["!announce"] = WaffleBotCommandAnnounce
	commandHandlers["!roll"] = WaffleBotCommandRoll
	commandHandlers["!stats"] = WaffleBotCommandBanchoStatistics
	commandHandlers["!rank"] = WaffleBotCommandRank
	commandHandlers["!leaderboard"] = WaffleBotCommandLeaderboards
	commandHandlers["!leaderboards"] = WaffleBotCommandLeaderboards
	commandHandlers["!silence"] = WaffleBotCommandSilence
	commandHandlers["!mp"] = WaffleBotCommandMultiplayer
}

func (client *WaffleBot) WaffleBotHandleCommand(sender client_manager.WaffleClient, message base_packet_structures.Message) {
	publicCommand := message.Target[0] == '#'

	var command string
	var arguments []string

	splitMessage := strings.Split(message.Message, " ")

	if len(splitMessage) == 0 {
		return
	}

	command = splitMessage[0]
	arguments = splitMessage[1:]

	handler, exists := commandHandlers[command]

	if !exists {
		return
	}

	result := handler(sender, arguments)

	for _, messageString := range result {
		if publicCommand {
			if message.Target == "#multiplayer" {
				senderLobby := sender.GetMultiplayerLobby()

				if senderLobby != nil {
					senderLobby.MultiChannel.SendMessage(WaffleBotInstance, messageString, message.Target)
				}
			} else {
				channel, exists := chat.GetChannelByName(message.Target)

				if exists {
					channel.SendMessage(WaffleBotInstance, messageString, message.Target)
				}
			}
		} else {
			sender.BanchoIrcMessage(base_packet_structures.Message{
				Sender:  "WaffleBot",
				Message: messageString,
				Target:  message.Target,
			})
		}
	}
}
