package randlib

func Choice(list []any) any {
	if (len(list) == 0) {
		return nil
	}

	return list[RandInt(0, len(list) - 1)]
}
