// Package docstraffic declares the Docs consumer's traffic vocabulary.
package docstraffic

import (
	"strings"

	"github.com/yueli-official/foundation/go/traffic"
)

const ResourceDocument traffic.ResourceKind = "document"

func Definition(timeZone string) traffic.Definition {
	return traffic.Definition{
		Version:  traffic.DefinitionVersion,
		TimeZone: strings.TrimSpace(timeZone),
		ResourceKinds: []traffic.ResourceKindDefinition{
			{Key: ResourceDocument},
		},
	}
}
