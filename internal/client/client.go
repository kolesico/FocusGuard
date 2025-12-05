package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kolesico/FocusGuard/internal/events"
)

func SendRequest(serverUri *string, event events.Events) {

	data, err := createPostRequest(event)
	if err != nil {
		log.Fatal("Ошибка при кодировании Event в json", err)
		return
	}

	req, err := http.NewRequest("POST", *serverUri, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Ошибка при создании запроса ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Ошибка при отправке запроса", err)
		return
	}

	defer resp.Body.Close()

	log.Printf("Ответ от сервера: %T\n", resp.Status)
}

func createPostRequest(event events.Events) ([]byte, error) {
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return data, nil
}