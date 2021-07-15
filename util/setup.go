package util

func Setup() error {
	if err := setupZsh(); err != nil {
		return err
	}

	if err := setupPluginLocaton(); err != nil {
		return err
	}
	return nil
}
