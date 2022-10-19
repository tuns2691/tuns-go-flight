package request

type CreateCustomerRequest struct {
	Role           int32  `json:"role" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	PhoneNumber    string `json:"phoneNumber" binding:"required"`
	DateOfBith     string `json:"dateOfBith" binding:"required"`
	IdentityCard   string `json:"identityCard" binding:"required"`
	Address        string `json:"address" binding:"max=256,min=6"`
	MembershipCard string `json:"membershipCard"`
	Password       string `json:"password"`
	Status         int32  `json:"status" binding:"required"`
}
