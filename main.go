package main

import (
	"./medios"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Declaration.
	var (
		opc       string
		condition = true
	)

	// Instances.
	image := medios.Image{ Title: "Gal", Format: "JPG", Channel: "RGB" }
	audio := medios.Audio{ Title: "Stand by me", Format: "MP3", Duration: 90 }
	video := medios.Video{ Title: "Batman", Format: "MP4", Frames: 10000 }

	data := medios.ContentWeb {
		Multimedia: []medios.Multimedia {
			&image,
			&audio,
			&video,
		},
	}

	// Menu.
	for exit := true; exit; exit = condition {
		fmt.Println("******* CONTENT WEB (MULTIMEDIA) *******")
		fmt.Println("1.-Add")
		fmt.Println("2.-Show")
		fmt.Println("3.-Exit")
		fmt.Print("-The Option is: ")
		fmt.Scan(&opc)

		switch opc {
			case "1":
				addData(&data)
			case "2":
				showData(&data)
			case "3":
				exited()
				condition = false
			default:
				invalidOptions()
		}
	}
}

// AddData function.
func addData(data *medios.ContentWeb) {
	// Declaration.
	var (
		opc2	   string
		quantity   int
		condition2 = true
		title      string
		format     string
		channel    string
		duration   int64
		frame      int64
	)

	// For strings.
	scanner := bufio.NewScanner(os.Stdin)

	// Two menu.
	for exit2 := true; exit2; exit2 = condition2 {
		fmt.Println("\n******* ADD DATA TO *******")
		fmt.Println("1.-Image")
		fmt.Println("2.-Audio")
		fmt.Println("3.-Video")
		fmt.Println("4.-Return")
		fmt.Print("-The Option is: ")
		fmt.Scan(&opc2)

		switch opc2 {
		case "1":
			fmt.Println("\n******* IMAGE *******")
			fmt.Print("\nHow many registrations do you want to make? ")
			fmt.Scan(&quantity)

			for i := 0; i < quantity; i++ {
				fmt.Println("\nDATA -", i + 1)
				fmt.Print("Add Title: ")
				scanner.Scan()
				title = scanner.Text()

				fmt.Print("Add Format: ")
				scanner.Scan()
				format = scanner.Text()

				fmt.Print("Add Channel: ")
				scanner.Scan()
				channel = scanner.Text()

				image := medios.Image{ Title: title, Format: format, Channel: channel }
				data.Multimedia = append(data.Multimedia, &image)
			}
		case "2":
			fmt.Println("\n******* AUDIO *******")
			fmt.Print("\nHow many registrations do you want to make? ")
			fmt.Scan(&quantity)

			for i := 0; i < quantity; i++ {
				fmt.Println("\nDATA -", i + 1)
				fmt.Print("Add Title: ")
				scanner.Scan()
				title = scanner.Text()

				fmt.Print("Add Format: ")
				scanner.Scan()
				format = scanner.Text()

				fmt.Print("Add Duration: ")
				fmt.Scan(&duration)

				audio := medios.Audio{ Title: title, Format: format, Duration: duration }
				data.Multimedia = append(data.Multimedia, &audio)
			}
		case "3":
			fmt.Println("\n******* VIDEO *******")
			fmt.Print("\nHow many registrations do you want to make? ")
			fmt.Scan(&quantity)

			for i := 0; i < quantity; i++ {
				fmt.Println("\nDATA -", i + 1)
				fmt.Print("Add Title: ")
				scanner.Scan()
				title = scanner.Text()

				fmt.Print("Add Format: ")
				scanner.Scan()
				format = scanner.Text()

				fmt.Print("Add Frame: ")
				fmt.Scan(&frame)

				video := medios.Video{ Title: title, Format: format, Frames: frame }
				data.Multimedia = append(data.Multimedia, &video)
			}
			case "4":
				returnMenu()
				condition2 = false
		default:
			invalidOptions()
		}
	}
}

// ShowData function.
func showData(data *medios.ContentWeb) {
	fmt.Println("\n******* SHOW DATA *******")
	fmt.Println(data.Show())
}

// InvalidOptions function.
func invalidOptions() {
	fmt.Print("\n-Invalid Option!\n\n")
}

// Exited function.
func returnMenu() {
	fmt.Print("\n-Back to Main Menu...\n\n")
}

// Exited function.
func exited() {
	fmt.Println("\n-System exited...")
}