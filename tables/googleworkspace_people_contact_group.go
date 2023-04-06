package tables

import (
	"context"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"google.golang.org/api/people/v1"
)

type TableGoogleworkspacePeopleContactGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspacePeopleContactGroupGenerator{}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetTableName() string {
	return "googleworkspace_people_contact_group"
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.PeopleService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxMembers := int64(2500)

			pageLimit := int64(200)

			var contactGroupNames [][]string
			resp := service.ContactGroups.List().PageSize(pageLimit)
			if err := resp.Pages(ctx, func(page *people.ListContactGroupsResponse) error {
				var resourceNames []string

				for _, contactGroup := range page.ContactGroups {
					resourceNames = append(resourceNames, contactGroup.ResourceName)

					if googleworkspace_client.IsCancelled(ctx) {
						page.NextPageToken = ""
						break
					}
				}
				if len(resourceNames) > 0 {
					contactGroupNames = append(contactGroupNames, resourceNames)
				}
				return nil
			}); err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, contactGroups := range contactGroupNames {
				data, err := service.ContactGroups.BatchGet().ResourceNames(contactGroups...).MaxMembers(maxMembers).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				if data.Responses != nil && len(data.Responses) > 0 {
					for _, i := range data.Responses {
						resultChannel <- i.ContactGroup
					}
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("formatted_name").ColumnType(schema.ColumnTypeString).Description("The name translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale for system groups names.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_count").ColumnType(schema.ColumnTypeInt).Description("The total number of contacts in the group irrespective of max members in specified in the request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the contact group resource has been deleted, or not.").
			Extractor(column_value_extractor.StructSelector("Metadata.Deleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time the group was last updated.").
			Extractor(column_value_extractor.StructSelector("Metadata.UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_data").ColumnType(schema.ColumnTypeJSON).Description("The group's client data.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_resource_names").ColumnType(schema.ColumnTypeJSON).Description("A list of contact person resource names that are members of the contact group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_name").ColumnType(schema.ColumnTypeString).Description("The resource name for the contact group, assigned by the server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The contact group name set by the group owner or a system provided name for system groups.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_type").ColumnType(schema.ColumnTypeString).Description("The contact group type.").Build(),
	}
}

func (x *TableGoogleworkspacePeopleContactGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
