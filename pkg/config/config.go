package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// app config holds the app config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Production    bool
	Session       *scs.SessionManager
}
