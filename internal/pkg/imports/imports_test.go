package imports_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/imports"
	"github.com/stretchr/testify/assert"
)

const (
	viperPkg = "github.com/spf13/viper"
)

func Test_imports_RegisterPrefixAlias(t *testing.T) {
	t.Run("Given scenario", func(t *testing.T) {
		i := imports.New()
		shortcut := "viper"
		expectedViper := "i0_spf13_viper"
		expectedRemote := "i1_viper_remote"

		if !assert.NoError(t, i.RegisterPrefixAlias(shortcut, viperPkg)) {
			return
		}

		assert.Empty(t, i.Imports()) // no import has been used yet

		// aliases for "viper" and "github.com/spf13/viper" should be equal
		assert.Equal(t, expectedViper, i.Alias(shortcut))
		assert.Equal(t, expectedViper, i.Alias(viperPkg))

		// aliases for "viper/remote" and "github.com/spf13/viper/remote" should be equal
		assert.Equal(t, expectedRemote, i.Alias(shortcut+"/remote"))
		assert.Equal(t, expectedRemote, i.Alias(viperPkg+"/remote"))

		assert.Equal(
			t,
			[]imports.Import{
				{Alias: "i0_spf13_viper", Path: "github.com/spf13/viper"},
				{Alias: "i1_viper_remote", Path: "github.com/spf13/viper/remote"},
			},
			i.Imports(),
		)
	})

	t.Run("Given error", func(t *testing.T) {
		i := imports.New()
		if !assert.NoError(t, i.RegisterPrefixAlias("viper", viperPkg)) {
			return
		}
		assert.EqualError(
			t,
			i.RegisterPrefixAlias("viper", viperPkg),
			`prefix is already registered: "viper"`,
		)
	})
}
