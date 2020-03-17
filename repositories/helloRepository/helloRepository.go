package helloRepository

import (
	"github.com/panupongdeve/gin-web-start/datasource/text/hello"
)

func GetHello() string {
	return hello.GetHelloText()
}
