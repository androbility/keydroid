package adbcmd

import (
	"strings"
	"unicode/utf8"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfigFile attempts to read a config-file from configDir.
// Failing that, LoadConfigFile parses defaultBindings for the
// configuration.
func LoadConfigFile(configDir, defaultBindings string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir)
	viper.SetDefault("keybindings", keymap)

	if err := viper.ReadInConfig(); err != nil {
		log.Debug("Configuration file not found; using defaults.")

		r := strings.NewReader(defaultBindings)
		if err = viper.ReadConfig(r); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("Loading default configuration failed.")
		}
	}

	// e.g., "h": "KEYCODE_HOME"
	for keycodeNewMapping, keycodeName := range viper.GetStringMapString("keybindings") {
		// Skip on empty value
		if len(keycodeNewMapping) == 0 {
			continue
		}
		keycodeName = strings.ToUpper(keycodeName)

		// Is the specified keycodeName valid?  If not, skip to the next.
		// e.g., "KEYCODE_HOME" is valid, "KEYCODE_HOMEZ" isn't
		newKeyCode, ok := keycodeLookupTable[keycodeName]
		if !ok {
			log.WithFields(log.Fields{
				"key": keycodeName,
			}).Warn("Invalid key")
			continue
		}

		// Extract the first rune from keycodeNewMapping.
		// That's the key we want to set.
		r, _ := utf8.DecodeRuneInString(keycodeNewMapping)

		// Set key to
		keymap[r] = newKeyCode
	}

	// If no keys are defined, fail.
	if len(keymap) == 0 {
		log.Fatal("No keybindings are defined; aborting.")
	}
}

var (
	keymap = map[rune]Keycode{}
)
