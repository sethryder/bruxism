package bruxism

import "errors"

// Message is a message interface, wraps a single message from a service.
type Message interface {
	Channel() string
	UserName() string
	UserID() string
	UserAvatar() string
	Message() string
	MessageID() string
	IsModerator() bool
}

var AlreadyJoinedError error = errors.New("Already joined.")

// Service is a service interface, wraps a single service such as YouTube or Discord.
type Service interface {
	Name() string
	UserName() string
	Open() (<-chan Message, error)
	IsMe(message Message) bool
	SendMessage(channel, message string) error
	DeleteMessage(channel, messageID string) error
	BanUser(channel, userID string, duration int) error
	UnbanUser(channel, userID string) error
	SetPlaying(game string) error
	Join(join string) error
	Typing(channel string) error
	PrivateMessage(userID, messageID string) error
	IsPrivate(message Message) bool
	SupportsMultiline() bool
	CommandPrefix() string
}

// LoadFunc is the function signature for a load handler.
type LoadFunc func(*Bot, Service, []byte) error

// SaveFunc is the function signature for a save handler.
type SaveFunc func() ([]byte, error)

// HelpFunc is the function signature for a help handler.
type HelpFunc func(*Bot, Service) []string

// MessageFunc is the function signature for a message handler.
type MessageFunc func(*Bot, Service, Message)

// Plugin is a plugin interface, supports loading and saving to a byte array and has help and message handlers.
type Plugin interface {
	Name() string
	Load(*Bot, Service, []byte) error
	Save() ([]byte, error)
	Help(*Bot, Service) []string
	Message(*Bot, Service, Message)
}
