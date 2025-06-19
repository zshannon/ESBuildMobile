package esbuildmobile

import "github.com/evanw/esbuild/pkg/api"

type TransformOptions struct {
	Color       api.StderrColor
	LogLevel    api.LogLevel
	LogLimit    int
	LogOverride map[string]api.LogLevel

	Sourcemap      api.SourceMap
	SourceRoot     string
	SourcesContent api.SourcesContent

	Target    api.Target
	Engines   []api.Engine
	Supported map[string]bool

	Platform   api.Platform
	Format     api.Format
	GlobalName string

	MangleProps       string
	ReserveProps      string
	MangleQuoted      api.MangleQuoted
	MangleCache       map[string]interface{}
	Drop              api.Drop
	DropLabels        []string
	MinifyWhitespace  bool
	MinifyIdentifiers bool
	MinifySyntax      bool
	LineLimit         int
	Charset           api.Charset
	TreeShaking       api.TreeShaking
	IgnoreAnnotations bool
	LegalComments     api.LegalComments

	JSX             api.JSX
	JSXFactory      string
	JSXFragment     string
	JSXImportSource string
	JSXDev          bool
	JSXSideEffects  bool

	TsconfigRaw string
	Banner      string
	Footer      string

	Define    map[string]string
	Pure      []string
	KeepNames bool

	Sourcefile string
	Loader     api.Loader
}

// Change `React.createElement` to custom name
func (t *TransformOptions) SetJSXFactory(newJSXFactory string) *TransformOptions {
	t.JSXFactory = newJSXFactory
	return t
}

// Change `React.Fragment` to custom name
func (t *TransformOptions) SetJSXFragment(newJSXFragment string) *TransformOptions {
	t.JSXFragment = newJSXFragment
	return t
}
