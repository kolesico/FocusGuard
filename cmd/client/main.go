package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	// "github.com/kolesico/FocusGuard/internal/client"
	"github.com/kolesico/FocusGuard/internal/monitor"
)

func main() {

	appName := flag.String("app", "Telegram.exe", "name of app for monitor")
	// serverUri := flag.String("server-uri", "http://127.0.0.1:8080/events", "server uri")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	events := monitor.RunMonitor(ctx, appName)

	for event := range events {
		log.Printf("%s: %s %s", event.Timestamp.Format(time.RFC3339), *appName, event.Event)
		// go client.SendRequest(serverUri, event)
	}
}