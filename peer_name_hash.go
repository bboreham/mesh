// +build peer_name_hash

package mesh

// Let peer names be SHA256 hashes of anything, provided they are unique.

import (
	"crypto/sha256"
	"encoding/hex"
)

// PeerName must be globally unique and usable as a map key.
type PeerName string

const (
	// PeerNameFlavour is the type of peer names we use.
	PeerNameFlavour = "hash"
	// NameSize is the number of bytes in a peer name.
	NameSize = sha256.Size >> 1
	// UnknownPeerName is used as a sentinel value.
	UnknownPeerName = PeerName("")
)

// PeerNameFromUserInput parses PeerName from a user-provided string.
// TODO(pb): does this need to be exported?
func PeerNameFromUserInput(userInput string) (PeerName, error) {
	// fixed-length identity
	nameByteAry := sha256.Sum256([]byte(userInput))
	return PeerNameFromBin(nameByteAry[:NameSize]), nil
}

// PeerNameFromString parses PeerName from a generic string.
// TODO(pb): does this need to be exported?
func PeerNameFromString(nameStr string) (PeerName, error) {
	if _, err := hex.DecodeString(nameStr); err != nil {
		return UnknownPeerName, err
	}
	return PeerName(nameStr), nil
}

// PeerNameFromBin parses PeerName from a byte slice.
// TODO(pb): does this need to be exported?
func PeerNameFromBin(nameByte []byte) PeerName {
	return PeerName(hex.EncodeToString(nameByte))
}

// Bin encodes PeerName as a byte slice.
func (name PeerName) Bin() []byte {
	res, err := hex.DecodeString(string(name))
	checkFatal(err)
	return res
}

// String encodes PeerName as a string.
func (name PeerName) String() string {
	return string(name)
}
