package alphabet

const _alphabetLetter_name = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type AlphabetLetter int32

var (
	Alphabet map[string]AlphabetLetter
)

func init() {
	Alphabet = make(map[string]AlphabetLetter, len(_alphabetLetter_name))
	for i, v := range _alphabetLetter_name {
		Alphabet[string(v)] = AlphabetLetter(i + 1)
	}
}
