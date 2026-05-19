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

    document := js.Global().Get("document")

    prompt := document.Call(
        "getElementById",
        "prompt",
    ).Get("value").String()

    payload := map[string]string{
        "prompt": prompt,
    }

    body, _ := json.Marshal(payload)

    req, _ := http.NewRequest(
        "POST",
        "http://localhost:8080/api/chat",
        bytes.NewBuffer(body),
    )

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}

    res, err := client.Do(req)

    if err != nil {
        return nil
    }

    defer res.Body.Close()

    responseBody, _ := io.ReadAll(res.Body)

    document.Call(
        "getElementById",
        "response",
    ).Set("innerText", string(responseBody))

    return nil
}

func main() {

    js.Global().Set(
        "sendPrompt",
        js.FuncOf(sendPrompt),
    )

    select {}
}