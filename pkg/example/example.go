package example

//go:generate mockgen -destination=../../mock/mock_example.go -package=mock -source=./example.go  Example

import "errors"

type Example interface {
	Get(key string) (value string, err error)
	Set(key string, value string) error
	Test()
}

type exampleImpl struct {
	values map[string]string
}

func New() Example {
	return &exampleImpl{make(map[string]string)}
}

func (e *exampleImpl) Set(key string, value string) error {
	e.values[key] = value
	return nil
}

func (e *exampleImpl) Get(key string) (value string, err error) {
	v, f := e.values[key]
	if !f {
		return "", errors.New("key not found")
	}
	return v, nil
}

func (e *exampleImpl) Test() {

}
