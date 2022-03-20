package utils

import (
	"net"
	"shorturl/internal/storage"
	"time"
)

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

func StartUpStorage() *storage.Storage {
	s := storage.Storage{FilePath: "C:\\Go\\Saved\\data.json"}
	s.Init()
	return &s
}
