package provider

import (
	"github.com/selefra/selefra-provider-googleworkspace/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspacePeopleContactGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspaceDriveMyFileGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspacePeopleDirectoryPeopleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspaceGmailMyMessageGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspaceGmailMySettingsGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspacePeopleContactGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspaceCalendarMyEventGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspaceDriveGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGoogleworkspaceGmailMyDraftGenerator{}),
	}
}
