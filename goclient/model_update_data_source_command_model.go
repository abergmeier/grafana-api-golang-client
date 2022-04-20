/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goclient

// Also acts as api DTO
type UpdateDataSourceCommandModel struct {
	Access            *DsAccessModel    `json:"access,omitempty"`
	BasicAuth         bool              `json:"basicAuth,omitempty"`
	BasicAuthPassword string            `json:"basicAuthPassword,omitempty"`
	BasicAuthUser     string            `json:"basicAuthUser,omitempty"`
	Database          string            `json:"database,omitempty"`
	IsDefault         bool              `json:"isDefault,omitempty"`
	JsonData          *JsonModel        `json:"jsonData,omitempty"`
	Name              string            `json:"name,omitempty"`
	Password          string            `json:"password,omitempty"`
	SecureJsonData    map[string]string `json:"secureJsonData,omitempty"`
	Type_             string            `json:"type,omitempty"`
	Uid               string            `json:"uid,omitempty"`
	Url               string            `json:"url,omitempty"`
	User              string            `json:"user,omitempty"`
	Version           int64             `json:"version,omitempty"`
	WithCredentials   bool              `json:"withCredentials,omitempty"`
}
