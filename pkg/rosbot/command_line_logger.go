package rosbot

import "fmt"

var commandLineMessages = [26]string{
	"Error: Cannot create Discord session.",
	"Error: Callback error.",
	"Error: Cannot open connection.",
	"Log: ROSbot is up.",
	"Error: Error while closing session.",
	"Error: Cannot log in.",
	"Log: Successfully logged in.",
	"Log: Playlist created successfully.",
	"Error: Cannot read user stats.",
	"Error: Unrecognised command.",
	"Log: Require login for action.",
	"Error: Cannot delete message.",
	"Error: Cannot find user.",
	"Log: Logging in.",
	"Error: Error getting Spotify token.",
	"Error: Cannot log in.",
	"Log: Login successful.",
	"Log: Creating playlist.",
	"Error: Cannot create playlist.",
	"Error: Cannot find information for keyword.",
	"Error: Cannot load playlist image.",
	"Error: Cannot find tracks.",
	"Error: Cannot add tracks to playlist.",
	"Error: Cannot set playlist image.",
	"Error: Cannot get artist stats.",
	"Error: Cannot get track stats.",
}

// commandLineLogger Logs necessary information on the command line.
func commandLineLogger(messageCode int) {
	switch messageCode {
	case 1:
		fmt.Println(commandLineMessages[0])
	case 2:
		fmt.Println(commandLineMessages[1])
	case 3:
		fmt.Println(commandLineMessages[2])
	case 4:
		fmt.Println(commandLineMessages[3])
	case 5:
		fmt.Println(commandLineMessages[4])
	case 6:
		fmt.Println(commandLineMessages[5])
	case 7:
		fmt.Println(commandLineMessages[6])
	case 8:
		fmt.Println(commandLineMessages[7])
	case 9:
		fmt.Println(commandLineMessages[8])
	case 10:
		fmt.Println(commandLineMessages[9])
	case 11:
		fmt.Println(commandLineMessages[10])
	case 12:
		fmt.Println(commandLineMessages[11])
	case 13:
		fmt.Println(commandLineMessages[12])
	case 14:
		fmt.Println(commandLineMessages[13])
	case 15:
		fmt.Println(commandLineMessages[14])
	case 16:
		fmt.Println(commandLineMessages[15])
	case 17:
		fmt.Println(commandLineMessages[16])
	case 18:
		fmt.Println(commandLineMessages[17])
	case 19:
		fmt.Println(commandLineMessages[18])
	case 20:
		fmt.Println(commandLineMessages[19])
	case 21:
		fmt.Println(commandLineMessages[20])
	case 22:
		fmt.Println(commandLineMessages[21])
	case 23:
		fmt.Println(commandLineMessages[22])
	case 24:
		fmt.Println(commandLineMessages[23])
	case 25:
		fmt.Println(commandLineMessages[24])
	case 26:
		fmt.Println(commandLineMessages[25])
	}
}
