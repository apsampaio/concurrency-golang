package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("first", &wg)
	wg.Wait()
	printMessage()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	out := string(res)

	if !strings.Contains(out, "first") {
		t.Errorf("expected to find first, but is not there")
	}

}
