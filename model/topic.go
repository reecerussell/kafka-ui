package model

// Topic represents the structure of topic data in the config store.
type Topic struct {
	Name        string  `json:"name"`
	DisplayName *string `json:"displayName,omitempty"`
}
