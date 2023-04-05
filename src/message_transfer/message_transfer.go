package message_transfer

import (
	"encoding/json"
	"errors"

	"github.com/HyperloopUPV-H8/Backend-H8/message_transfer/models"
	"github.com/rs/zerolog"
	trace "github.com/rs/zerolog/log"
)

const (
	MESSAGE_TRANSFER_HANDLER_NAME = "messageTransfer"
)

var (
	messageTransfer       *MessageTransfer
	messageTransferConfig = MessageTransferConfig{
		updateTopic: "message/update",
	}
)

type MessageTransferConfig struct {
	updateTopic string `toml:"update_topic"`
}

func SetConfig(config MessageTransferConfig) {
	messageTransferConfig = config
}

func Get() *MessageTransfer {
	if messageTransfer == nil {
		initMessageTransfer()
	}
	trace.Debug().Msg("get message transfer")
	return messageTransfer
}

func initMessageTransfer() {
	trace.Info().Msg("init message transfer")
	messageTransfer = &MessageTransfer{
		updateTopic: messageTransferConfig.updateTopic,
		sendMessage: defaultSendMessage,
		trace:       trace.With().Str("component", MESSAGE_TRANSFER_HANDLER_NAME).Logger(),
	}
}

type MessageTransfer struct {
	updateTopic string
	sendMessage func(topic string, payload any, targets ...string) error
	trace       zerolog.Logger
}

func (messageTransfer *MessageTransfer) SendMessage(message models.Message) error {
	messageTransfer.trace.Warn().Uint16("id", message.ID).Str("type", message.Type).Str("desc", message.Description).Msg("send message")
	return messageTransfer.sendMessage(messageTransfer.updateTopic, message)
}

func (messageTransfer *MessageTransfer) UpdateMessage(topic string, payload json.RawMessage, source string) {
	messageTransfer.trace.Warn().Str("source", source).Str("topic", topic).Msg("got message")
}

func (messageTransfer *MessageTransfer) SetSendMessage(sendMessage func(topic string, payload any, targets ...string) error) {
	messageTransfer.trace.Debug().Msg("set send message")
	messageTransfer.sendMessage = sendMessage
}

func (messageTransfer *MessageTransfer) HandlerName() string {
	return MESSAGE_TRANSFER_HANDLER_NAME
}

func defaultSendMessage(string, any, ...string) error {
	return errors.New("message transfer must be registered before using")
}
