package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSitesChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"wwww.google.com",
		"wwww.netflix.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"wwww.google.com":         true,
		"wwww.netflix.com":        true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebSitesChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "an url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
