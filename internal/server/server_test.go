package server

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer(address string) {
	s := ShortURLServer{Address: address}
	s.Init()
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestRunning(t *testing.T) {
	go StartServer("localhost:3000")

	err := CheckServerState()
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGet(t *testing.T) {
	go func() {
		s := ShortURLServer{Address: "localhost:3001"}
		s.Init()
		s.AddGET("/g", func(context *gin.Context) {
			context.Data(http.StatusOK, "text/plain", []byte("hello"))
		})
		t.Error(s.Run())
		return
	}()

	resp, _ := http.Get("http://localhost:3001/g/")
	body, _ := io.ReadAll(resp.Body)

	if strings.Compare(string(body), "hello") != 0 {
		t.Fatal()
	}
}

type Value struct {
	Name string `json:"name"`
}

func TestPost(t *testing.T) {
	values := map[string]string{"name": "go"}
	jsonData, _ := json.Marshal(values)
	go func() {
		s := ShortURLServer{Address: "localhost:3002"}
		s.Init()
		s.AddPOST("/p", func(context *gin.Context) {
			var data Value
			if err := context.BindJSON(&data); err != nil {
				t.Error(err)
				return
			}
			if data.Name != values["name"] {
				t.Error()
				return
			}
		})
		t.Error(s.Run())
		return
	}()

	_, _ = http.Post("http://localhost:3002/p", "text/plain", bytes.NewBuffer(jsonData))
}
func TestPut(t *testing.T) {
	values := map[string]string{"name": "go"}
	jsonData, _ := json.Marshal(values)
	go func() {
		s := ShortURLServer{Address: "localhost:3003"}
		s.Init()
		s.AddPUT("/put", func(context *gin.Context) {
			var data Value
			if err := context.BindJSON(&data); err != nil {
				t.Error(err)
				return
			}
			if data.Name != values["name"] {
				t.Error()
				return
			}
		})
		t.Error(s.Run())
		return
	}()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3003/put", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, _ = client.Do(req)
}

type Index struct {
	Id string `json:"Index"`
}

func TestDelete(t *testing.T) {
	values := map[string]string{"Index": "1"}
	jsonData, _ := json.Marshal(values)
	go func() {
		s := ShortURLServer{Address: "localhost:3004"}
		s.Init()
		s.AddDELETE("/del", func(context *gin.Context) {
			var data Index
			if err := context.BindJSON(&data); err != nil {
				t.Error(err)
				return
			}
			i, _ := strconv.Atoi(values["Index"])
			j, _ := strconv.Atoi(data.Id)
			if j != i {
				t.Error()
				return
			}
		})
		t.Error(s.Run())
		return
	}()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3004/del", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, _ = client.Do(req)
}

func TestRouters(t *testing.T) {
	go StartServer("localhost:3005")

	_, err := http.Get("http://localhost:3005/s/twitter.com")

	if err != nil {
		t.Fatal(err)
	}

	_, err = http.Get("http://localhost:3005/u/699caeda1ab7477b")

	if err != nil {
		t.Fatal(err)
	}
}

func CheckServerState() error {
	host := "localhost"
	port := "3000"
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		return err
	}
	return nil
}
