package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool{
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T){
	websites := []string{
		"http://google.com",
		"waat://furhurterwe.geds",
		"http://facebook.com",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://facebook.com": true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}