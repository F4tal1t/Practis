package main

import "fmt"

// Notification interface defines the contract for all notification types
type Notification interface {
	Send(message string)
	GetType() string
}

// EmailNotification implements Notification interface
type EmailNotification struct {
	recipient string
}

func (e *EmailNotification) Send(message string) {
	fmt.Printf("📧 Email sent to %s: %s\n", e.recipient, message)
}

func (e *EmailNotification) GetType() string {
	return "EMAIL"
}

// SMSNotification implements Notification interface
type SMSNotification struct {
	phoneNumber string
}

func (s *SMSNotification) Send(message string) {
	fmt.Printf("📱 SMS sent to %s: %s\n", s.phoneNumber, message)
}

func (s *SMSNotification) GetType() string {
	return "SMS"
}

// PushNotification implements Notification interface
type PushNotification struct {
	deviceToken string
}

func (p *PushNotification) Send(message string) {
	fmt.Printf("🔔 Push notification sent to %s: %s\n", p.deviceToken, message)
}

func (p *PushNotification) GetType() string {
	return "PUSH"
}

// SlackNotification implements Notification interface
type SlackNotification struct {
	channel string
}

func (sl *SlackNotification) Send(message string) {
	fmt.Printf("💬 Slack message sent to #%s: %s\n", sl.channel, message)
}

func (sl *SlackNotification) GetType() string {
	return "SLACK"
}

// NotificationFactory creates notification instances based on type
func CreateNotification(notificationType string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{recipient: "user@example.com"}
	case "sms":
		return &SMSNotification{phoneNumber: "+1234567890"}
	case "push":
		return &PushNotification{deviceToken: "device_token_123"}
	case "slack":
		return &SlackNotification{channel: "general"}
	default:
		fmt.Printf("❌ Unknown notification type: %s\n", notificationType)
		return nil
	}
}

// NotificationManager manages multiple notification types
type NotificationManager struct {
	notifications []Notification
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		notifications: make([]Notification, 0),
	}
}

func (nm *NotificationManager) AddNotification(notificationType string) {
	notification := CreateNotification(notificationType)
	if notification != nil {
		nm.notifications = append(nm.notifications, notification)
		fmt.Printf("✅ Added %s notification\n", notification.GetType())
	}
}

func (nm *NotificationManager) SendToAll(message string) {
	fmt.Println("📢 Broadcasting message to all notification channels:")
	for _, notification := range nm.notifications {
		notification.Send(message)
	}
}
