package tables

import (
	"context"

	"github.com/selefra/selefra-provider-googleworkspace/googleworkspace_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"google.golang.org/api/calendar/v3"
)

type TableGoogleworkspaceCalendarMyEventGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspaceCalendarMyEventGenerator{}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetTableName() string {
	return "googleworkspace_calendar_my_event"
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			service, err := googleworkspace_client.CalendarService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxResult := int64(2500)

			var query string

			resp := service.Events.List("primary").ShowDeleted(false).SingleEvents(true).Q(query).MaxResults(maxResult)

			if err := resp.Pages(ctx, func(page *calendar.Events) error {
				for _, event := range page.Items {
					resultChannel <- calendarEvent{*event, page.Summary}

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

type calendarEvent = struct {
	calendar.Event
	CalendarId string
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the event end time.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("color_id").ColumnType(schema.ColumnTypeString).Description("The color of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A short user-defined description of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_copy").ColumnType(schema.ColumnTypeBool).Description("Indicates whether event propagation is disabled, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Opaque identifier of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeJSON).Description("Specifies the creator details of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_properties").ColumnType(schema.ColumnTypeJSON).Description("A list of extended properties of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Description("ETag of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments").ColumnType(schema.ColumnTypeJSON).Description("A list of file attachments for the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("Last modification time of the event.").
			Extractor(column_value_extractor.StructSelector("Updated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("calendar_id").ColumnType(schema.ColumnTypeString).Description("Identifier of the calendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hangout_link").ColumnType(schema.ColumnTypeString).Description("An absolute link to the Google Hangout associated with this event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_type").ColumnType(schema.ColumnTypeString).Description("Specifies the type of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_link").ColumnType(schema.ColumnTypeString).Description("An absolute link to this event in the Google Calendar Web UI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Creation time of the event.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guests_can_modify").ColumnType(schema.ColumnTypeBool).Description("Indicates whether attendees other than the organizer can modify the event, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ical_uid").ColumnType(schema.ColumnTypeString).Description("Specifies the event unique identifier as defined in RFC5545. It is used to uniquely identify events accross calendaring systems and must be supplied when importing events via the import method.").
			Extractor(column_value_extractor.StructSelector("ICalUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("Specifies the title of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attendees").ColumnType(schema.ColumnTypeJSON).Description("A list of attendees of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conference_data").ColumnType(schema.ColumnTypeJSON).Description("The conference-related information, such as details of a Google Meet conference.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locked").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this is a locked event copy where no changes can be made to the main event fields \"summary\", \"description\", \"location\", \"start\", \"end\" or \"recurrence\".").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeString).Description("Visibility of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organizer").ColumnType(schema.ColumnTypeJSON).Description("Specifies the organizer details of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Description("Geographic location of the event as free-form text.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attendees_omitted").ColumnType(schema.ColumnTypeBool).Description("Indicates whether attendees may have been omitted from the event's representation, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time_unspecified").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the end time is actually unspecified, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guests_can_invite_others").ColumnType(schema.ColumnTypeBool).Description("Indicates whether attendees other than the organizer can invite others to the event, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).Description("The time zone of the calendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurrence").ColumnType(schema.ColumnTypeJSON).Description("A list of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("day").ColumnType(schema.ColumnTypeString).Description("Specifies the day of a week.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transparency").ColumnType(schema.ColumnTypeString).Description("Indicates whether the event blocks time on the calendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reminders").ColumnType(schema.ColumnTypeJSON).Description("Information about the event's reminders for the authenticated user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurring_event_id").ColumnType(schema.ColumnTypeString).Description("For an instance of a recurring event, this is the id of the recurring event to which this instance belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guests_can_see_other_guests").ColumnType(schema.ColumnTypeBool).Description("Indicates whether attendees other than the organizer can modify the event, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sequence").ColumnType(schema.ColumnTypeInt).Description("Sequence number as per iCalendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_start_time").ColumnType(schema.ColumnTypeJSON).Description("For an instance of a recurring event, this is the time at which this event would start according to the recurrence data in the recurring event identified by recurringEventId.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeJSON).Description("Source from which the event was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the event start time.").Build(),
	}
}

func (x *TableGoogleworkspaceCalendarMyEventGenerator) GetSubTables() []*schema.Table {
	return nil
}
