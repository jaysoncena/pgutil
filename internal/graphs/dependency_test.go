package graphs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/graphs"
)

func TestNewDependencyGraph(t *testing.T) {
	require := require.New(t)

	var expected error = &graphs.InvalidEdgeError{Node: "foo", Edge: "bar"}
	nodes := map[string][]string{
		"foo": {"bar"},
	}
	g, err := graphs.NewDependencyGraph(nodes)
	require.Nil(g)
	require.Equal(expected, err)
	require.NotEmpty(err.Error())

	expected = &graphs.CircularDependencyError{
		Cycle: []string{"foo", "bar"},
	}
	nodes = map[string][]string{
		"foo": {"bar"},
		"bar": {"foo"},
	}
	g, err = graphs.NewDependencyGraph(nodes)
	require.Nil(g)
	require.Equal(expected, err)
	require.NotEmpty(err.Error())
}

func TestDependencyGraph(t *testing.T) {
	require := require.New(t)

	expected := map[string][]string{
		"":  {"g", "r"},
		"g": {"a", "p"},
		"r": nil,
		"a": nil,
		"p": {"h"},
	}
	nodes := map[string][]string{
		"a": {"g"},
		"g": {},
		"h": {"g", "p", "r"},
		"p": {"g"},
		"r": {},
	}
	g, err := graphs.NewDependencyGraph(nodes)
	require.NoError(err)

	completed := []string{""}
	for g.HasPending() {
		for len(completed) > 0 {
			n := completed[0]
			actual := g.Next(n)
			require.Equal(expected[n], actual)
			completed = append(completed[1:], actual...)
		}
	}
}
