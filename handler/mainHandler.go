package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

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
			// v.Info.Timestamp
			// add checker to only reply to recent messages
			duration, _ := time.ParseDuration("-100s")
			startTime := time.Now().Add(duration)

			// v.Info.Timestamp.




			fmt.Println("Its a Direct message ...")
			if(v.Info.Timestamp.Before(startTime)){
				// before 100 seconds of script starting
				// 
				fmt.Println("But message is before starting script")
			}else{
				

				fmt.Println(" ...deploying handler")

				whatsappClient.textMessageHandler(v)
			}

		}

    }
}




func (whatsappClient *WhatsappClient)textMessageHandler(message *events.Message){

	match, _ := regexp.MatchString("[A-Z]{3}2[A-Z]{3} [0-9]+([.][0-9])*", strings.ToUpper(message.Message.GetConversation()))
	fmt.Println(match)

	// match for currency convert FX
	if(match){
		re := regexp.MustCompile(`[A-Z]{3}2[A-Z]{3} [0-9]+([.][0-9])*`)
		message_text := re.FindStringSubmatch(strings.ToUpper(message.Message.GetConversation()))
		fmt.Println("\n Matched currency convertion ",message_text)
		conv_from:=message_text[0][0:3]
		conv_to:=message_text[0][4:7]
		count_str:=message_text[0][8:]
		if count, err := strconv.ParseFloat(count_str, 32); err == nil {
			fmt.Println((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
			response:= ((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
			response=response*float32(count)
			whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
		}else{
			response:="Error in the value given"
			whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
		}

		
		
	}else{
		fmt.Println("No matching hadlers deployed yet")
		fmt.Println(message.Info)
		// response:="fickle is away \n ~q_me"
		// whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))

	}
	

}


