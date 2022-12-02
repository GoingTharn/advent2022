package problems

import "os"

func getInput(path string) ([]byte, error) {
	//Get the response bytes from the url
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return contents, err
}
