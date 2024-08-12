package pokeapi

type RespShallowLocations struct {
	Count int `json:"count"`
	Next int `json:"next"`
	Previous int `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}