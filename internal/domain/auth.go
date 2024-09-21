package domain

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    // Add other relevant fields
}
