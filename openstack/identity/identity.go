package identity

import (
    "fmt"
    "strings"
)

//Version defaults to V2
const VERSION = "2.0"

type Creds struct {
    Project string
    Username    string
    Password    string
}

type Token struct {
    Project string
    Token   string
}

func (creds *Creds) GetClient (url string, version string) (idf IdentityIFace, err error) {
    if version == "" {
        version = VERSION
    }

    if strings.HasPrefix(version, "2") {
        idf, err = creds.getToken(url)
        return
    } else if strings.HasPrefix(version, "3") {
        fmt.Println("Not Implemented V3 API")
        panic("Not Implemented")
    } else {
        panic("Unsupported versions. Identity supports 2.x or 3.x")
    }
}

func (token *Token) GetClient (url string, version string) (idf IdentityIFace) {
    panic("Not Implemented")
}
