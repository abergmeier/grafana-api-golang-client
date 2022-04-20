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

type TokenModel struct {
	Account                   string            `json:"account,omitempty"`
	Company                   string            `json:"company,omitempty"`
	DetailsUrl                string            `json:"details_url,omitempty"`
	Exp                       int64             `json:"exp,omitempty"`
	Iat                       int64             `json:"iat,omitempty"`
	IncludedAdmins            int64             `json:"included_admins,omitempty"`
	IncludedUsers             int64             `json:"included_users,omitempty"`
	IncludedViewers           int64             `json:"included_viewers,omitempty"`
	Iss                       string            `json:"iss,omitempty"`
	Jti                       string            `json:"jti,omitempty"`
	Lexp                      int64             `json:"lexp,omitempty"`
	LicExpWarnDays            int64             `json:"lic_exp_warn_days,omitempty"`
	Lid                       string            `json:"lid,omitempty"`
	LimitBy                   string            `json:"limit_by,omitempty"`
	MaxConcurrentUserSessions int64             `json:"max_concurrent_user_sessions,omitempty"`
	Nbf                       int64             `json:"nbf,omitempty"`
	Prod                      []string          `json:"prod,omitempty"`
	Slug                      string            `json:"slug,omitempty"`
	Status                    *TokenStatusModel `json:"status,omitempty"`
	Sub                       string            `json:"sub,omitempty"`
	TokExpWarnDays            int64             `json:"tok_exp_warn_days,omitempty"`
	Trial                     bool              `json:"trial,omitempty"`
	TrialExp                  int64             `json:"trial_exp,omitempty"`
	UpdateDays                int64             `json:"update_days,omitempty"`
	UsageBilling              bool              `json:"usage_billing,omitempty"`
}
