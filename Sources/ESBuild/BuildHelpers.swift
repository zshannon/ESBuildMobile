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
        self.internalOptions = opts
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
        get { .browser }
        set { internalOptions.configurePlatform(by: newValue.rawValue) }
    }

    public var format: OutputFormat {
        get { .default }
        set { internalOptions.configureFormat(by: newValue.rawValue) }
    }

    public var target: Target {
        get { .esnext }
        set { internalOptions.configureTarget(by: newValue.rawValue) }
    }

    public var sourcemap: SourceMap {
        get { .none }
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
        get { .info }
        set { internalOptions.configureLogLevel(by: newValue.rawValue) }
    }

    public var logLimit: Int {
        get { Int(internalOptions.logLimit) }
        set { internalOptions.configureLogLimit(newValue) }
    }

    public var loader: Loader {
        get { .js }
        set { internalOptions.configureLoader(by: newValue.rawValue) }
    }

    public var sourcefile: String {
        get { internalOptions.sourcefile }
        set { internalOptions.configureSourcefile(newValue) }
    }

    private var defineMap: [String: String] = [:]
    public var define: [String: String] {
        get { defineMap }
        set {
            defineMap = newValue
            for (k, v) in newValue { internalOptions.configureDefineEntry(k, value: v) }
        }
    }

    internal var _internal: EsbuildmobileBuildOptions { internalOptions }
}

// MARK: - Build interface

public enum BuildError: Error, LocalizedError {
    case buildFailed(String)

    public var errorDescription: String? {
        switch self {
        case .buildFailed(let msg):
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

extension String {
    public func buildJS(with options: BuildOptions = BuildOptions()) throws -> String {
        return try Builder(options).build(self)
    }
}
