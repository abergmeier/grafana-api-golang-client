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

import (
	"time"
)

// DashboardVersionMeta extends the dashboard version model with the names associated with the UserIds, overriding the field with the same name from the DashboardVersion model.
type DashboardVersionMetaModel struct {
	Created       time.Time  `json:"created,omitempty"`
	CreatedBy     string     `json:"createdBy,omitempty"`
	DashboardId   int64      `json:"dashboardId,omitempty"`
	Data          *JsonModel `json:"data,omitempty"`
	Id            int64      `json:"id,omitempty"`
	Message       string     `json:"message,omitempty"`
	ParentVersion int64      `json:"parentVersion,omitempty"`
	RestoredFrom  int64      `json:"restoredFrom,omitempty"`
	Version       int64      `json:"version,omitempty"`
}
