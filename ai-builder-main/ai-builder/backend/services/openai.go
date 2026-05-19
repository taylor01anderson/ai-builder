package services

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "os"
)

func GenerateResponse(prompt string) (string, error) {

    apiKey := os.Getenv("OPENAI_API_KEY")

    payload := map[string]interface{}{
        "model": "gpt-4.1-mini",
        "messages": []map[string]string{
            {
                "role": "user",
                "content": prompt,
            },
        },
    }

    body, _ := json.Marshal(payload)

    req, _ := http.NewRequest(
        "POST",
        "https://api.openai.com/v1/chat/completions",
        bytes.NewBuffer(body),
    )

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}

    res, err := client.Do(req)

    if err != nil {
        return "", err
    }

    defer res.Body.Close()

    responseBody, _ := io.ReadAll(res.Body)

    return string(responseBody), nil
}