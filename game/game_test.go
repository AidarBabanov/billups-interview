package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetChoice(t *testing.T) {
	game := NewGame()
	for i := 0; i < 100; i++ {
		id, name := game.GetChoice()
		require.True(t, id <= 5 && id >= 1)
		require.Equal(t, choices[id-1], name)
	}
}

func TestPlay(t *testing.T) {
	pl1, pl2, res, desc := play(1, 2)
	require.Equal(t, 1, pl1)
	require.Equal(t, 2, pl2)
	require.Equal(t, "lose", res)
	require.Equal(t, "Paper covers rock.", desc)
}
