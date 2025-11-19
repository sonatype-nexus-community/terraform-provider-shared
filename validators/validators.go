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

package validators

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// StatusValidators returns validators for status fields
func StatusValidators(validStatuses ...string) []validator.String {
	return []validator.String{
		stringvalidator.OneOf(validStatuses...),
	}
}

// SourceTypeValidators returns validators for source type fields
// Common values: local, ldap, crowd, saml
func SourceTypeValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("local", "ldap", "crowd", "saml"),
	}
}

// RealmValidators returns validators for realm fields
// Common values: internal, external
func RealmValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("internal", "external"),
	}
}

// BooleanStringValidators returns validators for string booleans
// Accepts: "true", "false"
func BooleanStringValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("true", "false", "True", "False", "TRUE", "FALSE"),
	}
}

// ProviderValidators returns validators for provider type fields
// Common values: github, gitlab, bitbucket, etc.
func ProviderValidators(providers ...string) []validator.String {
	return []validator.String{
		stringvalidator.OneOf(providers...),
	}
}

// RoleValidators returns validators for role fields
// Common values: admin, user, viewer, editor, etc.
func RoleValidators(roles ...string) []validator.String {
	return []validator.String{
		stringvalidator.OneOf(roles...),
	}
}

// URLProtocolValidators returns validators for URL protocols
func URLProtocolValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("http", "https", "ftp", "sftp"),
	}
}

// AuthenticationMethodValidators returns validators for authentication methods
func AuthenticationMethodValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("basic", "token", "apikey", "oauth2", "saml", "ldap"),
	}
}

// EnvironmentValidators returns validators for environment fields
func EnvironmentValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("dev", "development", "test", "testing", "staging", "stage", "production", "prod"),
	}
}

// SortOrderValidators returns validators for sort order fields
func SortOrderValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("asc", "ascending", "desc", "descending"),
	}
}

// SeverityValidators returns validators for severity/priority levels
func SeverityValidators() []validator.String {
	return []validator.String{
		stringvalidator.OneOf("critical", "high", "medium", "low", "info"),
	}
}
