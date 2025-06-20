import ESBuildMobile
import Foundation

// MARK: - Plugin Class

/// Represents an ESBuild plugin
public class Plugin {
    private let internalPlugin: EsbuildmobilePlugin

    public init(name: String) {
        guard let plugin = EsbuildmobileNewPlugin(name) else {
            fatalError("Failed to create plugin")
        }
        internalPlugin = plugin
    }

    init(_ goPlugin: EsbuildmobilePlugin) {
        internalPlugin = goPlugin
    }

    public var name: String {
        get { internalPlugin.getName() }
        set { internalPlugin.setPluginName(newValue) }
    }

    var _internal: EsbuildmobilePlugin { internalPlugin }

    // MARK: - OnStart callback

    /// Adds a simple onStart callback that prints a message
    public func addOnStartPrint(_ message: String) {
        guard let callback = EsbuildmobileNewSimpleOnStartCallback(message) else {
            fatalError("Failed to create onStart callback")
        }
        internalPlugin.onStart(callback)
    }

    // MARK: - OnResolve callbacks

    /// Adds an onResolve rule with the given filter and result
    public func addOnResolve(
        filter: String, namespace: String? = nil, path: String, toNamespace: String? = nil,
        external: Bool = false
    ) {
        guard let options = EsbuildmobileNewOnResolveOptions() else { return }
        options.setResolveFilter(filter)
        if let ns = namespace {
            options.setResolveNamespace(ns)
        }

        let callback = EsbuildmobileSimpleResolveCallback()
        callback.path = path
        if let ns = toNamespace {
            callback.namespace = ns
        }
        callback.external = external

        internalPlugin.onResolve(options, callback: callback)
    }

    // MARK: - OnLoad callbacks

    /// Adds an onLoad rule with the given filter and content
    public func addOnLoad(
        filter: String, namespace: String? = nil, contents: String, loader: Int
    ) {
        let options = EsbuildmobileNewOnLoadOptions()
        options?.setLoadFilter(filter)
        if let ns = namespace {
            options?.setLoadNamespace(ns)
        }

        let callback = EsbuildmobileSimpleLoadCallback()
        callback.contents = contents
        callback.loader = loader

        internalPlugin.onLoad(options, callback: callback)
    }
}

// MARK: - Filter Constants

public enum PluginFilters {
    /// Common filter patterns
    public static let react = "^react$"
    public static let reactDOM = "^react-dom$"
    public static let lodash = "^lodash$"
    public static let jquery = "^jquery$"
    public static let nodeModules = "^[^./]+"
    public static let relativeImports = "^\\."
    public static let absoluteImports = "^/"
    public static let allFiles = ".*"
    public static let jsFiles = "\\.js$"
    public static let tsFiles = "\\.ts$"
    public static let cssFiles = "\\.css$"
    public static let scssFiles = "\\.scss$"
    public static let lessFiles = "\\.less$"
    public static let jsonFiles = "\\.json$"
    public static let imageFiles = "\\.(png|jpg|jpeg|gif|svg)$"
    public static let fontFiles = "\\.(woff|woff2|eot|ttf|otf)$"
}

// MARK: - Namespace Constants

public enum PluginNamespaces {
    public static let file = "file"
    public static let http = "http"
    public static let https = "https"
    public static let virtual = "virtual"
    public static let generated = "generated"
    public static let transformed = "transformed"
    public static let external = "external"
}

// MARK: - Plugin Builders

public extension Plugin {
    /// Creates a plugin that transforms React imports to use a global variable
    static func reactGlobalTransform(globalName: String = "_FLICKCORE_$REACT") -> Plugin {
        let plugin = Plugin(name: "react-global-transform")

        // Transform react imports
        plugin.addOnResolve(
            filter: PluginFilters.react,
            path: "react",
            toNamespace: "use-flick-react-global"
        )

        // Load the transformed module
        plugin.addOnLoad(
            filter: PluginFilters.allFiles,
            namespace: "use-flick-react-global",
            contents: "module.exports = \(globalName)",
            loader: Int(EsbuildmobileGetLoaderJS())
        )

        return plugin
    }

    /// Creates a plugin that marks all node_modules as external
    static func externalizeNodeModules() -> Plugin {
        let plugin = Plugin(name: "externalize-node-modules")

        plugin.addOnResolve(
            filter: PluginFilters.nodeModules,
            path: "", // Empty path means keep original
            external: true
        )

        return plugin
    }
}

// MARK: - Example Usage

/*
 Example usage:

 // Create a basic plugin
 let myPlugin = Plugin(name: "my-plugin")
 myPlugin.addOnStartPrint("Plugin started!")

 // Use pre-built plugins
 let reactPlugin = Plugin.reactGlobalTransform()

 // Add plugins to build options
 let options = BuildOptions()
 options.addPlugin(myPlugin)
 options.addPlugin(reactPlugin)

 // Or set multiple plugins at once
 options.plugins = [myPlugin, reactPlugin]

 // Build with plugins
 let result = try "console.log('Hello, world!')".buildJS(with: options)
 */
