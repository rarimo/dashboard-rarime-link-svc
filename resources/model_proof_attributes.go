/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ProofAttributes struct {
	// The date and time when the proof was created in the timestamp format
	CreatedAt string `json:"created_at"`
	// The ID of the user who created the proof
	Creator string `json:"creator"`
	// The ID of the organization that issued the proof's claim
	OrgId string `json:"org_id"`
	// The proof object in JSON string format
	Proof string `json:"proof"`
	// The type of the proof
	ProofType string `json:"proof_type"`
}
