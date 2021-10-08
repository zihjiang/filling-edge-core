
package common

const (
	CurrentOffsetVersion = 2
	PollSourceOffsetKey  = "$com.streamsets.sdc2go.pollsource.offset$"
)

var emptyOffset = ""

type SourceOffset struct {
	Version int
	Offset  map[string]*string
}

func GetDefaultOffset() SourceOffset {
	return SourceOffset{
		Version: CurrentOffsetVersion,
		Offset:  map[string]*string{PollSourceOffsetKey: &emptyOffset},
	}
}
