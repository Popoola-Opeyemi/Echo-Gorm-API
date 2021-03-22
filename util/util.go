package util

func getError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
