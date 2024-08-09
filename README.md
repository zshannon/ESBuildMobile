# ESBuildMobile
 Swift and Kotlin ESBuild bindings for use on Android, iOS, and MacOS


## Swift Example

```swift
import ESBuildMobileHelpers
import ESBuildMobile

let opts = EsbuildmobileTransformOptions()

// Set your desired JSX factory function
opts.setJSXFactory("exports.text")

do {
    // Transform JSX using the helper function
    let result = try transformJSX("""
    <text>Hello World!</text>
    """, options: opts)
    
    print("Result: ", result)
    
} catch let error as TransformError {
    switch error {
    case .transformationFailed(let nsError):
        print("Transformation failed with error: \(nsError)")
    }
} catch {
    print("An unexpected error occurred: \(error)")
}
```