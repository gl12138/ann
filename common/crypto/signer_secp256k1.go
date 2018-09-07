package crypto

import (
	"crypto/sha256"
	"github.com/annchain/OG/types"
	secp256k1 "github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ripemd160"
)

type SignerSecp256k1 struct {
}

func (s *SignerSecp256k1) GetCryptoType() CryptoType {
	return CryptoTypeSecp256k1
}

func (s *SignerSecp256k1) Sign(privKey PrivateKey, msg []byte) Signature {
	priv, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privKey.Bytes)
	sig, err := priv.Sign(Sha256(msg))
	if err != nil {
		panic(err)
	}

	return SignatureFromBytes(CryptoTypeSecp256k1, sig.Serialize())
}

func (s *SignerSecp256k1) PubKey(privKey PrivateKey) PublicKey {
	_, pub__ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privKey.Bytes)
	pub := [64]byte{}
	copy(pub[:], pub__.SerializeUncompressed()[1:])
	return PublicKeyFromBytes(CryptoTypeSecp256k1, pub[:])
}

func (s *SignerSecp256k1) Verify(pubKey PublicKey, signature Signature, msg []byte) bool {
	return ed25519.Verify(pubKey.Bytes, msg, signature.Bytes)

	pub__, err := secp256k1.ParsePubKey(pubKey.Bytes, secp256k1.S256())
	if err != nil {
		return false
	}
	sig__, err := secp256k1.ParseDERSignature(signature.Bytes, secp256k1.S256())
	if err != nil {
		return false
	}
	return sig__.Verify(Sha256(msg), pub__)
}

func (s *SignerSecp256k1) RandomKeyPair() (publicKey PublicKey, privateKey PrivateKey, err error) {
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return
	}
	publicKey = PublicKeyFromBytes(CryptoTypeSecp256k1, public)
	privateKey = PrivateKeyFromBytes(CryptoTypeSecp256k1, private)
	return
}

// Address calculate the address from the pubkey
func (s *SignerSecp256k1) Address(pubKey PublicKey) types.Address {

	hasherSHA256 := sha256.New()
	hasherSHA256.Write([]byte{byte(pubKey.Type)})
	hasherSHA256.Write(pubKey.Bytes) // does not error
	sha := hasherSHA256.Sum(nil)

	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha) // does not error
	result := hasherRIPEMD160.Sum(nil)
	return types.BytesToAddress(result)
}