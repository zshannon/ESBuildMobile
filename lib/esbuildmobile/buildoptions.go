package esbuildmobile

import "github.com/evanw/esbuild/pkg/api"

// BuildOptions mirrors a subset of esbuild BuildOptions for gomobile bindings
// Only the most common options are included for brevity
// More options can be added as needed

type BuildOptions struct {
	Bundle            bool
	Write             bool
	Outfile           string
	Platform          api.Platform
	Format            api.Format
	Target            api.Target
	Sourcemap         api.SourceMap
	MinifyWhitespace  bool
	MinifyIdentifiers bool
	MinifySyntax      bool
	GlobalName        string
	Define            map[string]string
	LogLevel          api.LogLevel
	LogLimit          int
	Loader            api.Loader
	Sourcefile        string
}

// NewBuildOptions returns BuildOptions with defaults
func NewBuildOptions() *BuildOptions {
	return &BuildOptions{}
}

func (b *BuildOptions) ConfigureBundle(bundle bool)             { b.Bundle = bundle }
func (b *BuildOptions) ConfigureWrite(write bool)               { b.Write = write }
func (b *BuildOptions) ConfigureOutfile(outfile string)         { b.Outfile = outfile }
func (b *BuildOptions) ConfigurePlatform(platform api.Platform) { b.Platform = platform }
func (b *BuildOptions) ConfigureFormat(format api.Format)       { b.Format = format }
func (b *BuildOptions) ConfigureTarget(target api.Target)       { b.Target = target }
func (b *BuildOptions) ConfigureSourcemap(sm api.SourceMap)     { b.Sourcemap = sm }
func (b *BuildOptions) ConfigureMinifyWhitespace(v bool)        { b.MinifyWhitespace = v }
func (b *BuildOptions) ConfigureMinifyIdentifiers(v bool)       { b.MinifyIdentifiers = v }
func (b *BuildOptions) ConfigureMinifySyntax(v bool)            { b.MinifySyntax = v }
func (b *BuildOptions) ConfigureGlobalName(n string)            { b.GlobalName = n }
func (b *BuildOptions) ConfigureLogLevel(l api.LogLevel)        { b.LogLevel = l }
func (b *BuildOptions) ConfigureLogLimit(l int)                 { b.LogLimit = l }
func (b *BuildOptions) ConfigureLoader(loader api.Loader)       { b.Loader = loader }
func (b *BuildOptions) ConfigureSourcefile(sf string)           { b.Sourcefile = sf }

func (b *BuildOptions) ConfigureDefineEntry(key, value string) {
	if b.Define == nil {
		b.Define = make(map[string]string)
	}
	b.Define[key] = value
}

// Helper config methods using string values for gomobile compatibility
func (b *BuildOptions) ConfigurePlatformByString(platform string) {
	switch platform {
	case "browser":
		b.Platform = api.PlatformBrowser
	case "node":
		b.Platform = api.PlatformNode
	case "neutral":
		b.Platform = api.PlatformNeutral
	}
}

func (b *BuildOptions) ConfigureFormatByString(format string) {
	switch format {
	case "iife":
		b.Format = api.FormatIIFE
	case "cjs":
		b.Format = api.FormatCommonJS
	case "esm":
		b.Format = api.FormatESModule
	default:
		b.Format = api.FormatDefault
	}
}

func (b *BuildOptions) ConfigureTargetByString(target string) {
	switch target {
	case "esnext":
		b.Target = api.ESNext
	case "es2020":
		b.Target = api.ES2020
	case "es2019":
		b.Target = api.ES2019
	case "es2018":
		b.Target = api.ES2018
	case "es2017":
		b.Target = api.ES2017
	case "es2016":
		b.Target = api.ES2016
	case "es2015":
		b.Target = api.ES2015
	case "es5":
		b.Target = api.ES5
	default:
		b.Target = api.ESNext
	}
}

func (b *BuildOptions) ConfigureSourcemapByString(sm string) {
	switch sm {
	case "none":
		b.Sourcemap = api.SourceMapNone
	case "inline":
		b.Sourcemap = api.SourceMapInline
	case "external":
		b.Sourcemap = api.SourceMapExternal
	case "linked":
		b.Sourcemap = api.SourceMapLinked
	case "both":
		b.Sourcemap = api.SourceMapInlineAndExternal
	default:
		b.Sourcemap = api.SourceMapNone
	}
}

func (b *BuildOptions) ConfigureLogLevelByString(level string) {
	switch level {
	case "verbose":
		b.LogLevel = api.LogLevelVerbose
	case "debug":
		b.LogLevel = api.LogLevelDebug
	case "info":
		b.LogLevel = api.LogLevelInfo
	case "warning":
		b.LogLevel = api.LogLevelWarning
	case "error":
		b.LogLevel = api.LogLevelError
	case "silent":
		b.LogLevel = api.LogLevelSilent
	default:
		b.LogLevel = api.LogLevelInfo
	}
}

func (b *BuildOptions) ConfigureLoaderByString(loader string) {
	switch loader {
	case "js":
		b.Loader = api.LoaderJS
	case "jsx":
		b.Loader = api.LoaderJSX
	case "ts":
		b.Loader = api.LoaderTS
	case "tsx":
		b.Loader = api.LoaderTSX
	case "css":
		b.Loader = api.LoaderCSS
	case "json":
		b.Loader = api.LoaderJSON
	case "text":
		b.Loader = api.LoaderText
	case "base64":
		b.Loader = api.LoaderBase64
	case "dataurl":
		b.Loader = api.LoaderDataURL
	case "file":
		b.Loader = api.LoaderFile
	case "binary":
		b.Loader = api.LoaderBinary
	}
}
