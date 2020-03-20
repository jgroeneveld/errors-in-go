package main

func main() {

}

type ValidationError struct {
}

func (err ValidationError) Error() string {
	return "a validation error"
}
