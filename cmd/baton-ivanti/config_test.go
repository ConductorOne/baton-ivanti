package main

import (
	"testing"

	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/conductorone/baton-sdk/pkg/test"
	"github.com/spf13/viper"
)

func TestConfigs(t *testing.T) {
	configurationSchema := field.NewConfiguration(
		ConfigurationFields,
		FieldRelationships...,
	)

	testCases := []test.TestCase{
		// Add test cases here.
	}

	validate := func(v *viper.Viper) error {
		return ValidateConfig(t.Context(), v)
	}
	test.ExerciseTestCases(t, configurationSchema, validate, testCases)
}
