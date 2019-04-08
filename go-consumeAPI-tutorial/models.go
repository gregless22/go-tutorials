package main

type Response struct {
	Name    string
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int     `json:"entry_number"`
	Species Species `json:"pokemon_species"`
}

type Species struct {
	Name string `json:"name"`
}
