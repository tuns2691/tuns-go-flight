package response

type CustomerResponse struct {
	Id             string `json:"id"`
	Role           int32  `json:"role"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	DateOfBith     string `json:"dateOfBith"`
	IdentityCard   string `json:"identityCard"`
	Address        string `json:"address"`
	MembershipCard string `json:"membershipCard"`
	Password       string `json:"password"`
	Status         int32  `json:"status"`
}
