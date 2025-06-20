import ESBuild
import ESBuildMobile

// MARK: - Basic Plugin Examples

func basicPluginExample() throws {
    // Example 1: Simple plugin with onStart callback
    let plugin = Plugin(name: "my-plugin")
    plugin.addOnStartPrint("Plugin is starting...")

    let options = BuildOptions(bundle: true)
    options.addPlugin(plugin)

    let result = try Builder(options).build("console.log('Hello, world!')")
    print(result)
}

// MARK: - React Global Transform

func reactGlobalTransformExample() throws {
    // Example 2: Transform React imports to use a global variable
    let jsx = """
        import React, { useState } from 'react'

        export default function Counter() {
            const [count, setCount] = useState(0)
            return (
                <div>
                    <h1>Count: {count}</h1>
                    <button onClick={() => setCount(count + 1)}>
                        Increment
                    </button>
                </div>
            )
        }
        """

    let options = BuildOptions(
        bundle: true,
        loader: .jsx
    )

    // Use the pre-built React global transform plugin
    let reactPlugin = Plugin.reactGlobalTransform()
    options.addPlugin(reactPlugin)

    let result = try Builder(options).build(jsx)
    // Result will have React imports transformed to use _FLICKCORE_$REACT global
}

// MARK: - Custom Global Variable

func customGlobalVariableExample() throws {
    // Example 3: Use a custom global variable name
    let jsx = """
        import React from 'react'
        const App = () => <div>Hello</div>
        """

    let options = BuildOptions(
        bundle: true,
        loader: .jsx
    )

    // Use window.MyReact instead of the default
    let reactPlugin = Plugin.reactGlobalTransform(globalName: "window.MyReact")
    options.addPlugin(reactPlugin)

    let result = try Builder(options).build(jsx)
    // React will be replaced with window.MyReact
}

// MARK: - Externalize Dependencies

func externalizeExample() throws {
    // Example 4: Mark node_modules as external
    let code = """
        import React from 'react'
        import lodash from 'lodash'
        import { format } from 'date-fns'
        import local from './local-module'

        console.log(React, lodash, format, local)
        """

    let options = BuildOptions(bundle: true)

    // This plugin marks all node_modules as external
    let externalPlugin = Plugin.externalizeNodeModules()
    options.addPlugin(externalPlugin)

    let result = try Builder(options).build(code)
    // node_modules imports will be marked as external
    // local imports will still be bundled
}

// MARK: - Custom Virtual Module

func virtualModuleExample() throws {
    // Example 5: Create a virtual module
    let plugin = Plugin(name: "virtual-module-plugin")

    // Resolve virtual imports to a custom namespace
    plugin.addOnResolve(
        filter: "^virtual:config$",
        path: "virtual-config",
        toNamespace: "virtual"
    )

    // Provide content for the virtual module
    plugin.addOnLoad(
        filter: ".*",
        namespace: "virtual",
        contents: """
            export const config = {
                apiUrl: 'https://api.example.com',
                version: '1.0.0',
                features: {
                    darkMode: true,
                    analytics: false
                }
            }
            """,
        loader: Int(EsbuildmobileGetLoaderJS())
    )

    let code = """
        import { config } from 'virtual:config'
        console.log('API URL:', config.apiUrl)
        console.log('Version:', config.version)
        """

    let options = BuildOptions(bundle: true)
    options.addPlugin(plugin)

    let result = try Builder(options).build(code)
    // The virtual module will be included in the bundle
}

// MARK: - Multiple Plugins

func multiplePluginsExample() throws {
    // Example 6: Use multiple plugins together
    let code = """
        import React from 'react'
        import lodash from 'lodash'
        import { config } from 'virtual:env'

        const App = () => {
            return <div>{config.appName}</div>
        }

        export default App
        """

    // Create a virtual environment plugin
    let envPlugin = Plugin(name: "env-plugin")
    envPlugin.addOnResolve(
        filter: "^virtual:env$",
        path: "env",
        toNamespace: "virtual-env"
    )
    envPlugin.addOnLoad(
        filter: ".*",
        namespace: "virtual-env",
        contents: "export const config = { appName: 'My App' }",
        loader: Int(EsbuildmobileGetLoaderJS())
    )

    let options = BuildOptions(
        bundle: true,
        loader: .jsx
    )

    // Add multiple plugins
    options.plugins = [
        Plugin.reactGlobalTransform(),
        Plugin.externalizeNodeModules(),
        envPlugin,
    ]

    let result = try Builder(options).build(code)
}

// MARK: - CSS Module Transform

func cssModuleExample() throws {
    // Example 7: Transform CSS module imports
    let plugin = Plugin(name: "css-modules")

    // Intercept .module.css imports
    plugin.addOnResolve(
        filter: "\\.module\\.css$",
        path: "css-module",
        toNamespace: "css-module"
    )

    // Return a JavaScript object for CSS modules
    plugin.addOnLoad(
        filter: ".*",
        namespace: "css-module",
        contents: """
            export default {
                container: 'container_abc123',
                title: 'title_def456',
                button: 'button_ghi789'
            }
            """,
        loader: Int(EsbuildmobileGetLoaderJS())
    )

    let code = """
        import styles from './styles.module.css'

        console.log(styles.container)
        console.log(styles.title)
        """

    let options = BuildOptions(bundle: true)
    options.addPlugin(plugin)

    let result = try Builder(options).build(code)
}

// MARK: - Path Alias Plugin

func pathAliasExample() throws {
    // Example 8: Create path aliases
    let plugin = Plugin(name: "path-alias")

    // Resolve @ imports to src directory
    plugin.addOnResolve(
        filter: "^@/",
        path: "src/",
        toNamespace: "file"
    )

    // Note: In a real implementation, you'd need to handle the actual file resolution
    // This is a simplified example

    let code = """
        import { utils } from '@/utils'
        import { Button } from '@/components/Button'

        console.log(utils, Button)
        """

    let options = BuildOptions(bundle: true)
    options.addPlugin(plugin)

    // This would resolve @/utils to src/utils, etc.
}

// MARK: - Build Event Callbacks

func buildEventExample() throws {
    // Example 9: Multiple plugins with different callbacks
    let startPlugin = Plugin(name: "start-plugin")
    startPlugin.addOnStartPrint("Build starting...")

    let resolvePlugin = Plugin(name: "resolve-logger")
    // This logs when specific modules are resolved
    resolvePlugin.addOnStartPrint("Resolve plugin initialized")

    let options = BuildOptions(bundle: true)
    options.plugins = [startPlugin, resolvePlugin]

    let result = try Builder(options).build("console.log('test')")
}

// MARK: - Environment Variables Plugin

func envVariablesExample() throws {
    // Example 10: Inject environment variables
    let envPlugin = Plugin(name: "env-vars")

    let envVars = [
        "NODE_ENV": "production",
        "API_KEY": "secret-key-123",
        "VERSION": "1.2.3",
    ]

    // Create virtual modules for each env var
    for (key, value) in envVars {
        envPlugin.addOnResolve(
            filter: "^process\\.env\\.\(key)$",
            path: "env-\(key)",
            toNamespace: "env-var"
        )

        envPlugin.addOnLoad(
            filter: "env-\(key)",
            namespace: "env-var",
            contents: "export default '\(value)'",
            loader: Int(EsbuildmobileGetLoaderJS())
        )
    }

    let code = """
        import NODE_ENV from 'process.env.NODE_ENV'
        import API_KEY from 'process.env.API_KEY'

        console.log('Environment:', NODE_ENV)
        console.log('API Key:', API_KEY)
        """

    let options = BuildOptions(bundle: true)
    options.addPlugin(envPlugin)

    let result = try Builder(options).build(code)
}

// MARK: - Usage with Builder Extensions

func builderExtensionExample() throws {
    // Example 11: Use plugins with string extensions
    let jsx = """
        import React from 'react'
        export default () => <h1>Hello World</h1>
        """

    // Create custom options with plugins
    let options = BuildOptions(
        bundle: true,
        loader: .jsx
    )
    options.addPlugin(Plugin.reactGlobalTransform())

    // Use the string extension with custom options
    let result = try jsx.buildJS(with: options)

    // Or use the Builder directly
    let builder = Builder(options)
    let result2 = try builder.build(jsx)
}
