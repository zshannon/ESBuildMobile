import ESBuildMobile
import Foundation

/// A Swift-native wrapper for ESBuild transform options with idiomatic property access
public class TransformOptions {
    private let internalOptions: EsbuildmobileTransformOptions

    /// Initialize with default options
    public init() {
        guard let options = EsbuildmobileNewTransformOptions() else {
            fatalError("Failed to create ESBuild transform options")
        }
        internalOptions = options
    }

    /// Initialize with custom options
    public init(
        jsxFactory: String = "React.createElement",
        jsxFragment: String = "React.Fragment",
        jsxImportSource: String = "",
        jsxDev: Bool = false,
        jsxSideEffects: Bool = false,
        jsx: JSXMode = .transform,
        minifyWhitespace: Bool = false,
        minifyIdentifiers: Bool = false,
        minifySyntax: Bool = false,
        globalName: String = "",
        format: OutputFormat = .default,
        platform: Platform = .browser,
        target: Target = .esnext,
        sourceRoot: String = "",
        sourcefile: String = "",
        sourcemap: SourceMap = .none,
        sourcesContent: SourcesContent = .include,
        keepNames: Bool = false,
        logLevel: LogLevel = .info,
        logLimit: Int = 0,
        tsconfigRaw: String = "",
        banner: String = "",
        footer: String = "",
        lineLimit: Int = 80,
        ignoreAnnotations: Bool = false,
        mangleProps: String = "",
        reserveProps: String = "",
        mangleQuoted: Bool = false,
        color: ColorMode = .auto,
        charset: Charset = .utf8,
        treeShaking: Bool = false,
        legalComments: LegalComments = .default,
        drop: DropMode = .none,
        loader: Loader = .js,
        define: [String: String] = [:]
    ) {
        guard let options = EsbuildmobileNewTransformOptions() else {
            fatalError("Failed to create ESBuild transform options")
        }
        internalOptions = options

        // Configure all properties
        self.jsxFactory = jsxFactory
        self.jsxFragment = jsxFragment
        self.jsxImportSource = jsxImportSource
        self.jsxDev = jsxDev
        self.jsxSideEffects = jsxSideEffects
        self.jsx = jsx
        self.minifyWhitespace = minifyWhitespace
        self.minifyIdentifiers = minifyIdentifiers
        self.minifySyntax = minifySyntax
        self.globalName = globalName
        self.format = format
        self.platform = platform
        self.target = target
        self.sourceRoot = sourceRoot
        self.sourcefile = sourcefile
        self.sourcemap = sourcemap
        self.sourcesContent = sourcesContent
        self.keepNames = keepNames
        self.logLevel = logLevel
        self.logLimit = logLimit
        self.tsconfigRaw = tsconfigRaw
        self.banner = banner
        self.footer = footer
        self.lineLimit = lineLimit
        self.ignoreAnnotations = ignoreAnnotations
        self.mangleProps = mangleProps
        self.reserveProps = reserveProps
        self.mangleQuoted = mangleQuoted
        self.color = color
        self.charset = charset
        self.treeShaking = treeShaking
        self.legalComments = legalComments
        self.drop = drop
        self.loader = loader
        self.define = define
    }

    // MARK: - JSX Configuration

    /// The JSX factory function name (e.g., "React.createElement", "h")
    public var jsxFactory: String {
        get { internalOptions.jsxFactory }
        set { internalOptions.configureJSXFactory(newValue) }
    }

    /// The JSX fragment function name (e.g., "React.Fragment", "Fragment")
    public var jsxFragment: String {
        get { internalOptions.jsxFragment }
        set { internalOptions.configureJSXFragment(newValue) }
    }

    /// The JSX import source for automatic JSX runtime
    public var jsxImportSource: String {
        get { internalOptions.jsxImportSource }
        set { internalOptions.configureJSXImportSource(newValue) }
    }

    /// Whether to use JSX development mode
    public var jsxDev: Bool {
        get { internalOptions.jsxDev }
        set { internalOptions.configureJSXDev(newValue) }
    }

    /// Whether JSX has side effects
    public var jsxSideEffects: Bool {
        get { internalOptions.jsxSideEffects }
        set { internalOptions.configureJSXSideEffects(newValue) }
    }

    /// JSX transformation mode
    public var jsx: JSXMode {
        get { .transform } // Default fallback since we can't read the enum
        set { internalOptions.configureJSX(by: newValue.rawValue) }
    }

    // MARK: - Minification Options

    /// Whether to minify whitespace
    public var minifyWhitespace: Bool {
        get { internalOptions.minifyWhitespace }
        set { internalOptions.configureMinifyWhitespace(newValue) }
    }

    /// Whether to minify identifiers
    public var minifyIdentifiers: Bool {
        get { internalOptions.minifyIdentifiers }
        set { internalOptions.configureMinifyIdentifiers(newValue) }
    }

    /// Whether to minify syntax
    public var minifySyntax: Bool {
        get { internalOptions.minifySyntax }
        set { internalOptions.configureMinifySyntax(newValue) }
    }

    // MARK: - Output Configuration

    /// Global name for UMD bundles
    public var globalName: String {
        get { internalOptions.globalName }
        set { internalOptions.configureGlobalName(newValue) }
    }

    /// Output format
    public var format: OutputFormat {
        get { .default } // Default fallback since we can't read the enum
        set { internalOptions.configureFormat(by: newValue.rawValue) }
    }

    /// Target platform
    public var platform: Platform {
        get { .browser } // Default fallback since we can't read the enum
        set { internalOptions.configurePlatform(by: newValue.rawValue) }
    }

    /// Target ECMAScript version
    public var target: Target {
        get { .esnext } // Default fallback since we can't read the enum
        set { internalOptions.configureTarget(by: newValue.rawValue) }
    }

    // MARK: - Source Configuration

    /// Source root directory
    public var sourceRoot: String {
        get { internalOptions.sourceRoot }
        set { internalOptions.configureSourceRoot(newValue) }
    }

    /// Source file name for error reporting
    public var sourcefile: String {
        get { internalOptions.sourcefile }
        set { internalOptions.configureSourcefile(newValue) }
    }

    /// Sourcemap generation mode
    public var sourcemap: SourceMap {
        get { .none } // Default fallback since we can't read the enum
        set { internalOptions.configureSourcemap(by: newValue.rawValue) }
    }

    /// Sources content inclusion
    public var sourcesContent: SourcesContent {
        get { .include } // Default fallback since we can't read the enum
        set { internalOptions.configureSourcesContent(by: newValue.rawValue) }
    }

    // MARK: - Advanced Options

    /// Whether to keep original names
    public var keepNames: Bool {
        get { internalOptions.keepNames }
        set { internalOptions.configureKeepNames(newValue) }
    }

    /// Log level for error reporting
    public var logLevel: LogLevel {
        get { .info } // Default fallback since we can't read the enum
        set { internalOptions.configureLogLevel(by: newValue.rawValue) }
    }

    /// Log limit for error reporting
    public var logLimit: Int {
        get { Int(internalOptions.logLimit) }
        set { internalOptions.configureLogLimit(newValue) }
    }

    /// Raw tsconfig.json content
    public var tsconfigRaw: String {
        get { internalOptions.tsconfigRaw }
        set { internalOptions.configureTsconfigRaw(newValue) }
    }

    /// Banner text to prepend to output
    public var banner: String {
        get { internalOptions.banner }
        set { internalOptions.configureBanner(newValue) }
    }

    /// Footer text to append to output
    public var footer: String {
        get { internalOptions.footer }
        set { internalOptions.configureFooter(newValue) }
    }

    /// Line limit for output formatting
    public var lineLimit: Int {
        get { Int(internalOptions.lineLimit) }
        set { internalOptions.configureLineLimit(newValue) }
    }

    /// Whether to ignore annotations
    public var ignoreAnnotations: Bool {
        get { internalOptions.ignoreAnnotations }
        set { internalOptions.configureIgnoreAnnotations(newValue) }
    }

    /// Mangle properties pattern
    public var mangleProps: String {
        get { internalOptions.mangleProps }
        set { internalOptions.configureMangleProps(newValue) }
    }

    /// Reserve properties pattern
    public var reserveProps: String {
        get { internalOptions.reserveProps }
        set { internalOptions.configureReserveProps(newValue) }
    }

    /// Mangle quoted properties
    public var mangleQuoted: Bool {
        get { false } // Default fallback since we can't read the enum
        set { internalOptions.configureMangleQuoted(by: newValue ? "true" : "false") }
    }

    /// Color output mode
    public var color: ColorMode {
        get { .auto } // Default fallback since we can't read the enum
        set { internalOptions.configureColor(by: newValue.rawValue) }
    }

    /// Character set for output
    public var charset: Charset {
        get { .utf8 } // Default fallback since we can't read the enum
        set { internalOptions.configureCharset(by: newValue.rawValue) }
    }

    /// Tree shaking mode
    public var treeShaking: Bool {
        get { false } // Default fallback since we can't read the enum
        set { internalOptions.configureTreeShaking(by: newValue ? "true" : "false") }
    }

    /// Legal comments handling
    public var legalComments: LegalComments {
        get { .default } // Default fallback since we can't read the enum
        set { internalOptions.configureLegalComments(by: newValue.rawValue) }
    }

    /// Drop mode for console/debugger statements
    public var drop: DropMode {
        get { .none } // Default fallback since we can't read the enum
        set { internalOptions.configureDrop(by: newValue.rawValue) }
    }

    /// File loader type
    public var loader: Loader {
        get { .js } // Default fallback since we can't read the enum
        set { internalOptions.configureLoader(by: newValue.rawValue) }
    }

    // MARK: - Map-based properties (using helper methods)

    private var defineMap: [String: String] = [:]

    /// Define replacement values
    public var define: [String: String] {
        get { defineMap }
        set {
            defineMap = newValue
            for (key, value) in newValue {
                internalOptions.configureDefineEntry(key, value: value)
            }
        }
    }

    /// Add a single define entry
    public func setDefine(_ key: String, value: String) {
        defineMap[key] = value
        internalOptions.configureDefineEntry(key, value: value)
    }

    // MARK: - Method Chaining Support

    /// Configure JSX factory and fragment in one call
    @discardableResult
    public func jsx(factory: String, fragment: String) -> TransformOptions {
        jsxFactory = factory
        jsxFragment = fragment
        return self
    }

    /// Enable all minification options
    @discardableResult
    public func minify() -> TransformOptions {
        minifyWhitespace = true
        minifyIdentifiers = true
        minifySyntax = true
        return self
    }

    /// Configure for production builds
    @discardableResult
    public func production() -> TransformOptions {
        return minify()
    }

    /// Configure for development builds
    @discardableResult
    public func development() -> TransformOptions {
        jsxDev = true
        keepNames = true
        return self
    }

    // MARK: - Internal access

    var _internal: EsbuildmobileTransformOptions {
        return internalOptions
    }
}

// MARK: - Enum Types

public enum JSXMode: String, CaseIterable {
    case transform
    case preserve
    case automatic
}

public enum OutputFormat: String, CaseIterable {
    case `default` = ""
    case iife
    case cjs
    case esm
}

public enum Platform: String, CaseIterable {
    case browser
    case node
    case neutral
}

public enum Target: String, CaseIterable {
    case esnext
    case es2020
    case es2019
    case es2018
    case es2017
    case es2016
    case es2015
    case es6
    case es5
}

public enum SourceMap: String, CaseIterable {
    case none
    case inline
    case linked
    case external
    case both
}

public enum SourcesContent: String, CaseIterable {
    case include
    case exclude
}

public enum LogLevel: String, CaseIterable {
    case verbose
    case debug
    case info
    case warning
    case error
    case silent
}

public enum ColorMode: String, CaseIterable {
    case auto
    case always
    case never
}

public enum Charset: String, CaseIterable {
    case ascii
    case utf8
}

public enum LegalComments: String, CaseIterable {
    case `default` = ""
    case none
    case inline
    case eof
    case linked
    case external
}

public enum DropMode: String, CaseIterable {
    case none = ""
    case console
    case debugger
}

public enum Loader: String, CaseIterable {
    case js
    case jsx
    case ts
    case tsx
    case css
    case json
    case text
    case base64
    case dataurl
    case file
    case binary
}

// MARK: - Error Types

/// An error type representing possible failures during JSX transformation.
public enum TransformError: Error, LocalizedError {
    case transformationFailed(String)
    case invalidOptions(String)
    case internalError(String)

    public var errorDescription: String? {
        switch self {
        case let .transformationFailed(message):
            return "JSX transformation failed: \(message)"
        case let .invalidOptions(message):
            return "Invalid transform options: \(message)"
        case let .internalError(message):
            return "Internal error: \(message)"
        }
    }
}

/// Result of a JSX transformation
public struct TransformResult {
    /// The transformed JavaScript code
    public let code: String

    /// Any warnings generated during transformation
    public let warnings: [String]

    /// Creates a new transform result
    init(code: String, warnings: [String] = []) {
        self.code = code
        self.warnings = warnings
    }
}

// MARK: - Main Transformer

/// Main interface for JSX transformation
public struct Transform {
    private let options: TransformOptions

    /// Initialize with transform options
    /// - Parameter options: Transform options (uses default if nil)
    public init(_ options: TransformOptions? = nil) {
        self.options = options ?? TransformOptions()
    }

    /// Transforms code using the configured options
    /// - Parameter code: The code string to transform
    /// - Returns: The transformed code
    /// - Throws: TransformError if transformation fails
    public func transform(_ code: String) throws -> String {
        let result = try transformWithResult(code)
        return result.code
    }

    /// Transforms code and returns detailed results
    /// - Parameter code: The code string to transform
    /// - Returns: TransformResult containing code and any warnings
    /// - Throws: TransformError if transformation fails
    public func transformWithResult(_ code: String) throws -> TransformResult {
        // Perform the transformation using the internal options
        var error: NSError?
        let result = EsbuildmobileTransformJSX(code, options._internal, &error)

        if let error = error {
            throw TransformError.transformationFailed(error.localizedDescription)
        }

        return TransformResult(code: result)
    }
}

// MARK: - Convenience Extensions

public extension String {
    /// Transforms this string using default settings
    /// - Returns: The transformed code
    /// - Throws: TransformError if transformation fails
    func transform() throws -> String {
        let transformer = Transform()
        return try transformer.transform(self)
    }

    /// Transforms this string using the specified options
    /// - Parameter options: Transform options
    /// - Returns: The transformed code
    /// - Throws: TransformError if transformation fails
    func transform(with options: TransformOptions) throws -> String {
        let transformer = Transform(options)
        return try transformer.transform(self)
    }
}

// MARK: - Preset Factory Methods

public extension TransformOptions {
    /// Creates TransformOptions configured for React
    static func react() -> TransformOptions {
        let options = TransformOptions()
        options.jsxFactory = "React.createElement"
        options.jsxFragment = "React.Fragment"
        return options
    }

    /// Creates TransformOptions configured for Preact
    static func preact() -> TransformOptions {
        let options = TransformOptions()
        options.jsxFactory = "h"
        options.jsxFragment = "Fragment"
        return options
    }

    /// Creates TransformOptions configured for Vue 3
    static func vue() -> TransformOptions {
        let options = TransformOptions()
        options.jsxFactory = "h"
        options.jsxFragment = "Fragment"
        return options
    }

    /// Creates TransformOptions with custom JSX configuration
    static func custom(factory: String, fragment: String) -> TransformOptions {
        let options = TransformOptions()
        options.jsxFactory = factory
        options.jsxFragment = fragment
        return options
    }

    /// Creates TransformOptions configured for production (minified)
    static func production() -> TransformOptions {
        return TransformOptions().production()
    }

    /// Creates TransformOptions configured for development
    static func development() -> TransformOptions {
        return TransformOptions().development()
    }

    /// Preset for React with TypeScript
    static func reactTypeScript() -> TransformOptions {
        let options = TransformOptions.react()
        options.loader = .tsx
        options.keepNames = true
        return options
    }
}

// MARK: - Usage Examples in Documentation

/*
 Example usage:

 // Basic transformation with React defaults
 let result = try JSXTransformer.transform("<div>Hello World</div>")

 // Or using String extension
 let result = try "<div>Hello</div>".transformJSX()

 // With custom options using idiomatic Swift properties
 let options = TransformOptions()
 options.jsxFactory = "h"
 options.jsxFragment = "Fragment"
 options.minifyWhitespace = true
 options.platform = .browser
 options.target = .es2020
 let result = try JSXTransformer.transform(jsx, options: options)

 // Using presets
 let result = try jsx.transformJSX(with: .preact().production())

 // Method chaining
 let options = TransformOptions()
     .jsx(factory: "createElement", fragment: "Fragment")
     .minify()
     .development()
 let result = try jsx.transformJSX(with: options)

 // Configure advanced options
 let options = TransformOptions()
 options.format = .esm
 options.platform = .node
 options.target = .es2018
 options.sourcemap = .inline
 options.loader = .tsx
 options.setDefine("process.env.NODE_ENV", value: "production")
 */
