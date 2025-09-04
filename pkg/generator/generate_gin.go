package generator

import "github.com/Ayomits/go-boilerplate/pkg/generator/utils"

type GenerateGinProject struct{}

func NewGinProjectGenerator() utils.ProjectGenerator {
	return &GenerateGinProject{}
}

func (g *GenerateGinProject) Generate() {
	utils.GeneratTemplateSafe(utils.GinType)
}
