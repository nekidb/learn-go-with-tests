package concurrency

import (
		"testing"
		"reflect"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://vk.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://doodle.com",
		"http://vk.com",
		"http://habr.ru",
	}
	
	want := map[string]bool{
		"http://doodle.com": 		true,
		"http://vk.com":		false,
		"http://habr.ru":		true,
	}
	
	got := CheckWebsites(mockWebsiteChecker, websites)
	
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
