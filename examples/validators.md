# Validators Examples

Validators enforce constraints on resource attribute values to ensure data integrity.

## Enum Validation

Restrict attribute values to a predefined set of options:

```go
import (
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/sonatype-nexus-community/terraform-provider-shared/validators"
)

schema.Schema{
    Attributes: map[string]schema.Attribute{
        // Environment must be one of: development, staging, production
        "environment": schema.StringAttribute{
            Description: "The deployment environment",
            Required:    true,
            Validators:  validators.StringOneOfValidator("development", "staging", "production"),
        },

        // Status must be one of: active, inactive, pending
        "status": schema.StringAttribute{
            Description: "The status of the resource",
            Optional:    true,
            Validators:  validators.StringOneOfValidator("active", "inactive", "pending"),
        },

        // Storage class options
        "storage_class": schema.StringAttribute{
            Description: "The storage class for the resource",
            Optional:    true,
            Validators: validators.StringOneOfValidator(
                "standard",
                "fast",
                "archive",
            ),
        },
    },
}
```

## Common Use Cases

- **Environment selectors**: development, staging, production
- **Status fields**: active, inactive, pending, archived
- **Resource types**: repository, application, policy
- **Log levels**: debug, info, warning, error
- **Authentication methods**: basic, oauth, saml

## Error Messages

When validation fails, users see clear error messages:

```
The value of "environment" must be one of: ["development", "staging", "production"], got: "test"
```
