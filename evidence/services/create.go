package services

import "path"

const (
	evidenceCreateAPI = "api/v1/subject"
)

type createEvidenceOperation struct {
	evidenceBody EvidenceCreationBody
}

func (c *createEvidenceOperation) getOperationRestApi() string {
	return path.Join(evidenceCreateAPI, c.evidenceBody.SubjectUri)
}

func (c *createEvidenceOperation) getRequestBody() *[]byte {
	return c.evidenceBody.DSSEFileRaw
}

func (c *createEvidenceOperation) getOperationSuccessfulMsg() string {
	return "Evidence is successfully created"
}

func (rbs *EvidenceService) UploadEvidence(evidenceDetails EvidenceDetails) ([]byte, error) {
	operation := createEvidenceOperation{
		evidenceBody: EvidenceCreationBody{
			EvidenceDetails: evidenceDetails,
		},
	}
	body, err := rbs.doOperation(&operation)
	return body, err
}

type SourceBuildDetails struct {
	BuildName   string
	BuildNumber string
	ProjectKey  string
}

type EvidenceCreationBody struct {
	EvidenceDetails
}
