package main

import "errors"

func validation(i *info) (*info, error) {
	if i.user == "" {
		return nil, errors.New("user is empty")
	}
	if i.pass == "" {
		return nil, errors.New("pass is empty")
	}
	return i, nil

}

type info struct {
	user string
	pass string
}

func main() {
	i := &info{
		user: "",
		pass: "",
	}
	_, err := validation(i)
	if err != nil {
		panic(err)
	}

}
