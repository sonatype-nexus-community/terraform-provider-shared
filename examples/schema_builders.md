# Schema Builders Examples

Schema builders provide utility functions for creating consistent Terraform resource schema attributes.

## Basic Attributes

```go
import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/sonatype-nexus-community/terraform-provider-shared/schema/attributes"
)

func SchemaBuilderExample(ctx context.Context) schema.Schema {
    return schema.Schema{
        Attributes: map[string]schema.Attribute{
            // Standard ID attribute (read-only, computed)
            "id": attributes.StandardID(),

            // Required string attribute
            "name": attributes.RequiredString("The name of the resource"),

            // Optional string attribute
            "description": attributes.OptionalString("A description of the resource"),

            // Required boolean attribute with default
            "enabled": attributes.RequiredBool("Whether the resource is enabled", nil),

            // Optional integer attribute
            "port": attributes.OptionalInt64("The port number"),

            // Optional float attribute
            "timeout": attributes.OptionalFloat64("Timeout in seconds"),

            // Computed timestamp (server-managed)
            "last_updated": attributes.Timestamp(),

            // String attribute with sensitive data (passwords, tokens)
            "api_token": attributes.SensitiveString("The API token for authentication"),
        },
    }
}
```

## Collection Attributes

```go
// List of strings
"tags": attributes.StringList("A list of tags"),

// Set of strings (unordered, unique values)
"allowed_hosts": attributes.StringSet("A set of allowed hostnames"),

// Map of strings (key-value pairs)
"metadata": attributes.StringMap("Custom metadata as key-value pairs"),
```

## Benefits

- **Consistency**: Ensures all attributes follow the same patterns across your provider
- **Less Boilerplate**: Pre-configured attributes with sensible defaults
- **Type Safety**: Strongly typed attribute builders prevent common mistakes
- **Maintainability**: Updates to attribute patterns apply across the entire provider
