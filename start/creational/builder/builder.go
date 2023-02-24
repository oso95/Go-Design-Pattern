package main

import "fmt"

type NotificationBuilder struct {
	Title string
	Subtitle string
	Message string
	Image string
	Icon string
	Priority int
	NotType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubtitle(subtitle string) {
	nb.Subtitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	if nb.Icon != "" && nb.Subtitle == "" {
		return nil, fmt.Errorf("Subtitle is required when icon is set")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("Priority can't be greater than 5")
	}

	return &Notification{
		title: nb.Title,
		subtitle: nb.Subtitle,
		message: nb.Message,
		image: nb.Image,
		icon: nb.Icon,
		priority: nb.Priority,
		notType: nb.NotType,
	}, nil
}
