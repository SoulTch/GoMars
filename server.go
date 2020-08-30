package main

import (
	"math/rand"

	grpc "./core/github.com/SoulTch/gomars_go"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var userID = map[string]int{}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type gomarsServer struct {
}

func (s *gomarsServer) Action(stream grpc.GoMars_ActionServer) error {

}

func (s *gomarsServer) Notive(stream grpc.GoMars_NoticeServer) error {

}
