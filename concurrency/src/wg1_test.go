 package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething("epic", &wg)
	wg.Wait()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "epic") {
		t.Errorf("Expected to find epic , but not there")
	}
}
