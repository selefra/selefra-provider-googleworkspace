package tables

import (
	"context"
	"regexp"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
	"google.golang.org/api/gmail/v1"
)

type TableGoogleworkspaceGmailMyMessageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspaceGmailMyMessageGenerator{}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetTableName() string {
	return "googleworkspace_gmail_my_message"
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			var query string

			maxResults := int64(500)

			resp := service.Users.Messages.List("me").Q(query).MaxResults(maxResults)
			if err := resp.Pages(ctx, func(page *gmail.ListMessagesResponse) error {
				for _, message := range page.Messages {
					resultChannel <- message

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

func extractMessageSender(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(*gmail.Message)
	if data.Payload == nil {
		return nil, nil
	}

	for _, payloadHeader := range data.Payload.Headers {
		if payloadHeader.Name == "From" {
			regexExp := regexp.MustCompile(`\<(.*?) *\>`)
			senderEmail := regexExp.FindStringSubmatch(payloadHeader.Value)
			if len(senderEmail) > 1 {
				return senderEmail[1], nil
			}
		}
	}

	return nil, nil
}
func getGmailMyMessage(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	service, err := googleworkspace_client.GmailService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	var messageID string
	if result != nil {
		messageID = result.(*gmail.Message).Id
	}

	if messageID == "" {
		return nil, nil
	}

	resp, err := service.Users.Messages.Get("me", messageID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The immutable ID of the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("raw").ColumnType(schema.ColumnTypeString).Description("The entire email message in an RFC 2822 formatted and base64url encoded string.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snippet").ColumnType(schema.ColumnTypeString).Description("A short part of the message text.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of IDs of labels applied to this message.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payload").ColumnType(schema.ColumnTypeJSON).Description("The parsed email structure in the message parts.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thread_id").ColumnType(schema.ColumnTypeString).Description("The ID of the thread the message belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("history_id").ColumnType(schema.ColumnTypeString).Description("The ID of the last history record that modified this message.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sender_email").ColumnType(schema.ColumnTypeString).Description("Specifies the email address of the sender.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				if reflect_util.IsNil(r) {
					return nil, nil
				}

				r, err = extractMessageSender(ctx, clientMeta, taskClient, task, row, column, r)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internal_date").ColumnType(schema.ColumnTypeTimestamp).Description("The internal message creation timestamp which determines ordering in the inbox.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("InternalDate")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_estimate").ColumnType(schema.ColumnTypeInt).Description("Estimated size in bytes of the message.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getGmailMyMessage(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableGoogleworkspaceGmailMyMessageGenerator) GetSubTables() []*schema.Table {
	return nil
}
