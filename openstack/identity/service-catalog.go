package identity

type ServiceCatalogEntry struct {
    Name    string  `json:"name"`
    Type    string  `json:"type"`
    Endpoints   []ServiceEndpoint   `json:"endpoints"`
    // Endpoints []map[string]string `json:"endpoints"`
}

type ServiceEndpoint struct {
    Type    string  `json:"type"`
    Region  string  `json:"region"`
    PublicURL   string  `json:"publicurl"`
    AdminURL    string  `json:"adminurl"`
    InternalURL string  `json:"internalurl"`
    VersionID   string  `json:"versionid"`
}

