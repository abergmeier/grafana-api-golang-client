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

// DataTimeSeries -- this structure is deprecated, all new work should use DataFrames from the SDK
type DataTimeSeriesModel struct {
	Name   string                     `json:"name,omitempty"`
	Points *DataTimeSeriesPointsModel `json:"points,omitempty"`
	Tags   map[string]string          `json:"tags,omitempty"`
}
