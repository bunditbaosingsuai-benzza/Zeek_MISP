package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"project-finals/mispfetch"
	"project-finals/receiver"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found or failed to load:", err)
	}

	mispURL := os.Getenv("MISP_URL")
	mispAPIKey := os.Getenv("MISP_API_KEY")

	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	receiver.StartLogReceiver("5050", rdb)

	go func() {
		for {
			mispfetch.FetchMISPData(rdb, mispURL, mispAPIKey)
			time.Sleep(5 * time.Minute)
		}
	}()

	for logLine := range receiver.LogChannel {
		checkLogAgainstRedis(rdb, logLine)
	}
}

func checkLogAgainstRedis(rdb *redis.Client, log string) {
	keys, err := rdb.HKeys(ctx, "api_misp").Result()
	if err != nil {
		return
	}

	for _, key := range keys {
		if strings.Contains(log, key) {
			val, err := rdb.HGet(ctx, "api_misp", key).Result()
			if err == nil {
				var raw map[string]string
				json.Unmarshal([]byte(val), &raw)

				if isDangerousType(raw["type"]) {
					loc, _ := time.LoadLocation("Asia/Bangkok")
					entry := struct {
						Timestamp string `json:"timestamp"`
						ID        string `json:"id"`
						Type      string `json:"type"`
						Value     string `json:"value"`
					}{
						Timestamp: time.Now().In(loc).Format("2006-01-02T15:04:05-0700"),
						ID:        raw["id"],
						Type:      raw["type"],
						Value:     raw["value"],
					}

					jsonLog, _ := json.Marshal(entry)
					fmt.Println(string(jsonLog))
					appendToLogFile("/alert.log", string(jsonLog)+"\n")
				}
			}
		}
	}
}

func isDangerousType(t string) bool {
	switch t {
	case "ip", "ip-src", "ip-dst", "domain", "url", "link", "md5", "sha1", "sha256", "host":
		return true
	default:
		return false
	}
}

func appendToLogFile(path string, log string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("# Unable to log :", err)
		return
	}
	defer f.Close()
	f.WriteString(log)
}
