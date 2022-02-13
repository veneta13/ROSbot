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
)

// Gets message text for the bot Discord messages.
func logger(messageCode int, arguments ...[]string) string {
	if messageCode == 1 {
		return message1
	}

	if messageCode == 2 {
		return message2
	}

	return ""
}
