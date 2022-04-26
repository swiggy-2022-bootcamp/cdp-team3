package configs

import "testing"

// expectedGot is a constant string to format error messages.
const expectedGot string = "Expected a non empty string, got %s"

func TestEnvs(t *testing.T) {
	// Arrange and Act
	got := []string{
		EnvAccessKey(),
		EnvCartHost(),
		EnvCartServiceGRPCPort(),
		EnvKafkaBrokerAddress(),
		EnvKafkaUserCreatedTopic(),
		EnvKafkaUserDeletedTopic(),
		EnvRegion(),
		EnvSecretKey(),
		EnvServicePort(),
	}

	// Assert 
	for _, env := range got {
		if env == "" {
			t.Errorf(expectedGot, env)
		}
	}
}
