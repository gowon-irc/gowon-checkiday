package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	checkidayAPIURL = "https://checkiday.com/api/4/?d"
)

type daysJSON struct {
	Days        []Event `json:"events"`
	MultiEvents []Event `json:"multiday_ongoing"`
}

type Event struct {
	Name string `json:"name"`
}

func (ds daysJSON) listDays() (out []string) {
	for _, d := range ds.Days {
		out = append(out, d.Name)
	}
	return
}

func (ds daysJSON) listMdays() (out []string) {
	for _, d := range ds.MultiEvents {
		out = append(out, d.Name)
	}
	return
}

func (e Event) String() string {
	return e.Name
}

func colourList(in []string) (out []string) {
	colours := []string{"green", "red", "blue", "orange", "magenta", "cyan", "yellow"}
	cl := len(colours)

	for n, i := range in {
		c := colours[n%cl]
		o := fmt.Sprintf("{%s}%s{clear}", c, i)
		out = append(out, o)
	}

	return out
}

func checkiday() (msg string, err error) {
	data := &daysJSON{}

	res, err := http.Get(checkidayAPIURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", err
	}

	cl := colourList(data.listDays())

	return strings.Join(cl, ", "), nil
}

func checkmday() (msg string, err error) {
	data := &daysJSON{}

	res, err := http.Get(checkidayAPIURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", err
	}

	cl := colourList(data.listMdays())

	return strings.Join(cl, ", "), nil
}
