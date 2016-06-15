package identity

type AccessType struct {
    Token   Token   `json:"token"`
    User    interface{} `json:"id"`
    ServiceCatalog  []ServiceCatalogEntry   `json:"servicecatalog"`
}

type AuthToken struct {
    Access  AccessType  `json:"access"`
}

type Token struct {
    ID  string  `json:"id"`
    Expires time.Time   `json:"expires"`
    Project struct  {
        ID  string  `json:"id"`
        Name    string  `json:"name"`
    }   `json:"tenant"`
}

func (t *AuthToken) GetEndpoint(service string, region string) (string, error) {
    fmt.Println("Not Implemented")
}

func (t *AuthToken) GetExpiration() time.Time {
    fmt.Println("Not Implemented")
}

type AuthIFace interface {
    GetEndpoint()   string
    GetExpiration() time.Time
    //Create calls to follow
}

