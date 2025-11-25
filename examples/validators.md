# Validators Examples

Validators enforce constraints on attribute values to ensure data integrity. The schema package provides both Resource and DataSource validators with consistent naming conventions.

## Function Naming Convention

- **Resource Validators**: Prefixed with `Resource`
  - `ResourceRequiredString...()`, `ResourceOptionalString...()`, etc.
  
- **DataSource Validators**: Prefixed with `DataSource`
  - `DataSourceRequiredString...()`, `DataSourceOptionalString...()`, etc.

## String Validators with Regex

Enforce pattern matching on string attributes:

### Resource Example

```go
import (
    "regexp"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/sonatype-nexus-community/terraform-provider-shared/schema"
)

schema.Schema{
    Attributes: map[string]schema.Attribute{
        // Repository name must match pattern
        "name": schema.ResourceRequiredStringWithRegex(
            "The name of the repository (alphanumeric and hyphens only)",
            regexp.MustCompile(`^[a-zA-Z0-9-]+$`),
            "must contain only alphanumeric characters and hyphens",
        ),

        // Optional email with pattern validation
        "admin_email": schema.ResourceOptionalStringWithRegex(
            "Administrator email address",
            regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
            "must be a valid email address",
        ),
    },
}
```

### DataSource Example

```go
import (
    "regexp"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/sonatype-nexus-community/terraform-provider-shared/schema"
)

schema.Schema{
    Attributes: map[string]schema.Attribute{
        // Filter pattern for data source queries
        "filter": schema.DataSourceRequiredStringWithRegex(
            "Filter expression for search",
            regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`),
            "must contain only alphanumeric characters, hyphens, and underscores",
        ),
    },
}
```

## String Validators with Length

Enforce length constraints on string attributes:

### Resource Example

```go
"username": schema.ResourceRequiredStringWithLengthBetween(
    "The username for authentication",
    3,  // minimum
    32, // maximum
),

"api_key": schema.ResourceRequiredSensitiveStringWithLengthAtLeast(
    "The API key for authentication",
    20, // minimum length
),

"description": schema.ResourceOptionalStringWithLengthAtMost(
    "Optional resource description",
    500, // maximum length
),
```

### DataSource Example

```go
"search_term": schema.DataSourceRequiredStringWithLengthAtLeast(
    "Search term for filtering",
    1, // minimum length
),

"limit": schema.DataSourceOptionalStringWithLengthBetween(
    "Result limit expression",
    1,   // minimum
    100, // maximum
),
```

## Combined Regex and Length Validation

Validate both pattern and length constraints:

### Resource Example

```go
"repo_name": schema.ResourceRequiredStringWithRegexAndLength(
    "Repository name (alphanumeric, 3-64 characters)",
    regexp.MustCompile(`^[a-zA-Z0-9-]+$`),
    "must contain only alphanumeric characters and hyphens",
    3,  // minimum length
    64, // maximum length
),
```

### DataSource Example

```go
"filter_query": schema.DataSourceOptionalStringWithRegexAndLength(
    "Filter query expression",
    regexp.MustCompile(`^[a-zA-Z0-9\-_\s]+$`),
    "must contain only alphanumeric characters, hyphens, underscores, and spaces",
    1,    // minimum length
    1000, // maximum length
),
```

## Integer Validators

Enforce numeric range constraints:

### Resource Example

```go
// Port number validation
"port": schema.ResourceRequiredInt32WithRange(
    "Server port number",
    1024,  // minimum
    65535, // maximum
),

// Timeout in seconds
"timeout_seconds": schema.ResourceOptionalInt64WithRange(
    "Request timeout in seconds",
    1,     // minimum
    3600,  // maximum (1 hour)
),
```

### DataSource Example

```go
// Page number for pagination
"page": schema.DataSourceOptionalInt32WithRange(
    "Page number for results",
    1,     // minimum
    10000, // maximum
),

// Result limit
"limit": schema.DataSourceOptionalInt64WithRange(
    "Maximum number of results",
    1,      // minimum
    100000, // maximum
),
```

## Float Validators

Enforce numeric range constraints for float values:

### Resource Example

```go
// Storage quota in GB
"quota_gb": schema.ResourceRequiredFloat64WithRange(
    "Storage quota in gigabytes",
    0.1,      // minimum
    10000.0,  // maximum
),

// CPU percentage
"cpu_limit": schema.ResourceOptionalFloat64WithRange(
    "CPU usage limit as percentage",
    0.0,   // minimum
    100.0, // maximum
),
```

### DataSource Example

```go
// Score threshold
"min_score": schema.DataSourceOptionalFloat64WithRange(
    "Minimum score threshold for filtering",
    0.0,   // minimum
    100.0, // maximum
),
```

## Custom Validators

Use custom validator functions when needed:

### Resource Example

```go
import (
    "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
    "github.com/sonatype-nexus-community/terraform-provider-shared/schema"
)

"environment": schema.ResourceStringWithValidators(
    "Deployment environment",
    stringvalidator.OneOf("development", "staging", "production"),
),
```

### DataSource Example

```go
"format": schema.DataSourceStringWithValidators(
    "Output format",
    stringvalidator.OneOf("json", "xml", "csv"),
),
```

## Common Use Cases

### Resource Attributes

- **Required credentials**: `ResourceRequiredSensitiveStringWithLengthAtLeast()`
- **Repository names**: `ResourceRequiredStringWithRegex()` with alphanumeric pattern
- **Configuration values**: `ResourceOptionalStringWithLengthBetween()`
- **Numeric settings**: `ResourceRequiredInt32WithRange()` or `ResourceRequiredFloat64WithRange()`

### DataSource Attributes

- **Query filters**: `DataSourceRequiredStringWithRegex()` with search pattern
- **Pagination**: `DataSourceOptionalInt64WithRange()` for page/limit
- **Search terms**: `DataSourceOptionalStringWithLengthAtLeast()`
- **Numeric filters**: `DataSourceOptionalInt64WithRange()`

## Error Messages

When validation fails, users see clear error messages:

### Regex Validation Error
```
Error: Resource validation error
  with sonatype_repository.main,
  on main.tf line 5, in resource "sonatype_repository" "main":
   5:   name = "invalid@name"

The name attribute must contain only alphanumeric characters and hyphens
```

### Range Validation Error
```
Error: Resource validation error
  with sonatype_server.main,
  on main.tf line 8, in resource "sonatype_server" "main":
   8:   port = 99999

The port attribute must be between 1024 and 65535
```

### Length Validation Error
```
Error: Attribute validation error
  with sonatype_user.admin,
  on main.tf line 3, in resource "sonatype_user" "admin":
   3:   password = "abc"

The password attribute must have a minimum length of 8
```

## Sensitive Attributes

For sensitive data like passwords and API keys:

```go
"password": schema.ResourceRequiredSensitiveStringWithLengthAtLeast(
    "User password",
    8, // minimum length
),

"api_token": schema.ResourceRequiredSensitiveStringWithLengthAtLeast(
    "API authentication token",
    32, // minimum length
),
```

## Best Practices

1. **Use specific validators**: Choose the most specific validator for your use case
2. **Clear descriptions**: Provide helpful MarkdownDescription values
3. **Meaningful error messages**: Custom error messages should explain what values are accepted
4. **Consistent patterns**: Use consistent regex patterns across similar attributes
5. **Sensitive handling**: Always use sensitive validators for credentials and tokens
6. **Explicit intent**: Use the Resource/DataSource prefix to make schema type clear
