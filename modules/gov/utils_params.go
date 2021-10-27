package gov

import (
	"fmt"

	"github.com/forbole/bdjuno/v2/types"
)

// updateParams updates the governance parameters for the given height
func (m *Module) updateParams(height int64) error {
	depositParams, err := m.source.DepositParams(height)
	if err != nil {
		return fmt.Errorf("error while getting gov deposit params: %s", err)
	}

	votingParams, err := m.source.VotingParams(height)
	if err != nil {
		return fmt.Errorf("error while getting gov voting params: %s", err)
	}

	tallyParams, err := m.source.TallyParams(height)
	if err != nil {
		return fmt.Errorf("error while getting gov tally params: %s", err)
	}

	return m.db.SaveGovParams(types.NewGovParams(
		types.NewVotingParams(votingParams),
		types.NewDepositParam(depositParams),
		types.NewTallyParams(tallyParams),
		height,
	))
}
