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

type Request struct {
    Url, Data string
}

func (r *Request) SetURL(parts ...string) {
	var url strings.Builder

    for i := 0; i < len(parts); i++ {
        url.WriteString(parts[i])
    }

	r.Url = url.String()
}

func (r *Request) SetJSONData(jsonMap map[string]string) {
    jsonString, _ := json.Marshal(jsonMap)
    r.Data = string(jsonString)
}

func (r *Request) Get() string {
    timeout := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeout,
    }

    request, err := http.NewRequest("GET", r.Url, bytes.NewBufferString(r.Data))

    if err != nil {
		panic(err)
	}

    response, err := client.Do(request)

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

func (r *Request) Post() string {
    returned := ""
    var postData = []byte("")
    var postType = "text/plain"

    if isJSON(r.Data) {
        rawIn := json.RawMessage(r.Data)
        jsonData, err := rawIn.MarshalJSON()
        postData = jsonData
        postType = "application/json"

        if err != nil {
            panic(err)
        }
    } else {
        postData = []byte(url.QueryEscape(r.Data))
    }

    request, err := http.NewRequest(http.MethodPost, r.Url, bytes.NewReader(postData))

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
