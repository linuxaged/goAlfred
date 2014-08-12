package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "os"
)

type Query struct {
	Data string `xml:",chardata"`
}

type Paragraph struct {
	Data string `xml:",chardata"`
}

type Translation struct {
	Paragraph Paragraph `xml:"paragraph"`
}

// type basic
type Phonetic struct {
	Data string `xml:",chardata"`
}

type US_Phonetic struct {
	Data string `xml:",chardata"`
}

type UK_Phonetic struct {
	Data string `xml:",chardata"`
}

type Ex struct {
	Data string `xml:",chardata"`
}

type Explains struct {
	Ex Ex `xml:"ex"` // TODO
}

type Basic struct {
	Phonetic    Phonetic    `xml:"phonetic"`
	US_Phonetic US_Phonetic `xml:"us-phonetic"`
	UK_Phonetic UK_Phonetic `xml:"uk-phonetic"`
	Explains    Explains    `xml:"explains"`
}

// type web
type Key struct {
	Data string `xml:",chardata"`
}
type Value struct {
	Ex Ex `xml:"ex"` // TODO
}
type Explain struct {
	Key   Key   `xml:"key"`
	Value Value `xml:"value"`
}
type Web struct {
	Explain Explain `xml:"explain"` // TODO
}

// type youdao-fanyi
type youdao_fanyi struct {
	XMLName     xml.Name    `xml:"youdao-fanyi"`
	ErrorCode   int         `xml:"errorCode"`
	Query       Query       `xml:"query"`
	Translation Translation `xml:"translation"`
	Basic       Basic       `xml:"basic"`
	Web         Web         `xml:"web"`
}

func main() {
	var url_yd *url.URL
	url_yd, err := url.Parse("http://fanyi.youdao.com")
	if err != nil {
		panic(err.Error())
	}
	url_yd.Path += "/openapi.do?"
	param := url.Values{}
	param.Add("keyfrom", "rediffuse")
	param.Add("key", "1698275791")
	param.Add("type", "data")
	param.Add("doctype", "xml")
	param.Add("version", "1.1")
	param.Add("q", "fuck")
	url_yd.RawQuery = param.Encode()
	// if len(os.Args) == 2 {
	result, err := http.Get(url_yd.String())
	if err != nil {
		panic(err.Error())
	} else {
		defer result.Body.Close()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			panic(err.Error())
		} else {
			yd_fy := youdao_fanyi{}
			err := xml.Unmarshal(body, &yd_fy)
			if err != nil {
				panic(err.Error())
			} else {
				fmt.Println(yd_fy.Query.Data)
			}
			// fmt.Println(body)
		}
	}
	// }

}
