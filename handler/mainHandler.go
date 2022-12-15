package handler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/C-m3-Codin/q_me/unitcovertion"
	"go.mau.fi/whatsmeow/types/events"
)

func  (whatsappClient *WhatsappClient)EventHandler(evt interface{}) {
    switch v := evt.(type) {
    case *events.Message:
        fmt.Println("Received a message!", v.Message.GetConversation())
		// v.Message.
        unitcovertion.GetCurrencyUnit("USD")
		whatsappClient.textMessageHandler(v)
	
	case *events.CallOffer:
		// handle call
	case *events.JoinedGroup:
		// joined group

    }
}


func (whatsappClient *WhatsappClient)textMessageHandler(message *events.Message){

	match, _ := regexp.MatchString("[A-Z]{3}2[A-Z]{3}", strings.ToUpper(message.Message.GetConversation()))
	fmt.Println(match)

	if(match){
		// convert currency
		fmt.Println("\n\n Matched ftx convert lol  \n\n")
		// fmt.Println(message.Message.GetConversation()[0:3])
		message_text:=strings.ToUpper(message.Message.GetConversation())
		conv_from:=message_text[0:3]
		conv_to:=message_text[4:7]
		fmt.Println(unitcovertion.GetCurrencyUnit(conv_from))
		fmt.Println(unitcovertion.GetCurrencyUnit(conv_to))
		fmt.Println(message.Info.Sender)
		fmt.Println((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
		response:=(1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to)
		// var message *proto.Message
		whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
		
		
	}else{
		fmt.Println("Not matched")
	}
	

}



