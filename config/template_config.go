package config

type TemplateConfigI interface {
	Config() TemplateConfig
}

type TemplateConfigImp struct{}

type TemplateConfig struct {
	Title   string
	Content string
}

func GetTemplateConfig() TemplateConfigI {
	return &TemplateConfigImp{}
}

func (t *TemplateConfigImp) Config() TemplateConfig {
	return TemplateConfig{
		Title:   "",
		Content: "",
	}
}
