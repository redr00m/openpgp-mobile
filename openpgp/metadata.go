package openpgp

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp"
)

func getIdentities(input map[string]*openpgp.Identity) []Identity {

	identities := []Identity{}
	for _, identity := range input {
		if identity == nil || identity.UserId == nil {
			continue
		}
		identities = append(identities, Identity{
			ID:      identity.UserId.Id,
			Comment: identity.UserId.Comment,
			Email:   identity.UserId.Email,
			Name:    identity.UserId.Name,
		})
	}
	return identities
}

func (o *FastOpenPGP) GetPublicKeyMetadata(key string) (*PublicKeyMetadata, error) {
	entityList, err := o.readArmoredKeyRing(key, openpgp.PublicKeyType)
	if err != nil {
		return nil, fmt.Errorf("publicKey error: %w", err)
	}
	if len(entityList) < 1 {
		return nil, fmt.Errorf("publicKey error: %w", errors.New("no key found"))
	}
	entity := entityList[0]
	publicKey := entity.PrimaryKey
	if publicKey == nil {
		return nil, fmt.Errorf("publicKey error: %w", errors.New("no publicKey found"))
	}

	var byteIDs []string
	for _, byteID := range publicKey.Fingerprint {
		byteIDs = append(byteIDs, fmt.Sprint(byteID))
	}

	return &PublicKeyMetadata{
		KeyID:        publicKey.KeyIdString(),
		KeyIDShort:   publicKey.KeyIdShortString(),
		KeyIDNumeric: fmt.Sprintf("%d", publicKey.KeyId),
		CreationTime: publicKey.CreationTime.Format(time.RFC3339),
		Fingerprint:  strings.Join(byteIDs, ":"),
		IsSubKey:     publicKey.IsSubkey,
		Identities:   getIdentities(entity.Identities),
	}, nil
}

func (o *FastOpenPGP) GetPrivateKeyMetadata(key string) (*PrivateKeyMetadata, error) {
	entityList, err := o.readArmoredKeyRing(key, openpgp.PrivateKeyType)
	if err != nil {
		return nil, fmt.Errorf("privateKey error: %w", err)
	}
	if len(entityList) < 1 {
		return nil, fmt.Errorf("privateKey error: %w", errors.New("no key found"))
	}
	entity := entityList[0]
	privateKey := entity.PrivateKey
	if privateKey == nil {
		return nil, fmt.Errorf("privateKey error: %w", errors.New("no privateKey found"))
	}

	var byteIDs []string
	for _, byteID := range privateKey.Fingerprint {
		byteIDs = append(byteIDs, fmt.Sprint(byteID))
	}

	return &PrivateKeyMetadata{
		KeyID:        privateKey.KeyIdString(),
		KeyIDShort:   privateKey.KeyIdShortString(),
		KeyIDNumeric: fmt.Sprintf("%d", privateKey.KeyId),
		CreationTime: privateKey.CreationTime.Format(time.RFC3339),
		Fingerprint:  strings.Join(byteIDs, ":"),
		IsSubKey:     privateKey.IsSubkey,
		Encrypted:    privateKey.Encrypted,
		Identities:   getIdentities(entity.Identities),
	}, nil
}
