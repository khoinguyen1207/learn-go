package template

type TemplateService interface {
	Render(name string, data any) (string, error)
}
