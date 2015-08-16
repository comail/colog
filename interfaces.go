package colog

// Hook is the interface to be implemented by event hooks
type Hook interface {
	Levels() []Level   // returns the set of levels for which the hook should be triggered
	Fire(*Entry) error // triggers the hook, this function will be called for every eligible log entry
}

// Formatter interface must be implemented by message formatters
// Format(*Entry) will be called and the resulting bytes sent to output
type Formatter interface {
	Format(*Entry) ([]byte, error) // The actual formatter called every time
	DisableColors()                // DisableColors is called when the output doesn't support them
	SetFlags(flags int)            // Like the standard log.SetFlags(flags int)
	Flags() int                    // Like the standard log.Flags() int
}

// Extractor interface must be implemented by data extractors
// the extractor reads the message and tries to extract key-value
// pairs from the message and sets the in the entry
type Extractor interface {
	Extract(*Entry) error
}
