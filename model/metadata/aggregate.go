package metadata

import (
	"time"

	"github.com/mishudark/eventhus"
)

// TODO: Do we need to put Aggregate, Command, Event in the type name?
// TODO: Is it appropriate to just call it SiteAggregate?

type SiteMetadataAggregate struct {
	eventhus.BaseAggregate
	Name string
	Description string
	DateInstalled time.Time
}

type CreateSiteMetadataCommand struct {
	eventhus.BaseCommand
}

type SiteMetadataCreatedEvent struct {
}

// ?
type SetupAggregate struct {
	eventhus.BaseAggregate
}
