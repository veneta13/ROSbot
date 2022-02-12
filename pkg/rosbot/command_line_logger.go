package rosbot

import "fmt"

//commandLineLogger Logs necessary information on the command line.
func commandLineLogger (messageCode int) {
	switch messageCode {
	case 1:
		fmt.Println("Error: Cannot create Discord session.")
	case 2:
		fmt.Println("Error: Callback error.")
	case 3:
		fmt.Println("Error: Cannot open connection.")
	case 4:
		fmt.Println("Log: ROSbot is up.")
	case 5:
		fmt.Println("Error: Error while closing session.")
	case 6:
		fmt.Println("Error: Cannot log in.")
	case 7:
		fmt.Println("Log: Successfully logged in.")
	case 8:
		fmt.Println("Log: Playlist created successfully.")
	case 9:
		fmt.Println("Error: Cannot read user stats.")
	case 10:
		fmt.Println("Error: Unrecognised command.")
	case 11:
		fmt.Println("Log: Require login for action.")
	case 12:
		fmt.Println("Error: Cannot delete message.")
	case 13:
		fmt.Println("Error: Cannot find user.")
	case 14:
		fmt.Println("Log: Logging in.")
	case 15:
		fmt.Println("Error: Error getting Spotify token.")
	case 16:
		fmt.Println("Error: Cannot log in.")
	case 17:
		fmt.Println("Log: Login successful.")
	case 18:
		fmt.Println("Log: Creating playlist.")
	case 19:
		fmt.Println("Error: Cannot create playlist.")
	case 20:
		fmt.Println("Error: Cannot find information for keyword.")
	case 21:
		fmt.Println("Error: Cannot load playlist image.")
	case 22:
		fmt.Println("Error: Cannot find tracks.")
	case 23:
		fmt.Println("Error: Cannot add tracks to playlist.")
	case 24:
		fmt.Println("Error: Cannot set playlist image.")
	case 25:
		fmt.Println("Error: Cannot get artist stats.")
	case 26:
		fmt.Println("Error: Cannot get track stats.")
	}
}
