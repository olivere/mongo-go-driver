package description

// TopologyKind represents a specific topology configuration.
type TopologyKind uint32

// These constants are the available topology configurations.
const (
	Single                TopologyKind = 1
	ReplicaSet            TopologyKind = 2
	ReplicaSetNoPrimary   TopologyKind = 4 + ReplicaSet
	ReplicaSetWithPrimary TopologyKind = 8 + ReplicaSet
	Sharded               TopologyKind = 256
)

// String implements the fmt.Stringer interface.
func (kind TopologyKind) String() string {
	switch kind {
	case Single:
		return "Single"
	case ReplicaSet:
		return "ReplicaSet"
	case ReplicaSetNoPrimary:
		return "ReplicaSetNoPrimary"
	case ReplicaSetWithPrimary:
		return "ReplicaSetWithPrimary"
	case Sharded:
		return "Sharded"
	}

	return "Unknown"
}