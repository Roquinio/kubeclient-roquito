package main

import (
	"os/exec"
	"testing"
)

func TestEmptyCMD(t *testing.T) {
	exec.Command("./roquito")
}

func TestHelpCMD(t *testing.T) {
	exec.Command("./roquito -h")
}

func TestGetVerbs(t *testing.T) {
	verbs := []string{"ns", "namespaces", "dp", "deployment", "pods", "p", "svc", "services"}
	cmd := "./roquito get "
	for _, s := range verbs {
		exec.Command(cmd + s)
		// fmt.Println(cmd + s)
	}

}
