package file

import (
	"io/ioutil"

	sdk "github.com/hashicorp/sentinel-sdk"
	framework "github.com/hashicorp/sentinel-sdk/framework"
)

func New() sdk.Import {
	return &framework.Import{
		Root: &root{},
	}
}

type root struct {
}

// framework.Root impl.
func (m *root) Configure(raw map[string]interface{}) error {
	return nil
}

// framework.Namespace impl.
func (m *root) Get(key string) (interface{}, error) {
	return nil, nil
}

// framework.Call impl.
func (m *root) Func(key string) interface{} {
	switch key {
	case "open":
		return func(path string) (interface{}, error) {
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}
			return string(contents), nil
		}
	}
	return nil
}
