package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/OKuharenok/go-counter/readers/filereader"
	"github.com/OKuharenok/go-counter/readers/urlreader"
	"github.com/OKuharenok/go-counter/types"
)

func main() {
	var k int
	flag.IntVar(&k, "k", 5, "Max goroutines")
	flag.Parse()

	if k < 1 {
		fmt.Printf("Invalid k: %d\n", k)
		os.Exit(1)
	}

	inProcess := make(chan struct{}, k)
	done := make(chan struct{})
	result := make(chan types.Result, k)
	total := 0
	wg := &sync.WaitGroup{}
	input := bufio.NewScanner(os.Stdin)

	go func() {
		for res := range result {
			if res.Error != nil {
				fmt.Printf("Error result for %s: %s\n", res.Path, res.Error)
				continue
			}
			total += res.Count
			fmt.Printf("Count for %s: %d\n", res.Path, res.Count)
		}
		done <- struct{}{}
	}()

	for input.Scan() {
		path := input.Text()
		inProcess <- struct{}{}
		wg.Add(1)
		go handle(path, inProcess, result, wg)
	}

	wg.Wait()
	close(result)
	<-done
	fmt.Printf("Total: %d\n", total)
}

func handle(path string, inProcess chan struct{}, result chan types.Result, wg *sync.WaitGroup) {
	defer func() {
		<-inProcess
		wg.Done()
	}()

	var reader types.Reader

	u, err := url.Parse(path)
	if err != nil {
		result <- types.Result{
			Path:  path,
			Error: err,
		}
		return
	}

	if u.Scheme == "" {
		reader = filereader.NewReader(path)
	} else {
		reader = urlreader.NewReader(path)
	}

	data, err := reader.Read()
	if err != nil {
		result <- types.Result{
			Path:  path,
			Error: err,
		}
		return
	}

	result <- types.Result{
		Path:  path,
		Count: strings.Count(string(data), "Go"),
		Error: nil,
	}
}
