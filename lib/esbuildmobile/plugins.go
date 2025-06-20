package esbuildmobile

import "github.com/evanw/esbuild/pkg/api"

// Callback interfaces for mobile compatibility
type OnResolveCallback interface {
	Call(args *OnResolveArgs) *OnResolveResult
}

type OnLoadCallback interface {
	Call(args *OnLoadArgs) *OnLoadResult
}

type OnStartCallback interface {
	Call() *OnStartResult
}

type OnEndCallback interface {
	Call(result *BuildResult) *OnEndResult
}

// Plugin represents an ESBuild plugin
type Plugin struct {
	name            string
	onResolveRules  []onResolveRule
	onLoadRules     []onLoadRule
	onStartCallback OnStartCallback
	onEndCallback   OnEndCallback
}

type onResolveRule struct {
	Options  *OnResolveOptions
	Callback OnResolveCallback
}

type onLoadRule struct {
	Options  *OnLoadOptions
	Callback OnLoadCallback
}

// BuildResult represents the result of a build (simplified for mobile)
type BuildResult struct {
	Errors   []api.Message
	Warnings []api.Message
}

// SideEffects represents the side effects setting for a module
type SideEffects uint8

const (
	SideEffectsTrue SideEffects = iota
	SideEffectsFalse
)

// ResolveKind represents the type of import being resolved
type ResolveKind uint8

const (
	ResolveNone ResolveKind = iota
	ResolveEntryPoint
	ResolveJSImportStatement
	ResolveJSRequireCall
	ResolveJSDynamicImport
	ResolveJSRequireResolve
	ResolveCSSImportRule
	ResolveCSSComposesFrom
	ResolveCSSURLToken
)

// ResolveOptions contains options for resolving a module
type ResolveOptions struct {
	PluginName string
	Importer   string
	Namespace  string
	ResolveDir string
	Kind       ResolveKind
	// Note: PluginData and With are omitted as they can't be exported to mobile
}

// ResolveResult contains the result of resolving a module
type ResolveResult struct {
	Errors   []api.Message
	Warnings []api.Message

	Path        string
	External    bool
	SideEffects bool
	Namespace   string
	Suffix      string
	// Note: PluginData is omitted as it can't be exported to mobile
}

// OnStartResult contains the result from an OnStart callback
type OnStartResult struct {
	Errors   []api.Message
	Warnings []api.Message
}

// OnEndResult contains the result from an OnEnd callback
type OnEndResult struct {
	Errors   []api.Message
	Warnings []api.Message
}

// OnResolveOptions contains options for the OnResolve callback
type OnResolveOptions struct {
	Filter    string
	Namespace string
}

// OnResolveArgs contains arguments passed to the OnResolve callback
type OnResolveArgs struct {
	Path       string
	Importer   string
	Namespace  string
	ResolveDir string
	Kind       ResolveKind
	// Note: PluginData and With are omitted as they can't be exported to mobile
}

// OnResolveResult contains the result from an OnResolve callback
type OnResolveResult struct {
	PluginName string

	Errors   []api.Message
	Warnings []api.Message

	Path        string
	External    bool
	SideEffects SideEffects
	Namespace   string
	Suffix      string
	// Note: PluginData is omitted as it can't be exported to mobile

	WatchFiles []string
	WatchDirs  []string
}

// OnLoadOptions contains options for the OnLoad callback
type OnLoadOptions struct {
	Filter    string
	Namespace string
}

// OnLoadArgs contains arguments passed to the OnLoad callback
type OnLoadArgs struct {
	Path      string
	Namespace string
	Suffix    string
	// Note: PluginData and With are omitted as they can't be exported to mobile
}

// OnLoadResult contains the result from an OnLoad callback
type OnLoadResult struct {
	PluginName string

	Errors   []api.Message
	Warnings []api.Message

	Contents   *string
	ResolveDir string
	Loader     api.Loader
	// Note: PluginData is omitted as it can't be exported to mobile

	WatchFiles []string
	WatchDirs  []string
}

// Conversion functions to/from API types

func (s SideEffects) ToAPI() api.SideEffects {
	switch s {
	case SideEffectsTrue:
		return api.SideEffectsTrue
	case SideEffectsFalse:
		return api.SideEffectsFalse
	default:
		return api.SideEffectsTrue
	}
}

func SideEffectsFromAPI(s api.SideEffects) SideEffects {
	switch s {
	case api.SideEffectsTrue:
		return SideEffectsTrue
	case api.SideEffectsFalse:
		return SideEffectsFalse
	default:
		return SideEffectsTrue
	}
}

func (r ResolveKind) ToAPI() api.ResolveKind {
	switch r {
	case ResolveNone:
		return api.ResolveNone
	case ResolveEntryPoint:
		return api.ResolveEntryPoint
	case ResolveJSImportStatement:
		return api.ResolveJSImportStatement
	case ResolveJSRequireCall:
		return api.ResolveJSRequireCall
	case ResolveJSDynamicImport:
		return api.ResolveJSDynamicImport
	case ResolveJSRequireResolve:
		return api.ResolveJSRequireResolve
	case ResolveCSSImportRule:
		return api.ResolveCSSImportRule
	case ResolveCSSComposesFrom:
		return api.ResolveCSSComposesFrom
	case ResolveCSSURLToken:
		return api.ResolveCSSURLToken
	default:
		return api.ResolveNone
	}
}

func ResolveKindFromAPI(r api.ResolveKind) ResolveKind {
	switch r {
	case api.ResolveNone:
		return ResolveNone
	case api.ResolveEntryPoint:
		return ResolveEntryPoint
	case api.ResolveJSImportStatement:
		return ResolveJSImportStatement
	case api.ResolveJSRequireCall:
		return ResolveJSRequireCall
	case api.ResolveJSDynamicImport:
		return ResolveJSDynamicImport
	case api.ResolveJSRequireResolve:
		return ResolveJSRequireResolve
	case api.ResolveCSSImportRule:
		return ResolveCSSImportRule
	case api.ResolveCSSComposesFrom:
		return ResolveCSSComposesFrom
	case api.ResolveCSSURLToken:
		return ResolveCSSURLToken
	default:
		return ResolveNone
	}
}

func (o *OnResolveOptions) ToAPI() api.OnResolveOptions {
	return api.OnResolveOptions{
		Filter:    o.Filter,
		Namespace: o.Namespace,
	}
}

func OnResolveOptionsFromAPI(o api.OnResolveOptions) *OnResolveOptions {
	return &OnResolveOptions{
		Filter:    o.Filter,
		Namespace: o.Namespace,
	}
}

func (o *OnLoadOptions) ToAPI() api.OnLoadOptions {
	return api.OnLoadOptions{
		Filter:    o.Filter,
		Namespace: o.Namespace,
	}
}

func OnLoadOptionsFromAPI(o api.OnLoadOptions) *OnLoadOptions {
	return &OnLoadOptions{
		Filter:    o.Filter,
		Namespace: o.Namespace,
	}
}

// Getter functions for constants (gomobile friendly)

func GetSideEffectsTrue() SideEffects  { return SideEffectsTrue }
func GetSideEffectsFalse() SideEffects { return SideEffectsFalse }

func GetResolveNone() ResolveKind              { return ResolveNone }
func GetResolveEntryPoint() ResolveKind        { return ResolveEntryPoint }
func GetResolveJSImportStatement() ResolveKind { return ResolveJSImportStatement }
func GetResolveJSRequireCall() ResolveKind     { return ResolveJSRequireCall }
func GetResolveJSDynamicImport() ResolveKind   { return ResolveJSDynamicImport }
func GetResolveJSRequireResolve() ResolveKind  { return ResolveJSRequireResolve }
func GetResolveCSSImportRule() ResolveKind     { return ResolveCSSImportRule }
func GetResolveCSSComposesFrom() ResolveKind   { return ResolveCSSComposesFrom }
func GetResolveCSSURLToken() ResolveKind       { return ResolveCSSURLToken }

// Constructor functions

func NewOnResolveOptions() *OnResolveOptions {
	return &OnResolveOptions{}
}

func NewOnLoadOptions() *OnLoadOptions {
	return &OnLoadOptions{}
}

func NewOnResolveResult() *OnResolveResult {
	return &OnResolveResult{
		Errors:     make([]api.Message, 0),
		Warnings:   make([]api.Message, 0),
		WatchFiles: make([]string, 0),
		WatchDirs:  make([]string, 0),
	}
}

func NewOnLoadResult() *OnLoadResult {
	return &OnLoadResult{
		Errors:     make([]api.Message, 0),
		Warnings:   make([]api.Message, 0),
		WatchFiles: make([]string, 0),
		WatchDirs:  make([]string, 0),
	}
}

func NewOnStartResult() *OnStartResult {
	return &OnStartResult{
		Errors:   make([]api.Message, 0),
		Warnings: make([]api.Message, 0),
	}
}

func NewOnEndResult() *OnEndResult {
	return &OnEndResult{
		Errors:   make([]api.Message, 0),
		Warnings: make([]api.Message, 0),
	}
}

// Plugin constructor and configuration methods

func NewPlugin(name string) *Plugin {
	return &Plugin{
		name:           name,
		onResolveRules: make([]onResolveRule, 0),
		onLoadRules:    make([]onLoadRule, 0),
	}
}

func (p *Plugin) GetName() string {
	return p.name
}

func (p *Plugin) SetPluginName(name string) {
	p.name = name
}

// OnResolve adds a resolve callback
func (p *Plugin) OnResolve(options *OnResolveOptions, callback OnResolveCallback) {
	p.onResolveRules = append(p.onResolveRules, onResolveRule{
		Options:  options,
		Callback: callback,
	})
}

// OnLoad adds a load callback
func (p *Plugin) OnLoad(options *OnLoadOptions, callback OnLoadCallback) {
	p.onLoadRules = append(p.onLoadRules, onLoadRule{
		Options:  options,
		Callback: callback,
	})
}

// OnStart sets the start callback
func (p *Plugin) OnStart(callback OnStartCallback) {
	p.onStartCallback = callback
}

// OnEnd sets the end callback
func (p *Plugin) OnEnd(callback OnEndCallback) {
	p.onEndCallback = callback
}

// Getter methods for rules
func (p *Plugin) GetOnResolveRulesCount() int {
	return len(p.onResolveRules)
}

func (p *Plugin) GetOnLoadRulesCount() int {
	return len(p.onLoadRules)
}

func (p *Plugin) GetOnResolveRuleOptions(index int) *OnResolveOptions {
	if index >= 0 && index < len(p.onResolveRules) {
		return p.onResolveRules[index].Options
	}
	return nil
}

func (p *Plugin) GetOnResolveRuleCallback(index int) OnResolveCallback {
	if index >= 0 && index < len(p.onResolveRules) {
		return p.onResolveRules[index].Callback
	}
	return nil
}

func (p *Plugin) GetOnLoadRuleOptions(index int) *OnLoadOptions {
	if index >= 0 && index < len(p.onLoadRules) {
		return p.onLoadRules[index].Options
	}
	return nil
}

func (p *Plugin) GetOnLoadRuleCallback(index int) OnLoadCallback {
	if index >= 0 && index < len(p.onLoadRules) {
		return p.onLoadRules[index].Callback
	}
	return nil
}

func (p *Plugin) GetOnStartCallback() OnStartCallback {
	return p.onStartCallback
}

func (p *Plugin) GetOnEndCallback() OnEndCallback {
	return p.onEndCallback
}

func (p *Plugin) HasOnStartCallback() bool {
	return p.onStartCallback != nil
}

func (p *Plugin) HasOnEndCallback() bool {
	return p.onEndCallback != nil
}

// Helper methods for creating plugin results with common patterns

// CreateResolveResult creates a basic resolve result
func CreateResolveResult(path string, external bool, namespace string) *OnResolveResult {
	result := NewOnResolveResult()
	result.Path = path
	result.External = external
	result.Namespace = namespace
	return result
}

// CreateLoadResult creates a basic load result
func CreateLoadResult(contents string, loader api.Loader) *OnLoadResult {
	result := NewOnLoadResult()
	result.Contents = &contents
	result.Loader = loader
	return result
}

// CreateLoadResultWithResolveDir creates a load result with resolve directory
func CreateLoadResultWithResolveDir(contents string, loader api.Loader, resolveDir string) *OnLoadResult {
	result := NewOnLoadResult()
	result.Contents = &contents
	result.Loader = loader
	result.ResolveDir = resolveDir
	return result
}

// Loader helper functions for common use cases

// CreateJSLoadResult creates a load result with JS loader
func CreateJSLoadResult(contents string) *OnLoadResult {
	return CreateLoadResult(contents, api.LoaderJS)
}

// CreateTSLoadResult creates a load result with TypeScript loader
func CreateTSLoadResult(contents string) *OnLoadResult {
	return CreateLoadResult(contents, api.LoaderTS)
}

// CreateCSSLoadResult creates a load result with CSS loader
func CreateCSSLoadResult(contents string) *OnLoadResult {
	return CreateLoadResult(contents, api.LoaderCSS)
}

// CreateJSONLoadResult creates a load result with JSON loader
func CreateJSONLoadResult(contents string) *OnLoadResult {
	return CreateLoadResult(contents, api.LoaderJSON)
}

// CreateTextLoadResult creates a load result with text loader
func CreateTextLoadResult(contents string) *OnLoadResult {
	return CreateLoadResult(contents, api.LoaderText)
}

// Common filter pattern constants
const (
	FilterReact           = "^react$"
	FilterReactDOM        = "^react-dom$"
	FilterLodash          = "^lodash$"
	FilterJQuery          = "^jquery$"
	FilterNodeModules     = "^[^.].*/.*$"
	FilterRelativeImports = "^\\.\\.?/"
	FilterAbsoluteImports = "^/"
	FilterAllFiles        = ".*"
	FilterJSFiles         = "\\.jsx?$"
	FilterTSFiles         = "\\.tsx?$"
	FilterCSSFiles        = "\\.css$"
	FilterSCSSFiles       = "\\.scss$"
	FilterLESSFiles       = "\\.less$"
	FilterJSONFiles       = "\\.json$"
	FilterImageFiles      = "\\.(png|jpg|jpeg|gif|svg|webp)$"
	FilterFontFiles       = "\\.(woff|woff2|eot|ttf|otf)$"
)

// Common namespace constants
const (
	NamespaceFile        = "file"
	NamespaceHTTP        = "http"
	NamespaceHTTPS       = "https"
	NamespaceVirtual     = "virtual"
	NamespaceGenerated   = "generated"
	NamespaceTransformed = "transformed"
	NamespaceExternal    = "external"
)

// Getter functions for filter

// Common resolve patterns

// CreateExternalResolveResult creates a resolve result that marks the import as external
func CreateExternalResolveResult(path string) *OnResolveResult {
	return CreateResolveResult(path, true, "")
}

// CreateNamespaceResolveResult creates a resolve result that redirects to a specific namespace
func CreateNamespaceResolveResult(path string, namespace string) *OnResolveResult {
	return CreateResolveResult(path, false, namespace)
}

// Filter helper functions for creating common filter patterns

// CreateFilterForPath creates an OnResolveOptions with a path filter
func CreateFilterForPath(pathPattern string) *OnResolveOptions {
	options := NewOnResolveOptions()
	options.SetResolveFilter(pathPattern)
	return options
}

// CreateFilterForNamespace creates an OnLoadOptions with a namespace filter
func CreateFilterForNamespace(namespace string) *OnLoadOptions {
	options := NewOnLoadOptions()
	options.SetLoadNamespace(namespace)
	return options
}

// CreateFilterForPathAndNamespace creates an OnResolveOptions with both path and namespace filters
func CreateFilterForPathAndNamespace(pathPattern string, namespace string) *OnResolveOptions {
	options := NewOnResolveOptions()
	options.SetResolveFilter(pathPattern)
	options.SetResolveNamespace(namespace)
	return options
}

// Common plugin creation patterns

// CreateSimpleTransformPlugin creates a plugin that transforms imports matching a pattern
func CreateSimpleTransformPlugin(name string, pathPattern string, namespace string, contents string, loader api.Loader) *Plugin {
	plugin := NewPlugin(name)

	// Add resolve rule to redirect matching imports to our namespace
	resolveOptions := CreateFilterForPath(pathPattern)
	plugin.OnResolve(resolveOptions, &SimpleResolveCallback{
		Path:      pathPattern,
		Namespace: namespace,
	})

	// Add load rule to provide the transformed contents
	loadOptions := CreateFilterForNamespace(namespace)
	plugin.OnLoad(loadOptions, &SimpleLoadCallback{
		Contents: contents,
		Loader:   int(loader),
	})

	return plugin
}

// Simple callback implementations for common use cases
type SimpleResolveCallback struct {
	Path      string
	Namespace string
	External  bool
}

func (c *SimpleResolveCallback) Call(args *OnResolveArgs) *OnResolveResult {
	if c.External {
		return CreateExternalResolveResult(c.Path)
	}
	if c.Namespace != "" {
		return CreateNamespaceResolveResult(c.Path, c.Namespace)
	}
	// Default: just return the path as-is
	result := NewOnResolveResult()
	result.Path = c.Path
	return result
}

type SimpleLoadCallback struct {
	Contents string
	Loader   int // Use int for gomobile compatibility
}

func (c *SimpleLoadCallback) Call(args *OnLoadArgs) *OnLoadResult {
	return CreateLoadResult(c.Contents, api.Loader(c.Loader))
}

// Configuration methods for OnResolveOptions

func (o *OnResolveOptions) SetResolveFilter(filter string) {
	o.Filter = filter
}

func (o *OnResolveOptions) SetResolveNamespace(namespace string) {
	o.Namespace = namespace
}

func (o *OnResolveOptions) GetResolveFilter() string {
	return o.Filter
}

func (o *OnResolveOptions) GetResolveNamespace() string {
	return o.Namespace
}

// Configuration methods for OnLoadOptions

func (o *OnLoadOptions) SetLoadFilter(filter string) {
	o.Filter = filter
}

func (o *OnLoadOptions) SetLoadNamespace(namespace string) {
	o.Namespace = namespace
}

func (o *OnLoadOptions) GetLoadFilter() string {
	return o.Filter
}

func (o *OnLoadOptions) GetLoadNamespace() string {
	return o.Namespace
}

// Configuration methods for OnResolveResult

func (o *OnResolveResult) SetResolvePluginName(name string) {
	o.PluginName = name
}

func (o *OnResolveResult) SetResolvePath(path string) {
	o.Path = path
}

func (o *OnResolveResult) SetResolveExternal(external bool) {
	o.External = external
}

func (o *OnResolveResult) SetResolveSideEffects(sideEffects SideEffects) {
	o.SideEffects = sideEffects
}

func (o *OnResolveResult) SetResolveNamespace(namespace string) {
	o.Namespace = namespace
}

func (o *OnResolveResult) SetResolveSuffix(suffix string) {
	o.Suffix = suffix
}

func (o *OnResolveResult) AddResolveWatchFile(file string) {
	o.WatchFiles = append(o.WatchFiles, file)
}

func (o *OnResolveResult) AddResolveWatchDir(dir string) {
	o.WatchDirs = append(o.WatchDirs, dir)
}

func (o *OnResolveResult) AddResolveError(message api.Message) {
	o.Errors = append(o.Errors, message)
}

func (o *OnResolveResult) AddResolveWarning(message api.Message) {
	o.Warnings = append(o.Warnings, message)
}

// Configuration methods for OnLoadResult

func (o *OnLoadResult) SetLoadPluginName(name string) {
	o.PluginName = name
}

func (o *OnLoadResult) SetLoadContents(contents string) {
	o.Contents = &contents
}

func (o *OnLoadResult) SetLoadResolveDir(dir string) {
	o.ResolveDir = dir
}

func (o *OnLoadResult) SetLoadLoader(loader api.Loader) {
	o.Loader = loader
}

func (o *OnLoadResult) AddLoadWatchFile(file string) {
	o.WatchFiles = append(o.WatchFiles, file)
}

func (o *OnLoadResult) AddLoadWatchDir(dir string) {
	o.WatchDirs = append(o.WatchDirs, dir)
}

func (o *OnLoadResult) AddLoadError(message api.Message) {
	o.Errors = append(o.Errors, message)
}

func (o *OnLoadResult) AddLoadWarning(message api.Message) {
	o.Warnings = append(o.Warnings, message)
}

// Getter methods

func (o *OnResolveResult) GetResolvePluginName() string       { return o.PluginName }
func (o *OnResolveResult) GetResolvePath() string             { return o.Path }
func (o *OnResolveResult) GetResolveExternal() bool           { return o.External }
func (o *OnResolveResult) GetResolveSideEffects() SideEffects { return o.SideEffects }
func (o *OnResolveResult) GetResolveNamespace() string        { return o.Namespace }
func (o *OnResolveResult) GetResolveSuffix() string           { return o.Suffix }
func (o *OnResolveResult) GetResolveWatchFilesCount() int     { return len(o.WatchFiles) }
func (o *OnResolveResult) GetResolveWatchDirsCount() int      { return len(o.WatchDirs) }
func (o *OnResolveResult) GetResolveErrorsCount() int         { return len(o.Errors) }
func (o *OnResolveResult) GetResolveWarningsCount() int       { return len(o.Warnings) }

func (o *OnResolveResult) GetResolveWatchFile(index int) string {
	if index >= 0 && index < len(o.WatchFiles) {
		return o.WatchFiles[index]
	}
	return ""
}

func (o *OnResolveResult) GetResolveWatchDir(index int) string {
	if index >= 0 && index < len(o.WatchDirs) {
		return o.WatchDirs[index]
	}
	return ""
}

func (o *OnResolveResult) GetResolveError(index int) api.Message {
	if index >= 0 && index < len(o.Errors) {
		return o.Errors[index]
	}
	return api.Message{}
}

func (o *OnResolveResult) GetResolveWarning(index int) api.Message {
	if index >= 0 && index < len(o.Warnings) {
		return o.Warnings[index]
	}
	return api.Message{}
}

func (o *OnLoadResult) GetLoadPluginName() string { return o.PluginName }
func (o *OnLoadResult) GetLoadContents() string {
	if o.Contents != nil {
		return *o.Contents
	}
	return ""
}
func (o *OnLoadResult) GetLoadResolveDir() string   { return o.ResolveDir }
func (o *OnLoadResult) GetLoadLoader() api.Loader   { return o.Loader }
func (o *OnLoadResult) GetLoadWatchFilesCount() int { return len(o.WatchFiles) }
func (o *OnLoadResult) GetLoadWatchDirsCount() int  { return len(o.WatchDirs) }
func (o *OnLoadResult) GetLoadErrorsCount() int     { return len(o.Errors) }
func (o *OnLoadResult) GetLoadWarningsCount() int   { return len(o.Warnings) }

func (o *OnLoadResult) GetLoadWatchFile(index int) string {
	if index >= 0 && index < len(o.WatchFiles) {
		return o.WatchFiles[index]
	}
	return ""
}

func (o *OnLoadResult) GetLoadWatchDir(index int) string {
	if index >= 0 && index < len(o.WatchDirs) {
		return o.WatchDirs[index]
	}
	return ""
}

func (o *OnLoadResult) GetLoadError(index int) api.Message {
	if index >= 0 && index < len(o.Errors) {
		return o.Errors[index]
	}
	return api.Message{}
}

func (o *OnLoadResult) GetLoadWarning(index int) api.Message {
	if index >= 0 && index < len(o.Warnings) {
		return o.Warnings[index]
	}
	return api.Message{}
}

// Swift Bridge Callbacks
// These types allow Swift to create plugin callbacks without subclassing

// SwiftOnStartCallback is a concrete implementation that can be used from Swift
type SwiftOnStartCallback struct {
	handler func() *OnStartResult
}

// NewSwiftOnStartCallback creates a callback that executes a simple handler
func NewSwiftOnStartCallback() *SwiftOnStartCallback {
	return &SwiftOnStartCallback{}
}

// SetHandler sets the handler function (for internal use)
func (c *SwiftOnStartCallback) SetHandler(handler func() *OnStartResult) {
	c.handler = handler
}

// Call implements the OnStartCallback interface
func (c *SwiftOnStartCallback) Call() *OnStartResult {
	if c.handler != nil {
		return c.handler()
	}
	return NewOnStartResult()
}

// SimpleOnStartCallback provides a basic implementation for Swift
type SimpleOnStartCallback struct {
	PrintMessage string
}

// NewSimpleOnStartCallback creates a simple callback for testing
func NewSimpleOnStartCallback(message string) *SimpleOnStartCallback {
	return &SimpleOnStartCallback{PrintMessage: message}
}

// Call implements the OnStartCallback interface
func (c *SimpleOnStartCallback) Call() *OnStartResult {
	if c.PrintMessage != "" {
		println(c.PrintMessage)
	}
	return NewOnStartResult()
}

// SwiftOnEndCallback is a concrete implementation that can be used from Swift
type SwiftOnEndCallback struct {
	handler func(*BuildResult) *OnEndResult
}

// NewSwiftOnEndCallback creates a callback that executes a handler
func NewSwiftOnEndCallback() *SwiftOnEndCallback {
	return &SwiftOnEndCallback{}
}

// SetHandler sets the handler function
func (c *SwiftOnEndCallback) SetHandler(handler func(*BuildResult) *OnEndResult) {
	c.handler = handler
}

// Call implements the OnEndCallback interface
func (c *SwiftOnEndCallback) Call(result *BuildResult) *OnEndResult {
	if c.handler != nil {
		return c.handler(result)
	}
	return NewOnEndResult()
}

// SwiftOnResolveCallback provides resolve handling for Swift
type SwiftOnResolveCallback struct {
	handler func(*OnResolveArgs) *OnResolveResult
}

// NewSwiftOnResolveCallback creates a new resolve callback
func NewSwiftOnResolveCallback() *SwiftOnResolveCallback {
	return &SwiftOnResolveCallback{}
}

// SetHandler sets the handler function
func (c *SwiftOnResolveCallback) SetHandler(handler func(*OnResolveArgs) *OnResolveResult) {
	c.handler = handler
}

// Call implements the OnResolveCallback interface
func (c *SwiftOnResolveCallback) Call(args *OnResolveArgs) *OnResolveResult {
	if c.handler != nil {
		return c.handler(args)
	}
	return NewOnResolveResult()
}

// SwiftOnLoadCallback provides load handling for Swift
type SwiftOnLoadCallback struct {
	handler func(*OnLoadArgs) *OnLoadResult
}

// NewSwiftOnLoadCallback creates a new load callback
func NewSwiftOnLoadCallback() *SwiftOnLoadCallback {
	return &SwiftOnLoadCallback{}
}

// SetHandler sets the handler function
func (c *SwiftOnLoadCallback) SetHandler(handler func(*OnLoadArgs) *OnLoadResult) {
	c.handler = handler
}

// Call implements the OnLoadCallback interface
func (c *SwiftOnLoadCallback) Call(args *OnLoadArgs) *OnLoadResult {
	if c.handler != nil {
		return c.handler(args)
	}
	return NewOnLoadResult()
}
