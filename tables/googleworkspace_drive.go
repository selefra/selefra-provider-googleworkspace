package tables

import (
	"context"
	"strings"
	"time"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableGoogleworkspaceDriveGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspaceDriveGenerator{}

func (x *TableGoogleworkspaceDriveGenerator) GetTableName() string {
	return "googleworkspace_drive"
}

func (x *TableGoogleworkspaceDriveGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspaceDriveGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspaceDriveGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspaceDriveGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.DriveService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			var queryFilter, query string
			var filter []string

			if queryFilter != "" {
				query = queryFilter
			} else if len(filter) > 0 {
				query = strings.Join(filter, " and ")
			}

			requiredFields := []googleapi.Field{}

			var useDomainAdminAccess bool

			pageSize := int64(100)

			resp := service.Drives.List().Fields(requiredFields...).Q(query).UseDomainAdminAccess(useDomainAdminAccess).PageSize(pageSize)
			if err := resp.Pages(ctx, func(page *drive.DriveList) error {
				for _, data := range page.Drives {
					parsedTime, _ := time.Parse(time.RFC3339, data.CreatedTime)
					data.CreatedTime = parsedTime.Format(time.RFC3339)
					resultChannel <- data

					if googleworkspace_client.IsCancelled(ctx) {
						page.NextPageToken = ""
						break
					}
				}
				return nil
			}); err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableGoogleworkspaceDriveGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspaceDriveGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("background_image_link").ColumnType(schema.ColumnTypeString).Description("A short-lived link to this shared drive's background image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("copy_requires_writer_permission").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the options to copy, print, or download files inside this shared drive, should be disabled for readers and commenters, or not.").
			Extractor(column_value_extractor.StructSelector("Restrictions.CopyRequiresWriterPermission")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_users_only").ColumnType(schema.ColumnTypeBool).Description("Indicates whether access to this shared drive and items inside this shared drive is restricted to users of the domain to which this shared drive belongs.").
			Extractor(column_value_extractor.StructSelector("Restrictions.DomainUsersOnly")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of this shared drive which is also the ID of the top level folder of this shared drive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of this shared drive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time at which the shared drive was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hidden").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the shared drive is hidden from default view, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("admin_managed_restrictions").ColumnType(schema.ColumnTypeBool).Description("Indicates whether administrative privileges on this shared drive are required to modify restrictions, or not.").
			Extractor(column_value_extractor.StructSelector("Restrictions.AdminManagedRestrictions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("theme_id").ColumnType(schema.ColumnTypeString).Description("The ID of the theme from which the background image and color will be set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("background_image_file").ColumnType(schema.ColumnTypeJSON).Description("An image file and cropping parameters from which a background image for this shared drive is set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("color_rgb").ColumnType(schema.ColumnTypeString).Description("The color of this shared drive as an RGB hex string.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("drive_members_only").ColumnType(schema.ColumnTypeBool).Description("Indicates whether access to items inside this shared drive is restricted to its members, or not.").
			Extractor(column_value_extractor.StructSelector("Restrictions.DriveMembersOnly")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capabilities").ColumnType(schema.ColumnTypeJSON).Description("Describes the capabilities the current user has on this shared drive.").Build(),
	}
}

func (x *TableGoogleworkspaceDriveGenerator) GetSubTables() []*schema.Table {
	return nil
}
