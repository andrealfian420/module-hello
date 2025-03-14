package modulehello

// gunakan awalan huruf depan pada fungsi atau variabel
// agar bisa diakses secara publik
func Greeting() string {
	return "Howdy!"
}

func GreetByName(name string) string {
	return "Howdy" + name + "!"
}
