package io

// NullWriter defines a writer that discards everything
type NullWriter struct{}

// Write implements the io.Writer interface
func (w *NullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
