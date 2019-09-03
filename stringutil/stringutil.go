package stringutil

// multiples functions

// FullName comment
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

// FullNameNakedReturn Commentfmt
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
