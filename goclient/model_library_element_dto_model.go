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

type LibraryElementDtoModel struct {
	Description string                      `json:"description,omitempty"`
	FolderId    int64                       `json:"folderId,omitempty"`
	Id          int64                       `json:"id,omitempty"`
	Kind        int64                       `json:"kind,omitempty"`
	Meta        *LibraryElementDtoMetaModel `json:"meta,omitempty"`
	Model       interface{}                 `json:"model,omitempty"`
	Name        string                      `json:"name,omitempty"`
	OrgId       int64                       `json:"orgId,omitempty"`
	Type_       string                      `json:"type,omitempty"`
	Uid         string                      `json:"uid,omitempty"`
	Version     int64                       `json:"version,omitempty"`
}
