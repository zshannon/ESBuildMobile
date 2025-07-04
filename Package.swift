// swift-tools-version:6.1
import PackageDescription

let package = Package(
    name: "ESBuildMobile",
    platforms: [.iOS(.v17), .macOS(.v14)],
    products: [
        .library(
            name: "ESBuild",
            targets: ["ESBuild"]
        )
    ],
    targets: [
        .binaryTarget(
            name: "ESBuildMobile",
            path: "./Sources/ESBuildMobile.xcframework"
        ),
        .target(
            name: "ESBuild",
            dependencies: [
                .target(name: "ESBuildMobile")
                // .byName(name: "ESBuildMobile")
            ]
        ),
        .testTarget(
            name: "ESBuildMobileTests",
            dependencies: [
                .target(name: "ESBuild"),
                .target(name: "ESBuildMobile"),
            ]
        ),
    ]
)
