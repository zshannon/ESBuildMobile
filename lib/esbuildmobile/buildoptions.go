package esbuildmobile

import "github.com/evanw/esbuild/pkg/api"

type BuildOptions struct {
	// Logging options
	Color       api.StderrColor         // Documentation: https://esbuild.github.io/api/#color
	LogLevel    api.LogLevel            // Documentation: https://esbuild.github.io/api/#log-level
	LogLimit    int                     // Documentation: https://esbuild.github.io/api/#log-limit
	LogOverride map[string]api.LogLevel // Documentation: https://esbuild.github.io/api/#log-override

	// Source map options
	Sourcemap      api.SourceMap      // Documentation: https://esbuild.github.io/api/#sourcemap
	SourceRoot     string             // Documentation: https://esbuild.github.io/api/#source-root
	SourcesContent api.SourcesContent // Documentation: https://esbuild.github.io/api/#sources-content

	// Target environment options
	Target    api.Target      // Documentation: https://esbuild.github.io/api/#target
	Engines   []api.Engine    // Documentation: https://esbuild.github.io/api/#target
	Supported map[string]bool // Documentation: https://esbuild.github.io/api/#supported

	// Minification and mangling options
	MangleProps       string                 // Documentation: https://esbuild.github.io/api/#mangle-props
	ReserveProps      string                 // Documentation: https://esbuild.github.io/api/#mangle-props
	MangleQuoted      api.MangleQuoted       // Documentation: https://esbuild.github.io/api/#mangle-props
	MangleCache       map[string]interface{} // Documentation: https://esbuild.github.io/api/#mangle-props
	Drop              api.Drop               // Documentation: https://esbuild.github.io/api/#drop
	DropLabels        []string               // Documentation: https://esbuild.github.io/api/#drop-labels
	MinifyWhitespace  bool                   // Documentation: https://esbuild.github.io/api/#minify
	MinifyIdentifiers bool                   // Documentation: https://esbuild.github.io/api/#minify
	MinifySyntax      bool                   // Documentation: https://esbuild.github.io/api/#minify
	LineLimit         int                    // Documentation: https://esbuild.github.io/api/#line-limit
	Charset           api.Charset            // Documentation: https://esbuild.github.io/api/#charset
	TreeShaking       api.TreeShaking        // Documentation: https://esbuild.github.io/api/#tree-shaking
	IgnoreAnnotations bool                   // Documentation: https://esbuild.github.io/api/#ignore-annotations
	LegalComments     api.LegalComments      // Documentation: https://esbuild.github.io/api/#legal-comments

	// JSX options
	JSX             api.JSX // Documentation: https://esbuild.github.io/api/#jsx-mode
	JSXFactory      string  // Documentation: https://esbuild.github.io/api/#jsx-factory
	JSXFragment     string  // Documentation: https://esbuild.github.io/api/#jsx-fragment
	JSXImportSource string  // Documentation: https://esbuild.github.io/api/#jsx-import-source
	JSXDev          bool    // Documentation: https://esbuild.github.io/api/#jsx-dev
	JSXSideEffects  bool    // Documentation: https://esbuild.github.io/api/#jsx-side-effects

	// Transform options
	Define    map[string]string // Documentation: https://esbuild.github.io/api/#define
	Pure      []string          // Documentation: https://esbuild.github.io/api/#pure
	KeepNames bool              // Documentation: https://esbuild.github.io/api/#keep-names

	// Bundle options
	GlobalName        string                // Documentation: https://esbuild.github.io/api/#global-name
	Bundle            bool                  // Documentation: https://esbuild.github.io/api/#bundle
	PreserveSymlinks  bool                  // Documentation: https://esbuild.github.io/api/#preserve-symlinks
	Splitting         bool                  // Documentation: https://esbuild.github.io/api/#splitting
	Outfile           string                // Documentation: https://esbuild.github.io/api/#outfile
	Metafile          bool                  // Documentation: https://esbuild.github.io/api/#metafile
	Outdir            string                // Documentation: https://esbuild.github.io/api/#outdir
	Outbase           string                // Documentation: https://esbuild.github.io/api/#outbase
	AbsWorkingDir     string                // Documentation: https://esbuild.github.io/api/#working-directory
	Platform          api.Platform          // Documentation: https://esbuild.github.io/api/#platform
	Format            api.Format            // Documentation: https://esbuild.github.io/api/#format
	External          []string              // Documentation: https://esbuild.github.io/api/#external
	Packages          api.Packages          // Documentation: https://esbuild.github.io/api/#packages
	Alias             map[string]string     // Documentation: https://esbuild.github.io/api/#alias
	MainFields        []string              // Documentation: https://esbuild.github.io/api/#main-fields
	Conditions        []string              // Documentation: https://esbuild.github.io/api/#conditions
	Loader            map[string]api.Loader // Documentation: https://esbuild.github.io/api/#loader
	ResolveExtensions []string              // Documentation: https://esbuild.github.io/api/#resolve-extensions
	Tsconfig          string                // Documentation: https://esbuild.github.io/api/#tsconfig
	TsconfigRaw       string                // Documentation: https://esbuild.github.io/api/#tsconfig-raw
	OutExtension      map[string]string     // Documentation: https://esbuild.github.io/api/#out-extension
	PublicPath        string                // Documentation: https://esbuild.github.io/api/#public-path
	Inject            []string              // Documentation: https://esbuild.github.io/api/#inject
	Banner            map[string]string     // Documentation: https://esbuild.github.io/api/#banner
	Footer            map[string]string     // Documentation: https://esbuild.github.io/api/#footer
	NodePaths         []string              // Documentation: https://esbuild.github.io/api/#node-paths

	// Output naming options
	EntryNames string // Documentation: https://esbuild.github.io/api/#entry-names
	ChunkNames string // Documentation: https://esbuild.github.io/api/#chunk-names
	AssetNames string // Documentation: https://esbuild.github.io/api/#asset-names

	// Entry points
	EntryPoints         []string         // Documentation: https://esbuild.github.io/api/#entry-points
	EntryPointsAdvanced []api.EntryPoint // Documentation: https://esbuild.github.io/api/#entry-points

	// Input/output options
	Stdin          *api.StdinOptions // Documentation: https://esbuild.github.io/api/#stdin
	Write          bool              // Documentation: https://esbuild.github.io/api/#write
	AllowOverwrite bool              // Documentation: https://esbuild.github.io/api/#allow-overwrite
	Plugins        []*Plugin         // Documentation: https://esbuild.github.io/plugins/

	// Legacy properties for backwards compatibility
	Sourcefile   string     // For single file transforms
	LoaderSingle api.Loader // For single file loader
}

// NewBuildOptions returns BuildOptions with defaults
func NewBuildOptions() *BuildOptions {
	return &BuildOptions{
		Define:       make(map[string]string),
		LogOverride:  make(map[string]api.LogLevel),
		Alias:        make(map[string]string),
		Loader:       make(map[string]api.Loader),
		OutExtension: make(map[string]string),
		Banner:       make(map[string]string),
		Footer:       make(map[string]string),
		Supported:    make(map[string]bool),
		MangleCache:  make(map[string]interface{}),
		Plugins:      make([]*Plugin, 0),
	}
}

// Convert to esbuild API BuildOptions
func (b *BuildOptions) ToAPIBuildOptions() api.BuildOptions {
	return api.BuildOptions{
		Color:       b.Color,
		LogLevel:    b.LogLevel,
		LogLimit:    b.LogLimit,
		LogOverride: b.LogOverride,

		Sourcemap:      b.Sourcemap,
		SourceRoot:     b.SourceRoot,
		SourcesContent: b.SourcesContent,

		Target:    b.Target,
		Engines:   b.Engines,
		Supported: b.Supported,

		MangleProps:       b.MangleProps,
		ReserveProps:      b.ReserveProps,
		MangleQuoted:      b.MangleQuoted,
		MangleCache:       b.MangleCache,
		Drop:              b.Drop,
		DropLabels:        b.DropLabels,
		MinifyWhitespace:  b.MinifyWhitespace,
		MinifyIdentifiers: b.MinifyIdentifiers,
		MinifySyntax:      b.MinifySyntax,
		LineLimit:         b.LineLimit,
		Charset:           b.Charset,
		TreeShaking:       b.TreeShaking,
		IgnoreAnnotations: b.IgnoreAnnotations,
		LegalComments:     b.LegalComments,

		JSX:             b.JSX,
		JSXFactory:      b.JSXFactory,
		JSXFragment:     b.JSXFragment,
		JSXImportSource: b.JSXImportSource,
		JSXDev:          b.JSXDev,
		JSXSideEffects:  b.JSXSideEffects,

		Define:    b.Define,
		Pure:      b.Pure,
		KeepNames: b.KeepNames,

		GlobalName:        b.GlobalName,
		Bundle:            b.Bundle,
		PreserveSymlinks:  b.PreserveSymlinks,
		Splitting:         b.Splitting,
		Outfile:           b.Outfile,
		Metafile:          b.Metafile,
		Outdir:            b.Outdir,
		Outbase:           b.Outbase,
		AbsWorkingDir:     b.AbsWorkingDir,
		Platform:          b.Platform,
		Format:            b.Format,
		External:          b.External,
		Packages:          b.Packages,
		Alias:             b.Alias,
		MainFields:        b.MainFields,
		Conditions:        b.Conditions,
		Loader:            b.Loader,
		ResolveExtensions: b.ResolveExtensions,
		Tsconfig:          b.Tsconfig,
		TsconfigRaw:       b.TsconfigRaw,
		OutExtension:      b.OutExtension,
		PublicPath:        b.PublicPath,
		Inject:            b.Inject,
		Banner:            b.Banner,
		Footer:            b.Footer,
		NodePaths:         b.NodePaths,

		EntryNames: b.EntryNames,
		ChunkNames: b.ChunkNames,
		AssetNames: b.AssetNames,

		EntryPoints:         b.EntryPoints,
		EntryPointsAdvanced: b.EntryPointsAdvanced,

		Stdin:          b.Stdin,
		Write:          b.Write,
		AllowOverwrite: b.AllowOverwrite,
		Plugins:        b.convertPlugins(),
	}
}

// convertPlugins converts our mobile-friendly plugins to esbuild API plugins
func (b *BuildOptions) convertPlugins() []api.Plugin {
	apiPlugins := make([]api.Plugin, len(b.Plugins))
	for i, plugin := range b.Plugins {
		apiPlugins[i] = b.convertPlugin(plugin)
	}
	return apiPlugins
}

// convertPlugin converts a single plugin to esbuild API plugin
func (b *BuildOptions) convertPlugin(plugin *Plugin) api.Plugin {
	return api.Plugin{
		Name: plugin.GetName(),
		Setup: func(build api.PluginBuild) {
			// Register onResolve callbacks
			for _, rule := range plugin.onResolveRules {
				callback := rule.Callback
				build.OnResolve(rule.Options.ToAPI(), func(args api.OnResolveArgs) (api.OnResolveResult, error) {
					mobileArgs := &OnResolveArgs{
						Path:       args.Path,
						Importer:   args.Importer,
						Namespace:  args.Namespace,
						ResolveDir: args.ResolveDir,
						Kind:       ResolveKindFromAPI(args.Kind),
					}
					result := callback.Call(mobileArgs)
					return api.OnResolveResult{
						PluginName:  result.PluginName,
						Errors:      result.Errors,
						Warnings:    result.Warnings,
						Path:        result.Path,
						External:    result.External,
						SideEffects: result.SideEffects.ToAPI(),
						Namespace:   result.Namespace,
						Suffix:      result.Suffix,
						WatchFiles:  result.WatchFiles,
						WatchDirs:   result.WatchDirs,
					}, nil
				})
			}

			// Register onLoad callbacks
			for _, rule := range plugin.onLoadRules {
				callback := rule.Callback
				build.OnLoad(rule.Options.ToAPI(), func(args api.OnLoadArgs) (api.OnLoadResult, error) {
					mobileArgs := &OnLoadArgs{
						Path:      args.Path,
						Namespace: args.Namespace,
						Suffix:    args.Suffix,
					}
					result := callback.Call(mobileArgs)
					return api.OnLoadResult{
						PluginName: result.PluginName,
						Errors:     result.Errors,
						Warnings:   result.Warnings,
						Contents:   result.Contents,
						ResolveDir: result.ResolveDir,
						Loader:     result.Loader,
						WatchFiles: result.WatchFiles,
						WatchDirs:  result.WatchDirs,
					}, nil
				})
			}

			// Register onStart callback
			if plugin.onStartCallback != nil {
				callback := plugin.onStartCallback
				build.OnStart(func() (api.OnStartResult, error) {
					result := callback.Call()
					return api.OnStartResult{
						Errors:   result.Errors,
						Warnings: result.Warnings,
					}, nil
				})
			}

			// Register onEnd callback
			if plugin.onEndCallback != nil {
				callback := plugin.onEndCallback
				build.OnEnd(func(result *api.BuildResult) (api.OnEndResult, error) {
					mobileResult := &BuildResult{
						Errors:   result.Errors,
						Warnings: result.Warnings,
					}
					endResult := callback.Call(mobileResult)
					return api.OnEndResult{
						Errors:   endResult.Errors,
						Warnings: endResult.Warnings,
					}, nil
				})
			}
		},
	}
}

// Basic configuration methods
func (b *BuildOptions) ConfigureBundle(bundle bool)             { b.Bundle = bundle }
func (b *BuildOptions) ConfigureWrite(write bool)               { b.Write = write }
func (b *BuildOptions) ConfigureOutfile(outfile string)         { b.Outfile = outfile }
func (b *BuildOptions) ConfigureOutdir(outdir string)           { b.Outdir = outdir }
func (b *BuildOptions) ConfigureOutbase(outbase string)         { b.Outbase = outbase }
func (b *BuildOptions) ConfigurePlatform(platform api.Platform) { b.Platform = platform }
func (b *BuildOptions) ConfigureFormat(format api.Format)       { b.Format = format }
func (b *BuildOptions) ConfigureTarget(target api.Target)       { b.Target = target }
func (b *BuildOptions) ConfigureSourcemap(sm api.SourceMap)     { b.Sourcemap = sm }
func (b *BuildOptions) ConfigureSourceRoot(root string)         { b.SourceRoot = root }
func (b *BuildOptions) ConfigureGlobalName(n string)            { b.GlobalName = n }
func (b *BuildOptions) ConfigureLogLevel(l api.LogLevel)        { b.LogLevel = l }
func (b *BuildOptions) ConfigureLogLimit(l int)                 { b.LogLimit = l }
func (b *BuildOptions) ConfigureSourcefile(sf string)           { b.Sourcefile = sf }
func (b *BuildOptions) ConfigurePublicPath(path string)         { b.PublicPath = path }
func (b *BuildOptions) ConfigureAbsWorkingDir(dir string)       { b.AbsWorkingDir = dir }

// Minification options
func (b *BuildOptions) ConfigureMinifyWhitespace(v bool)        { b.MinifyWhitespace = v }
func (b *BuildOptions) ConfigureMinifyIdentifiers(v bool)       { b.MinifyIdentifiers = v }
func (b *BuildOptions) ConfigureMinifySyntax(v bool)            { b.MinifySyntax = v }
func (b *BuildOptions) ConfigureTreeShaking(ts api.TreeShaking) { b.TreeShaking = ts }
func (b *BuildOptions) ConfigureIgnoreAnnotations(v bool)       { b.IgnoreAnnotations = v }
func (b *BuildOptions) ConfigureKeepNames(v bool)               { b.KeepNames = v }
func (b *BuildOptions) ConfigureLineLimit(limit int)            { b.LineLimit = limit }

// Advanced bundling options
func (b *BuildOptions) ConfigurePreserveSymlinks(v bool) { b.PreserveSymlinks = v }
func (b *BuildOptions) ConfigureSplitting(v bool)        { b.Splitting = v }
func (b *BuildOptions) ConfigureMetafile(v bool)         { b.Metafile = v }
func (b *BuildOptions) ConfigureAllowOverwrite(v bool)   { b.AllowOverwrite = v }
func (b *BuildOptions) ConfigurePackages(p api.Packages) { b.Packages = p }

// JSX configuration
func (b *BuildOptions) ConfigureJSX(jsx api.JSX)               { b.JSX = jsx }
func (b *BuildOptions) ConfigureJSXFactory(factory string)     { b.JSXFactory = factory }
func (b *BuildOptions) ConfigureJSXFragment(fragment string)   { b.JSXFragment = fragment }
func (b *BuildOptions) ConfigureJSXImportSource(source string) { b.JSXImportSource = source }
func (b *BuildOptions) ConfigureJSXDev(dev bool)               { b.JSXDev = dev }
func (b *BuildOptions) ConfigureJSXSideEffects(effects bool)   { b.JSXSideEffects = effects }

// Mangle configuration
func (b *BuildOptions) ConfigureMangleProps(props string)         { b.MangleProps = props }
func (b *BuildOptions) ConfigureReserveProps(props string)        { b.ReserveProps = props }
func (b *BuildOptions) ConfigureMangleQuoted(mq api.MangleQuoted) { b.MangleQuoted = mq }

// Drop configuration
func (b *BuildOptions) ConfigureDrop(drop api.Drop)                   { b.Drop = drop }
func (b *BuildOptions) ConfigureCharset(charset api.Charset)          { b.Charset = charset }
func (b *BuildOptions) ConfigureLegalComments(lc api.LegalComments)   { b.LegalComments = lc }
func (b *BuildOptions) ConfigureSourcesContent(sc api.SourcesContent) { b.SourcesContent = sc }
func (b *BuildOptions) ConfigureColor(color api.StderrColor)          { b.Color = color }

// TypeScript configuration
func (b *BuildOptions) ConfigureTsconfig(config string) { b.Tsconfig = config }
func (b *BuildOptions) ConfigureTsconfigRaw(raw string) { b.TsconfigRaw = raw }

// Output naming configuration
func (b *BuildOptions) ConfigureEntryNames(names string) { b.EntryNames = names }
func (b *BuildOptions) ConfigureChunkNames(names string) { b.ChunkNames = names }
func (b *BuildOptions) ConfigureAssetNames(names string) { b.AssetNames = names }

// Array/slice configuration methods
func (b *BuildOptions) AddEntryPoint(path string) {
	b.EntryPoints = append(b.EntryPoints, path)
}

func (b *BuildOptions) AddAdvancedEntryPoint(input, output string) {
	b.EntryPointsAdvanced = append(b.EntryPointsAdvanced, api.EntryPoint{
		InputPath:  input,
		OutputPath: output,
	})
}

func (b *BuildOptions) AddExternal(module string) {
	b.External = append(b.External, module)
}

func (b *BuildOptions) AddMainField(field string) {
	b.MainFields = append(b.MainFields, field)
}

func (b *BuildOptions) AddCondition(condition string) {
	b.Conditions = append(b.Conditions, condition)
}

func (b *BuildOptions) AddResolveExtension(ext string) {
	b.ResolveExtensions = append(b.ResolveExtensions, ext)
}

func (b *BuildOptions) AddInject(file string) {
	b.Inject = append(b.Inject, file)
}

func (b *BuildOptions) AddNodePath(path string) {
	b.NodePaths = append(b.NodePaths, path)
}

func (b *BuildOptions) AddPure(fn string) {
	b.Pure = append(b.Pure, fn)
}

func (b *BuildOptions) AddDropLabel(label string) {
	b.DropLabels = append(b.DropLabels, label)
}

func (b *BuildOptions) AddPlugin(plugin *Plugin) {
	b.Plugins = append(b.Plugins, plugin)
}

func (b *BuildOptions) AddEngine(name string, version string) {
	engineName := EngineNameFromString(name)
	customEngine := Engine{
		Name:    engineName,
		Version: version,
	}
	// Also add to the API engines list for compatibility
	b.Engines = append(b.Engines, customEngine.ToAPIEngine())
}

// Map configuration methods
func (b *BuildOptions) ConfigureDefineEntry(key, value string) {
	if b.Define == nil {
		b.Define = make(map[string]string)
	}
	b.Define[key] = value
}

func (b *BuildOptions) ConfigureLogOverrideEntry(key string, level api.LogLevel) {
	if b.LogOverride == nil {
		b.LogOverride = make(map[string]api.LogLevel)
	}
	b.LogOverride[key] = level
}

func (b *BuildOptions) ConfigureAliasEntry(from, to string) {
	if b.Alias == nil {
		b.Alias = make(map[string]string)
	}
	b.Alias[from] = to
}

func (b *BuildOptions) ConfigureLoaderEntry(ext string, loader api.Loader) {
	if b.Loader == nil {
		b.Loader = make(map[string]api.Loader)
	}
	b.Loader[ext] = loader
}

func (b *BuildOptions) ConfigureOutExtensionEntry(ext, newExt string) {
	if b.OutExtension == nil {
		b.OutExtension = make(map[string]string)
	}
	b.OutExtension[ext] = newExt
}

func (b *BuildOptions) ConfigureBannerEntry(fileType, banner string) {
	if b.Banner == nil {
		b.Banner = make(map[string]string)
	}
	b.Banner[fileType] = banner
}

func (b *BuildOptions) ConfigureFooterEntry(fileType, footer string) {
	if b.Footer == nil {
		b.Footer = make(map[string]string)
	}
	b.Footer[fileType] = footer
}

func (b *BuildOptions) ConfigureSupportedEntry(feature string, supported bool) {
	if b.Supported == nil {
		b.Supported = make(map[string]bool)
	}
	b.Supported[feature] = supported
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
	default:
		b.Platform = api.PlatformDefault
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
	case "es2024":
		b.Target = api.ES2024
	case "es2023":
		b.Target = api.ES2023
	case "es2022":
		b.Target = api.ES2022
	case "es2021":
		b.Target = api.ES2021
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
		b.LoaderSingle = api.LoaderJS
	case "jsx":
		b.LoaderSingle = api.LoaderJSX
	case "ts":
		b.LoaderSingle = api.LoaderTS
	case "tsx":
		b.LoaderSingle = api.LoaderTSX
	case "css":
		b.LoaderSingle = api.LoaderCSS
	case "json":
		b.LoaderSingle = api.LoaderJSON
	case "text":
		b.LoaderSingle = api.LoaderText
	case "base64":
		b.LoaderSingle = api.LoaderBase64
	case "dataurl":
		b.LoaderSingle = api.LoaderDataURL
	case "file":
		b.LoaderSingle = api.LoaderFile
	case "binary":
		b.LoaderSingle = api.LoaderBinary
	case "copy":
		b.LoaderSingle = api.LoaderCopy
	case "empty":
		b.LoaderSingle = api.LoaderEmpty
	case "globalcss":
		b.LoaderSingle = api.LoaderGlobalCSS
	case "localcss":
		b.LoaderSingle = api.LoaderLocalCSS
	default:
		b.LoaderSingle = api.LoaderDefault
	}
}

func (b *BuildOptions) ConfigurePackagesByString(packages string) {
	switch packages {
	case "bundle":
		b.Packages = api.PackagesBundle
	case "external":
		b.Packages = api.PackagesExternal
	default:
		b.Packages = api.PackagesDefault
	}
}

func (b *BuildOptions) ConfigureTreeShakingByString(treeShaking string) {
	switch treeShaking {
	case "true":
		b.TreeShaking = api.TreeShakingTrue
	case "false":
		b.TreeShaking = api.TreeShakingFalse
	default:
		b.TreeShaking = api.TreeShakingDefault
	}
}

func (b *BuildOptions) ConfigureJSXByString(jsx string) {
	switch jsx {
	case "transform":
		b.JSX = api.JSXTransform
	case "preserve":
		b.JSX = api.JSXPreserve
	case "automatic":
		b.JSX = api.JSXAutomatic
	default:
		b.JSX = api.JSXTransform
	}
}

func (b *BuildOptions) ConfigureLegalCommentsByString(comments string) {
	switch comments {
	case "none":
		b.LegalComments = api.LegalCommentsNone
	case "inline":
		b.LegalComments = api.LegalCommentsInline
	case "eof":
		b.LegalComments = api.LegalCommentsEndOfFile
	case "linked":
		b.LegalComments = api.LegalCommentsLinked
	case "external":
		b.LegalComments = api.LegalCommentsExternal
	default:
		b.LegalComments = api.LegalCommentsDefault
	}
}

func (b *BuildOptions) ConfigureCharsetByString(charset string) {
	switch charset {
	case "ascii":
		b.Charset = api.CharsetASCII
	case "utf8":
		b.Charset = api.CharsetUTF8
	default:
		b.Charset = api.CharsetDefault
	}
}

func (b *BuildOptions) ConfigureColorByString(color string) {
	switch color {
	case "never":
		b.Color = api.ColorNever
	case "always":
		b.Color = api.ColorAlways
	default:
		b.Color = api.ColorIfTerminal
	}
}

func (b *BuildOptions) ConfigureSourcesContentByString(content string) {
	switch content {
	case "include":
		b.SourcesContent = api.SourcesContentInclude
	case "exclude":
		b.SourcesContent = api.SourcesContentExclude
	default:
		b.SourcesContent = api.SourcesContentInclude
	}
}

func (b *BuildOptions) ConfigureMangleQuotedByString(quoted string) {
	switch quoted {
	case "true":
		b.MangleQuoted = api.MangleQuotedTrue
	case "false":
		b.MangleQuoted = api.MangleQuotedFalse
	default:
		b.MangleQuoted = api.MangleQuotedFalse
	}
}

// Drop configuration helpers
func (b *BuildOptions) ConfigureDropConsole(drop bool) {
	if drop {
		b.Drop |= api.DropConsole
	} else {
		b.Drop &^= api.DropConsole
	}
}

func (b *BuildOptions) ConfigureDropDebugger(drop bool) {
	if drop {
		b.Drop |= api.DropDebugger
	} else {
		b.Drop &^= api.DropDebugger
	}
}

// Stdin configuration
func (b *BuildOptions) ConfigureStdin(contents, resolveDir, sourcefile string, loader api.Loader) {
	b.Stdin = &api.StdinOptions{
		Contents:   contents,
		ResolveDir: resolveDir,
		Sourcefile: sourcefile,
		Loader:     loader,
	}
}

func (b *BuildOptions) ConfigureStdinByString(contents, resolveDir, sourcefile, loader string) {
	var loaderEnum api.Loader
	switch loader {
	case "js":
		loaderEnum = api.LoaderJS
	case "jsx":
		loaderEnum = api.LoaderJSX
	case "ts":
		loaderEnum = api.LoaderTS
	case "tsx":
		loaderEnum = api.LoaderTSX
	case "css":
		loaderEnum = api.LoaderCSS
	case "json":
		loaderEnum = api.LoaderJSON
	case "text":
		loaderEnum = api.LoaderText
	case "base64":
		loaderEnum = api.LoaderBase64
	case "dataurl":
		loaderEnum = api.LoaderDataURL
	case "file":
		loaderEnum = api.LoaderFile
	case "binary":
		loaderEnum = api.LoaderBinary
	case "copy":
		loaderEnum = api.LoaderCopy
	case "empty":
		loaderEnum = api.LoaderEmpty
	case "globalcss":
		loaderEnum = api.LoaderGlobalCSS
	case "localcss":
		loaderEnum = api.LoaderLocalCSS
	default:
		loaderEnum = api.LoaderDefault
	}

	b.Stdin = &api.StdinOptions{
		Contents:   contents,
		ResolveDir: resolveDir,
		Sourcefile: sourcefile,
		Loader:     loaderEnum,
	}
}

// Enum value getters for Swift interoperability

// Platform enum values
func GetPlatformDefault() int { return int(api.PlatformDefault) }
func GetPlatformBrowser() int { return int(api.PlatformBrowser) }
func GetPlatformNode() int    { return int(api.PlatformNode) }
func GetPlatformNeutral() int { return int(api.PlatformNeutral) }

// Format enum values
func GetFormatDefault() int  { return int(api.FormatDefault) }
func GetFormatIIFE() int     { return int(api.FormatIIFE) }
func GetFormatCommonJS() int { return int(api.FormatCommonJS) }
func GetFormatESModule() int { return int(api.FormatESModule) }

// Target enum values
func GetTargetDefault() int { return int(api.DefaultTarget) }
func GetTargetESNext() int  { return int(api.ESNext) }
func GetTargetES5() int     { return int(api.ES5) }
func GetTargetES2015() int  { return int(api.ES2015) }
func GetTargetES2016() int  { return int(api.ES2016) }
func GetTargetES2017() int  { return int(api.ES2017) }
func GetTargetES2018() int  { return int(api.ES2018) }
func GetTargetES2019() int  { return int(api.ES2019) }
func GetTargetES2020() int  { return int(api.ES2020) }
func GetTargetES2021() int  { return int(api.ES2021) }
func GetTargetES2022() int  { return int(api.ES2022) }
func GetTargetES2023() int  { return int(api.ES2023) }
func GetTargetES2024() int  { return int(api.ES2024) }

// SourceMap enum values
func GetSourceMapNone() int              { return int(api.SourceMapNone) }
func GetSourceMapInline() int            { return int(api.SourceMapInline) }
func GetSourceMapLinked() int            { return int(api.SourceMapLinked) }
func GetSourceMapExternal() int          { return int(api.SourceMapExternal) }
func GetSourceMapInlineAndExternal() int { return int(api.SourceMapInlineAndExternal) }

// LogLevel enum values
func GetLogLevelSilent() int  { return int(api.LogLevelSilent) }
func GetLogLevelVerbose() int { return int(api.LogLevelVerbose) }
func GetLogLevelDebug() int   { return int(api.LogLevelDebug) }
func GetLogLevelInfo() int    { return int(api.LogLevelInfo) }
func GetLogLevelWarning() int { return int(api.LogLevelWarning) }
func GetLogLevelError() int   { return int(api.LogLevelError) }

// Loader enum values
func GetLoaderNone() int      { return int(api.LoaderNone) }
func GetLoaderJS() int        { return int(api.LoaderJS) }
func GetLoaderJSX() int       { return int(api.LoaderJSX) }
func GetLoaderTS() int        { return int(api.LoaderTS) }
func GetLoaderTSX() int       { return int(api.LoaderTSX) }
func GetLoaderCSS() int       { return int(api.LoaderCSS) }
func GetLoaderJSON() int      { return int(api.LoaderJSON) }
func GetLoaderText() int      { return int(api.LoaderText) }
func GetLoaderBase64() int    { return int(api.LoaderBase64) }
func GetLoaderDataURL() int   { return int(api.LoaderDataURL) }
func GetLoaderFile() int      { return int(api.LoaderFile) }
func GetLoaderBinary() int    { return int(api.LoaderBinary) }
func GetLoaderCopy() int      { return int(api.LoaderCopy) }
func GetLoaderDefault() int   { return int(api.LoaderDefault) }
func GetLoaderEmpty() int     { return int(api.LoaderEmpty) }
func GetLoaderGlobalCSS() int { return int(api.LoaderGlobalCSS) }
func GetLoaderLocalCSS() int  { return int(api.LoaderLocalCSS) }

// StderrColor enum values
func GetColorIfTerminal() int { return int(api.ColorIfTerminal) }
func GetColorNever() int      { return int(api.ColorNever) }
func GetColorAlways() int     { return int(api.ColorAlways) }

// SourcesContent enum values
func GetSourcesContentInclude() int { return int(api.SourcesContentInclude) }
func GetSourcesContentExclude() int { return int(api.SourcesContentExclude) }

// MangleQuoted enum values
func GetMangleQuotedFalse() int { return int(api.MangleQuotedFalse) }
func GetMangleQuotedTrue() int  { return int(api.MangleQuotedTrue) }

// Drop enum values (these are flags, so they can be combined)
func GetDropConsole() int  { return int(api.DropConsole) }
func GetDropDebugger() int { return int(api.DropDebugger) }

// Charset enum values
func GetCharsetDefault() int { return int(api.CharsetDefault) }
func GetCharsetASCII() int   { return int(api.CharsetASCII) }
func GetCharsetUTF8() int    { return int(api.CharsetUTF8) }

// TreeShaking enum values
func GetTreeShakingDefault() int { return int(api.TreeShakingDefault) }
func GetTreeShakingFalse() int   { return int(api.TreeShakingFalse) }
func GetTreeShakingTrue() int    { return int(api.TreeShakingTrue) }

// LegalComments enum values
func GetLegalCommentsDefault() int   { return int(api.LegalCommentsDefault) }
func GetLegalCommentsNone() int      { return int(api.LegalCommentsNone) }
func GetLegalCommentsInline() int    { return int(api.LegalCommentsInline) }
func GetLegalCommentsEndOfFile() int { return int(api.LegalCommentsEndOfFile) }
func GetLegalCommentsLinked() int    { return int(api.LegalCommentsLinked) }
func GetLegalCommentsExternal() int  { return int(api.LegalCommentsExternal) }

// JSX enum values
func GetJSXTransform() int { return int(api.JSXTransform) }
func GetJSXPreserve() int  { return int(api.JSXPreserve) }
func GetJSXAutomatic() int { return int(api.JSXAutomatic) }

// Packages enum values
func GetPackagesDefault() int  { return int(api.PackagesDefault) }
func GetPackagesBundle() int   { return int(api.PackagesBundle) }
func GetPackagesExternal() int { return int(api.PackagesExternal) }

// Getter methods for current BuildOptions values
func (b *BuildOptions) GetPlatform() int       { return int(b.Platform) }
func (b *BuildOptions) GetFormat() int         { return int(b.Format) }
func (b *BuildOptions) GetTarget() int         { return int(b.Target) }
func (b *BuildOptions) GetSourcemap() int      { return int(b.Sourcemap) }
func (b *BuildOptions) GetLogLevel() int       { return int(b.LogLevel) }
func (b *BuildOptions) GetLoaderSingle() int   { return int(b.LoaderSingle) }
func (b *BuildOptions) GetColor() int          { return int(b.Color) }
func (b *BuildOptions) GetSourcesContent() int { return int(b.SourcesContent) }
func (b *BuildOptions) GetMangleQuoted() int   { return int(b.MangleQuoted) }
func (b *BuildOptions) GetDrop() int           { return int(b.Drop) }
func (b *BuildOptions) GetCharset() int        { return int(b.Charset) }
func (b *BuildOptions) GetTreeShaking() int    { return int(b.TreeShaking) }
func (b *BuildOptions) GetLegalComments() int  { return int(b.LegalComments) }
func (b *BuildOptions) GetJSX() int            { return int(b.JSX) }
func (b *BuildOptions) GetPackages() int       { return int(b.Packages) }

// Set methods using enum values as integers
func (b *BuildOptions) SetPlatformByInt(platform int)   { b.Platform = api.Platform(platform) }
func (b *BuildOptions) SetFormatByInt(format int)       { b.Format = api.Format(format) }
func (b *BuildOptions) SetTargetByInt(target int)       { b.Target = api.Target(target) }
func (b *BuildOptions) SetSourcemapByInt(sourcemap int) { b.Sourcemap = api.SourceMap(sourcemap) }
func (b *BuildOptions) SetLogLevelByInt(logLevel int)   { b.LogLevel = api.LogLevel(logLevel) }
func (b *BuildOptions) SetLoaderSingleByInt(loader int) { b.LoaderSingle = api.Loader(loader) }
func (b *BuildOptions) SetColorByInt(color int)         { b.Color = api.StderrColor(color) }
func (b *BuildOptions) SetSourcesContentByInt(content int) {
	b.SourcesContent = api.SourcesContent(content)
}
func (b *BuildOptions) SetMangleQuotedByInt(quoted int) { b.MangleQuoted = api.MangleQuoted(quoted) }
func (b *BuildOptions) SetDropByInt(drop int)           { b.Drop = api.Drop(drop) }
func (b *BuildOptions) SetCharsetByInt(charset int)     { b.Charset = api.Charset(charset) }
func (b *BuildOptions) SetTreeShakingByInt(shaking int) { b.TreeShaking = api.TreeShaking(shaking) }
func (b *BuildOptions) SetLegalCommentsByInt(comments int) {
	b.LegalComments = api.LegalComments(comments)
}
func (b *BuildOptions) SetJSXByInt(jsx int)           { b.JSX = api.JSX(jsx) }
func (b *BuildOptions) SetPackagesByInt(packages int) { b.Packages = api.Packages(packages) }

// Getter methods for slices and maps (return length/size info)
func (b *BuildOptions) GetEntryPointsCount() int         { return len(b.EntryPoints) }
func (b *BuildOptions) GetEntryPointsAdvancedCount() int { return len(b.EntryPointsAdvanced) }
func (b *BuildOptions) GetExternalCount() int            { return len(b.External) }
func (b *BuildOptions) GetMainFieldsCount() int          { return len(b.MainFields) }
func (b *BuildOptions) GetConditionsCount() int          { return len(b.Conditions) }
func (b *BuildOptions) GetResolveExtensionsCount() int   { return len(b.ResolveExtensions) }
func (b *BuildOptions) GetInjectCount() int              { return len(b.Inject) }
func (b *BuildOptions) GetNodePathsCount() int           { return len(b.NodePaths) }
func (b *BuildOptions) GetPureCount() int                { return len(b.Pure) }
func (b *BuildOptions) GetDropLabelsCount() int          { return len(b.DropLabels) }
func (b *BuildOptions) GetPluginsCount() int             { return len(b.Plugins) }

// Getter methods for string array elements
func (b *BuildOptions) GetEntryPoint(index int) string {
	if index >= 0 && index < len(b.EntryPoints) {
		return b.EntryPoints[index]
	}
	return ""
}

func (b *BuildOptions) GetExternal(index int) string {
	if index >= 0 && index < len(b.External) {
		return b.External[index]
	}
	return ""
}

func (b *BuildOptions) GetMainField(index int) string {
	if index >= 0 && index < len(b.MainFields) {
		return b.MainFields[index]
	}
	return ""
}

func (b *BuildOptions) GetCondition(index int) string {
	if index >= 0 && index < len(b.Conditions) {
		return b.Conditions[index]
	}
	return ""
}

func (b *BuildOptions) GetResolveExtension(index int) string {
	if index >= 0 && index < len(b.ResolveExtensions) {
		return b.ResolveExtensions[index]
	}
	return ""
}

func (b *BuildOptions) GetInject(index int) string {
	if index >= 0 && index < len(b.Inject) {
		return b.Inject[index]
	}
	return ""
}

func (b *BuildOptions) GetNodePath(index int) string {
	if index >= 0 && index < len(b.NodePaths) {
		return b.NodePaths[index]
	}
	return ""
}

func (b *BuildOptions) GetPure(index int) string {
	if index >= 0 && index < len(b.Pure) {
		return b.Pure[index]
	}
	return ""
}

func (b *BuildOptions) GetDropLabel(index int) string {
	if index >= 0 && index < len(b.DropLabels) {
		return b.DropLabels[index]
	}
	return ""
}

func (b *BuildOptions) GetPlugin(index int) *Plugin {
	if index >= 0 && index < len(b.Plugins) {
		return b.Plugins[index]
	}
	return nil
}

// Getter methods for map sizes
func (b *BuildOptions) GetDefineCount() int       { return len(b.Define) }
func (b *BuildOptions) GetLogOverrideCount() int  { return len(b.LogOverride) }
func (b *BuildOptions) GetAliasCount() int        { return len(b.Alias) }
func (b *BuildOptions) GetLoaderCount() int       { return len(b.Loader) }
func (b *BuildOptions) GetOutExtensionCount() int { return len(b.OutExtension) }
func (b *BuildOptions) GetBannerCount() int       { return len(b.Banner) }
func (b *BuildOptions) GetFooterCount() int       { return len(b.Footer) }
func (b *BuildOptions) GetSupportedCount() int    { return len(b.Supported) }
func (b *BuildOptions) GetMangleCacheCount() int  { return len(b.MangleCache) }

// Methods to get all keys from maps (for iteration in Swift)
func (b *BuildOptions) GetDefineKeys() []string {
	keys := make([]string, 0, len(b.Define))
	for k := range b.Define {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetLogOverrideKeys() []string {
	keys := make([]string, 0, len(b.LogOverride))
	for k := range b.LogOverride {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetAliasKeys() []string {
	keys := make([]string, 0, len(b.Alias))
	for k := range b.Alias {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetLoaderKeys() []string {
	keys := make([]string, 0, len(b.Loader))
	for k := range b.Loader {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetOutExtensionKeys() []string {
	keys := make([]string, 0, len(b.OutExtension))
	for k := range b.OutExtension {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetBannerKeys() []string {
	keys := make([]string, 0, len(b.Banner))
	for k := range b.Banner {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetFooterKeys() []string {
	keys := make([]string, 0, len(b.Footer))
	for k := range b.Footer {
		keys = append(keys, k)
	}
	return keys
}

func (b *BuildOptions) GetSupportedKeys() []string {
	keys := make([]string, 0, len(b.Supported))
	for k := range b.Supported {
		keys = append(keys, k)
	}
	return keys
}

// Methods to get values from maps by key
func (b *BuildOptions) GetDefineValue(key string) string {
	if b.Define == nil {
		return ""
	}
	return b.Define[key]
}

func (b *BuildOptions) GetLogOverrideValue(key string) int {
	if b.LogOverride == nil {
		return int(api.LogLevelInfo)
	}
	return int(b.LogOverride[key])
}

func (b *BuildOptions) GetAliasValue(key string) string {
	if b.Alias == nil {
		return ""
	}
	return b.Alias[key]
}

func (b *BuildOptions) GetLoaderValue(key string) int {
	if b.Loader == nil {
		return int(api.LoaderDefault)
	}
	return int(b.Loader[key])
}

func (b *BuildOptions) GetOutExtensionValue(key string) string {
	if b.OutExtension == nil {
		return ""
	}
	return b.OutExtension[key]
}

func (b *BuildOptions) GetBannerValue(key string) string {
	if b.Banner == nil {
		return ""
	}
	return b.Banner[key]
}

func (b *BuildOptions) GetFooterValue(key string) string {
	if b.Footer == nil {
		return ""
	}
	return b.Footer[key]
}

func (b *BuildOptions) GetSupportedValue(key string) bool {
	if b.Supported == nil {
		return false
	}
	return b.Supported[key]
}

// Methods to set map values by key with integer enum values
func (b *BuildOptions) SetLogOverrideValueByInt(key string, value int) {
	if b.LogOverride == nil {
		b.LogOverride = make(map[string]api.LogLevel)
	}
	b.LogOverride[key] = api.LogLevel(value)
}

func (b *BuildOptions) SetLoaderValueByInt(key string, value int) {
	if b.Loader == nil {
		b.Loader = make(map[string]api.Loader)
	}
	b.Loader[key] = api.Loader(value)
}

// Methods for EntryPointsAdvanced (complex type)
func (b *BuildOptions) GetAdvancedEntryPointInput(index int) string {
	if index >= 0 && index < len(b.EntryPointsAdvanced) {
		return b.EntryPointsAdvanced[index].InputPath
	}
	return ""
}

func (b *BuildOptions) GetAdvancedEntryPointOutput(index int) string {
	if index >= 0 && index < len(b.EntryPointsAdvanced) {
		return b.EntryPointsAdvanced[index].OutputPath
	}
	return ""
}

// Clear methods for arrays and maps
func (b *BuildOptions) ClearEntryPoints()         { b.EntryPoints = nil }
func (b *BuildOptions) ClearEntryPointsAdvanced() { b.EntryPointsAdvanced = nil }
func (b *BuildOptions) ClearExternal()            { b.External = nil }
func (b *BuildOptions) ClearMainFields()          { b.MainFields = nil }
func (b *BuildOptions) ClearConditions()          { b.Conditions = nil }
func (b *BuildOptions) ClearResolveExtensions()   { b.ResolveExtensions = nil }
func (b *BuildOptions) ClearInject()              { b.Inject = nil }
func (b *BuildOptions) ClearNodePaths()           { b.NodePaths = nil }
func (b *BuildOptions) ClearPure()                { b.Pure = nil }
func (b *BuildOptions) ClearDropLabels()          { b.DropLabels = nil }
func (b *BuildOptions) ClearEngines()             { b.Engines = nil }
func (b *BuildOptions) ClearPlugins()             { b.Plugins = make([]*Plugin, 0) }
func (b *BuildOptions) ClearDefine()              { b.Define = nil }
func (b *BuildOptions) ClearLogOverride()         { b.LogOverride = nil }
func (b *BuildOptions) ClearAlias()               { b.Alias = nil }
func (b *BuildOptions) ClearLoader()              { b.Loader = nil }
func (b *BuildOptions) ClearOutExtension()        { b.OutExtension = nil }
func (b *BuildOptions) ClearBanner()              { b.Banner = nil }
func (b *BuildOptions) ClearFooter()              { b.Footer = nil }
func (b *BuildOptions) ClearSupported()           { b.Supported = nil }
func (b *BuildOptions) ClearMangleCache()         { b.MangleCache = nil }
