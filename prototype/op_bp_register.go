package prototype

import "github.com/pkg/errors"


func (m *BpRegisterOperation) GetSigner(auths *map[string]bool) {
	(*auths)[m.Owner.Value] = true
}

func (m *BpRegisterOperation) Validate() error {
	if m == nil {
		return ErrNpe
	}

	if err := m.Owner.Validate(); err != nil {
		return errors.WithMessage(err, "Owner error")
	}
	if err := m.BlockSigningKey.Validate(); err != nil {
		return errors.WithMessage(err, "BlockSigningKey error")
	}
	if m.Props == nil {
		return ErrNpe
	}
	if err := AtMost1KChars(m.Url); err != nil {
		return errors.WithMessage(err, "invalid url")
	}
	if err := AtMost4KChars(m.Desc); err != nil {
		return errors.WithMessage(err, "invalid description")
	}

	// TODO chain property valid check
	return nil
}

func (m *BpRegisterOperation) GetAffectedProps(props *map[string]bool) {
	(*props)["*"] = true
}

func init() {
	registerOperation("bp_register", (*Operation_Op3)(nil), (*BpRegisterOperation)(nil));
}
