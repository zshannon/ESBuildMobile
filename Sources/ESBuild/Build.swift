import ESBuildMobile
import Foundation

/// A Swift-native wrapper around BuildOptions from the Go layer
public class BuildOptions {
    private let internalOptions: EsbuildmobileBuildOptions

    /// Create build options with defaults
    public init() {
        guard let opts = EsbuildmobileNewBuildOptions() else {
            fatalError("Failed to create build options")
        }
        internalOptions = opts
    }

    public init(
        bundle: Bool = false,
        outfile: String? = nil,
        platform: Platform = .browser,
        format: OutputFormat = .esm,
        target: Target = .es2020,
        sourcemap: SourceMap = .none,
        jsxFactory: String = "React.createElement",
        jsxFragment: String = "React.Fragment",
        minifySyntax: Bool = false,
        minifyWhitespace: Bool = false,
        minifyIdentifiers: Bool = false,
        globalName: String? = nil,
        logLevel: LogLevel = .info,
        logLimit: Int = 5,
        loader: Loader? = nil,
        sourcefile: String? = nil,
        define: [String: String] = [:]
    ) {
        guard let opts = EsbuildmobileNewBuildOptions() else {
            fatalError("Failed to create build options")
        }
        internalOptions = opts
        self.bundle = bundle
        if let outfile = outfile {
            self.outfile = outfile
        }
        self.platform = platform
        self.format = format
        self.jsxFactory = jsxFactory
        self.jsxFragment = jsxFragment
        self.target = target
        self.sourcemap = sourcemap
        self.minifySyntax = minifySyntax
        self.minifyWhitespace = minifyWhitespace
        self.minifyIdentifiers = minifyIdentifiers
        if let loader = loader {
            self.loader = loader
        }
        if let sourcefile = sourcefile {
            self.sourcefile = sourcefile
        }
        if let globalName = globalName {
            self.globalName = globalName
        }
        self.logLevel = logLevel
        self.logLimit = logLimit
        self.define = define
    }

    // MARK: - Basic options

    public var bundle: Bool {
        get { internalOptions.bundle }
        set { internalOptions.configureBundle(newValue) }
    }

    public var outfile: String {
        get { internalOptions.outfile }
        set { internalOptions.configureOutfile(newValue) }
    }

    public var platform: Platform {
        get { Platform(fromGoValue: internalOptions.getPlatform()) ?? .browser }
        set { internalOptions.configurePlatform(by: newValue.rawValue) }
    }

    public var format: OutputFormat {
        get { OutputFormat(fromGoValue: internalOptions.getFormat()) ?? .default }
        set { internalOptions.configureFormat(by: newValue.rawValue) }
    }

    public var target: Target {
        get { Target(fromGoValue: internalOptions.getTarget()) ?? .esnext }
        set { internalOptions.configureTarget(by: newValue.rawValue) }
    }

    public var sourcemap: SourceMap {
        get { SourceMap(fromGoValue: internalOptions.getSourcemap()) ?? .none }
        set { internalOptions.configureSourcemap(by: newValue.rawValue) }
    }

    public var minifyWhitespace: Bool {
        get { internalOptions.minifyWhitespace }
        set { internalOptions.configureMinifyWhitespace(newValue) }
    }

    public var minifyIdentifiers: Bool {
        get { internalOptions.minifyIdentifiers }
        set { internalOptions.configureMinifyIdentifiers(newValue) }
    }

    public var minifySyntax: Bool {
        get { internalOptions.minifySyntax }
        set { internalOptions.configureMinifySyntax(newValue) }
    }

    public var globalName: String {
        get { internalOptions.globalName }
        set { internalOptions.configureGlobalName(newValue) }
    }

    public var logLevel: LogLevel {
        get { LogLevel(fromGoValue: internalOptions.getLogLevel()) ?? .info }
        set { internalOptions.configureLogLevel(by: newValue.rawValue) }
    }

    public var logLimit: Int {
        get { Int(internalOptions.logLimit) }
        set { internalOptions.configureLogLimit(newValue) }
    }

    public var loader: Loader {
        get { Loader(fromGoValue: internalOptions.getLoaderSingle()) ?? .js }
        set { internalOptions.configureLoader(by: newValue.rawValue) }
    }

    public var sourcefile: String {
        get { internalOptions.sourcefile }
        set { internalOptions.configureSourcefile(newValue) }
    }

    // MARK: - Additional Build Options

    public var write: Bool {
        get { internalOptions.write }
        set { internalOptions.configureWrite(newValue) }
    }

    public var outdir: String {
        get { internalOptions.outdir }
        set { internalOptions.configureOutdir(newValue) }
    }

    public var outbase: String {
        get { internalOptions.outbase }
        set { internalOptions.configureOutbase(newValue) }
    }

    public var absWorkingDir: String {
        get { internalOptions.absWorkingDir }
        set { internalOptions.configureAbsWorkingDir(newValue) }
    }

    public var publicPath: String {
        get { internalOptions.publicPath }
        set { internalOptions.configurePublicPath(newValue) }
    }

    public var sourceRoot: String {
        get { internalOptions.sourceRoot }
        set { internalOptions.configureSourceRoot(newValue) }
    }

    public var sourcesContent: SourcesContent {
        get { SourcesContent(fromGoValue: internalOptions.getSourcesContent()) ?? .include }
        set { internalOptions.setSourcesContentBy(newValue.goValue) }
    }

    public var color: ColorMode {
        get { ColorMode(fromGoValue: internalOptions.getColor()) ?? .auto }
        set { internalOptions.setColorBy(newValue.goValue) }
    }

    // MARK: - Advanced Options

    public var preserveSymlinks: Bool {
        get { internalOptions.preserveSymlinks }
        set { internalOptions.configurePreserveSymlinks(newValue) }
    }

    public var splitting: Bool {
        get { internalOptions.splitting }
        set { internalOptions.configureSplitting(newValue) }
    }

    public var metafile: Bool {
        get { internalOptions.metafile }
        set { internalOptions.configureMetafile(newValue) }
    }

    public var allowOverwrite: Bool {
        get { internalOptions.allowOverwrite }
        set { internalOptions.configureAllowOverwrite(newValue) }
    }

    public var packages: Packages {
        get { Packages(fromGoValue: internalOptions.getPackages()) ?? .default }
        set { internalOptions.setPackagesBy(newValue.goValue) }
    }

    public var treeShaking: TreeShaking {
        get { TreeShaking(fromGoValue: internalOptions.getTreeShaking()) ?? .default }
        set { internalOptions.setTreeShakingBy(newValue.goValue) }
    }

    public var ignoreAnnotations: Bool {
        get { internalOptions.ignoreAnnotations }
        set { internalOptions.configureIgnoreAnnotations(newValue) }
    }

    public var keepNames: Bool {
        get { internalOptions.keepNames }
        set { internalOptions.configureKeepNames(newValue) }
    }

    public var lineLimit: Int {
        get { Int(internalOptions.lineLimit) }
        set { internalOptions.configureLineLimit(newValue) }
    }

    public var charset: Charset {
        get { Charset(fromGoValue: internalOptions.getCharset()) ?? .utf8 }
        set { internalOptions.setCharsetBy(newValue.goValue) }
    }

    public var legalComments: LegalComments {
        get { LegalComments(fromGoValue: internalOptions.getLegalComments()) ?? .default }
        set { internalOptions.setLegalCommentsBy(newValue.goValue) }
    }

    // MARK: - JSX Options

    public var jsx: JSXMode {
        get { JSXMode(fromGoValue: internalOptions.getJSX()) ?? .transform }
        set { internalOptions.setJSXBy(newValue.goValue) }
    }

    public var jsxFactory: String {
        get { internalOptions.jsxFactory }
        set { internalOptions.configureJSXFactory(newValue) }
    }

    public var jsxFragment: String {
        get { internalOptions.jsxFragment }
        set { internalOptions.configureJSXFragment(newValue) }
    }

    public var jsxImportSource: String {
        get { internalOptions.jsxImportSource }
        set { internalOptions.configureJSXImportSource(newValue) }
    }

    public var jsxDev: Bool {
        get { internalOptions.jsxDev }
        set { internalOptions.configureJSXDev(newValue) }
    }

    public var jsxSideEffects: Bool {
        get { internalOptions.jsxSideEffects }
        set { internalOptions.configureJSXSideEffects(newValue) }
    }

    // MARK: - Output Naming

    public var entryNames: String {
        get { internalOptions.entryNames }
        set { internalOptions.configureEntryNames(newValue) }
    }

    public var chunkNames: String {
        get { internalOptions.chunkNames }
        set { internalOptions.configureChunkNames(newValue) }
    }

    public var assetNames: String {
        get { internalOptions.assetNames }
        set { internalOptions.configureAssetNames(newValue) }
    }

    // MARK: - TypeScript

    public var tsconfig: String {
        get { internalOptions.tsconfig }
        set { internalOptions.configureTsconfig(newValue) }
    }

    public var tsconfigRaw: String {
        get { internalOptions.tsconfigRaw }
        set { internalOptions.configureTsconfigRaw(newValue) }
    }

    // MARK: - Mangling

    public var mangleProps: String {
        get { internalOptions.mangleProps }
        set { internalOptions.configureMangleProps(newValue) }
    }

    public var reserveProps: String {
        get { internalOptions.reserveProps }
        set { internalOptions.configureReserveProps(newValue) }
    }

    public var mangleQuoted: MangleQuoted {
        get { MangleQuoted(fromGoValue: internalOptions.getMangleQuoted()) ?? .false }
        set { internalOptions.setMangleQuotedBy(newValue.goValue) }
    }

    // MARK: - Collections

    private var defineMap: [String: String] = [:]
    public var define: [String: String] {
        get { defineMap }
        set {
            defineMap = newValue
            for (k, v) in newValue {
                internalOptions.configureDefineEntry(k, value: v)
            }
        }
    }

    public var entryPoints: [String] {
        get {
            let count = internalOptions.getEntryPointsCount()
            return (0 ..< count).compactMap { internalOptions.getEntryPoint($0) }
        }
        set {
            internalOptions.clearEntryPoints()
            for entry in newValue {
                internalOptions.addEntryPoint(entry)
            }
        }
    }

    public var external: [String] {
        get {
            let count = internalOptions.getExternalCount()
            return (0 ..< count).compactMap { internalOptions.getExternal($0) }
        }
        set {
            internalOptions.clearExternal()
            for ext in newValue {
                internalOptions.addExternal(ext)
            }
        }
    }

    public var inject: [String] {
        get {
            let count = internalOptions.getInjectCount()
            return (0 ..< count).compactMap { internalOptions.getInject($0) }
        }
        set {
            internalOptions.clearInject()
            for inj in newValue {
                internalOptions.addInject(inj)
            }
        }
    }

    public var pure: [String] {
        get {
            let count = internalOptions.getPureCount()
            return (0 ..< count).compactMap { internalOptions.getPure($0) }
        }
        set {
            internalOptions.clearPure()
            for p in newValue {
                internalOptions.addPure(p)
            }
        }
    }

    // MARK: - Plugin Support

    public var plugins: [Plugin] {
        get {
            let count = internalOptions.getPluginsCount()
            return (0 ..< count).compactMap { index in
                guard let plugin = internalOptions.getPlugin(index) else { return nil }
                return Plugin(plugin)
            }
        }
        set {
            internalOptions.clearPlugins()
            for plugin in newValue {
                internalOptions.add(plugin._internal)
            }
        }
    }

    /// Add a plugin to the build options
    public func addPlugin(_ plugin: Plugin) {
        internalOptions.add(plugin._internal)
    }

    /// Remove all plugins
    public func clearPlugins() {
        internalOptions.clearPlugins()
    }

    var _internal: EsbuildmobileBuildOptions { internalOptions }
}

// MARK: - Build interface

public enum BuildError: Error, LocalizedError {
    case buildFailed(String)

    public var errorDescription: String? {
        switch self {
        case let .buildFailed(msg):
            return "ESBuild build failed: \(msg)"
        }
    }
}

public struct Builder {
    private let options: BuildOptions

    public init(_ options: BuildOptions? = nil) {
        self.options = options ?? BuildOptions()
    }

    public func build(_ code: String) throws -> String {
        var error: NSError?
        let result = EsbuildmobileBuild(code, options._internal, &error)
        if let err = error {
            throw BuildError.buildFailed(err.localizedDescription)
        }
        return result
    }
}

public extension String {
    func buildJS(with options: BuildOptions = BuildOptions()) throws -> String {
        return try Builder(options).build(self)
    }
}
