package utils

import (
	"Skyline/internal/utils"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestConfiguration(t *testing.T) {
	// CASE 0 Initialize
	appSetting, err := utils.LoadAppConfig("../../../internal/configs")

	// CASE 1 LoadAppConfig
	t.Run("LoadAppConfig", func(t *testing.T) {
		require.NoError(t, err)
		require.NotEmpty(t, appSetting)

		require.NotNil(t, appSetting.Environment)
		require.NotNil(t, appSetting.DbConnection)
	})

	err = utils.InitDB(appSetting.DbConnection)

	// CASE 1 InitDB
	t.Run("InitDB", func(t *testing.T) {
		require.NoError(t, err)
		require.NotEmpty(t, utils.DB)
		require.Equal(t, "*gorm.DB", reflect.TypeOf(utils.DB).String())
	})
}
