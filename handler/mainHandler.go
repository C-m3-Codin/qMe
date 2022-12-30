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

	text_message := message.Message.GetConversation();

	match_currency, _ := regexp.MatchString("[A-Z]{3}2[A-Z]{3} [0-9]+([.][0-9])*", strings.ToUpper(text_message))
	match_unit, _ := regexp.MatchString("\\w*2\\w*.unit \\d*[.]?\\d*", strings.ToLower(text_message))
	match_whatsapp, _ := regexp.MatchString("whatsapp [1-9][\\d]{9}", strings.ToLower(text_message))
	match_reminder,_ :=regexp.MatchString("remindme[\\w* *\\w*]*\b(today|tomorrow)\b",strings.ToLower(text_message))

	fmt.Println("matched Currency",match_currency)
	fmt.Println("matched unit :",match_unit)
	fmt.Println("matched whatsapp :",match_whatsapp)
	fmt.Println("matched reminder :",match_reminder);

	// match for currency convert FX
	if(match_currency){
		response:=convertCurrency2Currency(text_message)
		whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))		
	}else if (match_unit){
		response:=convertUnit2Unit(text_message)
		whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
	}else if(match_whatsapp){
		response:=getWhatsappLink(text_message)
		whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
	}else if(match_reminder){
		// handle reminder match
		response:=setReminder(text_message)
		whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))

	}else{


		fmt.Println("No matching hadlers deployed yet")
		fmt.Println(message.Info)
		// response:="fickle is away \n ~q_me"
		// whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))

	}
	

}


func setReminder(message string)string{
// implement reminder

	return "Reminder set"
}



func getWhatsappLink(message string)string{
	
	re := regexp.MustCompile(`[1-9][\d]{9}`)
	whatsapp_number := re.FindStringSubmatch(message)
	message="https://wa.me/+91"+whatsapp_number[0]
	return message
}


func convertUnit2Unit(message string)string{

	re := regexp.MustCompile(`\w*2\w*.unit \d*[.]?\d*`)
	message_text := re.FindStringSubmatch(message)
	fmt.Println(message_text)
	message_text[0]=strings.ReplaceAll(message_text[0], " ", "")
	midPoint := strings.Index(message_text[0], "2")
	units:= strings.Index(message_text[0], ".unit")
	conv_from:=message_text[0][0:midPoint]
	conv_to:=message_text[0][midPoint+1:units]
	count:=message_text[0][units+5:]
	fmt.Println("from:",conv_from)
	fmt.Println("to:",conv_to)
	fmt.Println("count:",count)
	return unitcovertion.GetConverted(conv_from,conv_to,count)


	// return "Response after converting"
}

func convertCurrency2Currency(message string)string{
	re := regexp.MustCompile(`[A-Z]{3}2[A-Z]{3} [0-9]+([.][0-9])*`)
	message_text := re.FindStringSubmatch(strings.ToUpper(message))
	fmt.Println("\n Matched currency convertion ",message_text)
	conv_from:=message_text[0][0:3]
	conv_to:=message_text[0][4:7]
	count_str:=message_text[0][8:]
	if count, err := strconv.ParseFloat(count_str, 32); err == nil {
		fmt.Println((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
		response:= ((1/unitcovertion.GetCurrencyUnit(conv_from))*unitcovertion.GetCurrencyUnit(conv_to))
		response=response*float32(count)
		return fmt.Sprintf("%v",response)
		// whatsappClient.SendMessagetoWhatsapp(message.Info.Sender,"",fmt.Sprintf("%v",response))
	}else{
		response:="Error in the value given"
		return response
	}

}



