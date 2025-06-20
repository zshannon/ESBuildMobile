extension String {
    /// Transforms this string using default settings
    /// - Returns: The transformed code
    /// - Throws: TransformError if transformation fails
    public func transform() throws -> String {
        let transformer = Transform()
        return try transformer.transform(self)
    }

    /// Transforms this string using the specified options
    /// - Parameter options: Transform options
    /// - Returns: The transformed code
    /// - Throws: TransformError if transformation fails
    public func transform(with options: TransformOptions) throws -> String {
        let transformer = Transform(options)
        return try transformer.transform(self)
    }
}
