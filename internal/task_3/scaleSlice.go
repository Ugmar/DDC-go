package task3

import "math"

func ScaleSlice(slice *[]int, scaleFactor uint32) error{
	if slice == nil{
		return nil
	}
	if uint64(scaleFactor) * uint64(len(*slice)) > math.MaxUint32{
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
