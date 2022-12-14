package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/C-m3-Codin/q_me/handler"
	"github.com/C-m3-Codin/q_me/services"
	"github.com/C-m3-Codin/q_me/unitcovertion"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)





func main() {
    dbLog := waLog.Stdout("Database", "DEBUG", true)
    // Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
    container, err := sqlstore.New("sqlite3", "./names.db", dbLog)
    if err != nil {
        panic(err)
    }
    // If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
    deviceStore, err := container.GetFirstDevice()
    if err != nil {
        panic(err)
    }
    clientLog := waLog.Stdout("Client", "ERROR", true)
    client := whatsmeow.NewClient(deviceStore, clientLog)
    whatsappHandler:=handler.SetClient(client)
    
    client.AddEventHandler(whatsappHandler.EventHandler)

    if client.Store.ID == nil {
        // No ID stored, new login
        qrChan, _ := client.GetQRChannel(context.Background())
        err = client.Connect()
        if err != nil {
            panic(err)
        }
        for evt := range qrChan {
            if evt.Event == "code" {
                // Render the QR code here
                // e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
                // or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("\n\n\n\n\n\n\n start")
                fmt.Println("QR code:\n -", evt.Code,"-")
				fmt.Println("\n\n\n\n\n\n\n end","apple")
            } else {
                fmt.Println("Login event:", evt.Event)
            }
        }
    } else {
        // Already logged in, just connect
        err = client.Connect()
        if err != nil {
            panic(err)
        }
        go services.ScheduleCurrency()
        go unitcovertion.LoadUnitConversions("./unitConversion/units.csv")
    }


    
    // Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c

    client.Disconnect()
}


