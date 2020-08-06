package utils

import nimo "github.com/gugemichael/nimo4go"

var (
	HttpApi *nimo.HttpRestProvider
)

func InitHttpApi(port int) {
	HttpApi = nimo.NewHttpRestProvider(port)
}