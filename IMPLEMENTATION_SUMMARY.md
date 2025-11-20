# String Attribute Plan Modifier Helpers - Implementation Summary

## Overview
Implemented missing helper functions for string attributes with plan modifiers in `terraform-provider-shared/schema/attributes.go` to support plan modification logic without requiring direct schema attribute construction.

## Problem Addressed
PATTERN 2 identified multiple locations where string attributes with plan modifiers were constructed directly using `schema.StringAttribute{}` blocks instead of leveraging helper functions. This mirrors the pattern fix applied to bool attributes.

## Implemented Functions

### 1. `ResourceOptionalStringWithPlanModifier`
```go
func ResourceOptionalStringWithPlanModifier(description string, planMods ...planmodifier.String) resourceschema.StringAttribute
```
- **Purpose**: Optional string attribute with custom plan modifiers
- **Use Case**: Attributes that are optional and need plan modification logic
- **Example**: Optional fields with `UseStateForUnknown()` modifier

### 2. `ResourceOptionalStringWithDefaultAndPlanModifier`
```go
func ResourceOptionalStringWithDefaultAndPlanModifier(description string, defaultValue string, planMods ...planmodifier.String) resourceschema.StringAttribute
```
- **Purpose**: Optional string attribute with default value and plan modifiers
- **Use Case**: Optional fields with computed defaults and plan modification logic
- **Flags**: `Optional=true`, `Computed=true`

### 3. `ResourceComputedOptionalStringWithPlanModifier`
```go
func ResourceComputedOptionalStringWithPlanModifier(description string, planMods ...planmodifier.String) resourceschema.StringAttribute
```
- **Purpose**: Computed optional string attribute with plan modifiers
- **Use Case**: Fields that are both computed and optional, preserving state for unknown values
- **Flags**: `Optional=true`, `Computed=true`
- **Pattern**: Matches bool variant `ResourceComputedOptionalBoolWithPlanModifier()`

### 4. `ResourceComputedOptionalStringWithDefaultAndPlanModifier`
```go
func ResourceComputedOptionalStringWithDefaultAndPlanModifier(description string, defaultValue string, planMods ...planmodifier.String) resourceschema.StringAttribute
```
- **Purpose**: Computed optional string attribute with default value and plan modifiers
- **Use Case**: Fields with both computed flags and default values
- **Flags**: `Optional=true`, `Computed=true`, `Default=StaticString()`

### 5. `ResourceComputedStringWithPlanModifier`
```go
func ResourceComputedStringWithPlanModifier(description string, planMods ...planmodifier.String) resourceschema.StringAttribute
```
- **Purpose**: Computed string attribute with plan modifiers
- **Use Case**: Read-only computed fields with plan modification logic
- **Example**: Resource IDs with `UseStateForUnknown()` modifier

### 6. `ResourceSensitiveOptionalStringWithPlanModifier`
```go
func ResourceSensitiveOptionalStringWithPlanModifier(description string, planMods ...planmodifier.String) resourceschema.StringAttribute
```
- **Purpose**: Optional sensitive string attribute with plan modifiers
- **Use Case**: Sensitive fields (passwords, tokens) that are optional with plan modifications
- **Flags**: `Optional=true`, `Sensitive=true`

### 7. `ResourceRequiredStringEnumWithPlanModifier`
```go
func ResourceRequiredStringEnumWithPlanModifier(description string, planMods []planmodifier.String, enumValues ...string) resourceschema.StringAttribute
```
- **Purpose**: Required string attribute with enum validation and plan modifiers
- **Use Case**: Required enum fields that need plan modification logic
- **Example**: Protocol versions that cannot be changed after creation

## Consistency with Bool Attribute Variants
The implementation follows the exact same pattern as bool attribute helpers:
- `ResourceOptionalBoolWithPlanModifier()` → `ResourceOptionalStringWithPlanModifier()`
- `ResourceComputedOptionalBoolWithPlanModifier()` → `ResourceComputedOptionalStringWithPlanModifier()`
- `ResourceComputedOptionalBoolWithDefaultAndPlanModifier()` → `ResourceComputedOptionalStringWithDefaultAndPlanModifier()`

## Integration Points
All functions use the existing `newResourceStringAttribute()` builder and `stringAttributeConfig` struct, maintaining consistency with the codebase architecture.

## Location
File: `/Users/phorton/Documents/GitHub/sonatype-nexus-community/terraform-provider-shared/schema/attributes.go`

Lines added:
- Lines 91-99: `ResourceSensitiveOptionalStringWithPlanModifier()`
- Lines 151-159: `ResourceRequiredStringEnumWithPlanModifier()`
- Lines 197-204: `ResourceOptionalStringWithPlanModifier()`
- Lines 206-215: `ResourceOptionalStringWithDefaultAndPlanModifier()`
- Lines 217-225: `ResourceComputedOptionalStringWithPlanModifier()`
- Lines 227-236: `ResourceComputedOptionalStringWithDefaultAndPlanModifier()`
- Lines 238-245: `ResourceComputedStringWithPlanModifier()`

## Impact
These helpers eliminate the need for direct schema attribute construction in 8+ locations across terraform-provider-sonatyperepo, enabling future code consolidation and reducing boilerplate while maintaining type safety and validation.
