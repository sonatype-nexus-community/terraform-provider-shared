# Validators Enhancement - v0.5.0

## Summary

Added comprehensive validator helper functions to `schema/validators.go` to reduce boilerplate code when defining Terraform schema attributes with validation rules.

## New Helper Functions Added

### String Validators

#### Regex Validators
- `RequiredStringWithRegex()` - Required string with regex pattern validation
- `OptionalStringWithRegex()` - Optional string with regex pattern validation
- `ComputedStringWithRegex()` - Computed string with regex pattern validation

#### Length Validators
- `RequiredStringWithLengthBetween()` - Required string with min/max length
- `OptionalStringWithLengthBetween()` - Optional string with min/max length
- `RequiredStringWithLengthAtLeast()` - Required string with minimum length
- `OptionalStringWithLengthAtLeast()` - Optional string with minimum length
- `RequiredStringWithLengthAtMost()` - Required string with maximum length
- `OptionalStringWithLengthAtMost()` - Optional string with maximum length

#### Combined Validators
- `RequiredStringWithRegexAndLength()` - Required string with regex + length constraints
- `OptionalStringWithRegexAndLength()` - Optional string with regex + length constraints

#### Generic Validators (Custom combinations)
- `StringWithValidators()` - Required string with custom validator list
- `OptionalStringWithValidators()` - Optional string with custom validator list
- `RequiredStringWithValidators()` - Alias for StringWithValidators()
- `ComputedStringWithValidators()` - Computed string with custom validators

#### Sensitive String Validators
- `OptionalSensitiveStringWithLengthAtLeast()` - Optional sensitive string with minimum length
- `RequiredSensitiveStringWithLengthAtLeast()` - Required sensitive string with minimum length

### Integer Validators

#### Range Validators
- `RequiredInt64WithRange()` - Required int64 with min/max range
- `OptionalInt64WithRange()` - Optional int64 with min/max range
- `ComputedInt64WithRange()` - Computed int64 with min/max range

#### Generic Validators (Custom combinations)
- `Int64WithValidators()` - Required int64 with custom validator list
- `OptionalInt64WithValidators()` - Optional int64 with custom validator list
- `RequiredInt64WithValidators()` - Alias for Int64WithValidators()
- `ComputedInt64WithValidators()` - Computed int64 with custom validators

## Example Usage

### Before (Verbose)
```go
"name": schema.StringAttribute{
	Description: "The name of the privilege",
	Required:    true,
	Validators: []validator.String{
		stringvalidator.RegexMatches(
			regexp.MustCompile(`^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`),
			`Name must match the regex pattern`,
		),
	},
},
```

### After (Concise)
```go
"name": tfschema.RequiredStringWithRegex(
	"The name of the privilege",
	regexp.MustCompile(`^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`),
	`Name must match the regex pattern`,
),
```

### Regex + Length Constraints
```go
"bucket_name": tfschema.RequiredStringWithRegexAndLength(
	"The name of the GCS bucket",
	regexp.MustCompile(`^[a-z0-9][a-z0-9\-]*[a-z0-9]$|^[a-z0-9]$`),
	"Bucket name must contain only lowercase letters, numbers, and hyphens",
	3,
	63,
),
```

## Benefits

1. **Reduced Boilerplate**: Schema definitions are significantly more concise
2. **Consistency**: Standardized approach across all Terraform providers
3. **Maintainability**: Easier to understand at a glance
4. **Reusability**: Common patterns extracted into helpers
5. **Type Safety**: All validators properly typed

## Testing

All new functions have unit tests covering:
- Correct attribute mode (Required/Optional/Computed)
- Correct description assignment
- Proper validator assignment (where inspectable)

## Files Changed

- `schema/validators.go` - New file with all validator helper functions
- `schema/validators_test.go` - Comprehensive unit tests for all validators

## Backward Compatibility

✅ Fully backward compatible - this is an additive enhancement only.

## Providers Using These Helpers

- `terraform-provider-sonatyperepo` - Initial adoption started (v0.5.0+)

## Next Steps for Provider Integration

1. Update `terraform-provider-sonatyperepo` to use new helpers in:
   - `blob_store_google_cloud_resource.go`
   - `routing_rule.go`
   - `cleanup_policies.go`
   - `content_selector_resource.go`
   - `capability_type/common.go`
   - `capability_type/base_url.go`
   - `capability_type/outreach.go`
   - `system/security_saml.go`
   - `privilege/privileges_data_source.go`
   - `provider.go`

2. Update `terraform-provider-sonatypeiq` similarly
