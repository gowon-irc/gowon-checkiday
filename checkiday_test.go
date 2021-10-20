package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	d1 = Day{
		Name: "test day",
	}

	d2 = Day{
		Name: "test day 2",
	}

	dj = daysJSON{
		Days: []Day{d1, d2},
	}
)

func TestDay(t *testing.T) {
	assert.Equal(t, d1.String(), "test day")
}

func TestDays(t *testing.T) {
	expected := []string{
		"test day",
		"test day 2",
	}

	assert.Equal(t, dj.List(), expected)
}
