package main

import (
	"testing"
	"time"
)

func BenchmarkCheckWebsites(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, 
