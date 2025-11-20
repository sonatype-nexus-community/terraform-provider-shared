# Validators Quick Reference Guide

## String Validators

### Regex Pattern Matching
```go
// Required
tfschema.RequiredStringWithRegex(description, pattern, errorMsg)

// Optional
tfschema.OptionalStringWithRegex(description, pattern, errorMsg)

// Computed
tfschema.ComputedStringWithRegex(description, pattern, errorMsg)
```

### Length Validation
```go
// Min-Max (Between)
tfschema.RequiredStringWithLengthBetween(description, minLen, maxLen)
tfschema.OptionalStringWithLengthBetween(description, minLen, maxLen)

// Minimum Only
tfschema.RequiredStringWithLengthAtLeast(description, minLen)
tfschema.OptionalStringWithLengthAtLeast(description, minLen)

// Maximum Only
tfschema.RequiredStringWithLengthAtMost(description, maxLen)
tfschema.OptionalStringWithLengthAtMost(description, maxLen)
```

### Combined Regex + Length
```go
// Required with regex and min/max length
tfschema.RequiredStringWithRegexAndLength(description, pattern, errorMsg, minLen, maxLen)

// Optional with regex and min/max length
tfschema.OptionalStringWithRegexAndLength(description, pattern, errorMsg, minLen, maxLen)
```

### Sensitive Strings
```go
// Optional sensitive with minimum length
tfschema.OptionalSensitiveStringWithLengthAtLeast(description, minLen)

// Required sensitive with minimum length
tfschema.RequiredSensitiveStringWithLengthAtLeast(description, minLen)
```

### Custom Validator Chains
```go
// For custom combinations of validators
tfschema.StringWithValidators(description, validator1, validator2, ...)
tfschema.OptionalStringWithValidators(description, validator1, validator2, ...)
tfschema.RequiredStringWithValidators(description, validator1, validator2, ...)
tfschema.ComputedStringWithValidators(description, validator1, validator2, ...)
```

## Integer Validators

### Range Validation
```go
// Required with min/max range
tfschema.RequiredInt64WithRange(description, minValue, maxValue)

// Optional with min/max range
tfschema.OptionalInt64WithRange(description, minValue, maxValue)

// Computed with min/max range
tfschema.ComputedInt64WithRange(description, minValue, maxValue)
```

### Custom Validator Chains
```go
// For custom combinations of validators
tfschema.Int64WithValidators(description, validator1, validator2, ...)
tfschema.OptionalInt64WithValidators(description, validator1, validator2, ...)
tfschema.RequiredInt64WithValidators(description, validator1, validator2, ...)
tfschema.ComputedInt64WithValidators(description, validator1, validator2, ...)
```

## Common Patterns

### Privilege/Resource Names
```go
const namePattern = `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`

"name": tfschema.RequiredStringWithRegex(
	"Name of the resource",
	regexp.MustCompile(namePattern),
	"Name must start with alphanumeric or hyphen, contain only alphanumeric, underscores, hyphens, and periods",
),
```

### GCS Bucket Names
```go
const gcsBucketPattern = `^[a-z0-9][a-z0-9\-]*[a-z0-9]$|^[a-z0-9]$`

"bucket_name": tfschema.RequiredStringWithRegexAndLength(
	"GCS bucket name",
	regexp.MustCompile(gcsBucketPattern),
	"Must contain only lowercase letters, numbers, and hyphens",
	3,
	63,
),
```

### URLs
```go
const urlPattern = `^https?://[^\s]+$`

"webhook_url": tfschema.RequiredStringWithRegex(
	"Webhook URL",
	regexp.MustCompile(urlPattern),
	"Must be a valid http:// or https:// URL",
),
```

### XML Content
```go
const xmlPattern = `(?s)^\s*<.*>\s*$`

"metadata_xml": tfschema.RequiredStringWithRegex(
	"XML metadata",
	regexp.MustCompile(xmlPattern),
	"Must be valid XML format",
),
```

### Sensitive JSON
```go
"credentials_json": tfschema.RequiredSensitiveStringWithLengthAtLeast(
	"Credentials JSON content",
	10, // Minimum length validation
),
```

## Migration Guide

### Step 1: Extract Common Patterns
```go
// Define constants at package level
const (
	privilegeNamePattern = `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`
	urlPattern           = `^https?://[^\s]+$`
)
```

### Step 2: Replace Verbose Definitions
```go
// Old (11 lines)
"name": schema.StringAttribute{
	Description: "Name",
	Required:    true,
	Validators: []validator.String{
		stringvalidator.RegexMatches(
			regexp.MustCompile(privilegeNamePattern),
			"error message",
		),
	},
},

// New (5 lines)
"name": tfschema.RequiredStringWithRegex(
	"Name",
	regexp.MustCompile(privilegeNamePattern),
	"error message",
),
```

### Step 3: Combine with Length When Needed
```go
// Combined validators in one call
"name": tfschema.RequiredStringWithRegexAndLength(
	"Name",
	regexp.MustCompile(privilegeNamePattern),
	"error message",
	1,
	255,
),
```

## When to Use Custom `*WithValidators` Functions

Use the generic `*WithValidators` functions when:
1. You need multiple validators that aren't pre-combined
2. You have a unique validation requirement
3. You're combining validators in an uncommon way

```go
import "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

"custom_field": tfschema.RequiredStringWithValidators(
	"Custom field",
	stringvalidator.LengthBetween(5, 100),
	stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Z]`), "must start with uppercase"),
	// Add more validators as needed
),
```

## Complete Example

```go
func (r *MyResource) Schema(...) {
	resp.Schema = schema.Schema{
		Description: "My resource",
		Attributes: map[string]schema.Attribute{
			// Simple required string
			"name": tfschema.RequiredString("Name of the resource"),
			
			// Required with regex validation
			"identifier": tfschema.RequiredStringWithRegex(
				"Unique identifier",
				regexp.MustCompile(`^[a-z0-9]+$`),
				"Must contain only lowercase letters and numbers",
			),
			
			// Required with combined regex and length
			"hostname": tfschema.RequiredStringWithRegexAndLength(
				"Hostname or domain",
				regexp.MustCompile(`^[a-zA-Z0-9.-]+$`),
				"Must be a valid hostname",
				1,
				255,
			),
			
			// Optional with length constraint
			"description": tfschema.OptionalStringWithLengthAtMost(
				"Description",
				1000,
			),
			
			// Sensitive credentials with validation
			"api_key": tfschema.RequiredSensitiveStringWithLengthAtLeast(
				"API key",
				20,
			),
			
			// Optional integer with range
			"port": tfschema.OptionalInt64WithRange(
				"Port number",
				1,
				65535,
			),
			
			// Computed field
			"created_at": tfschema.ComputedString("Creation timestamp"),
		},
	}
}
```
