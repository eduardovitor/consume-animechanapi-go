/*
This program must use some available routes on the animechan.xyz API to generate some output
*/

package main

import (
	"fmt"
    "net/http"
    "encoding/json"
    "io"
)

const (
    baseURL = "https://animechan.xyz/api/"
)

type Quote struct {
    Anime     string `json:"anime"`
    Character string `json:"character"`
    Quote     string `json:"quote"`
}

func getQuoteObj(origin_obj Quote, resp http.Response) Quote {
    defer resp.Body.Close() // Close the response body when done
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return origin_obj
    }
    err = json.Unmarshal(body, &origin_obj)
    if err != nil {
        fmt.Println(err)
        return origin_obj
    }
    return origin_obj
}

func getQuoteObjs(origin_objs []Quote, resp http.Response) []Quote {
    defer resp.Body.Close() // Close the response body when done
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return origin_objs
    }
    err = json.Unmarshal(body, &origin_objs)
    if err != nil {
        fmt.Println(err)
        return origin_objs
    }
    return origin_objs
}
func getRandomQuote() Quote {
    respQuote := Quote{}
    resp, err := http.Get(baseURL+"random")
    if err != nil {
        fmt.Println(err)
        return respQuote
    }
    respQuote = getQuoteObj(respQuote,*resp)
    return respQuote
}

func getQuoteByAnime(anime string) Quote {
    respQuote := Quote{}
    resp, err := http.Get(baseURL+"random/anime?title="+anime)
    if err != nil {
        fmt.Println(err)
        return respQuote
    }
    respQuote = getQuoteObj(respQuote,*resp)
    return respQuote
}

func getQuoteByCharacter(character string) Quote {
    respQuote := Quote{}
    resp, err := http.Get(baseURL+"random/character?name="+character)
    if err != nil {
        fmt.Println(err)
        return respQuote
    }
    respQuote = getQuoteObj(respQuote,*resp)
    return respQuote
}

func getTenQuotes() []Quote{
    respQuotes := []Quote{}
    resp, err := http.Get(baseURL+"quotes")
    if err != nil {
        fmt.Println(err)
        return respQuotes
    }
    respQuotes = getQuoteObjs(respQuotes,*resp)
    return respQuotes
}
func main() {
    quotes := getTenQuotes()
    for _, quote := range quotes {
        quote_print, err := json.Marshal(&quote)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(string(quote_print))
    }
}
