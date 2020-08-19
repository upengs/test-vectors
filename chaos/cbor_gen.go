// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package chaos

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufState = []byte{129}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)
	if len(t.Unmarshallable) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Unmarshallable was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Unmarshallable))); err != nil {
		return err
	}
	for _, v := range t.Unmarshallable {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Unmarshallable: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Unmarshallable = make([]*UnmarshallableCBOR, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v UnmarshallableCBOR
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Unmarshallable[i] = &v
	}

	return nil
}
