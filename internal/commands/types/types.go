package types

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type Location struct {
	Name string `json:"name"`
	//Url  string `json:"url"`
}
