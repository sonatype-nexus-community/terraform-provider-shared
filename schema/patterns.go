/*
 * Copyright (c) 2019-present Sonatype, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schema

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// StandardResourceAttributes returns a map of standard attributes for creatable/updatable resources
// Includes: id (computed), last_updated (computed)
func StandardResourceAttributes() map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"id":           ResourceStandardID(),
		"last_updated": ResourceLastUpdated(),
	}
}

// NamedResourceAttributes returns a map of standard attributes plus a required name field
func NamedResourceAttributes() map[string]resourceschema.Attribute {
	attrs := StandardResourceAttributes()
	attrs["name"] = ResourceRequiredString("Name of the resource")
	return attrs
}

// IdentityResourceAttributes returns a map for username/id based resources
func IdentityResourceAttributes() map[string]resourceschema.Attribute {
	attrs := StandardResourceAttributes()
	attrs["username"] = ResourceRequiredString("Username identifier")
	return attrs
}

// OwnershipResourceAttributes returns a map for resources with ownership tracking
func OwnershipResourceAttributes() map[string]resourceschema.Attribute {
	attrs := StandardResourceAttributes()
	attrs["owner_id"] = ResourceRequiredString("ID of the resource owner")
	attrs["owner_type"] = ResourceRequiredString("Type of the owner (user, group, organization, etc.)")
	return attrs
}

// AuditableResourceAttributes returns a map for resources with audit trail
// Includes: id, created_at, updated_at
func AuditableResourceAttributes() map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"id":         ResourceStandardID(),
		"created_at": ResourceComputedString("Timestamp when the resource was created"),
		"updated_at": ResourceLastUpdated(),
	}
}

// DataSourceIDAttribute returns attributes for data source filters
func DataSourceIDAttribute() datasourceschema.StringAttribute {
	return datasourceschema.StringAttribute{
		Description: "ID to lookup the data source",
		Optional:    true,
		Computed:    true,
	}
}

// SourceAndReadOnlyAttributes returns common computed attributes for resources managed externally
func SourceAndReadOnlyAttributes() map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"source":     ResourceComputedOptionalString("Source system that created/manages this resource"),
		"read_only":  ResourceComputedOptionalBool("Whether the resource is read-only"),
		"managed_by": ResourceComputedOptionalString("System managing this resource"),
	}
}
