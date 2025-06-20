import ESBuildMobile

public enum EngineName: String, CaseIterable {
    case chrome
    case deno
    case edge
    case firefox
    case hermes
    case ie
    case ios
    case node
    case opera
    case rhino
    case safari

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetEngineChrome(): self = .chrome
        case EsbuildmobileGetEngineDeno(): self = .deno
        case EsbuildmobileGetEngineEdge(): self = .edge
        case EsbuildmobileGetEngineFirefox(): self = .firefox
        case EsbuildmobileGetEngineHermes(): self = .hermes
        case EsbuildmobileGetEngineIE(): self = .ie
        case EsbuildmobileGetEngineIOS(): self = .ios
        case EsbuildmobileGetEngineNode(): self = .node
        case EsbuildmobileGetEngineOpera(): self = .opera
        case EsbuildmobileGetEngineRhino(): self = .rhino
        case EsbuildmobileGetEngineSafari(): self = .safari
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .chrome: return EsbuildmobileGetEngineChrome()
        case .deno: return EsbuildmobileGetEngineDeno()
        case .edge: return EsbuildmobileGetEngineEdge()
        case .firefox: return EsbuildmobileGetEngineFirefox()
        case .hermes: return EsbuildmobileGetEngineHermes()
        case .ie: return EsbuildmobileGetEngineIE()
        case .ios: return EsbuildmobileGetEngineIOS()
        case .node: return EsbuildmobileGetEngineNode()
        case .opera: return EsbuildmobileGetEngineOpera()
        case .rhino: return EsbuildmobileGetEngineRhino()
        case .safari: return EsbuildmobileGetEngineSafari()
        }
    }
}

public struct Engine {
    public let name: EngineName
    public let version: String

    public init(name: EngineName, version: String) {
        self.name = name
        self.version = version
    }
}

public enum JSXMode: String, CaseIterable {
    case transform
    case preserve
    case automatic

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetJSXTransform(): self = .transform
        case EsbuildmobileGetJSXPreserve(): self = .preserve
        case EsbuildmobileGetJSXAutomatic(): self = .automatic
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .transform: return EsbuildmobileGetJSXTransform()
        case .preserve: return EsbuildmobileGetJSXPreserve()
        case .automatic: return EsbuildmobileGetJSXAutomatic()
        }
    }
}

public enum OutputFormat: String, CaseIterable {
    case `default` = ""
    case iife
    case cjs
    case esm

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetFormatDefault(): self = .default
        case EsbuildmobileGetFormatIIFE(): self = .iife
        case EsbuildmobileGetFormatCommonJS(): self = .cjs
        case EsbuildmobileGetFormatESModule(): self = .esm
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .default: return EsbuildmobileGetFormatDefault()
        case .iife: return EsbuildmobileGetFormatIIFE()
        case .cjs: return EsbuildmobileGetFormatCommonJS()
        case .esm: return EsbuildmobileGetFormatESModule()
        }
    }
}

public enum Platform: String, CaseIterable {
    case browser
    case node
    case neutral

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetPlatformBrowser(): self = .browser
        case EsbuildmobileGetPlatformNode(): self = .node
        case EsbuildmobileGetPlatformNeutral(): self = .neutral
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .browser: return EsbuildmobileGetPlatformBrowser()
        case .node: return EsbuildmobileGetPlatformNode()
        case .neutral: return EsbuildmobileGetPlatformNeutral()
        }
    }
}

public enum Target: String, CaseIterable {
    case esnext
    case es2024
    case es2023
    case es2022
    case es2021
    case es2020
    case es2019
    case es2018
    case es2017
    case es2016
    case es2015
    case es6
    case es5

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetTargetESNext(): self = .esnext
        case EsbuildmobileGetTargetES2024(): self = .es2024
        case EsbuildmobileGetTargetES2023(): self = .es2023
        case EsbuildmobileGetTargetES2022(): self = .es2022
        case EsbuildmobileGetTargetES2021(): self = .es2021
        case EsbuildmobileGetTargetES2020(): self = .es2020
        case EsbuildmobileGetTargetES2019(): self = .es2019
        case EsbuildmobileGetTargetES2018(): self = .es2018
        case EsbuildmobileGetTargetES2017(): self = .es2017
        case EsbuildmobileGetTargetES2016(): self = .es2016
        case EsbuildmobileGetTargetES2015(): self = .es2015
        case EsbuildmobileGetTargetES5(): self = .es5
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .esnext: return EsbuildmobileGetTargetESNext()
        case .es2024: return EsbuildmobileGetTargetES2024()
        case .es2023: return EsbuildmobileGetTargetES2023()
        case .es2022: return EsbuildmobileGetTargetES2022()
        case .es2021: return EsbuildmobileGetTargetES2021()
        case .es2020: return EsbuildmobileGetTargetES2020()
        case .es2019: return EsbuildmobileGetTargetES2019()
        case .es2018: return EsbuildmobileGetTargetES2018()
        case .es2017: return EsbuildmobileGetTargetES2017()
        case .es2016: return EsbuildmobileGetTargetES2016()
        case .es2015: return EsbuildmobileGetTargetES2015()
        case .es6: return EsbuildmobileGetTargetES2015()
        case .es5: return EsbuildmobileGetTargetES5()
        }
    }
}

public enum SourceMap: String, CaseIterable {
    case none
    case inline
    case linked
    case external
    case both

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetSourceMapNone(): self = .none
        case EsbuildmobileGetSourceMapInline(): self = .inline
        case EsbuildmobileGetSourceMapLinked(): self = .linked
        case EsbuildmobileGetSourceMapExternal(): self = .external
        case EsbuildmobileGetSourceMapInlineAndExternal(): self = .both
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .none: return EsbuildmobileGetSourceMapNone()
        case .inline: return EsbuildmobileGetSourceMapInline()
        case .linked: return EsbuildmobileGetSourceMapLinked()
        case .external: return EsbuildmobileGetSourceMapExternal()
        case .both: return EsbuildmobileGetSourceMapInlineAndExternal()
        }
    }
}

public enum SourcesContent: String, CaseIterable {
    case include
    case exclude

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetSourcesContentInclude(): self = .include
        case EsbuildmobileGetSourcesContentExclude(): self = .exclude
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .include: return EsbuildmobileGetSourcesContentInclude()
        case .exclude: return EsbuildmobileGetSourcesContentExclude()
        }
    }
}

public enum LogLevel: String, CaseIterable {
    case verbose
    case debug
    case info
    case warning
    case error
    case silent

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetLogLevelVerbose(): self = .verbose
        case EsbuildmobileGetLogLevelDebug(): self = .debug
        case EsbuildmobileGetLogLevelInfo(): self = .info
        case EsbuildmobileGetLogLevelWarning(): self = .warning
        case EsbuildmobileGetLogLevelError(): self = .error
        case EsbuildmobileGetLogLevelSilent(): self = .silent
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .verbose: return EsbuildmobileGetLogLevelVerbose()
        case .debug: return EsbuildmobileGetLogLevelDebug()
        case .info: return EsbuildmobileGetLogLevelInfo()
        case .warning: return EsbuildmobileGetLogLevelWarning()
        case .error: return EsbuildmobileGetLogLevelError()
        case .silent: return EsbuildmobileGetLogLevelSilent()
        }
    }
}

public enum ColorMode: String, CaseIterable {
    case auto
    case always
    case never

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetColorIfTerminal(): self = .auto
        case EsbuildmobileGetColorAlways(): self = .always
        case EsbuildmobileGetColorNever(): self = .never
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .auto: return EsbuildmobileGetColorIfTerminal()
        case .always: return EsbuildmobileGetColorAlways()
        case .never: return EsbuildmobileGetColorNever()
        }
    }
}

public enum Charset: String, CaseIterable {
    case ascii
    case utf8

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetCharsetASCII(): self = .ascii
        case EsbuildmobileGetCharsetUTF8(): self = .utf8
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .ascii: return EsbuildmobileGetCharsetASCII()
        case .utf8: return EsbuildmobileGetCharsetUTF8()
        }
    }
}

public enum LegalComments: String, CaseIterable {
    case `default` = ""
    case none
    case inline
    case eof
    case linked
    case external

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetLegalCommentsDefault(): self = .default
        case EsbuildmobileGetLegalCommentsNone(): self = .none
        case EsbuildmobileGetLegalCommentsInline(): self = .inline
        case EsbuildmobileGetLegalCommentsEndOfFile(): self = .eof
        case EsbuildmobileGetLegalCommentsLinked(): self = .linked
        case EsbuildmobileGetLegalCommentsExternal(): self = .external
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .default: return EsbuildmobileGetLegalCommentsDefault()
        case .none: return EsbuildmobileGetLegalCommentsNone()
        case .inline: return EsbuildmobileGetLegalCommentsInline()
        case .eof: return EsbuildmobileGetLegalCommentsEndOfFile()
        case .linked: return EsbuildmobileGetLegalCommentsLinked()
        case .external: return EsbuildmobileGetLegalCommentsExternal()
        }
    }
}

public enum DropMode: String, CaseIterable {
    case none = ""
    case console
    case debugger

    init?(fromGoValue value: Int) {
        switch value {
        case 0: self = .none
        case EsbuildmobileGetDropConsole(): self = .console
        case EsbuildmobileGetDropDebugger(): self = .debugger
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .none: return 0
        case .console: return EsbuildmobileGetDropConsole()
        case .debugger: return EsbuildmobileGetDropDebugger()
        }
    }
}

public enum TreeShaking: String, CaseIterable {
    case `default` = ""
    case `true`
    case `false`

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetTreeShakingDefault(): self = .default
        case EsbuildmobileGetTreeShakingTrue(): self = .true
        case EsbuildmobileGetTreeShakingFalse(): self = .false
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .default: return EsbuildmobileGetTreeShakingDefault()
        case .true: return EsbuildmobileGetTreeShakingTrue()
        case .false: return EsbuildmobileGetTreeShakingFalse()
        }
    }
}

public enum MangleQuoted: String, CaseIterable {
    case `true`
    case `false`

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetMangleQuotedTrue(): self = .true
        case EsbuildmobileGetMangleQuotedFalse(): self = .false
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .true: return EsbuildmobileGetMangleQuotedTrue()
        case .false: return EsbuildmobileGetMangleQuotedFalse()
        }
    }
}

public enum Packages: String, CaseIterable {
    case `default` = ""
    case bundle
    case external

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetPackagesDefault(): self = .default
        case EsbuildmobileGetPackagesBundle(): self = .bundle
        case EsbuildmobileGetPackagesExternal(): self = .external
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .default: return EsbuildmobileGetPackagesDefault()
        case .bundle: return EsbuildmobileGetPackagesBundle()
        case .external: return EsbuildmobileGetPackagesExternal()
        }
    }
}

public enum Loader: String, CaseIterable {
    case js
    case jsx
    case ts
    case tsx
    case css
    case json
    case text
    case base64
    case dataurl
    case file
    case binary
    case copy
    case empty
    case globalcss
    case localcss

    init?(fromGoValue value: Int) {
        switch value {
        case EsbuildmobileGetLoaderJS(): self = .js
        case EsbuildmobileGetLoaderJSX(): self = .jsx
        case EsbuildmobileGetLoaderTS(): self = .ts
        case EsbuildmobileGetLoaderTSX(): self = .tsx
        case EsbuildmobileGetLoaderCSS(): self = .css
        case EsbuildmobileGetLoaderJSON(): self = .json
        case EsbuildmobileGetLoaderText(): self = .text
        case EsbuildmobileGetLoaderBase64(): self = .base64
        case EsbuildmobileGetLoaderDataURL(): self = .dataurl
        case EsbuildmobileGetLoaderFile(): self = .file
        case EsbuildmobileGetLoaderBinary(): self = .binary
        case EsbuildmobileGetLoaderCopy(): self = .copy
        case EsbuildmobileGetLoaderEmpty(): self = .empty
        case EsbuildmobileGetLoaderGlobalCSS(): self = .globalcss
        case EsbuildmobileGetLoaderLocalCSS(): self = .localcss
        default: return nil
        }
    }

    var goValue: Int {
        switch self {
        case .js: return EsbuildmobileGetLoaderJS()
        case .jsx: return EsbuildmobileGetLoaderJSX()
        case .ts: return EsbuildmobileGetLoaderTS()
        case .tsx: return EsbuildmobileGetLoaderTSX()
        case .css: return EsbuildmobileGetLoaderCSS()
        case .json: return EsbuildmobileGetLoaderJSON()
        case .text: return EsbuildmobileGetLoaderText()
        case .base64: return EsbuildmobileGetLoaderBase64()
        case .dataurl: return EsbuildmobileGetLoaderDataURL()
        case .file: return EsbuildmobileGetLoaderFile()
        case .binary: return EsbuildmobileGetLoaderBinary()
        case .copy: return EsbuildmobileGetLoaderCopy()
        case .empty: return EsbuildmobileGetLoaderEmpty()
        case .globalcss: return EsbuildmobileGetLoaderGlobalCSS()
        case .localcss: return EsbuildmobileGetLoaderLocalCSS()
        }
    }
}
