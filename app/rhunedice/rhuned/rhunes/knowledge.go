package rhunes

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/pkg/tools"
)

func NewKnowledgeRhune(value int) *api.Rhune {
	knowledgeRhune := api.NewRhune(
		api.KnowledgeName,
		api.KnowledgeShort,
		"Knowledge increase",
		api.ExtraRhune,
		api.ExtraBucket,
		func(avatar api.IAvatar) (any, error) {
			tools.Logger.WithField("module", "rhunes").
				WithField("rhune", "knowledge").
				WithField("method", "NewKnowledge").
				Tracef("add %d knowledge to %s", value, avatar.GetName())

			avatar.GetKnowledge().Inc(value)
			return nil, nil
		},
	)
	return knowledgeRhune
}
