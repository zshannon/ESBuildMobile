package esbuildmobile

import "github.com/evanw/esbuild/pkg/api"

// Engine represents a JavaScript engine with its version
// This is our custom wrapper that maps to esbuild's api.Engine
type Engine struct {
	Name    api.EngineName
	Version string
}

// ToAPIEngine converts our Engine type to esbuild's api.Engine
func (e *Engine) ToAPIEngine() api.Engine {
	return api.Engine{
		Name:    e.Name,
		Version: e.Version,
	}
}

// FromAPIEngine creates our Engine type from esbuild's api.Engine
func FromAPIEngine(apiEngine api.Engine) Engine {
	return Engine{
		Name:    apiEngine.Name,
		Version: apiEngine.Version,
	}
}

// EngineNameFromString converts a string to api.EngineName
func EngineNameFromString(name string) api.EngineName {
	switch name {
	case "chrome":
		return api.EngineChrome
	case "deno":
		return api.EngineDeno
	case "edge":
		return api.EngineEdge
	case "firefox":
		return api.EngineFirefox
	case "hermes":
		return api.EngineHermes
	case "ie":
		return api.EngineIE
	case "ios":
		return api.EngineIOS
	case "node":
		return api.EngineNode
	case "opera":
		return api.EngineOpera
	case "rhino":
		return api.EngineRhino
	case "safari":
		return api.EngineSafari
	default:
		return api.EngineChrome // fallback
	}
}

// EngineNameToString returns the string representation of api.EngineName
func EngineNameToString(e api.EngineName) string {
	switch e {
	case api.EngineChrome:
		return "chrome"
	case api.EngineDeno:
		return "deno"
	case api.EngineEdge:
		return "edge"
	case api.EngineFirefox:
		return "firefox"
	case api.EngineHermes:
		return "hermes"
	case api.EngineIE:
		return "ie"
	case api.EngineIOS:
		return "ios"
	case api.EngineNode:
		return "node"
	case api.EngineOpera:
		return "opera"
	case api.EngineRhino:
		return "rhino"
	case api.EngineSafari:
		return "safari"
	default:
		return "chrome"
	}
}

// Engine name constant getters for mobile bindings
func GetEngineChrome() int  { return int(api.EngineChrome) }
func GetEngineDeno() int    { return int(api.EngineDeno) }
func GetEngineEdge() int    { return int(api.EngineEdge) }
func GetEngineFirefox() int { return int(api.EngineFirefox) }
func GetEngineHermes() int  { return int(api.EngineHermes) }
func GetEngineIE() int      { return int(api.EngineIE) }
func GetEngineIOS() int     { return int(api.EngineIOS) }
func GetEngineNode() int    { return int(api.EngineNode) }
func GetEngineOpera() int   { return int(api.EngineOpera) }
func GetEngineRhino() int   { return int(api.EngineRhino) }
func GetEngineSafari() int  { return int(api.EngineSafari) }
