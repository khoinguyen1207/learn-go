package template

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sync"
)

type ForgotPasswordTemplateData struct {
	Username  string
	ResetLink string
}

type templateService struct {
	dir   string
	cache map[string]*template.Template
	mu    sync.RWMutex
}

func NewTemplateService(templateDir string) TemplateService {
	return &templateService{
		dir:   templateDir,
		cache: make(map[string]*template.Template),
	}
}

func (t *templateService) Render(name string, data any) (string, error) {
	tmpl, err := t.load(name)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("template: failed to render %q: %w", name, err)
	}

	return buf.String(), nil
}

func (t *templateService) load(name string) (*template.Template, error) {
	// Check cache first
	t.mu.RLock()
	tmpl, ok := t.cache[name]
	t.mu.RUnlock()
	if ok {
		return tmpl, nil
	}

	path := filepath.Join(t.dir, name)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("template: file %q not found: %w", path, err)
	}

	tmpl, err = template.New(name).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("template: failed to parse %q: %w", name, err)
	}

	// Store in cache
	t.mu.Lock()
	t.cache[name] = tmpl
	t.mu.Unlock()

	return tmpl, nil
}
