// config package is all about setting global
// configurations required for golang cli tool

package config

// ====================================
// EDITABLE BLOCK [ STARTED ]
// ====================================

// Modify this configuation as needed
var config = map[string]string{
	"cli.name":      "Geon",
	"cli.usage":     "Golang CLI Application",
	"cli.version":   "1.0.0",
	"cli.copyright": "Abhishek Mane",
	"cli.author":    "abhishekmane@outlook.in",
}

// ====================================
// EDITABLE BLOCK [ ENDED ]
// ====================================

// Get : Function to return the configuration outside of this package
func Get() map[string]string {
	return config
}
