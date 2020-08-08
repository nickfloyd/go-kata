package iokata

import "io/ioutil"

func WriteFile() error {
	data := []byte("go\nkata\n")
	err := ioutil.WriteFile("temp/data", data, 0644)
	return err
}

func ReadFile() {

}
