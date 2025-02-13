// Code generated by ent, DO NOT EDIT.

package ent

import (
	"core/ent/account"
	"core/ent/order"
	"core/ent/schema"
	"core/ent/station"
	"core/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescUsername is the schema descriptor for username field.
	accountDescUsername := accountFields[0].Descriptor()
	// account.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	account.UsernameValidator = accountDescUsername.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[1].Descriptor()
	// account.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	account.PasswordValidator = accountDescPassword.Validators[0].(func(string) error)
	// accountDescToken is the schema descriptor for token field.
	accountDescToken := accountFields[2].Descriptor()
	// account.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	account.TokenValidator = accountDescToken.Validators[0].(func(string) error)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[3].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[4].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescMaintOrderNo is the schema descriptor for maintOrderNo field.
	orderDescMaintOrderNo := orderFields[0].Descriptor()
	// order.MaintOrderNoValidator is a validator for the "maintOrderNo" field. It is called by the builders before save.
	order.MaintOrderNoValidator = orderDescMaintOrderNo.Validators[0].(func(string) error)
	// orderDescVehiclePlateNo is the schema descriptor for vehiclePlateNo field.
	orderDescVehiclePlateNo := orderFields[1].Descriptor()
	// order.VehiclePlateNoValidator is a validator for the "vehiclePlateNo" field. It is called by the builders before save.
	order.VehiclePlateNoValidator = orderDescVehiclePlateNo.Validators[0].(func(string) error)
	// orderDescVehicleTeamName is the schema descriptor for vehicleTeamName field.
	orderDescVehicleTeamName := orderFields[2].Descriptor()
	// order.VehicleTeamNameValidator is a validator for the "vehicleTeamName" field. It is called by the builders before save.
	order.VehicleTeamNameValidator = orderDescVehicleTeamName.Validators[0].(func(string) error)
	// orderDescMaintRequestType is the schema descriptor for maintRequestType field.
	orderDescMaintRequestType := orderFields[3].Descriptor()
	// order.MaintRequestTypeValidator is a validator for the "maintRequestType" field. It is called by the builders before save.
	order.MaintRequestTypeValidator = orderDescMaintRequestType.Validators[0].(func(string) error)
	// orderDescStationName is the schema descriptor for stationName field.
	orderDescStationName := orderFields[5].Descriptor()
	// order.StationNameValidator is a validator for the "stationName" field. It is called by the builders before save.
	order.StationNameValidator = orderDescStationName.Validators[0].(func(string) error)
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderFields[6].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderFields[7].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	stationFields := schema.Station{}.Fields()
	_ = stationFields
	// stationDescAgencyName is the schema descriptor for agencyName field.
	stationDescAgencyName := stationFields[1].Descriptor()
	// station.AgencyNameValidator is a validator for the "agencyName" field. It is called by the builders before save.
	station.AgencyNameValidator = stationDescAgencyName.Validators[0].(func(string) error)
	// stationDescCreatedAt is the schema descriptor for created_at field.
	stationDescCreatedAt := stationFields[2].Descriptor()
	// station.DefaultCreatedAt holds the default value on creation for the created_at field.
	station.DefaultCreatedAt = stationDescCreatedAt.Default.(func() time.Time)
	// stationDescUpdatedAt is the schema descriptor for updated_at field.
	stationDescUpdatedAt := stationFields[3].Descriptor()
	// station.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	station.DefaultUpdatedAt = stationDescUpdatedAt.Default.(func() time.Time)
	// station.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	station.UpdateDefaultUpdatedAt = stationDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[1].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[3].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = userDescPasswordHash.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
