func GetKeyStructure(url string) string {
baseKey := "go.response"
basePath := strings.Split(url, "/GF")[0]
keyBasePath := strings.Replace(basePath, "/", ".", len(basePath))
key := baseKey + keyBasePath
return key
}
