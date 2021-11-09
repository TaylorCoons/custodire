package settings

import (
	"os"
	"testing"
)

func TestParseSettings(t *testing.T) {
	settings := ParseSettings("../settings.json")
	if settings.DbConnection.Username != "user" {
		t.Error("Failed to parse db username")
	}
	if settings.DbConnection.Password != "password" {
		t.Error("Failed to parse db password")
	}
	if settings.DbConnection.Host != "loopback" {
		t.Error("Failed to parse db host")
	}
	if settings.DbConnection.Database != "schema" {
		t.Error("Failed to parse db schema")
	}
	if settings.Port != "8080" {
		t.Error("Failed to parse port")
	}
}

func TestGetSettingsFile(t *testing.T) {
	// Test Dev
	file := GetSettingsFile(Dev)
	if file != "settings.dev.json" {
		t.Error("Failed to get settings file for dev")
	}
	// Test Test
	file = GetSettingsFile(Test)
	if file != "settings.test.json" {
		t.Error("Failed to get settings file for test")
	}
	// Test Prod
	file = GetSettingsFile(Prod)
	if file != "settings.prod.json" {
		t.Error("Failed to get settings file for prod")
	}
}

func TestGetEnvironment(t *testing.T) {
	const buildEnv string = "CUSTODIRE_ENV"
	prevEnvVar := os.Getenv(buildEnv)

	// Test Dev
	err := os.Setenv(buildEnv, "Dev")
	if err != nil {
		panic(err.Error())
	}
	env := GetEnvironment()
	if env != Dev {
		t.Error("Failed to set environment to dev")
	}
	// Test Test
	err = os.Setenv(buildEnv, "test")
	if err != nil {
		panic(err.Error())
	}
	env = GetEnvironment()
	if env != Test {
		t.Error("Failed to set environment to test")
	}
	// Test Prod
	err = os.Setenv(buildEnv, "PROD")
	if err != nil {
		panic(err.Error())
	}
	env = GetEnvironment()
	if env != Prod {
		t.Error("Failed to set environment to prod")
	}
	// Test default
	err = os.Setenv(buildEnv, "")
	if err != nil {
		panic(err.Error())
	}
	env = GetEnvironment()
	if env != Dev {
		t.Error("Failed to set environment to by default")
	}

	os.Setenv(buildEnv, prevEnvVar)
}
