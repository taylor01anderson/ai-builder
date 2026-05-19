//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"syscall/js"
)

func loadAssistants(this js.Value, args []js.Value) interface{} {

    res, err := http.Get("http://localhost:8080/api/assistants")
    if err != nil {
        js.Global().Get("console").Call("error", err.Error())
        return nil
    }
    defer func() {
        if res != nil && res.Body != nil {
            res.Body.Close()
        }
    }()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        js.Global().Get("console").Call("error", err.Error())
        return nil
    }

    var assistants []map[string]interface{}
    if err := json.Unmarshal(body, &assistants); err != nil {
        js.Global().Get("console").Call("error", err.Error())
        return nil
    }

    document := js.Global().Get("document")

    container := document.Call(
        "getElementById",
        "assistantList",
    )

    for _, assistant := range assistants {

        div := document.Call("createElement", "div")

        div.Set(
            "innerHTML",
            assistant["Name"],
        )

        container.Call("appendChild", div)
    }

    return nil
}