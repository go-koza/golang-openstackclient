package identity

import (
    "fmt"
    "time"
)

type AccessToken struct {
    ID  string  `json:"id"`
    Expires time.Time   `json:"expires"`
    Project struct  {
        ID  string  `json:"id"`
        Name    string  `json:"name"`
    }   `json:"tenant"`
}

type Access struct {
    AccessToken   AccessToken   `json:"token"`
    User    interface{} `json:"id"`
    ServiceCatalog  []ServiceCatalogEntry   `json:"servicecatalog"`
}

type IdentityToken struct {
    Access  Access  `json:"access"`
}


func (t *IdentityToken) GetEndpoint() string {
    fmt.Println("Not Implemented")
    return ""
}

func (t *IdentityToken) GetExpiration() time.Time {
    fmt.Println("Not Implemented")
    return time.Time{}
}


type IdentityIFace interface {
    GetEndpoint()   string
    GetExpiration() time.Time
    //Create calls to follow
}

