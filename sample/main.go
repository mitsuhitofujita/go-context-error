package main

import (
	"fmt"
	"errors"

	"./error"
)

func main() {
	fmt.Println("Hello, world.")

	err1 := error.NewContextError(error.Context{"entry_id": 123}, errors.New("error message"))
	err2 := error.NewContextError(error.Context{"user_code": "u456"}, err1)
	fmt.Printf("err! %s\n", err2.Error())



	/*
	h := sha1.New()
	m := map[string]bool{}

	for i := 0; i < 100000000; i++ {
		s := fmt.Sprintf("%dhoge", i)
		h.Write([]byte(s))
		x := fmt.Sprintf("%x", h.Sum(nil))
		if _, ok := m[x]; ok {
			fmt.Printf("dup!\n")
		}
		m[x] = true
		if (i % 1000000) == 0 {
			fmt.Printf(".")
		}
		//fmt.Printf("%s\n", x)
	}
	fmt.Printf("\n")
	 */

	/*
	glob.Get(".", "*.go")

	err := glob.Walk(".", "*.go", func(path string, info os.FileInfo) error {
		fmt.Printf("walk: %q\n", info.Name())
		return nil
	})
	if err != nil {
		panic(1)
	}

	_, tmpfileErr := tmpfile.Make(".", "test", func(path string, file *os.File) error {
		fmt.Printf("make: %q\n", path)
		_, err = file.Write([]byte("This is temp file.\n"))

		return err
	})
	if tmpfileErr != nil {
		panic(2)
	}
	 */
}
