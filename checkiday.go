package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	checkidayAPIURL = "https://checkiday.com/api/3/?d"
)

type daysJSON struct {
	Days []Day `json:"holidays"`
}

func (ds daysJSON) List() (out []string) {
	for _, d := range ds.Days {
		out = append(out, d.Name)
	}
	return
}

// Day represents a single day result from checkiday
type Day struct {
	Name string `json:"name"`
}

func (d Day) String() string {
	return d.Name
}

func colourList(in []string) (out []string) {
	colours := []string{"green", "red", "blue", "magenta", "orange", "yellow", "cyan"}
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", err
	}

	cl := colourList(data.List())

	return strings.Join(cl, ", "), nil
}
