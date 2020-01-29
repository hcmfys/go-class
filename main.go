package main

import (
	"springbus.org/classReader/cls"
)

func main() {

	path := "/Volumes/f/ApiCode/target/classes/org/springbus/API.class"
	c := cls.ClassReader{}
	c.ReadClass(path)

}
