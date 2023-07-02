package plugins

import (
	"fmt"
	"path"
	"plugin"
	"reflect"

	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/types"
)

// Function to get a symbol from a plugin.
func GetSymbol(pluginName string, symbolName string) (plugin.Symbol, error) {
	// Open the plugin shared object.
	plugin, err := plugin.Open(path.Join(config.Get().PluginDir, pluginName))
	if err != nil {
		return nil, err
	}

	// Lookup the symbol in the plugin
	return plugin.Lookup(symbolName)
}

// Function to get a sampling function from a plugin.
func GetSamplingFunc(pluginName string, functionName string) (types.SamplingFunc, error) {
	// Get the symbol from the plugin.
	symbol, err := GetSymbol(pluginName, functionName)
	if err != nil {
		return nil, err
	}

	// Check if the symbol is of matching type.
	if symbolType := reflect.TypeOf(symbol); symbolType != types.SamplingFuncType {
		return nil, fmt.Errorf("symbol %v in module %v is not a sampling function", functionName, pluginName)
	}

	// Try to convert the symbol to a `SamplingFunc`.
	function, ok := symbol.(types.SamplingFunc)
	if !ok {
		return nil, fmt.Errorf("failed to convert symbol %v in module %v to sampling function", functionName, pluginName)
	}

	return function, nil
}
