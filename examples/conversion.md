# Type Conversion Examples

The conversion utilities simplify converting between Go types and Terraform framework types.

## Pointer Conversions

Convert between primitive types and their pointer equivalents:

```go
import "github.com/sonatype-nexus-community/terraform-provider-shared/util"

// String conversions
strPtr := util.StringToPtr("hello")              // string -> *string
strVal := util.SafeString(strPtr)                // *string -> string (handles nil)
emptyStr := util.EmptyIfNil(nil)                 // *string -> string (nil becomes "")
nilStr := util.NilIfEmpty("")                    // string -> *string (empty becomes nil)

// Boolean conversions
boolPtr := util.BoolToPtr(true)                  // bool -> *bool
boolVal := util.SafeBool(boolPtr)                // *bool -> bool (handles nil)

// Integer conversions
int64Ptr := util.Int64ToPtr(42)                  // int64 -> *int64
int64Val := util.SafeInt64(int64Ptr)             // *int64 -> int64 (handles nil)

int32Ptr := util.Int32ToPtr(10)                  // int32 -> *int32
int32Val := util.SafeInt32(int32Ptr)             // *int32 -> int32 (handles nil)
```

## String Parsing

Parse strings to numeric types:

```go
// String to integers
num, err := util.StringToInt64("123")            // string -> int64
num32, err := util.StringToInt32("10")           // string -> int32

// String to float
floatNum, err := util.StringToFloat("3.14")      // string -> float64

// String to boolean
boolVal, err := util.StringToBool("true")        // string -> bool
```

## Number to String

Convert numeric types to strings:

```go
str := util.Int64ToString(123)                   // int64 -> string
str32 := util.Int32ToString(10)                  // int32 -> string
floatStr := util.FloatToString(3.14, 2)          // float64 -> string with precision
boolStr := util.BoolToString(true)               // bool -> "true" or "false"
```

## Collection Conversions

Convert Go slices to Terraform framework types:

```go
// String slice to Set
stringSet := util.StringSliceToValue([]string{"a", "b", "c"})

// Integer slice to Set
intSet := util.Int64SliceToValue([]int64{1, 2, 3})

// Pointer slice to Set
strPtrSlice := []*string{
    util.StringToPtr("x"),
    util.StringToPtr("y"),
}
ptrSet := util.StringPtrSliceToValue(strPtrSlice)
```

## Safe Access Patterns

Handle nil pointers gracefully without panicking:

```go
// Instead of: value := *apiResponse.StatusCode (may panic if nil)
// Use:
status := util.SafeInt64(apiResponse.StatusCode)  // Returns 0 if nil

// Instead of: name := *user.Name (may panic if nil)
// Use:
name := util.SafeString(user.Name)                 // Returns "" if nil

// Instead of: active := *config.Enabled (may panic if nil)
// Use:
active := util.SafeBool(config.Enabled)            // Returns false if nil
```

## Common Patterns

### Reading from API and Converting to Terraform

```go
// API returns *string, Terraform expects types.String
apiResponse := &APIResponse{
    Name: util.StringToPtr("my-resource"),
    Port: util.Int32ToPtr(8080),
}

terraformData := map[string]interface{}{
    "name": apiResponse.Name,           // or use StringPtrToValue
    "port": util.Int32PtrToValue(apiResponse.Port),
}
```

### Converting from Terraform to API

```go
// Terraform sends types, API expects pointers
terraformConfig := map[string]interface{}{
    "timeout": "30",
    "enabled": true,
}

apiRequest := &APIRequest{
    Timeout: util.StringToInt64(terraformConfig["timeout"].(string)),
    Enabled: util.BoolToPtr(terraformConfig["enabled"].(bool)),
}
```
