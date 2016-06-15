package identity

import (
    "fmt"
    "strings"
    "time"
)

//Version defaults to V2
var VERSION = "2.0"

type Auth struct {
    // AuthUrl is always required
    AuthUrl string
    // Project required for all scoped lookups
    Project string
}

type AuthCreds struct {
    Auth
    Version     string
    //Username/Password required only for password based auth
    Username    string
    Password    string
}


func (ac *AuthCreds) GetClient (Version string) {
    if Version == nil {
        //Assign to default versin in nil
        Version := VERSION
    }

    if strings.HasPrefix(Version, "2") {
        return getV2Token(ac)
    } else if strings.HasPrefix(Version, "3") {
        fmt.Println("Not Implemented V3 API")
        return nil
    } else {
        panic("Unsupported versions. Identity supports 2.x or 3.x")
    }
}
