# Terraform Provider Shared - Helper Functions Reference

## Complete List of Implemented Functions

Functions are located in:
- `schema/attributes.go` - String attribute helpers with plan modifiers
- `schema/set_attributes.go` - Set attribute helpers (with and without validators)
- `schema/list_attributes.go` - List attribute helpers
- `schema/map_attributes.go` - Map attribute helpers

All helpers follow the builder pattern.

### Plan Modifier Support Matrix

| Function | Optional | Computed | Sensitive | Default | Validators | Plan Modifiers |
|----------|----------|----------|-----------|---------|------------|----------------|
| `ResourceOptionalStringWithPlanModifier` | ✓ | ✗ | ✗ | ✗ | ✗ | ✓ |
| `ResourceOptionalStringWithDefaultAndPlanModifier` | ✓ | ✓ | ✗ | ✓ | ✗ | ✓ |
| `ResourceComputedOptionalStringWithPlanModifier` | ✓ | ✓ | ✗ | ✗ | ✗ | ✓ |
| `ResourceComputedOptionalStringWithDefaultAndPlanModifier` | ✓ | ✓ | ✗ | ✓ | ✗ | ✓ |
| `ResourceComputedStringWithPlanModifier` | ✗ | ✓ | ✗ | ✗ | ✗ | ✓ |
| `ResourceSensitiveOptionalStringWithPlanModifier` | ✓ | ✗ | ✓ | ✗ | ✗ | ✓ |
| `ResourceRequiredStringEnumWithPlanModifier` | ✗ | ✗ | ✗ | ✗ | ✓ | ✓ |

## Usage Examples

### Basic Optional String with Plan Modifier
```go
tfschema.ResourceOptionalStringWithPlanModifier(
    "My optional field",
    stringplanmodifier.UseStateForUnknown(),
)
```

### Computed Optional String (Most Common)
```go
tfschema.ResourceComputedOptionalStringWithPlanModifier(
    "URL to access the Repository",
    stringplanmodifier.UseStateForUnknown(),
)
```

### Sensitive Optional String with Plan Modifier
```go
tfschema.ResourceSensitiveOptionalStringWithPlanModifier(
    "Password field",
    stringplanmodifier.UseStateForUnknown(),
)
```

### Required Enum with Plan Modifier
```go
tfschema.ResourceRequiredStringEnumWithPlanModifier(
    "Protocol version",
    []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
    "v1", "v2", "v3",
)
```

### Computed ID with Plan Modifier
```go
tfschema.ResourceComputedStringWithPlanModifier(
    "Resource identifier",
    stringplanmodifier.UseStateForUnknown(),
)
```

## Direct Usage Locations (Candidates for Refactoring)

The following files currently use direct schema construction instead of these helpers:

1. **repository/repository_common.go** (Line ~457)
   - Pattern: `schema.StringAttribute` with `UseStateForUnknown()`
   - Helper: `ResourceComputedOptionalStringWithPlanModifier()`

2. **system/security_realms.go** (Line ~76)
   - Pattern: Computed ID with `UseStateForUnknown()`
   - Helper: `ResourceComputedStringWithPlanModifier()`

3. **repository/format/yum.go** (Lines ~317, ~324)
   - Pattern: Optional string with `UseStateForUnknown()`
   - Helper: `ResourceOptionalStringWithPlanModifier()`
   - Pattern: Optional sensitive string with `UseStateForUnknown()`
   - Helper: `ResourceSensitiveOptionalStringWithPlanModifier()`

4. **repository/format/proxy.go** (Lines ~186, ~212)
   - Pattern: Sensitive optional string with `UseStateForUnknown()`
   - Helper: `ResourceSensitiveOptionalStringWithPlanModifier()`

5. **repository/format/conan.go** (Line ~248)
   - Pattern: Required enum string with `UseStateForUnknown()`
   - Helper: `ResourceRequiredStringEnumWithPlanModifier()`

6. **repository/format/docker.go** (Line ~346)
   - Pattern: Bool with `UseStateForUnknown()` - Already has bool helper

7. **repository/format/hosted.go** (Line ~53)
   - Pattern: Object with `UseStateForUnknown()` - No helper needed

## Parameter Signatures

### Standard Variadic Plan Modifiers
```go
func ResourceOptionalStringWithPlanModifier(
    description string,
    planMods ...planmodifier.String,
) resourceschema.StringAttribute
```

### Slice-based Plan Modifiers (for enums)
```go
func ResourceRequiredStringEnumWithPlanModifier(
    description string,
    planMods []planmodifier.String,
    enumValues ...string,
) resourceschema.StringAttribute
```

Note: Enum functions use slice-based plan modifiers to distinguish from enum values.

## Set Attribute Validator Helpers

### Available Functions

Set validators are available in `schema/set_attributes.go` for String, Int64, and Bool element types.

#### String Sets with Validators
```go
func ResourceRequiredStringSetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute

func ResourceOptionalStringSetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute

func ResourceComputedStringSetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute
```

#### Int64 Sets with Validators
```go
func ResourceRequiredInt64SetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute

func ResourceOptionalInt64SetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute

func ResourceComputedInt64SetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute
```

#### Bool Sets with Validators
```go
func ResourceRequiredBoolSetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute

func ResourceOptionalBoolSetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute

func ResourceComputedBoolSetWithValidator(
    description string,
    validators ...validator.Set,
) resourceschema.SetAttribute
```

### Usage Examples

#### Required String Set with Size Validator
```go
"matchers": schema.ResourceRequiredStringSetWithValidator(
    "Regular expressions used to identify request paths",
    setvalidator.SizeAtLeast(1),
)
```

#### Optional String Set with Multiple Validators
```go
"tags": schema.ResourceOptionalStringSetWithValidator(
    "Resource tags",
    setvalidator.SizeAtLeast(0),
    setvalidator.SizeAtMost(10),
)
```

## Integration Notes

- All helpers use the internal collection builders (`newResourceSetAttribute`, etc.)
- Configuration is passed through `collectionConfig` struct
- Follows the exact same pattern as other attribute helpers
- Maintains backward compatibility - no breaking changes
- All existing functions remain unchanged
- Validators are passed as variadic parameters for clean API
