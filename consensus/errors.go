package consensus

import (
	cmn "github.com/lianxiangcloud/linkchain/libs/common"
)

type (
	ErrInvalidBlock error
	ErrProxyAppConn error

	ErrUnknownBlock struct {
		Height uint64
	}

	ErrBlockHashMismatch struct {
		CoreHash []byte
		AppHash  []byte
		Height   uint64
	}

	ErrAppBlockHeightTooHigh struct {
		CoreHeight uint64
		AppHeight  uint64
	}

	ErrLastStateMismatch struct {
		Height uint64
		Core   []byte
		App    []byte
	}

	ErrStateMismatch struct {
		Got      *NewStatus
		Expected *NewStatus
	}

	ErrNoValSetForHeight struct {
		Height uint64
	}

	ErrNoConsensusParamsForHeight struct {
		Height uint64
	}

	ErrNoABCIResponsesForHeight struct {
		Height uint64
	}
)

func (e ErrUnknownBlock) Error() string {
	return cmn.Fmt("Could not find block #%d", e.Height)
}

func (e ErrBlockHashMismatch) Error() string {
	return cmn.Fmt("App block hash (%X) does not match core block hash (%X) for height %d", e.AppHash, e.CoreHash, e.Height)
}

func (e ErrAppBlockHeightTooHigh) Error() string {
	return cmn.Fmt("App block height (%d) is higher than core (%d)", e.AppHeight, e.CoreHeight)
}
func (e ErrLastStateMismatch) Error() string {
	return cmn.Fmt("Latest block (%d) LastAppHash (%X) does not match app's AppHash (%X)", e.Height, e.Core, e.App)
}

func (e ErrStateMismatch) Error() string {
	return cmn.Fmt("State after replay does not match saved state. Got ----\n%v\nExpected ----\n%v\n", e.Got, e.Expected)
}

func (e ErrNoValSetForHeight) Error() string {
	return cmn.Fmt("Could not find validator set for height #%d", e.Height)
}

func (e ErrNoConsensusParamsForHeight) Error() string {
	return cmn.Fmt("Could not find consensus params for height #%d", e.Height)
}

func (e ErrNoABCIResponsesForHeight) Error() string {
	return cmn.Fmt("Could not find results for height #%d", e.Height)
}
