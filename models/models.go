package models

// User represents a user entity
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	// Add more fields as needed
}

// CacheData represents data that can be cached
type CacheData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	// Add more fields as needed
}
