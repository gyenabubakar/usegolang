package main

type TemplateWriter struct {
	parsedValue []byte
}

func (tw *TemplateWriter) Write(p []byte) (int, error) {
	tw.parsedValue = p
	return len(p), nil
}