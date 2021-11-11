package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yyoshiki41/go-radiko"
)

func main() {

	// if an auth_token is not necessary.
	client, err := radiko.New("")
	if err != nil {
		log.Fatalf("Failed to construct a radiko Client. %s", err)
	}
	// // Get stations data
	// stations, err := client.GetNowPrograms(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = client.AuthorizeToken(ctx)
	if err != nil {
		log.Fatal(err)
	}

	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	start, err := time.ParseInLocation("20060102150405", "20211107010000", location)
	url, err := client.TimeshiftPlaylistM3U8(ctx, "LFR", start)
	chunklist, err := radiko.GetChunklistFromM3U8(url)

	fmt.Println(chunklist[0])
}
