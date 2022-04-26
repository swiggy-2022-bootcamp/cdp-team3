package configs

import "testing"

// expectedGot is a constant string to format error messages.
const expectedGot string = "Expected a non empty string, got %s"

func TestEnvs(t *testing.T) {
	// Arrange and Act
	got := []string{
		EnvAuthHost(),
		EnvAuthServiceGRPCPort(),
		EnvCartHost(),
		EnvCartServiceGRPCPort(),
		EnvCheckoutHost(),
		EnvOrderHost(),
		EnvOrderServiceGRPCPort(),
		EnvServicePort(),
		EnvShippingHost(),
		EnvShippingServiceGRPCPort(),
	}

	// Assert 
	for _, env := range got {
		if env == "" {
			t.Errorf(expectedGot, env)
		}
	}
}
