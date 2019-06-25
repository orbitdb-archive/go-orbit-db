package http

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"
    "strings"
    "bytes"
    "net/url"
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

func isJSON(str string) bool {
    var js json.RawMessage
    return json.Unmarshal([]byte(str), &js) == nil
}

func Post(uri string, data string) string {
    returned := ""
    var postData = []byte("")
    var postType = "text/plain"

    if isJSON(data) {
        rawIn := json.RawMessage(data)
        jsonData, err := rawIn.MarshalJSON()
        postData = jsonData
        postType = "application/json"

        if err != nil {
            panic(err)
        }
    } else {
        postData = []byte(url.QueryEscape(data))
    }

    request, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(postData))

    request.Header.Set("Content-Type", postType)

    timeout := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeout,
    }

    response, err := client.Do(request)

    if err != nil {
        panic(err)
    } else {
        body, _ := ioutil.ReadAll(response.Body)

        returned = string(body)
    }

    return returned
}
