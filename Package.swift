// swift-tools-version:5.3
import PackageDescription

let package = Package(
  name: "ESBuildMobile",
  platforms: [
    .macOS(.v10_14),
    .iOS(.v13),
  ],
  products: [
    .library(
      name: "ESBuildMobile",
      targets: ["ESBuildMobileWrapper"]
    )
  ],
  targets: [
    .binaryTarget(
      name: "ESBuildMobile",
      path: "./Sources/ESBuildMobile/ESBuildMobile.xcframework"
    ),
    .target(
      name: "ESBuildMobileWrapper",
      dependencies: [
        .target(name: "ESBuildMobile")
      ]
    ),
  ]
)
