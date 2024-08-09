package esbuildmobile

// TODO: make into gomobile pkg
import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

// Returns code and error
// `TransformOptions` is optional
func TransformJSX(input string, options *TransformOptions) (code string, err error) {

	if options == nil {
		options = &TransformOptions{}
	}

	result := api.Transform(input, api.TransformOptions{
		Loader:      api.LoaderJSX,
		JSXFactory:  options.jSXFactory,
		JSXFragment: options.jSXFragment,
	})

	code = string(result.Code)

	if len(result.Errors) != 0 {
		err = fmt.Errorf("error: %v", result.Errors[0].Text)
	}

	return code, err
}
