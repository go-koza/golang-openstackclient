package utils

import (
    "bytes"
    "errors"
    "encoding/json"
    "net/http"
)

func PostJSON(url string, input interface{}, output interface{}) (err error) {
    client := http.Client{}

    inputJson, err := json.Marshal(input)
    if err != nil {
        return err
    }

    request, err := http.NewRequest("POST", url, bytes.NewBuffer(inputJson))
    if err != nil {
        return err
    }

    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Accept", "application/json")
    request.Header.Set("X-Auth-Token", "")
    response, err := client.Do(request)
    if err != nil {
        return err
    }

    if response.StatusCode != 201 && response.StatusCode != 202 {
        err = errors.New("Error: status code != 201 or 202, actual status code '" + response.Status + "'")
        return err
    }

    contentType := response.Header.Get("Content-Type")
    if contentType != "application/json" {
        err = errors.New("Error: Expected a json payload but instead recieved '" + contentType + "'")
        return err
    }

    err = json.NewDecoder(response.Body).Decode(&output)
    defer response.Body.Close()
    if err != nil {
        return err
    }

    return nil
}
