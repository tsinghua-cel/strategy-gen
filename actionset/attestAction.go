package actionset

var (
	_ Action = storeSignedAttestAction{}
	_ Action = rePackAttestationAction{}
)

type storeSignedAttestAction struct{}

func (s storeSignedAttestAction) DefaultParam() []interface{} { return []interface{}{} }

func (s storeSignedAttestAction) RandomParam() []interface{} { return []interface{}{} }

func (s storeSignedAttestAction) Desc() string { return "# Store signed attestation" }

func (s storeSignedAttestAction) Name() string { return "storeSignedAttest" }

func (s storeSignedAttestAction) MaxParam() int { return 0 }

func (s storeSignedAttestAction) MinParam() int { return 0 }

func (s storeSignedAttestAction) ActionType() AType { return AttestAction }

type rePackAttestationAction struct{}

func (r rePackAttestationAction) DefaultParam() []interface{} { return []interface{}{} }

func (r rePackAttestationAction) RandomParam() []interface{} { return []interface{}{} }

func (r rePackAttestationAction) Desc() string { return "# Re-pack attestation" }

func (r rePackAttestationAction) Name() string { return "rePackAttestation" }

func (r rePackAttestationAction) MaxParam() int { return 0 }

func (r rePackAttestationAction) MinParam() int { return 0 }

func (r rePackAttestationAction) ActionType() AType { return AttestAction }
