package pokeapi

import "fmt"

func (c *Client) Machines(limit int, offset int) (v NamedAPIResourceList, err error) {
	err = c.doUncached(&v, fmt.Sprintf("machine?limit=%d&offset=%d", limit, offset))
	return
}

func (c *Client) Machine(idOrUrl string) (v Machine, err error) {
	err = c.do(&v, fmt.Sprintf("machine/%s", idOrUrl))
	return
}

type Machine struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The TM or HM item that corresponds to this machine.
	Item NamedAPIResource `json:"item"`
	// The move that is taught by this machine.
	Move NamedAPIResource `json:"move"`
	// The version group that this machine applies to.
	Version_group NamedAPIResource `json:"version_group"`
}
