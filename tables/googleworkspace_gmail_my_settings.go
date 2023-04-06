package tables

import (
	"context"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"google.golang.org/api/googleapi"
)

type TableGoogleworkspaceGmailMySettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspaceGmailMySettingsGenerator{}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetTableName() string {
	return "googleworkspace_gmail_my_settings"
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			resp, err := service.Users.GetProfile("me").Do()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- resp

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getGmailMyAutoForwardingSetting(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetAutoForwarding("me").Do()
	if err != nil {
		return nil, err
	}

	if resp != nil {
		result := map[string]interface{}{
			"disposition":  resp.Disposition,
			"emailAddress": resp.EmailAddress,
			"enabled":      resp.Enabled,
		}
		return result, nil
	}

	return nil, nil
}
func getGmailMyImapSetting(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetImap("me").Do()
	if err != nil {
		return nil, err
	}

	if resp != nil {
		result := map[string]interface{}{
			"autoExpunge":     resp.AutoExpunge,
			"enabled":         resp.Enabled,
			"expungeBehavior": resp.ExpungeBehavior,
			"maxFolderSize":   resp.MaxFolderSize,
		}
		return result, nil
	}

	return nil, nil
}
func getGmailMyLanguage(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetLanguage("me").Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func getGmailMyPopSetting(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetPop("me").Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func getGmailMyVacationSetting(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetVacation("me").Do()
	if err != nil {
		return nil, err
	}

	if resp != nil {
		result := map[string]interface{}{
			"enableAutoReply":    resp.EnableAutoReply,
			"responseSubject":    resp.ResponseSubject,
			"restrictToContacts": resp.RestrictToContacts,
			"restrictToDomain":   resp.RestrictToDomain,
		}
		return result, nil
	}

	return nil, nil
}
func listGmailMyDelegateSettings(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.Delegates.List("me").Do()
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok {

			if gerr.Code == 403 && gerr.Message == "Access restricted to service accounts that have been delegated domain-wide authority" {
				return nil, nil
			}
		}
		return nil, err
	}

	return resp.Delegates, nil
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_language").ColumnType(schema.ColumnTypeString).Description("Specifies the language settings for the specified account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyLanguage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_forwarding").ColumnType(schema.ColumnTypeJSON).Description("Describes the auto-forwarding setting for the specified account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getGmailMyAutoForwardingSetting(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delegates").ColumnType(schema.ColumnTypeJSON).Description("A list of delegates for the specified account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := listGmailMyDelegateSettings(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("imap").ColumnType(schema.ColumnTypeJSON).Description("Describes the IMAP setting for the specified account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getGmailMyImapSetting(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pop").ColumnType(schema.ColumnTypeJSON).Description("Describes the POP settings for the specified account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getGmailMyPopSetting(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vacation").ColumnType(schema.ColumnTypeJSON).Description("Describes the vacation responder settings for the specified account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getGmailMyVacationSetting(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_email").ColumnType(schema.ColumnTypeString).Description("The user's email address.").
			Extractor(column_value_extractor.StructSelector("EmailAddress")).Build(),
	}
}

func (x *TableGoogleworkspaceGmailMySettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
