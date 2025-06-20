# ESBuild Mobile Plugin Usage

This document explains how to create and use ESBuild plugins from Swift using the mobile-friendly plugin API.

## Basic Plugin Creation

### 1. Create a Plugin

```swift
let plugin = esbuildmobile.NewPlugin("my-plugin")
```

### 2. Add OnResolve Callback

```swift
// Create resolve options with a filter
let resolveOptions = esbuildmobile.NewOnResolveOptions()
resolveOptions.setResolveFilter("^react$")

// Create a callback class that implements OnResolveCallback
class MyResolveCallback: NSObject, esbuildmobileOnResolveCallback {
    func call(_ args: esbuildmobileOnResolveArgs?) -> esbuildmobileOnResolveResult? {
        let result = esbuildmobile.NewOnResolveResult()
        result.setResolvePath(args?.path ?? "")
        result.setResolveNamespace("my-namespace")
        return result
    }
}

let resolveCallback = MyResolveCallback()
plugin.onResolve(resolveOptions, resolveCallback)
```

### 3. Add OnLoad Callback

```swift
// Create load options with a namespace filter
let loadOptions = esbuildmobile.NewOnLoadOptions()
loadOptions.setLoadNamespace("my-namespace")

// Create a callback class that implements OnLoadCallback
class MyLoadCallback: NSObject, esbuildmobileOnLoadCallback {
    func call(_ args: esbuildmobileOnLoadArgs?) -> esbuildmobileOnLoadResult? {
        let result = esbuildmobile.NewOnLoadResult()
        result.setLoadContents("module.exports = _FLICKCORE_$REACT")
        result.setLoadLoader(esbuildmobile.GetLoaderJS())
        return result
    }
}

let loadCallback = MyLoadCallback()
plugin.onLoad(loadOptions, loadCallback)
```

### 4. Add Plugin to Build Options

```swift
let buildOptions = esbuildmobile.NewBuildOptions()
buildOptions.addPlugin(plugin)
```

## React Global Transform Example

Here's a complete example that replicates the TypeScript plugin you showed:

```swift
class ReactGlobalTransformPlugin {
    static func create() -> esbuildmobilePlugin {
        let plugin = esbuildmobile.NewPlugin("react-global-transform")
        
        // OnResolve: Intercept react imports
        let resolveOptions = esbuildmobile.NewOnResolveOptions()
        resolveOptions.setResolveFilter("^react$")
        
        class ReactResolveCallback: NSObject, esbuildmobileOnResolveCallback {
            func call(_ args: esbuildmobileOnResolveArgs?) -> esbuildmobileOnResolveResult? {
                let result = esbuildmobile.NewOnResolveResult()
                result.setResolvePath(args?.path ?? "")
                result.setResolveNamespace("use-flick-react-global")
                return result
            }
        }
        
        plugin.onResolve(resolveOptions, ReactResolveCallback())
        
        // OnLoad: Provide the transformed content
        let loadOptions = esbuildmobile.NewOnLoadOptions()
        loadOptions.setLoadFilter(".*")
        loadOptions.setLoadNamespace("use-flick-react-global")
        
        class ReactLoadCallback: NSObject, esbuildmobileOnLoadCallback {
            func call(_ args: esbuildmobileOnLoadArgs?) -> esbuildmobileOnLoadResult? {
                let result = esbuildmobile.NewOnLoadResult()
                result.setLoadContents("module.exports = _FLICKCORE_$REACT")
                result.setLoadLoader(esbuildmobile.GetLoaderJS())
                return result
            }
        }
        
        plugin.onLoad(loadOptions, ReactLoadCallback())
        
        return plugin
    }
}

// Usage
let buildOptions = esbuildmobile.NewBuildOptions()
let plugin = ReactGlobalTransformPlugin.create()
buildOptions.addPlugin(plugin)
```

## Helper Functions

The library provides several helper functions to make common plugin patterns easier:

### Simple Transform Plugin

```swift
// Creates a plugin that transforms all imports matching a pattern
let plugin = esbuildmobile.CreateSimpleTransformPlugin(
    "my-transform",
    "^react$",
    "my-namespace", 
    "module.exports = _FLICKCORE_$REACT",
    esbuildmobile.GetLoaderJS()
)
```

### Quick Result Creation

```swift
// Create resolve results
let externalResult = esbuildmobile.CreateExternalResolveResult("some-path")
let namespaceResult = esbuildmobile.CreateNamespaceResolveResult("some-path", "my-namespace")

// Create load results with specific loaders
let jsResult = esbuildmobile.CreateJSLoadResult("console.log('hello')")
let cssResult = esbuildmobile.CreateCSSLoadResult("body { color: red; }")
let jsonResult = esbuildmobile.CreateJSONLoadResult("{\"key\": \"value\"}")
```

### Filter Creation

```swift
// Create common filter patterns
let pathFilter = esbuildmobile.CreateFilterForPath("^react$")
let namespaceFilter = esbuildmobile.CreateFilterForNamespace("my-namespace")
let combinedFilter = esbuildmobile.CreateFilterForPathAndNamespace("^react$", "my-namespace")
```

## Available Loaders

Use these getter functions to specify loaders:

- `esbuildmobile.GetLoaderJS()` - JavaScript
- `esbuildmobile.GetLoaderTS()` - TypeScript  
- `esbuildmobile.GetLoaderTSX()` - TypeScript JSX
- `esbuildmobile.GetLoaderJSX()` - JavaScript JSX
- `esbuildmobile.GetLoaderCSS()` - CSS
- `esbuildmobile.GetLoaderJSON()` - JSON
- `esbuildmobile.GetLoaderText()` - Text
- `esbuildmobile.GetLoaderFile()` - File
- `esbuildmobile.GetLoaderBinary()` - Binary

## OnStart and OnEnd Callbacks

You can also add callbacks that run at the start and end of builds:

```swift
class StartCallback: NSObject, esbuildmobileOnStartCallback {
    func call() -> esbuildmobileOnStartResult? {
        print("Build starting...")
        return esbuildmobile.NewOnStartResult()
    }
}

class EndCallback: NSObject, esbuildmobileOnEndCallback {
    func call(_ result: esbuildmobileBuildResult?) -> esbuildmobileOnEndResult? {
        print("Build finished with \(result?.getErrorsCount() ?? 0) errors")
        return esbuildmobile.NewOnEndResult()
    }
}

plugin.onStart(StartCallback())
plugin.onEnd(EndCallback())
```

## Error Handling

Add errors and warnings to plugin results:

```swift
class ErrorHandlingCallback: NSObject, esbuildmobileOnLoadCallback {
    func call(_ args: esbuildmobileOnLoadArgs?) -> esbuildmobileOnLoadResult? {
        let result = esbuildmobile.NewOnLoadResult()
        
        // Add an error
        let errorMsg = esbuildmobileMessage()
        errorMsg.text = "Something went wrong"
        result.addLoadError(errorMsg)
        
        // Add a warning
        let warningMsg = esbuildmobileMessage()
        warningMsg.text = "This might be a problem"
        result.addLoadWarning(warningMsg)
        
        return result
    }
}
```

## Watch Files and Directories

Tell ESBuild to watch additional files and directories:

```swift
class WatchingCallback: NSObject, esbuildmobileOnLoadCallback {
    func call(_ args: esbuildmobileOnLoadArgs?) -> esbuildmobileOnLoadResult? {
        let result = esbuildmobile.NewOnLoadResult()
        result.setLoadContents("// generated content")
        result.setLoadLoader(esbuildmobile.GetLoaderJS())
        
        // Watch additional files
        result.addLoadWatchFile("/path/to/config.json")
        result.addLoadWatchDir("/path/to/templates/")
        
        return result
    }
}
```

## Best Practices

1. **Use namespaces** to avoid conflicts between plugins
2. **Be specific with filters** to avoid unnecessary processing
3. **Handle errors gracefully** by adding error messages instead of crashing
4. **Use helper functions** when possible to reduce boilerplate
5. **Watch relevant files** if your plugin depends on external resources

## Plugin Lifecycle

1. Plugin is created with `NewPlugin()`
2. OnResolve and OnLoad callbacks are registered
3. Plugin is added to BuildOptions with `addPlugin()`
4. During build:
   - OnStart callback runs (if set)
   - For each import, OnResolve callbacks are checked
   - For resolved modules, OnLoad callbacks are checked  
   - OnEnd callback runs (if set)

This mobile-friendly API allows you to create powerful ESBuild plugins directly from Swift while maintaining compatibility with the ESBuild ecosystem.