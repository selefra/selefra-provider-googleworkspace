package tables

import (
	"context"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"google.golang.org/api/people/v1"
)

type TableGoogleworkspacePeopleDirectoryPeopleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspacePeopleDirectoryPeopleGenerator{}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetTableName() string {
	return "googleworkspace_people_directory_people"
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.PeopleService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			personFields := "addresses,biographies,birthdays,calendarUrls,clientData,coverPhotos,emailAddresses,events,externalIds,genders,interests,locations,memberships,metadata,miscKeywords,names,nicknames,occupations,organizations,phoneNumbers,photos,relations,sipAddresses,skills,urls,userDefined"

			maxResult := int64(1000)

			resp := service.People.ListDirectoryPeople().ReadMask(personFields).Sources("DIRECTORY_SOURCE_TYPE_DOMAIN_PROFILE").PageSize(maxResult)
			if err := resp.Pages(ctx, func(page *people.ListDirectoryPeopleResponse) error {
				for _, people := range page.People {

					var conn contacts
					if people.Names != nil {
						conn.Name = *people.Names[0]
					}
					if people.Birthdays != nil {
						conn.Birthday = *people.Birthdays[0]
					}
					if people.Genders != nil {
						conn.Gender = *people.Genders[0]
					}
					if people.Biographies != nil {
						conn.Biography = *people.Biographies[0]
					}
					resultChannel <- contacts{
						conn.Name,
						conn.Birthday,
						conn.Gender,
						conn.Biography,
						*people,
					}

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

type contacts = struct {
	Name      people.Name
	Birthday  people.Birthday
	Gender    people.Gender
	Biography people.Biography
	people.Person
}

func extractPrimaryEmailAddress(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(contacts)

	emailAddresses := data.EmailAddresses

	for _, email := range emailAddresses {
		if email.Metadata != nil && email.Metadata.Primary {
			return email.Value, nil
		}
	}

	return nil, nil
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("gender").ColumnType(schema.ColumnTypeString).Description("The gender for the person.").
			Extractor(column_value_extractor.StructSelector("Gender.Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("biography").ColumnType(schema.ColumnTypeJSON).Description("The person's biography.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_data").ColumnType(schema.ColumnTypeJSON).Description("The person's client data.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cover_photos").ColumnType(schema.ColumnTypeJSON).Description("The person's cover photos.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events").ColumnType(schema.ColumnTypeJSON).Description("The person's events.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locations").ColumnType(schema.ColumnTypeJSON).Description("The person's locations.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memberships").ColumnType(schema.ColumnTypeJSON).Description("The person's group memberships.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary_email_address").ColumnType(schema.ColumnTypeString).Description("The primary email address of the user contact.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := extractPrimaryEmailAddress(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name formatted according to the locale specified by the viewer's account.").
			Extractor(column_value_extractor.StructSelector("Name.DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_addresses").ColumnType(schema.ColumnTypeJSON).Description("The person's email addresses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Metadata about the person.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organizations").ColumnType(schema.ColumnTypeJSON).Description("The person's past or current organizations.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_name").ColumnType(schema.ColumnTypeString).Description("The resource name for the contact group, assigned by the server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addresses").ColumnType(schema.ColumnTypeJSON).Description("The person's street addresses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("calendar_urls").ColumnType(schema.ColumnTypeJSON).Description("The person's calendar URLs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_ids").ColumnType(schema.ColumnTypeJSON).Description("The person's external IDs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("occupations").ColumnType(schema.ColumnTypeJSON).Description("The person's occupations.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone_numbers").ColumnType(schema.ColumnTypeJSON).Description("The person's phone numbers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photos").ColumnType(schema.ColumnTypeJSON).Description("The person's photos.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("birthday").ColumnType(schema.ColumnTypeJSON).Description("The date of the birthday.").
			Extractor(column_value_extractor.StructSelector("Birthday.Date")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interests").ColumnType(schema.ColumnTypeJSON).Description("The person's interests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nicknames").ColumnType(schema.ColumnTypeJSON).Description("The person's nicknames.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("given_name").ColumnType(schema.ColumnTypeString).Description("The given name of the user contact.").
			Extractor(column_value_extractor.StructSelector("Name.GivenName")).Build(),
	}
}

func (x *TableGoogleworkspacePeopleDirectoryPeopleGenerator) GetSubTables() []*schema.Table {
	return nil
}
