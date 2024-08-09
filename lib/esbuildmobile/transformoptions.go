package esbuildmobile

type TransformOptions struct {
	// Change `React.createElement` to custom name
	jSXFactory string
	// Change `React.Fragment` to custom name
	jSXFragment string
}

// Change `React.createElement` to custom name
func (t *TransformOptions) SetJSXFactory(newJSXFactory string) *TransformOptions {
	t.jSXFactory = newJSXFactory
	return t
}

// Change `React.createElement` to custom name
func (t *TransformOptions) SetJSXFragment(newJSXFragment string) *TransformOptions {
	t.jSXFactory = newJSXFragment
	return t
}
