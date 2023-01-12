package model

import (
	credsdk "github.com/tavroi/ssi-sdk/credential"
	"github.com/tavroi/ssi-sdk/credential/exchange"
	"github.com/tavroi/ssi-sdk/util"
	"github.com/tavroi/ssi-service/internal/credential"
	"github.com/tavroi/ssi-service/internal/keyaccess"
	"github.com/tavroi/ssi-service/pkg/service/presentation/storage"
	"go.einride.tech/aip/filtering"
)

type CreatePresentationDefinitionRequest struct {
	PresentationDefinition exchange.PresentationDefinition `json:"presentationDefinition" validate:"required"`
}

func (cpr CreatePresentationDefinitionRequest) IsValid() bool {
	return util.IsValidStruct(cpr) == nil
}

type CreatePresentationDefinitionResponse struct {
	PresentationDefinition exchange.PresentationDefinition `json:"presentationDefinition"`
}

type GetPresentationDefinitionRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetPresentationDefinitionResponse struct {
	ID                     string                          `json:"id"`
	PresentationDefinition exchange.PresentationDefinition `json:"presentationDefinition"`
}

type DeletePresentationDefinitionRequest struct {
	ID string `json:"id" validate:"required"`
}

type CreateSubmissionRequest struct {
	Presentation  credsdk.VerifiablePresentation  `json:"presentation" validate:"required"`
	SubmissionJWT keyaccess.JWT                   `json:"submissionJwt,omitempty" validate:"required"`
	Submission    exchange.PresentationSubmission `json:"submission" validate:"required"`
	Credentials   []credential.Container          `json:"credentials,omitempty"`
}

func (csr CreateSubmissionRequest) IsValid() bool {
	return util.IsValidStruct(csr) == nil
}

type CreateSubmissionResponse struct {
	Submission exchange.PresentationSubmission `json:"submission"`
}

type GetSubmissionRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetSubmissionResponse struct {
	Submission Submission `json:"submission"`
}

type DeleteSubmissionRequest struct {
	ID string `json:"id" validate:"required"`
}

type ListSubmissionRequest struct {
	Filter filtering.Filter
}

type Submission struct {
	// One of {`pending`, `approved`, `denied`, `cancelled`}.
	Status string `json:"status" validate:"required"`
	// The reason why the submission was approved or denied.
	Reason string `json:"reason"`
	*exchange.PresentationSubmission
}

type ListSubmissionResponse struct {
	Submissions []Submission `json:"submissions"`
}

type ReviewSubmissionRequest struct {
	ID       string `json:"id" validate:"required"`
	Approved bool   `json:"approved"`
	Reason   string `json:"reason"`
}

// Validate runs validation on the request struct and returns errors when it's invalid.
func (r ReviewSubmissionRequest) Validate() error {
	return util.NewValidator().Struct(r)
}

// ServiceModel creates a Submission from a given StoredSubmission.
func ServiceModel(storedSubmission *storage.StoredSubmission) Submission {
	return Submission{
		Status:                 storedSubmission.Status.String(),
		Reason:                 storedSubmission.Reason,
		PresentationSubmission: &storedSubmission.Submission,
	}
}
