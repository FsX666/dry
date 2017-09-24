package docker

import (
	"sort"

	"github.com/docker/docker/api/types"
)

//Allowed sort methods
const (
	NoSortNetworks SortMode = iota
	SortNetworksByID
	SortNetworksByName
	SortNetworksByDriver
)

type dockerNetworks []types.NetworkResource

func (s dockerNetworks) Len() int      { return len(s) }
func (s dockerNetworks) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type networksByID struct{ dockerNetworks }

func (s networksByID) Less(i, j int) bool { return s.dockerNetworks[i].ID < s.dockerNetworks[j].ID }

type networksByName struct{ dockerNetworks }

func (s networksByName) Less(i, j int) bool {
	if len(s.dockerNetworks[i].Name) > 0 {
		if len(s.dockerNetworks[j].Name) > 0 {
			return s.dockerNetworks[i].Name[0] < s.dockerNetworks[j].Name[0]
		}
		return true
	}
	return false
}

type networksByDriver struct{ dockerNetworks }

func (s networksByDriver) Less(i, j int) bool {
	return s.dockerNetworks[i].Driver < s.dockerNetworks[j].Driver
}

//SortNetworks sorts the given network slice using the given mode
func SortNetworks(networks []types.NetworkResource, mode SortMode) {
	switch mode {
	case SortNetworksByID:
		sort.Sort(networksByID{networks})
	case SortNetworksByName:
		sort.Sort(networksByName{networks})
	case SortNetworksByDriver:
		sort.Sort(networksByDriver{networks})
	}
}
