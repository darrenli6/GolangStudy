package sample_factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else {

		return &helloAPI{}
	}
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s", name)
}