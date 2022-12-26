package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func writeFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("error while marshling message: ", err)
		return
	}
	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("error while writing file: ", err)
		return
	}
	fmt.Println("message has been written on file")
}

func readFile(fname string, pb proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("error while reading file: ", err)
		return
	}
	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("error while read file: ", err)
		return
	}
}
