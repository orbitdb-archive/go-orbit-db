package api

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"
    "strings"
    "bytes"
)

func BuildURL(parts ...string) string {
	var urlBuilder strings.Builder

    for i := 0; i < len(parts); i++ {
        urlBuilder.WriteString(parts[i])
    }

	return urlBuilder.String()
}

func MapToJSONString(jsonMap map[string]string) string {
    jsonString, _ := json.Marshal(jsonMap)
    return string(jsonString)
}

func Get(url string) string {
    timeout := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeout,
    }

    response, err := client.Get(url)

    if err != nil {
        panic(err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        return string(data)
    }

    return ""
}

func Post(url string, params string) string {
    rawIn := json.RawMessage(params)
    jsonValue, err := rawIn.MarshalJSON()

    if err != nil {
        panic(err)
    }

    timeout := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeout,
    }

    response, err := client.Post(url, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        panic(err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        return string(data)
    }

    return ""
}
