package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	apiclient "s21calendar/api"
	auth "s21calendar/utils"
	"time"
)

const serverUrl string = "https://edu-api.21-school.ru/services/21-school/api"

func main() {

	log.SetFlags(log.Lshortfile)

	securityHandler := auth.CredentialsStorage{
		Username: os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}
	client, err := apiclient.NewClient(serverUrl, &securityHandler)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.GetEvents(context.TODO(), apiclient.GetEventsParams{
		From:   time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC),
		To:     time.Date(2024, time.June, 25, 0, 0, 0, 0, time.UTC),
		Limit:  apiclient.OptInt64{Value: 1, Set: true},
		Offset: apiclient.OptInt64{Value: 2, Set: true},
	})
	if err != nil {
		log.Fatalln(err)
	}

	switch p := res.(type) {
	case *apiclient.EventsV1DTO:
		data, err := p.MarshalJSON()
		if err != nil {
			log.Fatalln(err)
		}
		var out bytes.Buffer
		if err := json.Indent(&out, data, "", "  "); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(out.String())
	default:
		fmt.Println("error")
	}
}
