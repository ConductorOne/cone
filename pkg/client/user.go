package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/types/known/anypb"
)

type UserResponse struct {
	UserView UserView     `json:"user_view"`
	Expanded []*anypb.Any `json:"expanded"`
}

type UserView struct {
	User              User   `json:"user"`
	RolesPath         string `json:"roles_path"`
	ManagersPath      string `json:"managers_path"`
	DelegatedUserPath string `json:"delegated_user_path"`
	DirectoriesPath   string `json:"directories_path"`
}

type User struct {
	ID                      string   `json:"id"`
	DisplayName             string   `json:"display_name"`
	Email                   string   `json:"email"`
	Status                  string   `json:"status"`
	RoleIds                 []string `json:"role_ids"`
	ManagerIds              []string `json:"manager_ids"`
	ManagerSources          []string `json:"manager_sources"`
	CreatedAt               string   `json:"created_at"`
	UpdatedAt               string   `json:"updated_at"`
	DeletedAt               string   `json:"deleted_at"`
	DelegatedUserId         string   `json:"delegated_user_id"`
	DirectoryStatus         string   `json:"directory_status"`
	DirectoryStatusSources  []string `json:"directory_status_sources"`
	JobTitle                string   `json:"job_title"`
	JobTitleSources         []string `json:"job_title_sources"`
	Department              string   `json:"department"`
	DepartmentSources       []string `json:"department_sources"`
	EmploymentStatus        string   `json:"employment_status"`
	EmploymentStatusSources []string `json:"employment_status_sources"`
	EmploymentType          string   `json:"employment_type"`
	EmploymentTypeSources   []string `json:"employment_type_sources"`
	DirectoryIds            []string `json:"directory_ids"`
}

func (c *client) GetUser(ctx context.Context, userID string) (*UserResponse, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.apiHost(),
		Path:   fmt.Sprintf("/api/v1/user/get/%s", userID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	userResp := &UserResponse{}
	err = json.Unmarshal(body, userResp)
	if err != nil {
		return nil, err
	}

	return userResp, nil
}
