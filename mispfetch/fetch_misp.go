package mispfetch

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
	"unicode"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

type Attribute struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MISPResponse struct {
	Response struct {
		Attribute []Attribute `json:"Attribute"`
	} `json:"response"`
}

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func FetchMISPData(rdb *redis.Client, url string, apiKey string) {
	page := 1
	for {
		body := map[string]interface{}{
			"returnFormat": "json",
			"limit":        100,
			"page":         page,
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		req.Header.Set("Authorization", apiKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		client := &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		resp, err := client.Do(req)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		defer resp.Body.Close()

		var mispResp MISPResponse
		err = json.NewDecoder(resp.Body).Decode(&mispResp)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		attributes := mispResp.Response.Attribute
		if len(attributes) == 0 {
			page = 1
			time.Sleep(10 * time.Second)
			continue
		}

		for _, attr := range attributes {
			if attr.Value == "" || isNumeric(attr.Value) || attr.Value == attr.ID {
				continue
			}

			exist, err := rdb.HExists(Ctx, "api_misp", attr.Value).Result()
			if err != nil || exist {
				continue
			}

			data := map[string]string{
				"id":    attr.ID,
				"type":  attr.Type,
				"value": attr.Value,
			}
			jsonData, _ := json.Marshal(data)
			_ = rdb.HSet(Ctx, "api_misp", attr.Value, jsonData).Err()
		}

		page++
		time.Sleep(10 * time.Second)
	}
}
