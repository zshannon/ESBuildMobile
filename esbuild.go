package esbuildmobile

// TODO: make into gomobile pkg
import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

type TransformOptions struct {
	JSXFactory string
}

// `JSXFactory` will replace the creation
func NewTransformOptions(JSXFactory string) *TransformOptions {
	return &TransformOptions{}
}

// Returns code and error
// `TransformOptions` is optional
func Transform(input string, options *TransformOptions) (code string, err error) {
	JSXFactory := ""

	if options != nil {
		JSXFactory = options.JSXFactory
	}

	result := api.Transform(input, api.TransformOptions{
		Loader:     api.LoaderJSX,
		Target:     api.Target(api.EngineIOS),
		JSXFactory: JSXFactory,
	})

	code = string(result.Code)

	if len(result.Errors) != 0 {
		err = fmt.Errorf("error: %v", result.Errors[0].Text)
	}

	return code, err
}
