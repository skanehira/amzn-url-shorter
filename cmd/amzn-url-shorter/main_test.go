package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestShorter(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			url:  "https://www.amazon.co.jp/dp/4844336479",
			want: "https://amazon.co.jp/dp/4844336479",
		},
		{
			url:  "https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%82%B4%E3%83%AA%E3%83%A9%E3%81%AE%E3%82%82%E3%81%A3%E3%81%A8%EF%BC%81-%E5%85%83%E6%B0%97%E3%81%8C%E5%87%BA%E3%82%8B%E3%83%9E%E3%83%B3%E3%82%AC%E3%80%90%E9%9B%BB%E5%AD%90%E9%99%90%E5%AE%9A%E6%8F%8F%E3%81%8D%E4%B8%8B%E3%82%8D%E3%81%97%E4%BB%98%E3%81%8D%E3%80%91-%E3%82%AA%E3%82%BF%E3%82%AF%E3%81%A0%E3%82%88%EF%BC%81%E3%81%84%E3%81%8B%E3%82%B4%E3%83%AA%E3%83%A9%E3%81%AE%E5%85%83%E6%B0%97%E3%81%8C%E5%87%BA%E3%82%8B%E3%83%9E%E3%83%B3%E3%82%AC-%E3%83%9E%E3%83%BC%E3%82%AC%E3%83%AC%E3%83%83%E3%83%88%E3%82%B3%E3%83%9F%E3%83%83%E3%82%AF%E3%82%B9DIGITAL-%E3%81%84%E3%81%8B%E3%82%B4%E3%83%AA%E3%83%A9-ebook/dp/B087GHS748/ref=sr_1_1?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&dchild=1&keywords=%E3%82%B4%E3%83%AA%E3%83%A9&qid=1590497318&s=books&sr=1-1",
			want: "https://amazon.co.jp/dp/B087GHS748",
		},
		{
			url:  "https://www.amazon.co.jp/%E3%80%902020%E6%9C%80%E6%96%B0Bluetooth5-1-%E7%9E%AC%E9%96%93%E6%8E%A5%E7%B6%9A%E3%80%91Bluetooth-%E8%93%8B%E3%82%92%E9%96%8B%E3%81%91%E3%81%A6%E7%9E%AC%E9%96%93%E3%83%9A%E3%82%A2%E3%83%AA%E3%83%B3%E3%82%B0-CVC8-0%E3%83%8E%E3%82%A4%E3%82%BA%E3%82%AD%E3%83%A3%E3%83%B3%E3%82%BB%E3%83%AA%E3%83%B3%E3%82%B0-Android%E5%AF%BE%E5%BF%9C/dp/B083ZVZXSW/ref=sr_1_1_sspa?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&dchild=1&keywords=%E3%82%A4%E3%83%A4%E3%83%9B%E3%83%B3&qid=1590497355&sr=8-1-spons&psc=1&spLa=ZW5jcnlwdGVkUXVhbGlmaWVyPUEzR0g2VjFPMjZWQ1FXJmVuY3J5cHRlZElkPUEwOTAyMDgyMlZGTUJTUVhWUkJOQSZlbmNyeXB0ZWRBZElkPUEzUFQyUjlZSlpLNklCJndpZGdldE5hbWU9c3BfYXRmJmFjdGlvbj1jbGlja1JlZGlyZWN0JmRvTm90TG9nQ2xpY2s9dHJ1ZQ==",
			want: "https://amazon.co.jp/dp/B083ZVZXSW",
		},
	}

	for _, tt := range tests {
		got, err := shorter(tt.url)
		if err != nil {
			t.Errorf("failed get short url: error=%s", err)

		}

		if tt.want != got {
			t.Errorf("faield get short url: \nwant=%s \ngot=%s \n", tt.want, got)
		}
	}
}

func TestShorterFailed(t *testing.T) {
	tests := []struct {
		url  string
		want error
	}{
		{
			url:  "",
			want: fmt.Errorf("%s: %s", errInvalidURL, ""),
		},
		{
			url:  "https://www.amazon.co.jp",
			want: fmt.Errorf("%s: %s", errInvalidURL, "https://www.amazon.co.jp"),
		},
		{
			url:  "https://www.amazon.co.jp/dp/",
			want: fmt.Errorf("%s: %s", errInvalidURL, "https://www.amazon.co.jp/dp/"),
		},
	}

	for _, tt := range tests {
		_, err := shorter(tt.url)
		if !errors.As(err, &errInvalidURL) {
			t.Errorf("unexpected error: want=%v got=%v", tt.want, err)
		}
	}
}

func TestMultiShorter(t *testing.T) {
	wants := []string{
		"https://amazon.co.jp/dp/B087GHS748",
		"https://amazon.co.jp/dp/B083ZVZXSW",
	}

	f, err := os.Open("testdata/urls.txt")
	if err != nil {
		t.Fatalf("failed to read file %s", err)
	}

	urls, err := multishorter(f)
	if err != nil {
		t.Fatalf("failed to get short url: %s", err)
	}

	for i, url := range urls {
		if wants[i] != url {
			t.Errorf("unexpected url: \nwant=%s \ngot=%s \n", wants[i], url)
		}
	}
}
