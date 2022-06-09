package stubgen

import (
	"path/filepath"
	"syscall"

	"github.com/HaswinVidanage/gqlgen/internal/code"

	"github.com/HaswinVidanage/gqlgen/codegen"
	"github.com/HaswinVidanage/gqlgen/codegen/config"
	"github.com/HaswinVidanage/gqlgen/codegen/templates"
	"github.com/HaswinVidanage/gqlgen/plugin"
)

func New(filename string, typename string) plugin.Plugin {
	return &Plugin{filename: filename, typeName: typename}
}

type Plugin struct {
	filename string
	typeName string
}

var _ plugin.CodeGenerator = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}

func (m *Plugin) Name() string {
	return "stubgen"
}

func (m *Plugin) MutateConfig(cfg *config.Config) error {
	_ = syscall.Unlink(m.filename)
	return nil
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	pkgPath := code.ImportPathForDir(filepath.Dir(m.filename))
	pkgName := code.NameForPackage(pkgPath)

	return templates.Render(templates.Options{
		PackageName: pkgName,
		Filename:    m.filename,
		Data: &ResolverBuild{
			Data:     data,
			TypeName: m.typeName,
		},
		GeneratedHeader: true,
	})
}

type ResolverBuild struct {
	*codegen.Data

	TypeName string
}
