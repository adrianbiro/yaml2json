package main

import (
	"fmt"
	"io"
	"os"

	"sigs.k8s.io/yaml"
)

// convert YAML from stdin to JSON
func main() {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	if len(stdin) < 1 {
		fmt.Printf("Usage:\n\tyaml_from_stdin | %s\n", os.Args[0])
		return
	}
	y, err := yaml.JSONToYAML(stdin)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println(string(j2))

}
