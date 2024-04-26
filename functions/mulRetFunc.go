package main

// for multiple param of same type,
// func funcName(param_1,..n paramType) returnType { return something}
func mulRet(a, b int) (int, string) {
	if a == 0 || b == 0 {
		res, err := 0, "one of the numbers are zero!"
		return res, err
	}
	return a * b, "Done!"
}
