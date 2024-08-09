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
      targets: ["ESBuildMobile"]
    )
  ],
  targets: [
    .binaryTarget(
      name: "ESBuildMobileFramework",
      path: "./Sources/ESBuildMobile.xcframework"
    ),
    .target(
      name: "ESBuildMobile",
      dependencies: [
        .target(name: "ESBuildMobileFramework")
      ]
    ),
  ]
)
