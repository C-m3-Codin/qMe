package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/C-m3-Codin/q_me/unitcovertion"
	"go.mau.fi/whatsmeow/types/events"
)

func  (whatsappClient *WhatsappClient)EventHandler(evt interface{}) {
    switch v := evt.(type) {
	case *events.CallOffer:
		fmt.Println("Received a call offer")
		// handle call
	case *events.JoinedGroup:
		fmt.Println("Joined a group")
		// joined group
    case *events.Message:

		// checks if its a status update
		if( v.Info.MessageSource.Chat.User=="status") {
		
			fmt.Println("ITs a status update")



		}else if(v.Info.MessageSource.IsGroup){
			// checks if message is a group message
			fmt.Println("Its a Group message message")
			// check for tags
			

		}else{
			fmt.Println("Its a Direct message ...deploying handler")
			whatsappClient.textMessageHandler(v)

		}

    }
}




func (whatsappClient *WhatsappClient)textMessageHandler(message *events.Message){

	match, _ := regexp.MatchString("[A-Z]{3}2[A-Z]{3} [1-9]+([.][1-9])*", strings.ToUpper(message.Message.GetConversation()))
	fmt.Println(match)

	if(match){
		fmt.Println("\n\nMatched\n\n\n")

		re := regexp.MustCompile(`[A-Z]{3}2[A-Z]{3} [1-9]+([.][1-9])*`)
		message_text := re.FindStringSubmatch(strings.ToUpper(message.Message.GetConversation()))
		fmt.Println("\n\nMatched",message_text,"\n\n\n")
		// message_text:=strings.ToUpper(message.Message.GetConversation())
		conv_from:=message_text[0][0:3]
		conv_to:=message_text[0][4:7]
		count_str:=message_text[0][8:]
		// convert float string to float
		if count, err := strconv.ParseFloat(count_str, 32); err == nil {
			fmt.Println(count) // 3.1415927410125732
			fmt.Println(count)
			fmt.Println((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
			response:= ((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
			response=response*float32(count)
			whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
		}

		// var message *proto.Message
		
		
	}else{
		fmt.Println("\n\n\n\nNot matched not message handler for \n ")
		fmt.Println(message.Info)
		response:="fickle is away \n ~q_me"
		whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))

	}
	

}



