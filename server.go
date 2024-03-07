package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/webhook", func(c echo.Context) error {

		// Generate from https://webhook.site/
		WEBHOOK_URL := "https://webhook.site/8d9ae884-5860-4393-ae5d-f6ac935516e9?auth=123456"

		jsonData := []byte(`{
			"name": "morpheus",
			"job": "leader"
		}`)

		// Prepare webhook request
		req, err := http.NewRequest("POST", WEBHOOK_URL, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		// Send webhook request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer func(Body io.ReadCloser) {
			if err := Body.Close(); err != nil {
				log.Println("Error closing response body:", err)
			}
		}(resp.Body)

		// Check response code
		if resp.StatusCode == http.StatusOK {
			return c.String(http.StatusOK, "Webhook sent")
		} else if resp.StatusCode == http.StatusNotFound {
			return c.String(http.StatusNotFound, "Webhook not found")
		} else {
			return errors.New("unexpected response code")
		}

	})
	e.Logger.Fatal(e.Start(":1323"))
}
