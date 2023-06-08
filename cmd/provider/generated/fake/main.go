package main

import (
	"fmt"

	provider "github.com/external-secrets/external-secrets-provider-fake"
)

//go:generate ./generate.sh $GOFILE
func main() {
	p := provider.Provider{}
	fmt.Printf("provider cap: %#v\n", p.Capabilities())
}