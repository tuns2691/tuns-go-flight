package customer_request

import "github.com/google/uuid"

type CustomerRequest struct {
	Id             uuid.UUID
	Role           int32
	Name           string
	Email          string
	PhoneNumber    string
	DateOfBith     string
	IdentityCard   string
	Address        string
	MembershipCard string
	Password       string
	Status         int32
}
