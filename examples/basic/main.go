package main

import (
	"fmt"
	"os"

	apiclient "github.com/minchao/go-ptx/basic/v2/client"
	"github.com/minchao/go-ptx/basic/v2/client/basic"
	"github.com/minchao/go-ptx/pkg/auth"
	"github.com/minchao/go-ptx/pkg/transport"
)

func main() {
	tp := transport.New()
	tp.DefaultAuthentication = auth.NewAuthentication(os.Getenv("APP_ID"), os.Getenv("APP_KEY"))
	client := apiclient.New(tp, nil)

	params := basic.NewBasicAPIAuthorityParams().
		WithDollarFormat("JSON")
	result, err := client.Basic.BasicAPIAuthority(params)
	if err != nil {
		panic(err)
	}
	for _, authority := range result.Payload {
		fmt.Printf("AuthorityID: %d\n", authority.AuthorityID)
		fmt.Printf("  AuthorityName: %s\n", authority.AuthorityName)
	}
}
