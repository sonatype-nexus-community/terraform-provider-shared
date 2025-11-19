# Error Handling Examples

Standardized error handling ensures consistent, user-friendly error messages across your provider.

## Creating Error Messages

The error functions return a title and message that can be used directly or added to diagnostics:

```go
import "github.com/sonatype-nexus-community/terraform-provider-shared/errors"

// API error
title, message := errors.APIError("creating", "user", "username already exists")
// title: "Error creating user"
// message: "Could not create user: username already exists"

// Not found error
title, message := errors.NotFoundError("application", "app-123")
// title: "application Not Found"
// message: "The application with ID 'app-123' was not found. It may have been deleted outside of Terraform."

// Validation error
title, message := errors.ValidationError("email", "invalid format")
// title: "Invalid email"
// message: "The email value is invalid: invalid format"

// Conflict error
title, message := errors.ConflictError("policy", "policy with same name exists")
// title: "Conflict creating policy"
// message: "A conflict occurred: policy with same name exists"

// Unauthorized error
title, message := errors.UnauthorizedError("delete resource")
// title: "Unauthorized delete resource"
// message: "You do not have permission to delete resource. Please check your credentials and permissions."

// Timeout error
title, message := errors.TimeoutError("updating", "configuration")
// title: "Timeout updating"
// message: "Operation timed out while updating configuration. The resource may have been created but confirmation could not be received."
```

## Adding Errors to Diagnostics

The recommended approach is to add errors directly to the response diagnostics:

```go
import (
    "http"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    "github.com/sonatype-nexus-community/terraform-provider-shared/errors"
)

func (r *myResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    // ... read resource ...
    
    if response.StatusCode == http.StatusNotFound {
        errors.AddNotFoundDiagnostic(&resp.Diagnostics, "application", state.ID.ValueString())
        return
    }
    
    if err != nil {
        errors.AddAPIErrorDiagnostic(&resp.Diagnostics, "reading", "application", response, err)
        return
    }
}
```

Available diagnostic functions:

```go
// Add errors
errors.AddAPIErrorDiagnostic(&diags, "reading", "policy", response, originalErr)
errors.AddNotFoundDiagnostic(&diags, "organization", "org-456")
errors.AddValidationDiagnostic(&diags, "port", "must be between 1 and 65535")
errors.AddConflictDiagnostic(&diags, "repository", "repository name already in use")
errors.AddUnauthorizedDiagnostic(&diags, "access this resource")
errors.AddTimeoutDiagnostic(&diags, "creating", "application")

// Add warnings (non-critical issues)
errors.HandleAPIWarning("Resource state may be inconsistent", &err, response, &diags)
```

## HTTP Status Code Checking

Helper functions to check HTTP response status codes:

```go
// Specific status checks
isNotFound := errors.IsNotFound(response.StatusCode)      // 404
isForbidden := errors.IsForbidden(response.StatusCode)    // 403
isUnauth := errors.IsUnauthorized(response.StatusCode)    // 401
isConflict := errors.IsConflict(response.StatusCode)      // 409

// Broad category checks
isClientErr := errors.IsClientError(response.StatusCode)  // 4xx
isServerErr := errors.IsServerError(response.StatusCode)  // 5xx
```

## Parsing API Responses

Extract error details from HTTP responses:

```go
// Parse error from HTTP response body
errorDetail := errors.ParseAPIError(response)
// Returns formatted string: "404 Not Found: {response body}"
```

## Advanced Error Handling

Automatic network error detection distinguishes between network issues and API errors:

```go
// This handles both network errors (DNS, connection refused, timeouts)
// and API errors, detecting the appropriate error type automatically
var err error = networkRequest()
errors.HandleAPIError("Failed to read resource", &err, response, &diags)
```

## Complete Example

```go
func (r *organizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data organizationResourceModel
    resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    // Make API call
    createReq := &api.CreateOrgRequest{
        Name: data.Name.ValueString(),
    }
    
    response, err := r.client.CreateOrganization(ctx, createReq)
    if err != nil {
        errors.HandleAPIError("Failed to create organization", &err, response, &resp.Diagnostics)
        return
    }

    if response.StatusCode == http.StatusConflict {
        errors.AddConflictDiagnostic(&resp.Diagnostics, "organization", "organization name already exists")
        return
    }

    if response.StatusCode != http.StatusCreated {
        errors.AddAPIErrorDiagnostic(&resp.Diagnostics, "creating", "organization", response, nil)
        return
    }

    // Success - populate state
    data.ID = types.StringValue(response.ID)
    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
```
