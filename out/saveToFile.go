package out

import "os"

func writeToFile(fileName string, data []byte) {
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)

	if err != nil {
		panic(err)
	}

	defer file.Close()
}
