package actionset

var (
	defaultStoreSignedAttestAction Action = storeSignedAttestAction{
		config: ActionConfig{
			Name:                "storeSignedAttest",
			Random:              false,
			ParamCount:          0,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}
	defaultRePackAttestationAction Action = rePackAttestationAction{
		config: ActionConfig{
			Name:                "rePackAttestation",
			Random:              false,
			ParamCount:          0,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}
	_ Action = rePackAttestationAction{}
)

type storeSignedAttestAction struct {
	config ActionConfig
}

func (s storeSignedAttestAction) DefaultParam() []interface{} { return []interface{}{} }

func (s storeSignedAttestAction) RandomParam() []interface{} { return []interface{}{} }

func (s storeSignedAttestAction) Desc() string { return "# Store signed attestation" }

func (s storeSignedAttestAction) Name() string { return "storeSignedAttest" }

func (s storeSignedAttestAction) MaxParam() int { return 0 }

func (s storeSignedAttestAction) MinParam() int { return 0 }

func (s storeSignedAttestAction) GetConfig() ActionConfig { return s.config }

func (s storeSignedAttestAction) WithConfig(config ActionConfig) Action {
	s.config = config
	return s
}

func (s storeSignedAttestAction) ActionType() AType { return AttestAction }

type rePackAttestationAction struct {
	config ActionConfig
}

func (r rePackAttestationAction) DefaultParam() []interface{} { return []interface{}{} }

func (r rePackAttestationAction) RandomParam() []interface{} { return []interface{}{} }

func (r rePackAttestationAction) Desc() string { return "# Re-pack attestation" }

func (r rePackAttestationAction) Name() string { return "rePackAttestation" }

func (r rePackAttestationAction) MaxParam() int { return 0 }

func (r rePackAttestationAction) MinParam() int { return 0 }

func (r rePackAttestationAction) GetConfig() ActionConfig { return r.config }

func (r rePackAttestationAction) WithConfig(config ActionConfig) Action {
	r.config = config
	return r
}

func (r rePackAttestationAction) ActionType() AType { return AttestAction }
