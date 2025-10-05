package task3

func ScaleSlice(slice *[]int, scaleFactor uint32) error{
	if scaleFactor * uint32(len(*slice)) < uint32(len(*slice)){
		return ErrOverflow
	}

	res_len := scaleFactor * uint32(len(*slice))

	tmp_slice := make([]int, res_len)

	for i := 0; i < int(res_len); i++{
		tmp_slice[i] = (*slice)[i % len(*slice)]
	}

	*slice = tmp_slice

	return nil
}
