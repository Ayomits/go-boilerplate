package generator

import "github.com/Ayomits/go-boilerplate/pkg/generator/utils"

type GenerateFiberProject struct{}

func NewFiberProjectGenerator() utils.ProjectGenerator {
	return &GenerateFiberProject{}
}

func (g *GenerateFiberProject) Generate() {
	utils.GenerateTemplate(utils.FiberType)
}
