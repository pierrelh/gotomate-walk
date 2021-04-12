package notification

// CreateDatabinder Define the NotificationCreate parameters
type CreateDatabinder struct {
	Title        string
	TitleIsVar   bool
	Message      string
	MessageIsVar bool
}
