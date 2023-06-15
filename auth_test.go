package twitter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "dreamer",
				ConfirmPassword: "dreamer",
			},
			err: nil,
		},
		{
			name: "inValid",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "dreamer",
				ConfirmPassword: "dreamer",
			},
			err: ErrValidation,
		},

		{
			name: "too short user name",
			input: RegisterInput{
				Username:        "b",
				Email:           "bob@gmail.com",
				Password:        "dreamer",
				ConfirmPassword: "dreamer",
			},
			err: ErrValidation,
		},

		{
			name: "too short password",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "d",
				ConfirmPassword: "d",
			},
			err: ErrValidation,
		},

		{
			name: "confirm password don't match to password ",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "dreamer",
				ConfirmPassword: "1d",
			},
			err: ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
