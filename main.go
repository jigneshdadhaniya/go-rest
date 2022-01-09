package main

func main() {
	a := App{}
	a.Initialize("", "mysecretpassword", "postgres")

	a.Run(":8080")
}
