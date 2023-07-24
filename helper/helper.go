package helper

func PanicIfError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
