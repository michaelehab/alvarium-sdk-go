package contracts

type LayerType string

const (
	ApplicationLayer LayerType = "app"
	CiCdLayer        LayerType = "cicd"
	OsLayer          LayerType = "os"
	HostLayer        LayerType = "host"
)

func (l LayerType) Validate() bool {
	switch l {
	case ApplicationLayer, CiCdLayer, OsLayer, HostLayer:
		return true
	default:
		return false
	}
}
