package esbuildmobile

import (
	"strings"
	"testing"

	"github.com/evanw/esbuild/pkg/api"
)

func TestConfigureMethods(t *testing.T) {
	opts := &TransformOptions{}
	opts.ConfigureJSXFactory("h")
	opts.ConfigureJSXFragment("Frag")

	if opts.JSXFactory != "h" {
		t.Fatalf("JSXFactory expected h, got %s", opts.JSXFactory)
	}
	if opts.JSXFragment != "Frag" {
		t.Fatalf("JSXFragment expected Frag, got %s", opts.JSXFragment)
	}
}

func TestTransformDefaultOptions(t *testing.T) {
	code, err := TransformJSX("(<div>hi</div>)", nil)
	if err != nil {
		t.Fatalf("transform failed: %v", err)
	}
	if !strings.Contains(code, "React.createElement(\"div\", null, \"hi\")") {
		t.Fatalf("unexpected output: %s", code)
	}
}

func TestTransformCustomFactory(t *testing.T) {
	opts := &TransformOptions{}
	opts.ConfigureJSXFactory("h")
	opts.ConfigureJSXFragment("Frag")
	code, err := TransformJSX("(<span>hi</span>)", opts)
	if err != nil {
		t.Fatalf("transform failed: %v", err)
	}
	if !strings.Contains(code, "h(\"span\", null, \"hi\")") {
		t.Fatalf("unexpected output: %s", code)
	}
}

func TestBannerAndFooter(t *testing.T) {
	opts := &TransformOptions{
		Banner: "// banner",
		Footer: "// footer",
	}
	code, err := TransformJSX("(<div>hi</div>)", opts)
	if err != nil {
		t.Fatalf("transform failed: %v", err)
	}
	if !strings.HasPrefix(code, "// banner") {
		t.Fatalf("banner missing in output: %s", code)
	}
	trimmed := strings.TrimSpace(code)
	if !strings.HasSuffix(trimmed, "// footer") {
		t.Fatalf("footer missing in output: %s", code)
	}
}

func TestMinifyWhitespace(t *testing.T) {
	opts := &TransformOptions{MinifyWhitespace: true}
	code, err := TransformJSX("(<div>hi</div>)", opts)
	if err != nil {
		t.Fatalf("transform failed: %v", err)
	}
	if !strings.Contains(code, "React.createElement(\"div\",null,\"hi\")") {
		t.Fatalf("minified output unexpected: %s", code)
	}
}

func TestFormatIIFE(t *testing.T) {
	opts := &TransformOptions{Format: api.FormatIIFE, GlobalName: "MyLib"}
	code, err := TransformJSX("export const x = <div/>;", opts)
	if err != nil {
		t.Fatalf("transform failed: %v", err)
	}
	if !strings.Contains(code, "var MyLib = (() =>") {
		t.Fatalf("IIFE wrapper missing: %s", code)
	}
}
