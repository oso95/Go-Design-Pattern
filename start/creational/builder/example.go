package main

import "fmt"

func main() {
	// Create a new notification builder
	nb := newNotificationBuilder()
	// Set the notification title
	nb.SetTitle("New Notification")
	// Set the notification image
	nb.SetImage("image.jpg")
	// Set the notification icon
	nb.SetIcon("icon.png")
	// Set the notification subtitle
	nb.SetSubtitle("This is a subtitle")
	// Set the notification priority
	nb.SetPriority(0)
	// Set the notification message
	nb.SetMessage("This is a new notification")
	// Set the notification type
	nb.SetType("alert")

	// Build the notification
	n, err := nb.Build()
	if err != nil {
		fmt.Println(err)
	}else{
		// Print the notification
		fmt.Printf("--------------------\n")
		fmt.Printf("Title: %s\n", n.title)
		fmt.Printf("Subtitle: %s\n", n.subtitle)
		fmt.Printf("Message: %s\n", n.message)
		fmt.Printf("Image: %s\n", n.image)
		fmt.Printf("Icon: %s\n", n.icon)
		fmt.Printf("Priority: %d\n", n.priority)
		fmt.Printf("Type: %s\n", n.notType)
	}
	
}