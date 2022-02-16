package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCheckIfAddressEntityIsValid(t *testing.T) {
	tests := []struct {
		name         string
		args         *Address
		want         bool
		invalidField string
	}{
		{
			name: "[Address] Should everything be valid",
			args: &Address{
				ID:           50,
				Name:         "Success Test",
				Cep:          "15502306",
				AddressLine1: "Success Test",
				AddressLine2: "Success Test",
				City:         "Success Test",
				State:        "Success Test",
				Country:      "Success Test",
				UserID:       544,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         true,
			invalidField: "",
		},

		{
			name: "[Address] Should Cep be invalid",
			args: &Address{
				ID:           50,
				Name:         "Failed Test",
				Cep:          "",
				AddressLine1: "Failed Test",
				AddressLine2: "Failed Test",
				City:         "Failed Test",
				State:        "Failed Test",
				Country:      "Failed Test",
				UserID:       544,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         false,
			invalidField: "Cep",
		},

		{
			name: "[Address] Should AddressLine1 be invalid",
			args: &Address{
				ID:           50,
				Name:         "Failed Test",
				Cep:          "15502306",
				AddressLine1: "",
				AddressLine2: "Failed Test",
				City:         "Failed Test",
				State:        "Failed Test",
				Country:      "Failed Test",
				UserID:       544,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         false,
			invalidField: "AddressLine1",
		},

		{
			name: "[Address] Should City be invalid",
			args: &Address{
				ID:           50,
				Name:         "Failed Test",
				Cep:          "15502306",
				AddressLine1: "Failed Test",
				AddressLine2: "Failed Test",
				City:         "",
				State:        "Failed Test",
				Country:      "Failed Test",
				UserID:       544,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         false,
			invalidField: "City",
		},

		{
			name: "[Address] Should State be invalid",
			args: &Address{
				ID:           50,
				Name:         "Failed Test",
				Cep:          "15502306",
				AddressLine1: "Failed Test",
				AddressLine2: "Failed Test",
				City:         "Failed Test",
				State:        "",
				Country:      "Failed Test",
				UserID:       544,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         false,
			invalidField: "State",
		},

		{
			name: "[Address] Should Country be invalid",
			args: &Address{
				ID:           50,
				Name:         "Failed Test",
				Cep:          "15502306",
				AddressLine1: "Failed Test",
				AddressLine2: "Failed Test",
				City:         "Failed Test",
				State:        "Failed Test",
				Country:      "",
				UserID:       544,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         false,
			invalidField: "Country",
		},

		{
			name: "[Address] Should UserID be invalid",
			args: &Address{
				ID:           50,
				Name:         "Failed Test",
				Cep:          "15502306",
				AddressLine1: "Failed Test",
				AddressLine2: "Failed Test",
				City:         "Failed Test",
				State:        "Failed Test",
				Country:      "Failed Test",
				UserID:       0,
				MainAddress:  true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			want:         false,
			invalidField: "UserID",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testValidate, testInvalidField := CheckIfAddressEntityIsValid(test.args)
			assert.Equalf(t, test.want, testValidate, test.name)
			assert.Equalf(t, test.invalidField, testInvalidField, test.name)
		})
	}
}

func TestPrepareAddressToUpdate(t *testing.T) {
	address := &Address{
		ID:           50,
		Name:         "To be Tested",
		Cep:          "15502306--",
		AddressLine1: "To be Tested",
		AddressLine2: "To be Tested",
		City:         "To be Tested",
		State:        "To be Tested",
		Country:      "To be Tested",
		UserID:       544,
		MainAddress:  true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tests := []struct {
		name string
		args *UpdateAddress
	}{
		{
			name: "[Address] Should update all address fields",
			args: &UpdateAddress{
				Name:         "Success Test",
				Cep:          "Success Test",
				AddressLine1: "Success Test",
				AddressLine2: "Success Test",
				City:         "Success Test",
				State:        "Success Test",
				Country:      "Success Test",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			PrepareAddressToUpdate(&address, test.args)
			assert.Equalf(t, address.Name, test.args.Name, test.name)
			assert.Equalf(t, address.Cep, test.args.Cep, test.name)
			assert.Equalf(t, address.AddressLine1, test.args.AddressLine1, test.name)
			assert.Equalf(t, address.AddressLine2, test.args.AddressLine2, test.name)
			assert.Equalf(t, address.City, test.args.City, test.name)
			assert.Equalf(t, address.State, test.args.State, test.name)
			assert.Equalf(t, address.Country, test.args.Country, test.name)
		})
	}

}

func TestMessageAddress(t *testing.T) {

	type args struct {
		genericMessage string
	}

	tests := []struct {
		name string
		args *args
		want string
	}{
		{
			name: "[Address] Should do address message convert",
			args: &args{genericMessage: "%_% test"},
			want: "Address test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testValidate := MessageAddress(test.args.genericMessage)
			assert.Equalf(t, test.want, testValidate, test.name)
		})
	}

}
