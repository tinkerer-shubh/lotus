// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package sealer

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"

	sealtasks "github.com/filecoin-project/lotus/storage/sealer/sealtasks"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *Call) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.ID (storiface.CallID) (struct)
	if len("ID") > 8192 {
		return xerrors.Errorf("Value in field \"ID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ID"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ID")); err != nil {
		return err
	}

	if err := t.ID.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.State (sealer.CallState) (uint64)
	if len("State") > 8192 {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("State")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Result (sealer.ManyBytes) (struct)
	if len("Result") > 8192 {
		return xerrors.Errorf("Value in field \"Result\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Result"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Result")); err != nil {
		return err
	}

	if err := t.Result.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.RetType (sealer.ReturnType) (string)
	if len("RetType") > 8192 {
		return xerrors.Errorf("Value in field \"RetType\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RetType"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RetType")); err != nil {
		return err
	}

	if len(t.RetType) > 8192 {
		return xerrors.Errorf("Value in field t.RetType was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.RetType))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.RetType)); err != nil {
		return err
	}
	return nil
}

func (t *Call) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Call{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Call: map struct too large (%d)", extra)
	}

	n := extra

	nameBuf := make([]byte, 7)
	for i := uint64(0); i < n; i++ {
		nameLen, ok, err := cbg.ReadFullStringIntoBuf(cr, nameBuf, 8192)
		if err != nil {
			return err
		}

		if !ok {
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(cr, func(cid.Cid) {}); err != nil {
				return err
			}
			continue
		}

		switch string(nameBuf[:nameLen]) {
		// t.ID (storiface.CallID) (struct)
		case "ID":

			{

				if err := t.ID.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ID: %w", err)
				}

			}
			// t.State (sealer.CallState) (uint64)
		case "State":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.State = CallState(extra)

			}
			// t.Result (sealer.ManyBytes) (struct)
		case "Result":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Result = new(ManyBytes)
					if err := t.Result.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Result pointer: %w", err)
					}
				}

			}
			// t.RetType (sealer.ReturnType) (string)
		case "RetType":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.RetType = ReturnType(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(r, func(cid.Cid) {}); err != nil {
				return err
			}
		}
	}

	return nil
}
func (t *WorkState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{166}); err != nil {
		return err
	}

	// t.ID (sealer.WorkID) (struct)
	if len("ID") > 8192 {
		return xerrors.Errorf("Value in field \"ID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ID"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ID")); err != nil {
		return err
	}

	if err := t.ID.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Status (sealer.WorkStatus) (string)
	if len("Status") > 8192 {
		return xerrors.Errorf("Value in field \"Status\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Status"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Status")); err != nil {
		return err
	}

	if len(t.Status) > 8192 {
		return xerrors.Errorf("Value in field t.Status was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Status))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Status)); err != nil {
		return err
	}

	// t.StartTime (int64) (int64)
	if len("StartTime") > 8192 {
		return xerrors.Errorf("Value in field \"StartTime\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("StartTime"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("StartTime")); err != nil {
		return err
	}

	if t.StartTime >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.StartTime)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.StartTime-1)); err != nil {
			return err
		}
	}

	// t.WorkError (string) (string)
	if len("WorkError") > 8192 {
		return xerrors.Errorf("Value in field \"WorkError\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("WorkError"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("WorkError")); err != nil {
		return err
	}

	if len(t.WorkError) > 8192 {
		return xerrors.Errorf("Value in field t.WorkError was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.WorkError))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.WorkError)); err != nil {
		return err
	}

	// t.WorkerCall (storiface.CallID) (struct)
	if len("WorkerCall") > 8192 {
		return xerrors.Errorf("Value in field \"WorkerCall\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("WorkerCall"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("WorkerCall")); err != nil {
		return err
	}

	if err := t.WorkerCall.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.WorkerHostname (string) (string)
	if len("WorkerHostname") > 8192 {
		return xerrors.Errorf("Value in field \"WorkerHostname\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("WorkerHostname"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("WorkerHostname")); err != nil {
		return err
	}

	if len(t.WorkerHostname) > 8192 {
		return xerrors.Errorf("Value in field t.WorkerHostname was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.WorkerHostname))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.WorkerHostname)); err != nil {
		return err
	}
	return nil
}

func (t *WorkState) UnmarshalCBOR(r io.Reader) (err error) {
	*t = WorkState{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("WorkState: map struct too large (%d)", extra)
	}

	n := extra

	nameBuf := make([]byte, 14)
	for i := uint64(0); i < n; i++ {
		nameLen, ok, err := cbg.ReadFullStringIntoBuf(cr, nameBuf, 8192)
		if err != nil {
			return err
		}

		if !ok {
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(cr, func(cid.Cid) {}); err != nil {
				return err
			}
			continue
		}

		switch string(nameBuf[:nameLen]) {
		// t.ID (sealer.WorkID) (struct)
		case "ID":

			{

				if err := t.ID.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ID: %w", err)
				}

			}
			// t.Status (sealer.WorkStatus) (string)
		case "Status":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.Status = WorkStatus(sval)
			}
			// t.StartTime (int64) (int64)
		case "StartTime":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.StartTime = int64(extraI)
			}
			// t.WorkError (string) (string)
		case "WorkError":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.WorkError = string(sval)
			}
			// t.WorkerCall (storiface.CallID) (struct)
		case "WorkerCall":

			{

				if err := t.WorkerCall.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.WorkerCall: %w", err)
				}

			}
			// t.WorkerHostname (string) (string)
		case "WorkerHostname":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.WorkerHostname = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(r, func(cid.Cid) {}); err != nil {
				return err
			}
		}
	}

	return nil
}
func (t *WorkID) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.Method (sealtasks.TaskType) (string)
	if len("Method") > 8192 {
		return xerrors.Errorf("Value in field \"Method\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Method"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Method")); err != nil {
		return err
	}

	if len(t.Method) > 8192 {
		return xerrors.Errorf("Value in field t.Method was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Method))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Method)); err != nil {
		return err
	}

	// t.Params (string) (string)
	if len("Params") > 8192 {
		return xerrors.Errorf("Value in field \"Params\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Params"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Params")); err != nil {
		return err
	}

	if len(t.Params) > 8192 {
		return xerrors.Errorf("Value in field t.Params was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Params))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Params)); err != nil {
		return err
	}
	return nil
}

func (t *WorkID) UnmarshalCBOR(r io.Reader) (err error) {
	*t = WorkID{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("WorkID: map struct too large (%d)", extra)
	}

	n := extra

	nameBuf := make([]byte, 6)
	for i := uint64(0); i < n; i++ {
		nameLen, ok, err := cbg.ReadFullStringIntoBuf(cr, nameBuf, 8192)
		if err != nil {
			return err
		}

		if !ok {
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(cr, func(cid.Cid) {}); err != nil {
				return err
			}
			continue
		}

		switch string(nameBuf[:nameLen]) {
		// t.Method (sealtasks.TaskType) (string)
		case "Method":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.Method = sealtasks.TaskType(sval)
			}
			// t.Params (string) (string)
		case "Params":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.Params = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(r, func(cid.Cid) {}); err != nil {
				return err
			}
		}
	}

	return nil
}
