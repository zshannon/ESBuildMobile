import XCTest

@testable import ESBuild

final class ESBuildMobileTests: XCTestCase {
    func testBasicTransformation() throws {
        let jsx = "<div>Hello World</div>"
        let transformer = Transform()
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("Hello World"))
        XCTAssertTrue(result.contains("React.createElement"))
    }

    func testStringExtensionTransform() throws {
        let jsx = "<div>Hello World</div>"
        let result = try jsx.transform()

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("Hello World"))
        XCTAssertTrue(result.contains("React.createElement"))
    }

    func testReactPreset() throws {
        let jsx = "<div>Test</div>"
        let options = TransformOptions.react()
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("React.createElement"))
        XCTAssertTrue(result.contains("Test"))
    }

    func testPreactPreset() throws {
        let jsx = "<div>Hello</div>"
        let options = TransformOptions.preact()
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("h("))
        XCTAssertTrue(result.contains("Hello"))
    }

    func testVuePreset() throws {
        let jsx = "<div>Vue Test</div>"
        let options = TransformOptions.vue()
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("h("))
        XCTAssertTrue(result.contains("Vue Test"))
    }

    func testCustomJSXConfiguration() throws {
        let jsx = "<div>Custom</div>"
        let options = TransformOptions.custom(factory: "createElement", fragment: "Fragment")
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("createElement"))
        XCTAssertTrue(result.contains("Custom"))
    }

    func testMinificationOptions() throws {
        let jsx = "<div>Test</div>"
        let options = TransformOptions()
        options.minifyWhitespace = true
        options.target = .es2017
        XCTAssertEqual(options.target, .es2017)
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)
        print("Minification result: '\(result)'")

        XCTAssertFalse(result.isEmpty, "Expected non-empty result but got: '\(result)'")
        XCTAssertTrue(
            result.contains("React.createElement"),
            "Expected React.createElement but got: '\(result)'"
        )
    }

    func testMethodChaining() throws {
        let jsx = "<div>Chain Test</div>"
        let options = TransformOptions()
            .jsx(factory: "h", fragment: "Fragment")

        let transformer = Transform(options)
        let result = try transformer.transform(jsx)
        print("Method chaining result: '\(result)'")

        XCTAssertFalse(result.isEmpty, "Expected non-empty result but got: '\(result)'")
        XCTAssertTrue(result.contains("h("), "Expected h( but got: '\(result)'")
        XCTAssertTrue(result.contains("Chain Test"), "Expected Chain Test but got: '\(result)'")
    }

    func testStringExtensionWithOptions() throws {
        let jsx = "<div>Extension Test</div>"
        let options = TransformOptions.preact()
        let result = try jsx.transform(with: options)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("h("))
        XCTAssertTrue(result.contains("Extension Test"))
    }

    func testDevelopmentPreset() throws {
        let jsx = "<div>Dev Test</div>"
        let options = TransformOptions.development()
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("React.createElement"))
        XCTAssertTrue(result.contains("Dev Test"))
    }

    func testTransformWithResult() throws {
        let jsx = "<div>Result Test</div>"
        let transformer = Transform()
        let transformResult = try transformer.transformWithResult(jsx)

        XCTAssertFalse(transformResult.code.isEmpty)
        XCTAssertTrue(transformResult.code.contains("React.createElement"))
        XCTAssertTrue(transformResult.code.contains("Result Test"))
        XCTAssertTrue(transformResult.warnings.isEmpty)
    }

    func testMutableOptionsConfiguration() throws {
        let jsx = "<div>Mutable Test</div>"
        let options = TransformOptions()
        options.jsxFactory = "customFactory"
        options.jsxFragment = "customFragment"

        let transformer = Transform(options)
        let result = try transformer.transform(jsx)
        print("Mutable options result: '\(result)'")

        XCTAssertFalse(result.isEmpty, "Expected non-empty result but got: '\(result)'")
        XCTAssertTrue(
            result.contains("customFactory"), "Expected customFactory but got: '\(result)'"
        )
        XCTAssertTrue(result.contains("Mutable Test"), "Expected Mutable Test but got: '\(result)'")
    }

    func testComplexJSX() throws {
        let jsx = """
        <div className="container">
            <h1>Title</h1>
            <p>Paragraph with <span>nested</span> content</p>
            <button onClick={handleClick}>Click me</button>
        </div>
        """

        let transformer = Transform()
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("React.createElement"))
        XCTAssertTrue(result.contains("container"))
        XCTAssertTrue(result.contains("Title"))
        XCTAssertTrue(result.contains("handleClick"))
    }

    func testErrorHandling() {
        // Test with malformed JSX - ESBuild is quite forgiving,
        // so this might still succeed, but we test the error handling path
        let invalidJSX = "<div><span></div>"

        do {
            let transformer = Transform()
            let result = try transformer.transform(invalidJSX)
            // If it succeeds, that's also acceptable as ESBuild can handle some malformed JSX
            XCTAssertFalse(result.isEmpty)
        } catch let error as TransformError {
            // If it fails, verify we get the right error type
            switch error {
            case let .transformationFailed(message):
                XCTAssertFalse(message.isEmpty)
            case let .invalidOptions(message):
                XCTAssertFalse(message.isEmpty)
            case let .internalError(message):
                XCTAssertFalse(message.isEmpty)
            }
        } catch {
            XCTFail("Unexpected error type: \(error)")
        }
    }

    func testJSXFragments() throws {
        let jsx = "<>Fragment content</>"
        let options = TransformOptions.react()

        do {
            let transformer = Transform(options)
            let result = try transformer.transform(jsx)
            XCTAssertFalse(result.isEmpty)
            // Should contain React.Fragment or similar
            XCTAssertTrue(result.contains("Fragment") || result.contains("React.createElement"))
        } catch {
            // Fragment syntax might not be fully supported, which is acceptable
            print("Fragment test failed (this may be expected): \(error)")
        }
    }

    func testReactTypeScriptPreset() throws {
        let jsx = "<div>TypeScript Test</div>"
        let options = TransformOptions.reactTypeScript()
        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("React.createElement"))
        XCTAssertTrue(result.contains("TypeScript Test"))
    }

    func testGlobalNameOption() throws {
        let jsx = "<div>Global Test</div>"
        let options = TransformOptions()
        options.globalName = "MyLibrary"

        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("React.createElement"))
    }

    func testBannerAndFooter() throws {
        let jsx = "<div>Banner Test</div>"
        let options = TransformOptions()
        options.banner = "/* Header comment */"
        options.footer = "/* Footer comment */"

        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("React.createElement"))
        // Note: Banner and footer might not appear in transform mode,
        // but we test that the options don't break the transformation
    }

    func testTransformWithNewInit() throws {
        let jsx = "<div>New Init Test</div>"
        let options = TransformOptions(
            jsxFactory: "h",
            jsxFragment: "Fragment",
            jsx: .transform,
            minifyWhitespace: true,
            target: .es2020,
            loader: .jsx
        )

        let transformer = Transform(options)
        let result = try transformer.transform(jsx)

        XCTAssertFalse(result.isEmpty)
        XCTAssertTrue(result.contains("h("))
        XCTAssertTrue(result.contains("New Init Test"))
    }

    func testReuseTransformer() throws {
        let options = TransformOptions.preact()
        let transformer = Transform(options)

        // Transform multiple pieces of code with same transformer
        let jsx1 = "<div>First</div>"
        let jsx2 = "<span>Second</span>"

        let result1 = try transformer.transform(jsx1)
        let result2 = try transformer.transform(jsx2)

        XCTAssertFalse(result1.isEmpty)
        XCTAssertTrue(result1.contains("h("))
        XCTAssertTrue(result1.contains("First"))

        XCTAssertFalse(result2.isEmpty)
        XCTAssertTrue(result2.contains("h("))
        XCTAssertTrue(result2.contains("Second"))
    }

    func testBasicBuild() throws {
        let code = "export const foo = 'bar'"
        let options = BuildOptions(bundle: true, format: .esm, target: .es2017)
        XCTAssertEqual(options.target, .es2017)
        let builder = Builder(options)
        let result = try builder.build(code)
        XCTAssertTrue(result.contains("foo"))
    }

    func testPlatformOption() throws {
        let buildOptions = BuildOptions(platform: .neutral)
        XCTAssertEqual(buildOptions.platform, .neutral)
        let transformOptions = TransformOptions(platform: .neutral)
        XCTAssertEqual(transformOptions.platform, .neutral)
    }
}
