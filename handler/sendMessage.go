package handler

import (
	"context"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)


type WhatsappClient struct {
	Client *whatsmeow.Client
}

func SetClient(whatsappClient *whatsmeow.Client) *WhatsappClient {
	return &WhatsappClient{
		Client: whatsappClient,
	}
}

func(whatsappClient *WhatsappClient) SendMessagetoWhatsapp(to types.JID, id string,message string) {


	whatsappClient.Client.SendMessage(context.Background(), to, "", &waProto.Message{
		Conversation: proto.String(message),
	})
}

