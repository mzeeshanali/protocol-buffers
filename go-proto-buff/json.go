package main

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func ToJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("error while marshling message to json: ", err)
		return ""
	}

	return string(out)
}

func FromJson(in string, pb proto.Message) {
	opt := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := opt.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("error while unmarshal json to message: ", err)
		return
	}
}
