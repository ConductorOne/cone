// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// Outcome - The outcome of the certification.
type Outcome string

const (
	OutcomeCertifyOutcomeUnspecified  Outcome = "CERTIFY_OUTCOME_UNSPECIFIED"
	OutcomeCertifyOutcomeCertified    Outcome = "CERTIFY_OUTCOME_CERTIFIED"
	OutcomeCertifyOutcomeDecertified  Outcome = "CERTIFY_OUTCOME_DECERTIFIED"
	OutcomeCertifyOutcomeError        Outcome = "CERTIFY_OUTCOME_ERROR"
	OutcomeCertifyOutcomeCancelled    Outcome = "CERTIFY_OUTCOME_CANCELLED"
	OutcomeCertifyOutcomeWaitTimedOut Outcome = "CERTIFY_OUTCOME_WAIT_TIMED_OUT"
)

func (e Outcome) ToPointer() *Outcome {
	return &e
}
func (e *Outcome) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CERTIFY_OUTCOME_UNSPECIFIED":
		fallthrough
	case "CERTIFY_OUTCOME_CERTIFIED":
		fallthrough
	case "CERTIFY_OUTCOME_DECERTIFIED":
		fallthrough
	case "CERTIFY_OUTCOME_ERROR":
		fallthrough
	case "CERTIFY_OUTCOME_CANCELLED":
		fallthrough
	case "CERTIFY_OUTCOME_WAIT_TIMED_OUT":
		*e = Outcome(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Outcome: %v", v)
	}
}

// The TaskTypeCertify message indicates that a task is a certify task and all related details.
type TaskTypeCertify struct {
	// The ID of the access review.
	AccessReviewID *string `json:"accessReviewId,omitempty"`
	// The ID of the specific access review object that owns this certify task. This is also set on a revoke task if the revoke task is created from the denied outcome of a certify task.
	AccessReviewSelection *string `json:"accessReviewSelection,omitempty"`
	// The ID of the app entitlement.
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The ID of the app.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app user.
	AppUserID *string `json:"appUserId,omitempty"`
	// The ID of the user.
	IdentityUserID *string `json:"identityUserId,omitempty"`
	// The outcome of the certification.
	Outcome     *Outcome   `json:"outcome,omitempty"`
	OutcomeTime *time.Time `json:"outcomeTime,omitempty"`
}

func (t TaskTypeCertify) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskTypeCertify) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskTypeCertify) GetAccessReviewID() *string {
	if o == nil {
		return nil
	}
	return o.AccessReviewID
}

func (o *TaskTypeCertify) GetAccessReviewSelection() *string {
	if o == nil {
		return nil
	}
	return o.AccessReviewSelection
}

func (o *TaskTypeCertify) GetAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.AppEntitlementID
}

func (o *TaskTypeCertify) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *TaskTypeCertify) GetAppUserID() *string {
	if o == nil {
		return nil
	}
	return o.AppUserID
}

func (o *TaskTypeCertify) GetIdentityUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityUserID
}

func (o *TaskTypeCertify) GetOutcome() *Outcome {
	if o == nil {
		return nil
	}
	return o.Outcome
}

func (o *TaskTypeCertify) GetOutcomeTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.OutcomeTime
}
