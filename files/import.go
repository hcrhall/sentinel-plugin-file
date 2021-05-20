package files

import (
	"os"

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
	case "read":
		return func(path string) (interface{}, error) {
			contents, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}
			return string(contents), nil
		}

	case "list":
		return func(path string) (interface{}, error) {
			var files []map[string]interface{}

			fileInfo, err := os.ReadDir(path)
			if err != nil {
				return nil, err
			}

			for _, file := range fileInfo {
				m := map[string]interface{}{
					"name":      file.Name(),
					"directory": file.IsDir(),
				}
				files = append(files, m)
			}
			return files, nil
		}
	}

	return nil
}
