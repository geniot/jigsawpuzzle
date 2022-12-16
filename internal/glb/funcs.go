package glb

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func IsShufflePerfect(arr []int, rowLength int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == 1 {
			return false
		}
		if i < len(arr)-rowLength {
			if arr[i+rowLength]-arr[i] == rowLength {
				return false
			}
		}
	}
	return true
}
