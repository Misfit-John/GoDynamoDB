package main

import GoDynamoDB "./src"
import "os"
import "io/ioutil"

func main() {
	GoDynamoDB.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	GoDynamoDB.Trace.Println("I have something standard to say")
	GoDynamoDB.Info.Println("Special Information")
	GoDynamoDB.Warning.Println("There is something you need to know about")
	GoDynamoDB.Error.Println("Something has failed")
}
