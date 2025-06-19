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

	loader := options.Loader
	if loader == 0 {
		loader = api.LoaderJSX
	}

	result := api.Transform(input, api.TransformOptions{
		Color:             options.Color,
		LogLevel:          options.LogLevel,
		LogLimit:          options.LogLimit,
		LogOverride:       options.LogOverride,
		Sourcemap:         options.Sourcemap,
		SourceRoot:        options.SourceRoot,
		SourcesContent:    options.SourcesContent,
		Target:            options.Target,
		Engines:           options.Engines,
		Supported:         options.Supported,
		Platform:          options.Platform,
		Format:            options.Format,
		GlobalName:        options.GlobalName,
		MangleProps:       options.MangleProps,
		ReserveProps:      options.ReserveProps,
		MangleQuoted:      options.MangleQuoted,
		MangleCache:       options.MangleCache,
		Drop:              options.Drop,
		DropLabels:        options.DropLabels,
		MinifyWhitespace:  options.MinifyWhitespace,
		MinifyIdentifiers: options.MinifyIdentifiers,
		MinifySyntax:      options.MinifySyntax,
		LineLimit:         options.LineLimit,
		Charset:           options.Charset,
		TreeShaking:       options.TreeShaking,
		IgnoreAnnotations: options.IgnoreAnnotations,
		LegalComments:     options.LegalComments,
		JSX:               options.JSX,
		JSXFactory:        options.JSXFactory,
		JSXFragment:       options.JSXFragment,
		JSXImportSource:   options.JSXImportSource,
		JSXDev:            options.JSXDev,
		JSXSideEffects:    options.JSXSideEffects,
		TsconfigRaw:       options.TsconfigRaw,
		Banner:            options.Banner,
		Footer:            options.Footer,
		Define:            options.Define,
		Pure:              options.Pure,
		KeepNames:         options.KeepNames,
		Sourcefile:        options.Sourcefile,
		Loader:            loader,
	})

	code = string(result.Code)

	if len(result.Errors) != 0 {
		err = fmt.Errorf("error: %v", result.Errors[0].Text)
	}

	return code, err
}
