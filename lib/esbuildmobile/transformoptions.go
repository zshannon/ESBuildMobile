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

// NewTransformOptions creates a new TransformOptions with default values
func NewTransformOptions() *TransformOptions {
	return &TransformOptions{}
}

// WithJSXFactory sets the JSX factory function name (e.g., "React.createElement")
func (t *TransformOptions) WithJSXFactory(factoryName string) {
	t.JSXFactory = factoryName
}

// WithJSXFragment sets the JSX fragment function name (e.g., "React.Fragment")
func (t *TransformOptions) WithJSXFragment(fragmentName string) {
	t.JSXFragment = fragmentName
}

// Change `React.createElement` to custom name
func (t *TransformOptions) ConfigureJSXFactory(factoryName string) {
	t.JSXFactory = factoryName
}

// Change `React.Fragment` to custom name
func (t *TransformOptions) ConfigureJSXFragment(fragmentName string) {
	t.JSXFragment = fragmentName
}

// Configure JSX import source
func (t *TransformOptions) ConfigureJSXImportSource(importSource string) {
	t.JSXImportSource = importSource
}

// Configure JSX development mode
func (t *TransformOptions) ConfigureJSXDev(enabled bool) {
	t.JSXDev = enabled
}

// Configure JSX side effects
func (t *TransformOptions) ConfigureJSXSideEffects(enabled bool) {
	t.JSXSideEffects = enabled
}

// Configure whitespace minification
func (t *TransformOptions) ConfigureMinifyWhitespace(enabled bool) {
	t.MinifyWhitespace = enabled
}

// Configure identifier minification
func (t *TransformOptions) ConfigureMinifyIdentifiers(enabled bool) {
	t.MinifyIdentifiers = enabled
}

// Configure syntax minification
func (t *TransformOptions) ConfigureMinifySyntax(enabled bool) {
	t.MinifySyntax = enabled
}

// Configure global name for UMD bundles
func (t *TransformOptions) ConfigureGlobalName(name string) {
	t.GlobalName = name
}

// Configure source root directory
func (t *TransformOptions) ConfigureSourceRoot(root string) {
	t.SourceRoot = root
}

// Configure source file name
func (t *TransformOptions) ConfigureSourcefile(filename string) {
	t.Sourcefile = filename
}

// Configure keep names option
func (t *TransformOptions) ConfigureKeepNames(enabled bool) {
	t.KeepNames = enabled
}

// Configure log limit
func (t *TransformOptions) ConfigureLogLimit(limit int) {
	t.LogLimit = limit
}

// Configure raw tsconfig content
func (t *TransformOptions) ConfigureTsconfigRaw(config string) {
	t.TsconfigRaw = config
}

// Configure banner text
func (t *TransformOptions) ConfigureBanner(banner string) {
	t.Banner = banner
}

// Configure footer text
func (t *TransformOptions) ConfigureFooter(footer string) {
	t.Footer = footer
}

// Configure line limit
func (t *TransformOptions) ConfigureLineLimit(limit int) {
	t.LineLimit = limit
}

// Configure ignore annotations
func (t *TransformOptions) ConfigureIgnoreAnnotations(enabled bool) {
	t.IgnoreAnnotations = enabled
}

// Configure mangle properties pattern
func (t *TransformOptions) ConfigureMangleProps(pattern string) {
	t.MangleProps = pattern
}

// Configure reserve properties pattern
func (t *TransformOptions) ConfigureReserveProps(pattern string) {
	t.ReserveProps = pattern
}

// Configure color output
func (t *TransformOptions) ConfigureColor(color api.StderrColor) {
	t.Color = color
}

// Configure log level
func (t *TransformOptions) ConfigureLogLevel(level api.LogLevel) {
	t.LogLevel = level
}

// Configure log override map
func (t *TransformOptions) ConfigureLogOverride(overrides map[string]api.LogLevel) {
	t.LogOverride = overrides
}

// Configure sourcemap generation
func (t *TransformOptions) ConfigureSourcemap(sourcemap api.SourceMap) {
	t.Sourcemap = sourcemap
}

// Configure sources content
func (t *TransformOptions) ConfigureSourcesContent(content api.SourcesContent) {
	t.SourcesContent = content
}

// Configure target environment
func (t *TransformOptions) ConfigureTarget(target api.Target) {
	t.Target = target
}

// Configure engines
func (t *TransformOptions) ConfigureEngines(engines []api.Engine) {
	t.Engines = engines
}

// Configure supported features
func (t *TransformOptions) ConfigureSupported(supported map[string]bool) {
	t.Supported = supported
}

// Configure platform
func (t *TransformOptions) ConfigurePlatform(platform api.Platform) {
	t.Platform = platform
}

// Configure output format
func (t *TransformOptions) ConfigureFormat(format api.Format) {
	t.Format = format
}

// Configure mangle quoted properties
func (t *TransformOptions) ConfigureMangleQuoted(quoted api.MangleQuoted) {
	t.MangleQuoted = quoted
}

// Configure mangle cache
func (t *TransformOptions) ConfigureMangleCache(cache map[string]interface{}) {
	t.MangleCache = cache
}

// Configure drop options
func (t *TransformOptions) ConfigureDrop(drop api.Drop) {
	t.Drop = drop
}

// Configure drop labels
func (t *TransformOptions) ConfigureDropLabels(labels []string) {
	t.DropLabels = labels
}

// Configure charset
func (t *TransformOptions) ConfigureCharset(charset api.Charset) {
	t.Charset = charset
}

// Configure tree shaking
func (t *TransformOptions) ConfigureTreeShaking(shaking api.TreeShaking) {
	t.TreeShaking = shaking
}

// Configure legal comments
func (t *TransformOptions) ConfigureLegalComments(comments api.LegalComments) {
	t.LegalComments = comments
}

// Configure JSX mode
func (t *TransformOptions) ConfigureJSX(jsx api.JSX) {
	t.JSX = jsx
}

// Configure define map
func (t *TransformOptions) ConfigureDefine(define map[string]string) {
	t.Define = define
}

// Configure pure functions
func (t *TransformOptions) ConfigurePure(pure []string) {
	t.Pure = pure
}

// Configure loader
func (t *TransformOptions) ConfigureLoader(loader api.Loader) {
	t.Loader = loader
}

// Configure loader by string
func (t *TransformOptions) ConfigureLoaderByString(loader string) {
	switch loader {
	case "js":
		t.Loader = api.LoaderJS
	case "jsx":
		t.Loader = api.LoaderJSX
	case "ts":
		t.Loader = api.LoaderTS
	case "tsx":
		t.Loader = api.LoaderTSX
	case "css":
		t.Loader = api.LoaderCSS
	case "json":
		t.Loader = api.LoaderJSON
	case "text":
		t.Loader = api.LoaderText
	case "base64":
		t.Loader = api.LoaderBase64
	case "dataurl":
		t.Loader = api.LoaderDataURL
	case "file":
		t.Loader = api.LoaderFile
	case "binary":
		t.Loader = api.LoaderBinary
	default:
		t.Loader = api.LoaderJS
	}
}

// Configure platform by string
func (t *TransformOptions) ConfigurePlatformByString(platform string) {
	switch platform {
	case "browser":
		t.Platform = api.PlatformBrowser
	case "node":
		t.Platform = api.PlatformNode
	case "neutral":
		t.Platform = api.PlatformNeutral
	default:
		t.Platform = api.PlatformBrowser
	}
}

// Configure format by string
func (t *TransformOptions) ConfigureFormatByString(format string) {
	switch format {
	case "iife":
		t.Format = api.FormatIIFE
	case "cjs":
		t.Format = api.FormatCommonJS
	case "esm":
		t.Format = api.FormatESModule
	default:
		t.Format = api.FormatDefault
	}
}

// Configure target by string
func (t *TransformOptions) ConfigureTargetByString(target string) {
	switch target {
	case "esnext":
		t.Target = api.ESNext
	case "es2020":
		t.Target = api.ES2020
	case "es2019":
		t.Target = api.ES2019
	case "es2018":
		t.Target = api.ES2018
	case "es2017":
		t.Target = api.ES2017
	case "es2016":
		t.Target = api.ES2016
	case "es2015":
		t.Target = api.ES2015
	case "es6":
		t.Target = api.ES2015
	case "es5":
		t.Target = api.ES5
	default:
		t.Target = api.ESNext
	}
}

// Configure JSX mode by string
func (t *TransformOptions) ConfigureJSXByString(jsx string) {
	switch jsx {
	case "transform":
		t.JSX = api.JSXTransform
	case "preserve":
		t.JSX = api.JSXPreserve
	case "automatic":
		t.JSX = api.JSXAutomatic
	default:
		t.JSX = api.JSXTransform
	}
}

// Configure sourcemap by string
func (t *TransformOptions) ConfigureSourcemapByString(sourcemap string) {
	switch sourcemap {
	case "none":
		t.Sourcemap = api.SourceMapNone
	case "inline":
		t.Sourcemap = api.SourceMapInline
	case "linked":
		t.Sourcemap = api.SourceMapLinked
	case "external":
		t.Sourcemap = api.SourceMapExternal
	case "both":
		t.Sourcemap = api.SourceMapInlineAndExternal
	default:
		t.Sourcemap = api.SourceMapNone
	}
}

// Configure sources content by string
func (t *TransformOptions) ConfigureSourcesContentByString(content string) {
	switch content {
	case "include":
		t.SourcesContent = api.SourcesContentInclude
	case "exclude":
		t.SourcesContent = api.SourcesContentExclude
	default:
		t.SourcesContent = api.SourcesContentInclude
	}
}

// Configure log level by string
func (t *TransformOptions) ConfigureLogLevelByString(level string) {
	switch level {
	case "verbose":
		t.LogLevel = api.LogLevelVerbose
	case "debug":
		t.LogLevel = api.LogLevelDebug
	case "info":
		t.LogLevel = api.LogLevelInfo
	case "warning":
		t.LogLevel = api.LogLevelWarning
	case "error":
		t.LogLevel = api.LogLevelError
	case "silent":
		t.LogLevel = api.LogLevelSilent
	default:
		t.LogLevel = api.LogLevelInfo
	}
}

// Configure color by string
func (t *TransformOptions) ConfigureColorByString(color string) {
	switch color {
	case "auto":
		t.Color = api.ColorIfTerminal
	case "always":
		t.Color = api.ColorAlways
	case "never":
		t.Color = api.ColorNever
	default:
		t.Color = api.ColorIfTerminal
	}
}

// Configure charset by string
func (t *TransformOptions) ConfigureCharsetByString(charset string) {
	switch charset {
	case "ascii":
		t.Charset = api.CharsetASCII
	case "utf8":
		t.Charset = api.CharsetUTF8
	default:
		t.Charset = api.CharsetUTF8
	}
}

// Configure tree shaking by string
func (t *TransformOptions) ConfigureTreeShakingByString(shaking string) {
	switch shaking {
	case "true":
		t.TreeShaking = api.TreeShakingTrue
	case "false":
		t.TreeShaking = api.TreeShakingFalse
	default:
		t.TreeShaking = api.TreeShakingFalse
	}
}

// Configure legal comments by string
func (t *TransformOptions) ConfigureLegalCommentsByString(comments string) {
	switch comments {
	case "none":
		t.LegalComments = api.LegalCommentsNone
	case "inline":
		t.LegalComments = api.LegalCommentsInline
	case "eof":
		t.LegalComments = api.LegalCommentsEndOfFile
	case "linked":
		t.LegalComments = api.LegalCommentsLinked
	case "external":
		t.LegalComments = api.LegalCommentsExternal
	default:
		t.LegalComments = api.LegalCommentsDefault
	}
}

// Configure mangle quoted by string
func (t *TransformOptions) ConfigureMangleQuotedByString(quoted string) {
	switch quoted {
	case "true":
		t.MangleQuoted = api.MangleQuotedTrue
	case "false":
		t.MangleQuoted = api.MangleQuotedFalse
	default:
		t.MangleQuoted = api.MangleQuotedFalse
	}
}

// Configure drop by string
func (t *TransformOptions) ConfigureDropByString(drop string) {
	switch drop {
	case "console":
		t.Drop = api.DropConsole
	case "debugger":
		t.Drop = api.DropDebugger
	default:
		t.Drop = 0
	}
}

// Configure drop labels by slice
func (t *TransformOptions) ConfigureDropLabelsBySlice(labels []string) {
	t.DropLabels = labels
}

// Configure define by key-value pairs
func (t *TransformOptions) ConfigureDefineEntry(key, value string) {
	if t.Define == nil {
		t.Define = make(map[string]string)
	}
	t.Define[key] = value
}

// Configure pure functions by slice
func (t *TransformOptions) ConfigurePureBySlice(pure []string) {
	t.Pure = pure
}

// Configure supported feature
func (t *TransformOptions) ConfigureSupportedFeature(feature string, supported bool) {
	if t.Supported == nil {
		t.Supported = make(map[string]bool)
	}
	t.Supported[feature] = supported
}
