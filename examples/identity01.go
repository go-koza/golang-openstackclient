package main

import (
    "fmt"

    "github.com/go-koza/golang-openstackclient/openstack/identity"
)

func main() {

    creds := identity.Creds{
        Username: "admin",
        Password: "stack",
        Project: "admin",
    }

    client, _ := creds.GetClient("http://10.0.2.15:5000/v2.0", "")   //Version defaults to 2.0
    fmt.Println(client)
}
