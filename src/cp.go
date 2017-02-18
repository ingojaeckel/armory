package main

func cp(args []pair) error {
	for _, arg := range args {
		if err := execute([]string{}, "cp", "-rv", arg.Key, arg.Value); err != nil {
			return err
		}
	}
	return nil
}
