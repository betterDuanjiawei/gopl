package main

func main()  {

}

//func noneEmpty(x []string) []string {
//	for k, v := range x {
//		if v != "" {
//			x[k] = v
//		}
//	}
//
//	return x
//}


func noneEmpty(x []string) []string {
	i := 0
	for _, v := range x {
		if v != "" {
			x[i] = v
			i++
		}
	}

	return x[:i]
}

func noneEmpty2(strings []string) []string  {
	out := strings[:0]
	for _,k := range strings {
		if k != "" {
			out = append(out, k)
		}
	}

	return out
}