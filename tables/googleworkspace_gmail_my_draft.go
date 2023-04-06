package tables

import (
	"context"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"google.golang.org/api/gmail/v1"
)

type TableGoogleworkspaceGmailMyDraftGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspaceGmailMyDraftGenerator{}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetTableName() string {
	return "googleworkspace_gmail_my_draft"
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			var query string

			maxResults := int64(500)

			resp := service.Users.Drafts.List("me").Q(query).MaxResults(maxResults)
			if err := resp.Pages(ctx, func(page *gmail.ListDraftsResponse) error {
				for _, draft := range page.Drafts {
					resultChannel <- draft

					if googleworkspace_client.IsCancelled(ctx) {
						page.NextPageToken = ""
						break
					}
				}
				return nil
			}); err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getGmailMyDraft(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	var draftID string
	if result != nil {
		draftID = result.(*gmail.Draft).Id
	}

	if draftID == "" {
		return nil, nil
	}

	resp, err := service.Users.Drafts.Get("me", draftID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("message_thread_id").ColumnType(schema.ColumnTypeString).Description("The ID of the thread the message belongs to.").
			Extractor(column_value_extractor.StructSelector("Message.ThreadId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_raw").ColumnType(schema.ColumnTypeString).Description("The entire email message in an RFC 2822 formatted and base64url encoded string.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.Raw")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_snippet").ColumnType(schema.ColumnTypeString).Description("A short part of the message text.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.Snippet")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_internal_date").ColumnType(schema.ColumnTypeTimestamp).Description("The internal message creation timestamp which determines ordering in the inbox.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.InternalDate")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_size_estimate").ColumnType(schema.ColumnTypeInt).Description("Estimated size in bytes of the message.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.SizeEstimate")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_label_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of IDs of labels applied to this message.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.LabelIds")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_payload").ColumnType(schema.ColumnTypeJSON).Description("The parsed email structure in the message parts.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.Payload")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("draft_id").ColumnType(schema.ColumnTypeString).Description("The immutable ID of the draft.").
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_id").ColumnType(schema.ColumnTypeString).Description("The immutable ID of the message.").
			Extractor(column_value_extractor.StructSelector("Message.Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_history_id").ColumnType(schema.ColumnTypeString).Description("The ID of the last history record that modified this message.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyDraft(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Message.HistoryId")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableGoogleworkspaceGmailMyDraftGenerator) GetSubTables() []*schema.Table {
	return nil
}
