package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func login(uid, pwd string) error {
	if uid == "" {
		return errors.New("login error")
	} else {
		return nil
	}
}

func getData(file string) ([]byte, error) {
	return []byte{}, nil
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

//func main() {
//	data, err := LoginAndGetData("", "pwd", "file")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(data)
//}

func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{}
	}
	return genErr
}

//func main() {
//	err := GenerateError(true)
//	fmt.Println(err != nil)
//	err = GenerateError(false)
//	fmt.Println(err != nil)
//}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker!? %w", err)
	}
	f.Close()
	return nil
}

//func main() {
//	err := fileChecker("not_here.txt")
//	if err != nil {
//		fmt.Println(err)
//		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
//			fmt.Println(wrappedErr)
//		}
//	}
//}

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}
