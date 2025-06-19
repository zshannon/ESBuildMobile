# ESBuildMobile Swift API Examples

A comprehensive guide showcasing the idiomatic Swift API for ESBuildMobile.

## Overview

ESBuildMobile now provides a completely idiomatic Swift API that hides all the `gomobile` complexity. The API uses computed properties with custom setters that call the underlying Go Configure methods, making it feel like a native Swift library.

## Basic Usage

### Simple JSX Transformation

```swift
import ESBuildMobileHelpers

// Transform JSX with default React settings
let jsx = "<div>Hello World</div>"
let result = try JSXTransformer.transform(jsx)
print(result)
// Output: /* @__PURE__ */ React.createElement("div", null, "Hello World");

// Using String extension
let result2 = try jsx.transformJSX()
```

### Using Transform Options
