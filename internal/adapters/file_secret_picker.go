package adapters

type IRandomIntPicker interface {
	PickRandomInt(upto int) int
}

type FileSecretPicker struct {
	filePath        string
	randomIntPicker IRandomIntPicker
}

func (f FileSecretPicker) PickSecret() (string, error) {
	words, err := ReadWordsFromFile(f.filePath)
	if err != nil {
		return "", err
	}
	index := f.randomIntPicker.PickRandomInt(len(words))
	return words[index], nil
}

func NewFileSecretPicker(filePath string, randomIntPicker IRandomIntPicker) FileSecretPicker {
	return FileSecretPicker{filePath, randomIntPicker}
}
