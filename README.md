# ESBuild

A Swift package that provides ESBuild transformation capabilities for mobile development. This package uses `gomobile` to bridge ESBuild's Go API with native Swift, providing an idiomatic Swift API that hides all the underlying complexity.

## Features

- ✅ **Idiomatic Swift API** - Native Swift properties and method chaining
- ✅ **Complete ESBuild Support** - All transform options available through Swift
- ✅ **Type-Safe Configuration** - Swift enums for all ESBuild options
- ✅ **Framework Presets** - Built-in configurations for React, Preact, Vue
- ✅ **Method Chaining** - Fluent API for easy configuration
- ✅ **String Extensions** - Transform code directly on String instances
- ✅ **Comprehensive Error Handling** - Swift-native error types
- ✅ **Reusable Transformers** - Configure once, transform multiple times

## Installation

Add this package to your Swift Package Manager dependencies:

```swift
dependencies: [
    .package(url: "https://github.com/zshannon/ESBuildMobile", from: "1.2.0")
]
```

Then import the module in your Swift code:

```swift
import ESBuild
```

## Quick Start

### Basic Transformation

```swift
import ESBuild

// Transform with default settings
let code = "<div>Hello World</div>"
let transformer = Transform()
let result = try transformer.transform(code)
print(result)
// Output: /* @__PURE__ */ React.createElement("div", null, "Hello World");

// Or use the String extension
let result = try code.transform()
```

### Using Presets

```swift
// React (default)
let result = try code.transform(with: TransformOptions.react())

// Preact
let transformer = Transform(TransformOptions.preact())
let result = try transformer.transform(code)
// Output: /* @__PURE__ */ h("div", null, "Hello World");

// Vue 3
let transformer = Transform(TransformOptions.vue())
let result = try transformer.transform(code)

// Production build (minified)
let transformer = Transform(TransformOptions.production())
let result = try transformer.transform(code)
```

## Idiomatic Swift API

### Configure Once, Transform Many

```swift
// Configure your transformer once
let transformer = Transform(TransformOptions(
    jsx: .automatic,
    minifyWhitespace: true,
    target: .es2020,
    format: .esm
))

// Reuse for multiple transforms
let result1 = try transformer.transform(jsxCode1)
let result2 = try transformer.transform(jsxCode2)
let result3 = try transformer.transform(jsxCode3)
```

### Property-Based Configuration

```swift
let options = TransformOptions()

// JSX Configuration
options.jsxFactory = "h"
options.jsxFragment = "Fragment"
options.jsx = .automatic
options.jsxDev = true

// Output Configuration
options.format = .esm
options.platform = .browser
options.target = .es2020

// Minification
options.minifyWhitespace = true
options.minifyIdentifiers = true
options.minifySyntax = true

// Advanced Options
options.sourcemap = .inline
options.loader = .tsx
options.charset = .utf8
options.treeShaking = true

let transformer = Transform(options)
let result = try transformer.transform(code)
```

### Convenient Initialization

```swift
// Initialize with all options at once
let transformer = Transform(TransformOptions(
    jsxFactory: "h",
    jsxFragment: "Fragment",
    jsx: .automatic,
    minifyWhitespace: true,
    minifyIdentifiers: true,
    minifySyntax: true,
    format: .esm,
    platform: .browser,
    target: .es2020,
    sourcemap: .inline,
    loader: .tsx
))
```

### Method Chaining

```swift
let transformer = Transform(
    TransformOptions()
        .jsx(factory: "createElement", fragment: "Fragment")
        .minify()
        .development()
)

let result = try transformer.transform(code)
```

### Define Variables

```swift
let options = TransformOptions()
options.setDefine("process.env.NODE_ENV", value: "production")
options.setDefine("__VERSION__", value: "1.0.0")

// Or set multiple at once
options.define = [
    "process.env.NODE_ENV": "production",
    "__VERSION__": "1.0.0"
]

let transformer = Transform(options)
```

## Framework-Specific Examples

### React

```swift
let options = TransformOptions.react()
options.jsxDev = true  // For development
options.target = .es2020
options.format = .esm

let transformer = Transform(options)
let result = try transformer.transform("""
    function Welcome({name}) {
        return <div>Hello, {name}!</div>;
    }
""")
```

### React with TypeScript

```swift
let options = TransformOptions.reactTypeScript()
options.loader = .tsx
options.keepNames = true

let transformer = Transform(options)
let result = try transformer.transform("""
    interface Props {
        name: string;
    }

    const Welcome: React.FC<Props> = ({name}) => {
        return <div>Hello, {name}!</div>;
    };
""")
```

### Preact

```swift
let options = TransformOptions.preact()
options.format = .esm
options.platform = .browser

let transformer = Transform(options)
let result = try transformer.transform("""
    import { h } from 'preact';

    export function App() {
        return <div>Hello from Preact!</div>;
    }
""")
```

### Vue 3 JSX

```swift
let options = TransformOptions.vue()
options.target = .es2020
options.format = .esm

let transformer = Transform(options)
let result = try transformer.transform("""
    export default {
        render() {
            return <div>Hello from Vue JSX!</div>;
        }
    };
""")
```

## Advanced Configuration

### Custom Build Pipeline

```swift
let options = TransformOptions(
    // Configure for modern browsers
    target: .es2020,
    platform: .browser,
    format: .esm,
    
    // Enable all optimizations
    minifyWhitespace: true,
    minifyIdentifiers: true,
    minifySyntax: true,
    treeShaking: true,
    drop: .console,
    
    // Configure source maps for debugging
    sourcemap: .inline,
    sourcesContent: .include,
    
    // Set up defines for environment
    define: [
        "process.env.NODE_ENV": "production",
        "__DEV__": "false"
    ]
)

let transformer = Transform(options)
let result = try transformer.transform(complexCode)
```

### Multiple File Types

```swift
// TypeScript JSX transformer
let tsxTransformer = Transform(TransformOptions(
    loader: .tsx,
    target: .es2020
))

// Regular JavaScript transformer
let jsTransformer = Transform(TransformOptions(
    loader: .js,
    format: .cjs
))

// CSS-in-JS transformer
let cssTransformer = Transform(TransformOptions(
    loader: .css
))
```

## Error Handling

```swift
do {
    let transformer = Transform(options)
    let result = try transformer.transform(code)
    print("Success: \(result)")
} catch TransformError.transformationFailed(let message) {
    print("Transform failed: \(message)")
} catch TransformError.invalidOptions(let message) {
    print("Invalid options: \(message)")
} catch TransformError.internalError(let message) {
    print("Internal error: \(message)")
} catch {
    print("Unexpected error: \(error)")
}
```

## Available Configuration Options

### JSX Options
- `jsxFactory: String` - JSX factory function (default: "React.createElement")
- `jsxFragment: String` - JSX fragment function (default: "React.Fragment")
- `jsx: JSXMode` - JSX transformation mode (.transform, .preserve, .automatic)
- `jsxDev: Bool` - Enable JSX development mode
- `jsxSideEffects: Bool` - Whether JSX has side effects
- `jsxImportSource: String` - Import source for automatic JSX runtime

### Output Options
- `format: OutputFormat` - Output format (.default, .iife, .cjs, .esm)
- `platform: Platform` - Target platform (.browser, .node, .neutral)
- `target: Target` - ECMAScript target (.es5, .es2015, .es2020, .esnext, etc.)
- `globalName: String` - Global name for IIFE format

### Minification Options
- `minifyWhitespace: Bool` - Remove unnecessary whitespace
- `minifyIdentifiers: Bool` - Shorten variable names
- `minifySyntax: Bool` - Use shorter syntax forms
- `drop: DropMode` - Drop console/debugger statements (.none, .console, .debugger)
- `treeShaking: Bool` - Enable dead code elimination

### Source Map Options
- `sourcemap: SourceMap` - Source map generation (.none, .inline, .linked, .external, .both)
- `sourcesContent: SourcesContent` - Include sources content (.include, .exclude)
- `sourceRoot: String` - Source root directory
- `sourcefile: String` - Source file name

### Advanced Options
- `loader: Loader` - File loader type (.js, .jsx, .ts, .tsx, .css, .json, etc.)
- `charset: Charset` - Output character set (.ascii, .utf8)
- `legalComments: LegalComments` - Legal comment handling
- `banner: String` - Text to prepend to output
- `footer: String` - Text to append to output
- `define: [String: String]` - Global replacements
- `keepNames: Bool` - Preserve original names
- `logLevel: LogLevel` - Logging verbosity

## API Reference

### Transform Class
- `Transform()` - Initialize with default options
- `Transform(TransformOptions)` - Initialize with custom options
- `transform(_:)` - Transform code and return result string
- `transformWithResult(_:)` - Transform code and return detailed results

### TransformOptions Initialization
- `TransformOptions()` - Create with defaults
- `TransformOptions(...)` - Create with all options as parameters

### Factory Methods
- `TransformOptions.react()` - React preset
- `TransformOptions.preact()` - Preact preset
- `TransformOptions.vue()` - Vue 3 preset
- `TransformOptions.production()` - Production preset (minified)
- `TransformOptions.development()` - Development preset
- `TransformOptions.custom(factory:fragment:)` - Custom JSX preset

### Chaining Methods
- `.jsx(factory:fragment:)` - Configure JSX factory and fragment
- `.minify()` - Enable all minification options
- `.production()` - Configure for production
- `.development()` - Configure for development

### String Extensions
- `String.transform()` - Transform string with defaults
- `String.transform(with:)` - Transform string with options

## Performance Tips

1. **Reuse Transformers**: Create a `Transform` instance once and reuse it for multiple transforms with the same configuration.

```swift
let transformer = Transform(options)
let results = try codeFiles.map { try transformer.transform($0) }
```

2. **Batch Similar Transforms**: Group files with similar transform requirements together.

```swift
let jsxTransformer = Transform(TransformOptions.react())
let tsxTransformer = Transform(TransformOptions.reactTypeScript())

// Transform all JSX files
let jsxResults = try jsxFiles.map { try jsxTransformer.transform($0) }

// Transform all TSX files
let tsxResults = try tsxFiles.map { try tsxTransformer.transform($0) }
```

## Requirements

- iOS 12.0+ / macOS 10.14+
- Swift 5.3+
- Xcode 12.0+

## Architecture

This package uses a sophisticated architecture to provide an idiomatic Swift API:

1. **Go Layer**: ESBuild's native Go API with comprehensive option support
2. **Bridge Layer**: `gomobile` generated bindings with custom Configure methods for every option
3. **Swift Layer**: Idiomatic computed properties that call the underlying Configure methods
4. **Helper Layer**: Convenient presets, method chaining, and String extensions

The result is a completely native Swift experience that hides all the `gomobile` complexity while providing access to every ESBuild feature.

## License

This project is licensed under the MIT License - see the LICENSE file for details.