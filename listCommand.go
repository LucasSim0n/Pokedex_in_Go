package main

func commandList(cfg *config, name string) error {

	err := cfg.apiClient.ListPokemon()
	return err
}
