package pokeapi

import "fmt"

func (c *Client) EncounterMethods(limit int, offset int) (v NamedAPIResourceList, err error) {
	err = c.doUncached(&v, fmt.Sprintf("encounter-method?limit=%d&offset=%d", limit, offset))
	return
}

func (c *Client) EncounterMethod(nameOrIdOrUrl string) (v EncounterMethod, err error) {
	err = c.do(&v, fmt.Sprintf("encounter-method/%s", nameOrIdOrUrl))
	return
}

func (c *Client) EncounterConditions(limit int, offset int) (v NamedAPIResourceList, err error) {
	err = c.doUncached(&v, fmt.Sprintf("encounter-condition?limit=%d&offset=%d", limit, offset))
	return
}

func (c *Client) EncounterCondition(nameOrIdOrUrl string) (v EncounterCondition, err error) {
	err = c.do(&v, fmt.Sprintf("encounter-condition/%s", nameOrIdOrUrl))
	return
}

func (c *Client) EncounterConditionValues(limit int, offset int) (v NamedAPIResourceList, err error) {
	err = c.doUncached(&v, fmt.Sprintf("encounter-condition-value?limit=%d&offset=%d", limit, offset))
	return
}

func (c *Client) EncounterConditionValue(nameOrIdOrUrl string) (v EncounterConditionValue, err error) {
	err = c.do(&v, fmt.Sprintf("encounter-condition-value/%s", nameOrIdOrUrl))
	return
}

type EncounterMethod struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A good value for sorting.
	Order int `json:"order"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type EncounterCondition struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of possible values for this encounter condition.
	Values []NamedAPIResource `json:"values"`
}

type EncounterConditionValue struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The condition this encounter condition value pertains to.
	Condition NamedAPIResource `json:"condition"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}
