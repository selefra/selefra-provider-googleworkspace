package provider

import (
	"context"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-googleworkspace/constants"
	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Googleworkspace,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var googleworkspaceConfig googleworkspace_client.Config
				err := config.Unmarshal(&googleworkspaceConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				if googleworkspaceConfig.Credentials == "" {
					googleworkspaceConfig.Credentials = os.Getenv("GOOGLE_WORKSPACE_CREDENTIALS")
				}

				if googleworkspaceConfig.ImpersonatedUserEmail == "" {
					googleworkspaceConfig.ImpersonatedUserEmail = os.Getenv("GOOGLE_WORKSPACE_IMPERSONATED_USER_EMAIL")
				}

				if googleworkspaceConfig.TokenPath == "" {
					googleworkspaceConfig.TokenPath = os.Getenv("GOOGLE_WORKSPACE_TOKEN_PATH")
				}

				clients, err := googleworkspace_client.NewClients(googleworkspaceConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# token_path: <Google Workspace Token Path>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var googleworkspaceConfig googleworkspace_client.Config
				err := config.Unmarshal(&googleworkspaceConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_10,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
