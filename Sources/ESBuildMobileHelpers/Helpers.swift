import ESBuildMobile

/// An error type representing possible failures during JSX transformation.
public enum TransformError: Error {
    case transformationFailed(NSError)
}

/// Transforms a JSX string using the specified options.
///
/// - Parameters:
///   - jsx: The JSX string to be transformed.
///   - options: The `EsbuildmobileTransformOptions` object containing transformation options.
/// - Throws: `TransformError.transformationFailed` if the transformation fails, encapsulating the underlying `NSError`.
/// - Returns: The transformed JSX as a `String`.
public func transformJSX(_ jsx: String, options: EsbuildmobileTransformOptions) throws -> String {
    var error: NSError?
    let result = EsbuildmobileTransformJSX(jsx, options, &error)
    
    if let error = error {
        throw TransformError.transformationFailed(error)
    }
    
    return result ?? ""
}
