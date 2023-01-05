package pkg

import (
	"bytes"
	"io"
	"log"
	"os"
)

// Function to make sure that the file exist
func CheckRequiredFiles(path string, filename []string) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, name := range filename {
		_, err := os.Stat(path + name)

		// if path + fileName does not exist
		if os.IsNotExist(err) {
			log.Fatalf("%s is missing!", path+name)
		}
	}
}

// Function to check if the file exists and
// proceed to create the file if it doesn't exist
func CheckIfFilesExistAndCreateFile(path string, filename []string) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, name := range filename {
		_, err := os.Stat(path + name)

		// If path + fileName does not exist
		if os.IsNotExist(err) {
			log.Printf("%s not found! Now proceeding to create %s", path+name, path+name)
			// Create the file
			_, err := os.Create(path + name)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Successfully created %s!", path+name)
		}
	}
}

// Function to open the file and return the data inside
func OpenFile(path string, filename string) ([]byte, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Open the file
	file, err := os.Open(path + filename)
	if err != nil {
		log.Fatal(err)
	}

	// Successfully opened the file
	log.Printf("Successfully opened the file: %s", path+filename)
	defer file.Close()

	var buf bytes.Buffer
	io.Copy(&buf, file)

	return buf.Bytes(), err
}
