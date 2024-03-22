package main

const englishPrefixHello = "Hello"
const spanishPrefixHello = "Hola"
const frenchPrefixHello = "Bonjour"

func greetPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = spanishPrefixHello
	case "French":
		prefix = frenchPrefixHello
	default:
		prefix = englishPrefixHello
	}
	return
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetPrefix(lang) + ", " + name
}

// func main() {
// 	fmt.Println(Hello("World", ""))
// }
