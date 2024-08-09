package esbuildmobile

// TODO: make into gomobile pkg
import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

type TransformOptions struct {
	// Change `React.createElement` to custom name
	JSXFactory string
	// Change `React.Fragment` to custom name
	JSXFragment string
}

// Returns code and error
// `TransformOptions` is optional
func TransformJSX(input string, options *TransformOptions) (code string, err error) {

	if options == nil {
		options = &TransformOptions{}
	}

	result := api.Transform(input, api.TransformOptions{
		Loader:      api.LoaderJSX,
		JSXFactory:  options.JSXFactory,
		JSXFragment: options.JSXFragment,
	})

	code = string(result.Code)

	if len(result.Errors) != 0 {
		err = fmt.Errorf("error: %v", result.Errors[0].Text)
	}

	return code, err
}
