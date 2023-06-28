package plugins

import (
	"path"
	"plugin"

	"github.com/levintp/observer/internal/config"
)

// Function to get a plugin.
func GetSymbol(pluginName string, symbolName string) (plugin.Symbol, error) {
	// Open the plugin shared object.
	plugin, err := plugin.Open(path.Join(config.Get().PluginDir, pluginName))
	if err != nil {
		return nil, err
	}

	// Lookup the symbol in the plugin
	return plugin.Lookup(symbolName)
}
