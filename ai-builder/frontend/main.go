//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"syscall/js"
)

func sendPrompt(this js.Value, args []js.Value) interface{} {
    go func() {
        document := js.Global().Get("document")

        prompt := document.Call(
            "getElementById",
            "chatInput",
        ).Get("value").String()

        payload := map[string]string{
            "prompt": prompt,
        }

        body, err := json.Marshal(payload)
        if err != nil {
            js.Global().Call("console.error", "Failed to marshal payload:", err.Error())
            document.Call("getElementById", "response").Set("innerText", "Error: Failed to process request")
            return
        }

        req, err := http.NewRequest(
            "POST",
            "http://localhost:8080/api/chat",
            bytes.NewBuffer(body),
        )
        if err != nil {
            js.Global().Call("console.error", "Failed to create request:", err.Error())
            document.Call("getElementById", "response").Set("innerText", "Error: Failed to create request")
            return
        }

        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}

        res, err := client.Do(req)

        if err != nil {
            js.Global().Call("console.error", "Chat request failed:", err.Error())
            document.Call("getElementById", "response").Set("innerText", "Error: "+err.Error())
            return
        }

        defer res.Body.Close()

        responseBody, _ := io.ReadAll(res.Body)

        var data map[string]interface{}
        json.Unmarshal(responseBody, &data)

        var content string
        if response, ok := data["response"].(string); ok {
            var chatResponse map[string]interface{}
            json.Unmarshal([]byte(response), &chatResponse)
            if choices, ok := chatResponse["choices"].([]interface{}); ok && len(choices) > 0 {
                if choice, ok := choices[0].(map[string]interface{}); ok {
                    if message, ok := choice["message"].(map[string]interface{}); ok {
                        if msg, ok := message["content"].(string); ok {
                            content = msg
                        }
                    }
                }
            }
        }

        document.Call(
            "getElementById",
            "chatOutput",
        ).Set("innerText", content)
    }()

    return nil
}

func main() {

    js.Global().Set(
        "sendPrompt",
        js.FuncOf(sendPrompt),
    )

    select {}
}