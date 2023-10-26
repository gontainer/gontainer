package token_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestStrategyFactory_Prepend(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		f := token.NewStrategyFactory()

		// func is not registered yet
		tk, err := f.Create(`%getEnv("HOST")%`)
		assert.EqualError(t, err, `not supported token: %getEnv("HOST")%`)
		assert.Empty(t, tk)

		// let's register func
		f.Prepend(token.NewFactoryFunction(mockAliaser{}, "getEnv", "pkg", "GetEnv"))
		tk, err = f.Create(`%getEnv("HOST")%`)
		assert.NoError(t, err)
		assert.Equal(t, token.Token{
			Kind:      token.KindFunc,
			Raw:       `%getEnv("HOST")%`,
			DependsOn: nil,
			Code:      `func() (r interface{}, err error) { r, err = caller.CallProvider(pkg.GetEnv, []interface{}{"HOST"}, true); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %getEnv(\"HOST\")%", err) }; return }`,
		}, tk)
	})

	t.Run("RegisterFunc", func(t *testing.T) {
		f := token.NewStrategyFactory()

		// func is not registered yet
		tk, err := f.Create(`%secret("password")%`)
		assert.EqualError(t, err, `not supported token: %secret("password")%`)
		assert.Empty(t, tk)

		// let's register func using registerer
		token.NewFuncRegisterer(f, mockAliaser{}).RegisterFunc("secret", "secrets", "Secret")
		tk, err = f.Create(`%secret("password")%`)
		assert.NoError(t, err)
		assert.Equal(t, token.Token{
			Kind:      token.KindFunc,
			Raw:       `%secret("password")%`,
			DependsOn: nil,
			Code:      `func() (r interface{}, err error) { r, err = caller.CallProvider(secrets.Secret, []interface{}{"password"}, true); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %secret(\"password\")%", err) }; return }`,
		}, tk)

		// let's register another func using registerer,
		// since it will prepend existing slice, the newer func has the bigger priority
		token.NewFuncRegisterer(f, mockAliaser{}).RegisterFunc("secret", "anotherSecrets", "SuperSecret")
		tk, err = f.Create(`%secret("password")%`)
		assert.NoError(t, err)
		assert.Equal(t, token.Token{
			Kind:      token.KindFunc,
			Raw:       `%secret("password")%`,
			DependsOn: nil,
			Code:      `func() (r interface{}, err error) { r, err = caller.CallProvider(anotherSecrets.SuperSecret, []interface{}{"password"}, true); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %secret(\"password\")%", err) }; return }`,
		}, tk)
	})
}
