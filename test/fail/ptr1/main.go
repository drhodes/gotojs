package main

func ptr_square(n_ptr *int) {
	if n_ptr == nil {
		return
	}
	val := *n_ptr
	*n_ptr = val * val
}

func main() {	
	n := new(int)
	*n = 4
	ptr_square(n)
	console.log(*n)
}
