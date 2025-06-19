# ESBuild

A Swift package that provides ESBuild JSX transformation capabilities for mobile development. This package uses `gomobile` to bridge ESBuild's Go API with native Swift, providing an idiomatic Swift API that hides all the underlying complexity.

## Features

- ✅ **Idiomatic Swift API** - Native Swift properties and method chaining
- ✅ **Complete ESBuild Support** - All transform options available through Swift
- ✅ **Type-Safe Configuration** - Swift enums for all ESBuild options
- ✅ **Framework Presets** - Built-in configurations for React, Preact, Vue
- ✅ **Method Chaining** - Fluent API for easy configuration
- ✅ **String Extensions** - Transform JSX directly on String instances
- ✅ **Comprehensive Error Handling** - Swift-native error types

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

### Basic JSX Transformation

```swift
import ESBuild

// Transform JSX with default React settings
let jsx = "<div>Hello World</div>"
let result = try JSXTransformer.transform(jsx)
print(result)
// Output: /* @__PURE__ */ React.createElement("div", null, "Hello World");

// Or use the String extension
let result = try jsx.transformJSX()
```

### Using Presets

```swift
// React (default)
let result = try jsx.transformJSX(with: .react())

// Preact
let result = try jsx.transformJSX(with: .preact())
// Output: /* @__PURE__ */ h("div", null, "Hello World");

// Vue 3
let result = try jsx.transformJSX(with: .vue())

// Production build (minified)
let result = try jsx.transformJSX(with: .production())
```

## Idiomatic Swift API

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

let result = try JSXTransformer.transform(jsx, options: options)
```

### Method Chaining

```swift
let result = try jsx.transformJSX(with:
    TransformOptions()
        .jsx(factory: "createElement", fragment: "Fragment")
        .minify()
        .development()
)
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
```

## Framework-Specific Examples

### React

```swift
let options = TransformOptions.react()
options.jsxDev = true  // For development
options.target = .es2020
options.format = .esm

let result = try transformJSX("""
    function Welcome({name}) {
        return <div>Hello, {name}!</div>;
    }
""", with: options)
```

### React with TypeScript

```swift
let options = TransformOptions.reactTypeScript()
options.loader = .tsx
options.keepNames = true

let result = try transformJSX("""
    interface Props {
        name: string;
    }

    const Welcome: React.FC<Props> = ({name}) => {
        return <div>Hello, {name}!</div>;
    };
""", with: options)
```

### Preact

```swift
let options = TransformOptions.preact()
options.format = .esm
options.platform = .browser

let result = try transformJSX("""
    import { h } from 'preact';

    export function App() {
        return <div>Hello from Preact!</div>;
    }
""", with: options)
```

### Vue 3 JSX

```swift
let options = TransformOptions.vue()
options.target = .es2020
options.format = .esm

let result = try transformJSX("""
    export default {
        render() {
            return <div>Hello from Vue JSX!</div>;
        }
    };
""", with: options)
```

## Advanced Configuration

### Custom Build Pipeline

```swift
let options = TransformOptions()

// Configure for modern browsers
options.target = .es2020
options.platform = .browser
options.format = .esm

// Enable all optimizations
options.minify()
options.treeShaking = true
options.drop = .console

// Configure source maps for debugging
options.sourcemap = .inline
options.sourcesContent = .include

// Set up defines for environment
options.define = [
    "process.env.NODE_ENV": "production",
    "__DEV__": "false"
]

let result = try JSXTransformer.transform(complexJSX, options: options)
```

### Multiple File Types

```swift
// TypeScript JSX
let tsxOptions = TransformOptions()
tsxOptions.loader = .tsx
tsxOptions.target = .es2020

// Regular JavaScript
let jsOptions = TransformOptions()
jsOptions.loader = .js
jsOptions.format = .cjs

// CSS-in-JS
let cssOptions = TransformOptions()
cssOptions.loader = .css
```

## Error Handling

```swift
do {
    let result = try JSXTransformer.transform(jsx, options: options)
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

## Method Reference

### Factory Methods
- `TransformOptions()` - Create with defaults
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

### Transform Methods
- `JSXTransformer.transform(_:options:)` - Transform with options
- `JSXTransformer.transformWithResult(_:options:)` - Get detailed results
- `String.transformJSX()` - Transform string with defaults
- `String.transformJSX(with:)` - Transform string with options

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
