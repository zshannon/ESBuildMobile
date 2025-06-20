import ESBuildMobile
import XCTest

@testable import ESBuild

final class BuildTests: XCTestCase {
    func testBasic() throws {
        let jsx = """
            import { useEffect, useState } from 'react'
            export default function App() {
                const [count, setCount] = useState(0);
                useEffect(() => {
                    const interval = setInterval(() => {
                        setCount(count => count + 1);
                    }, 1000);
                    return () => clearInterval(interval);
                }, []);
                return <text>{`Hello World: ${count}`}</text>;
            }
            """
        let options = BuildOptions(
            bundle: true,
            platform: .neutral,
            format: .cjs,
            target: .es2015,
            jsxFactory: "_FLICKCORE_$REACT.createElement",
            jsxFragment: "_FLICKCORE_$REACT.Fragment",
            loader: .jsx,
        )

        // Use the pre-built React global transform plugin
        let reactPlugin = Plugin.reactGlobalTransform()
        options.addPlugin(reactPlugin)
        XCTAssertEqual(1, options.plugins.count)
        let result = try Builder(options).build(jsx)
        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(
            result.contains(
                """
                return /* @__PURE__ */ _FLICKCORE_$REACT.createElement("text", null, `Hello World: ${count}`);
                """))
        XCTAssertTrue(result.contains("import_react.useState"))
    }

    func testPluginWithCustomGlobalName() throws {
        let jsx = """
            import React from 'react'
            const element = <h1>Test</h1>;
            """

        let options = BuildOptions(
            bundle: true,
            loader: .jsx
        )

        // Use custom global name
        let reactPlugin = Plugin.reactGlobalTransform(globalName: "window.MyReact")
        options.addPlugin(reactPlugin)

        let result = try Builder(options).build(jsx)

        XCTAssertTrue(result.contains("module.exports = window.MyReact"))
    }

    func testPluginOnStartCallback() throws {
        let options = BuildOptions(
            bundle: false,
            loader: .js
        )

        // Create a plugin with onStart callback
        let plugin = Plugin(name: "test-plugin")
        plugin.addOnStartPrint("Test plugin started!")
        options.addPlugin(plugin)

        // This should print "Test plugin started!" to console
        let result = try Builder(options).build("console.log('test')")

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("console.log"))
    }

    func testMultiplePlugins() throws {
        let jsx = """
            import React, { useState } from 'react'
            import _ from 'lodash'

            export default function App() {
                const [data, setData] = useState([])
                return <div>{_.join(data, ', ')}</div>
            }
            """

        let options = BuildOptions(
            bundle: true,
            loader: .jsx
        )

        // Add multiple plugins
        let reactPlugin = Plugin.reactGlobalTransform()
        let externalPlugin = Plugin.externalizeNodeModules()

        options.plugins = [reactPlugin, externalPlugin]
        XCTAssertEqual(2, options.plugins.count)

        let result = try Builder(options).build(jsx)

        // React should be transformed to global
        XCTAssertTrue(result.contains("_FLICKCORE_$REACT"))

        // Lodash should be externalized (not bundled)
        XCTAssertTrue(result.contains("require(\"lodash\")") || result.contains("from \"lodash\""))
    }

    func testCustomPlugin() throws {
        let options = BuildOptions(
            bundle: true,
            loader: .js
        )

        // Create a custom plugin that replaces a specific import
        let plugin = Plugin(name: "custom-replacer")
        plugin.addOnResolve(
            filter: "^my-custom-module$",
            path: "virtual-module",
            toNamespace: "custom-namespace"
        )
        plugin.addOnLoad(
            filter: ".*",
            namespace: "custom-namespace",
            contents: "export const value = 'replaced value';",
            loader: Int(EsbuildmobileGetLoaderJS())
        )

        options.addPlugin(plugin)

        let code = """
            import { value } from 'my-custom-module'
            console.log(value);
            """

        let result = try Builder(options).build(code)

        XCTAssertTrue(result.contains("replaced value"))
    }
}
