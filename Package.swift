// swift-tools-version:5.3
import PackageDescription

let package = Package(
  name: "ESBuildMobile",
  platforms: [
    .macOS(.v10_14),
  ],
  products: [
    .library(
      name: "ESBuildMobile",
      targets: ["ESBuildMobileHelpers", "ESBuildMobile"]
    )
  ],
  targets: [
    .binaryTarget(
      name: "ESBuildMobile",
      path: "./Sources/ESBuildMobile.xcframework"
    ),
    .target(
      name: "ESBuildMobileHelpers",
      dependencies: [
        .target(name: "ESBuildMobile"),
        // .byName(name: "ESBuildMobile")
      ]
    ),
  ]
)
