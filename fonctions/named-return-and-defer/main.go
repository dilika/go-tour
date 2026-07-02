package main

import "fmt"


func processFile(filename string) (err error){
	file, err := openFile(filename)
	if err != nil {
		return err
	}

	defer file.close()

	fmt.Println("File processing in progress...")

	return nil
}


func openFile(name string) (*FakeFile, error) {
	return &FakeFile{}, nil
}

func (f *FakeFile) close() error {
	return nil
}

type  FakeFile  struct {}


func main(){
	processFile("data.txt")
}
