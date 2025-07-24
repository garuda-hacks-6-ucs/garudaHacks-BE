package domain

import "time"

// Profile represents the core user entity in the system.
type Profile struct {
	WalletAddress string      `bson:"_id"`
	Role          string      `bson:"role"`
	IsActive      bool        `bson:"is_active"`
	Details       interface{} `bson:"details"`
	CreatedAt     time.Time   `bson:"created_at"`
	UpdatedAt     time.Time   `bson:"updated_at"`
	DeletedAt     *time.Time  `bson:"deleted_at,omitempty"`
}

// GovernmentDetails contains specific fields for a government profile.
type GovernmentDetails struct {
	AgencyName       string `bson:"agency_name" json:"agency_name"`
	AgencyType       string `bson:"agency_type" json:"agency_type"`
	OrganizationCode string `bson:"organization_code" json:"organization_code"`
	OfficeAddress    string `bson:"office_address" json:"office_address"`
	ContactEmail     string `bson:"contact_email" json:"contact_email"`
	ContactNumber    string `bson:"contact_number" json:"contact_number"`
}

// VendorDetails contains specific fields for a vendor profile.
type VendorDetails struct {
	CompanyName     string `bson:"company_name" json:"company_name"`
	NIB             string `bson:"nib" json:"nib"`
	NPWP            string `bson:"npwp" json:"npwp"`
	OfficeAddress   string `bson:"office_address" json:"office_address"`
	DomicileAddress string `bson:"domicile_address" json:"domicile_address"`
	ContactEmail    string `bson:"contact_email" json:"contact_email"`
	ContactNumber   string `bson:"contact_number" json:"contact_number"`
}

// CitizenDetails contains specific fields for a citizen profile.
type CitizenDetails struct {
	FullName        string `bson:"full_name" json:"full_name"`
	KTPIdentityHash string `bson:"ktp_identity_hash" json:"ktp_identity_hash"`
	IsVerified      bool   `bson:"is_verified" json:"is_verified"`
}
