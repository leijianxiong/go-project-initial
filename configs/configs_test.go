package configs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDefault(t *testing.T) {
	conf := Default()
	s, _ := json.MarshalIndent(conf, "", "\t")
	fmt.Println("default conf:", string(s))
}

func TestParse(t *testing.T) {
	_ = Parse()
	s, _ := json.MarshalIndent(Conf, "", "\t")
	fmt.Println("parse conf:", string(s))
}

func TestProjectDir(t *testing.T) {
	fmt.Println("project dir:", ProjectDir())
}
