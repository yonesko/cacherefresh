package cacherefresh

import (
	"errors"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	t.Run("no err case", func(t *testing.T) {
		c, err := New(func() (string, error) {
			return "Hello", nil
		}, time.Hour)
		if err != nil {
			t.Fatal(err)
		}
		data, err := c.Get()
		if err != nil {
			t.Fatal(err)
		}
		if data != "Hello" {
			t.Errorf("got %s, want Hello", data)
		}
	})
	t.Run("err case", func(t *testing.T) {
		_, err := New(func() (string, error) {
			return "", errors.New("some error")
		}, time.Hour)
		if err == nil {
			t.Fatal(err)
		}
	})
	t.Run("success after err case", func(t *testing.T) {
		try := 0
		c, err := New(func() (string, error) {
			try++
			if try == 1 {
				return "Hello", nil
			}
			switch try {
			case 1:
				return "Hello", nil
			case 2:
				return "", errors.New("some error")
			default:
				return "Hello", nil
			}
		}, time.Hour)
		if err != nil {
			t.Fatal(err)
		}

		c.Refresh()

		data, err := c.Get()
		if err == nil {
			t.Fatal(err)
		}
		if data != "" {
			t.Errorf("got %s, want empty", data)
		}

		c.Refresh()

		data, err = c.Get()
		if err != nil {
			t.Fatal(err)
		}
		if data != "Hello" {
			t.Errorf("got %s, want Hello", data)
		}
	})
}
