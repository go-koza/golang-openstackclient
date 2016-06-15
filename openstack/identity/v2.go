package identity

import "errors"

type Credentials struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

type Auth struct {
    Credentials `json:"passwordCredentials"`
    Project     `json:"tenantName"`
}

type V2Payload struct {
    //AuthUrl string  `json:"-"`
    Auth            `json:"auth"`
}

func (creds *AuthCreds) newV2Payload() (V2Payload, error) {
    //Validate creds
    payload := V2Payload{
        Auth: Auth{
            Credentials:    Credentials{
                Username:   creds.Username
                password:   creds.Password
            }
            Project:    creds.Project
        }
    }
    return payload


func getV2Token(creds *AuthCreds) (AuthIFace, error) {
    token := AuthToken{}

    v2payload, err := newV2Payload(creds)
    if err != nil {
        err = errors.New("Failed to get auth options")
        return nil, err
    }

    _, err := POST(creds.AuthUrl + "/tokens", &v2payload, &token)
    if err != nil {
        return nil, err
    }

    return token
}
