package main

import (
	"encoding/json"
)

type SMWResponse struct {
	Query struct {
		Results    SMWResponseResult `json:"results"`
		Serializer string            `json:"serializer"`
		Version    int               `json:"version"`
		Meta       SMWResponseMeta   `json:"meta"`
	} `json:"query"`
}

type SMWResponseMeta struct {
	Hash   string `json:"hash"`
	Count  int    `json:"count"`
	Offset int    `json:"offset"`
	Source string `json:"source"`
	Time   string `json:"time"`
}

type SMWResponseResult struct {
	Entries []SMWResponseResultEntry
}

type SMWResponseResultEntry struct {
	Printouts struct {
		Code  []string `json:"code"`
		Color []string `json:"color"`
		Name  []string `json:"name"`
		Start []struct {
			Timestamp string `json:"timestamp"`
			Raw       string `json:"raw"`
		} `json:"start"`
		End []struct {
			Timestamp string `json:"timestamp"`
			Raw       string `json:"raw"`
		} `json:"end"`
		Desc []string `json:"desc"`
		Icon []struct {
			Fulltext     string `json:"fulltext"`
			Fullurl      string `json:"fullurl"`
			Namespace    int    `json:"namespace"`
			Exists       string `json:"exists"`
			Displaytitle string `json:"displaytitle"`
		} `json:"icon"`
	} `json:"printouts"`
	Fulltext     string `json:"fulltext"`
	Fullurl      string `json:"fullurl"`
	Namespace    int    `json:"namespace"`
	Exists       string `json:"exists"`
	Displaytitle string `json:"displaytitle"`
}

type ParseResponse struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Text   struct {
			Content string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
}

type ApiResult struct {
	Results []ApiResultEntry `json:"results"`
	Version string           `json:"version"`
	Meta    SMWResponseMeta  `json:"meta"`
}

type ApiResultEntry struct {
	Id       string `json:"id"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
	StartStr string `json:"startStr"`
	EndStr   string `json:"endStr"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Url      string `json:"url"`
	Icon     string `json:"icon,omitempty"`
	Code     string `json:"code,omitempty"`
	Color    string `json:"color,omitempty"`
}

func (r *SMWResponseResult) UnmarshalJSON(data []byte) error {
	var m map[string]SMWResponseResultEntry

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	r.Entries = make([]SMWResponseResultEntry, len(m))
	i := 0

	for _, v := range m {
		r.Entries[i] = v
		i++
	}

	return nil
}
