// Code generated by fastssz. DO NOT EDIT.
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockHeader object
func (b *BeaconBlockHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockHeader object to a target array
func (b *BeaconBlockHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, b.ProposerIndex)

	// Field (2) 'ParentRoot'
	if len(b.ParentRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, b.ParentRoot...)

	// Field (3) 'StateRoot'
	if len(b.StateRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, b.StateRoot...)

	// Field (4) 'BodyRoot'
	if len(b.BodyRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, b.BodyRoot...)

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockHeader object
func (b *BeaconBlockHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'ParentRoot'
	if cap(b.ParentRoot) == 0 {
		b.ParentRoot = make([]byte, 0, len(buf[16:48]))
	}
	b.ParentRoot = append(b.ParentRoot, buf[16:48]...)

	// Field (3) 'StateRoot'
	if cap(b.StateRoot) == 0 {
		b.StateRoot = make([]byte, 0, len(buf[48:80]))
	}
	b.StateRoot = append(b.StateRoot, buf[48:80]...)

	// Field (4) 'BodyRoot'
	if cap(b.BodyRoot) == 0 {
		b.BodyRoot = make([]byte, 0, len(buf[80:112]))
	}
	b.BodyRoot = append(b.BodyRoot, buf[80:112]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockHeader object
func (b *BeaconBlockHeader) SizeSSZ() (size int) {
	size = 112
	return
}

// HashTreeRoot ssz hashes the BeaconBlockHeader object
func (b *BeaconBlockHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockHeader object with a hasher
func (b *BeaconBlockHeader) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (1) 'ProposerIndex'
	hh.PutUint64(b.ProposerIndex)

	// Field (2) 'ParentRoot'
	if len(b.ParentRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(b.ParentRoot)

	// Field (3) 'StateRoot'
	if len(b.StateRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(b.StateRoot)

	// Field (4) 'BodyRoot'
	if len(b.BodyRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(b.BodyRoot)

	hh.Merkleize(indx)
	return
}
