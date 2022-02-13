package rosbot

// Bot Discord messages.
const (
	message1 = ":scroll: **COMMANDS**\n\n" +
		":point_right: use `!help` to get this list\n\n" +
		":point_right: use `!log-in ID={CLIENT_ID} SECRET={CLIENT_SECRET}` to connect your Spotify account\n\n" +
		":point_right: use `!say-hi` for a surprise\n\n" +
		":point_right: use `!get-stats {type} {time}` to get your Spotify stats\n\n" +
		":bar_chart: __Supported types__:\n\n" +
		":singer: artist\n\n" +
		":musical_note: track\n\n" +
		":alarm_clock: __Supported time periods__:\n\n" +
		":clock1: last month\n\n" +
		":clock2: 6 months\n\n" +
		":clock3: all time\n\n" +
		":point_right: use `!create-playlist {mood}` to create a playlist\n\n" +
		":performing_arts: __Supported moods__:\n\n" +
		":smile: happy\n\n" +
		":sob: sad\n\n" +
		":relaxed: relaxed\n\n" +
		":partying_face: party\n\n" +
		":face_with_monocle: focused\n\n" +
		":smiling_face_with_3_hearts: romantic\n\n" +
		":christmas_tree: holiday\n\n" +
		":blue_car: travel\n\n" +
		":raised_hands: motivated \n\n" +
		":sleeping: sleepy \n\n"

	message2 = "Log in via Spotify here :point_right: "

	message3 = "Please delete your login message:exclamation:"

	message4 = "Cannot create playlist :pensive: Please try again"

	message5 = "Key word not recognised :cry: Please try again"

	message7 = " created successfully :partying_face:"

	message8 = "You can access your playlist here :point_right:" +
			   "https://open.spotify.com/playlist/"

	message9 = "Please `!log-in` before creating playlists :wink:"

	message10 = "Getting your stats was unsuccessful :cry: Please try again"

	message11 = "I don't recognise this command :thinking: Please try again"

	message12 = "Please `!log-in` go get your stats :wink:"
)

// Gets message text for the bot Discord messages.
func logger(messageCode int, arguments ...string) string {
	switch messageCode {
	case 1:
		return message1
	case 2:
		return message2
	case 3:
		return "<@!" + arguments[0] + ">" + message3
	case 4:
		return message4
	case 5:
		return message5
	case 6:
		return  "Hello "+ arguments[0] +" :wave:"
	case 7:
		return arguments[0] + message7
	case 8:
		return message8 + arguments[0]
	case 9:
		return message9
	case 10:
		return message10
	case 11:
		return message11
	case 12:
		return message12
	}

	return ""
}
