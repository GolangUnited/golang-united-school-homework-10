package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	host = "localhost"
	port = 8081
)

var baseURL = fmt.Sprintf("http://%s:%d/", host, port)

var ErrTimeout = errors.New("exited by timeout")

func waitForIt(host string, port int, timeout int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	ch := make(chan struct{})
	//errCh := make(chan error)
	go func() {
		step := 0
		for {
			fmt.Printf("Asking %s %2d of %2d\n", addr, step, timeout)
			step++
			c, err := net.Dial("tcp", addr)
			if err == nil {
				c.Close()
				ch <- struct{}{}
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-ch:
		return nil
	case <-time.After(time.Duration(timeout) * time.Second):
		return ErrTimeout
	}
}

func testInit(m *testing.M) int {
	go Start(host, port)
	waitForIt(host, port, 30)
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testInit(m))
}

func getPage(method, url string, body io.Reader) (*http.Response, error) {
	path := baseURL + url
	r, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	c := http.Client{}
	w, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func Test_handleName_01(t *testing.T) {
	data, err := getPage(http.MethodGet, "name/Tester", nil)
	require.NoError(t, err)
	b, err := ioutil.ReadAll(data.Body)
	require.NoError(t, err)
	assert.Equal(t, "Hello, Tester!", string(b))
}

func Test_handleBad_01(t *testing.T) {
	data, err := getPage(http.MethodGet, "bad", nil)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, data.StatusCode)
}

func Test_handlePost_01(t *testing.T) {
	data, err := getPage(http.MethodPost, "data", bytes.NewBuffer([]byte("my message")))
	require.NoError(t, err)
	b, err := ioutil.ReadAll(data.Body)
	require.NoError(t, err)
	assert.Equal(t, "I got message:\nmy message", string(b))
}
