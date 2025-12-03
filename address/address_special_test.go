package address

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressSpecial(t *testing.T) {
	r := require.New(t)

	poll, err := FromBytes(PollProtocolAddrHash[:])
	r.NoError(err)
	r.Equal(PollProtocol, poll.String())

	rolldpos, err := FromBytes(RollDPoSProtocolAddrHash[:])
	r.NoError(err)
	r.Equal(RollDPoSProtocol, rolldpos.String())
}
