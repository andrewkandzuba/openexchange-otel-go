package utils

import (
	"os"
	"syscall"
	"testing"
)

func TestGetPortFromEnv(t *testing.T) {

	expected := "8081"
	os.Setenv(ENV_PORT, expected)

	actual := GetEnv(ENV_PORT, "8080")

	if actual != "8081" {
		t.Fatalf(`Port Env Variable (Actual = %v, Expected = %v)`, actual, expected)
	}
}

func TestGetPortDefault(t *testing.T) {

	expected := "8080"
	syscall.Unsetenv(ENV_PORT)
	actual := GetEnv(ENV_PORT, "8080")

	if actual != "8080" {
		t.Fatalf(`Port Env Variable (Actual = %v, Expected = %v)`, actual, expected)
	}
}
