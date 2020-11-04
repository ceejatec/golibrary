package golibrary

import (
	hello "github.com/ceejatec/hello-go-mod"
)

func Greeting() string {
	return hello.Hello() + " from library 1.0"
}
