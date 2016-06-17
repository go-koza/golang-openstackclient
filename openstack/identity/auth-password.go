package identity

import (
    "errors"

    "github.com/go-koza/golang-openstackclient/openstack/utils"
)

type Credentials struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

type Auth struct {
    Credentials `json:"passwordCredentials"`
    Project string  `json:"tenantName"`
}

type AuthCreds struct {
    Auth    `json:"auth"`
}


func newAuthCreds (creds *Creds) (authCreds AuthCreds, err error) {
    //Validate creds
    err = nil
    authCreds = AuthCreds{
        Auth: Auth{
        Credentials:    Credentials{
            Username:   creds.Username,
            Password:   creds.Password,
        },
        Project:    creds.Project,
        },
    }

    return authCreds, err
}

func (creds *Creds) getToken (url string) (IdentityIFace, error) {
    token := IdentityToken{}

    authCreds, err := newAuthCreds(creds)
    if err != nil {
        err = errors.New("Failed to get auth options")
        return nil, err
    }

    err = utils.PostJSON(url + "/tokens", &authCreds, &token)
    if err != nil {
        return nil, err
    }

    return &token, nil
}

