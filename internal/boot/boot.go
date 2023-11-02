package boot

var bootstraps = []Booter{
	Migration{},
}

func Boot() error {
	for _, bootstrap := range bootstraps {
		if err := bootstrap.Boot(); err != nil {
			return err
		}
	}

	return nil
}
