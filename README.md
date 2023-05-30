# go-servicenow

Go SDK for ServiceNow.

This SDK has been created primarily for use with the [Steampipe ServiceNow plugin](https://github.com/turbot/steampipe-plugin-servicenow), which only uses GET methods from the ServiceNow API. There are no current plans to add support for DELETE/PATCH/POST methods.

## Getting started

### Requirements

- Go 1.19+

### Installation

```shell
go get github.com/turbot/go-servicenow
```

### Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbot/go-servicenow/servicenow"
)

func main() {

	// The URL of ServiceNow instance
	instanceURL := "https://dev129225.service-now.com"

	// Method of authentication. Currently, SDK only supports password grant type
	grantType := "password"

	// User's username and password
	username := "admin"
	password := "j0t%ldweqwd3%"

	// Client credentials
	// See: https://support.servicenow.com/kb?id=kb_article_view&sysparm_article=KB0725643
	clientID := "9148ce343214xewqR0392c96f819dbd422"
	clientSecret := "#$sf3EauTd"

	// Initialize a new client
	client, err := servicenow.New(servicenow.Config{
		InstanceURL:  instanceURL,
		GrantType:    grantType,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
	})
	if err != nil {
		log.Fatalf("failed to initialize the ServiceNow API client: %+v", err)
	}

	// Example: List ten contacts
	limit := 10
	offset := 0
	contacts, err := client.NowContact.List(limit, offset)
	if err != nil {
		log.Fatalf("failed listing contacts: %+v", err)
	}
	for _, contact := range contacts.Result {
		fmt.Println(contact.Name)
	}
}
```

## Feedback

### Contributing

Please open pull requests for fixes or additions to this SDK. Please also see our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-aws/blob/main/LICENSE).

### Raise an issue

To provide feedback or report a bug, [please raise an issue on our issue tracker](https://github.com/turbot/go-servicenow/issues).
