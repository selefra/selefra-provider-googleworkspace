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

type TableGoogleworkspaceDriveMyFileGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGoogleworkspaceDriveMyFileGenerator{}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetTableName() string {
	return "googleworkspace_drive_my_file"
}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetDataSource() *schema.DataSource {
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

			maxResult := int64(1000)

			resp := service.Files.List().Fields(requiredFields...).Q(query).PageSize(maxResult)
			if err := resp.Pages(ctx, func(page *drive.FileList) error {
				for _, file := range page.Files {
					parsedTime, _ := time.Parse(time.RFC3339, file.CreatedTime)
					file.CreatedTime = parsedTime.Format(time.RFC3339)
					resultChannel <- file

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

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("trashed_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time that the item was trashed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_restrictions").ColumnType(schema.ColumnTypeJSON).Description("Restrictions for accessing the content of the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared_with_me_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time at which the file was shared with the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capabilities").ColumnType(schema.ColumnTypeJSON).Description("Describes capabilities the current user has on this file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).Description("A collection of arbitrary key-value pairs which are visible to all apps.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_app_authorized").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the file was created or opened by the requesting app, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mime_type").ColumnType(schema.ColumnTypeString).Description("The MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploaded content if no value is provided.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("viewed_by_me").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the the file has been viewed by this user, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("viewed_by_me_time").ColumnType(schema.ColumnTypeTimestamp).Description("The last time the file was viewed by the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_hints").ColumnType(schema.ColumnTypeJSON).Description("Additional information about the content of the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission_ids").ColumnType(schema.ColumnTypeJSON).Description("List of permission IDs for users with access to this file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("starred").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user has starred the file, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("writers_can_share").ColumnType(schema.ColumnTypeBool).Description("Indicates whether users with only writer permission can modify the file's permissions, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("video_media_metadata").ColumnType(schema.ColumnTypeJSON).Description("Additional metadata about video media.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quota_bytes_used").ColumnType(schema.ColumnTypeInt).Description("The number of storage quota bytes used by the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_augmented_permissions").ColumnType(schema.ColumnTypeBool).Description("Indicates whether there are permissions directly on this file, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_by_me").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the file has been modified by this user, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modifying_user").ColumnType(schema.ColumnTypeJSON).Description("The last user to modify the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Specifies the name of the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_extension").ColumnType(schema.ColumnTypeString).Description("The final component of fullFileExtension.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_file_name").ColumnType(schema.ColumnTypeString).Description("The original filename of the uploaded content if available, or else the original value of the name field.").
			Extractor(column_value_extractor.StructSelector("OriginalFilename")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_key").ColumnType(schema.ColumnTypeString).Description("A key needed to access the item via a shared link.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thumbnail_version").ColumnType(schema.ColumnTypeInt).Description("The thumbnail version for use in thumbnail cache invalidation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trashing_user").ColumnType(schema.ColumnTypeJSON).Description("Specifies the user who trashed the file explicitly.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owned_by_me").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user owns the file, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_media_metadata").ColumnType(schema.ColumnTypeJSON).Description("Additional metadata about image media, if available.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owners").ColumnType(schema.ColumnTypeJSON).Description("The owner of this file. Only certain legacy files may have more than one owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parents").ColumnType(schema.ColumnTypeJSON).Description("The IDs of the parent folders which contain the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_time").ColumnType(schema.ColumnTypeTimestamp).Description("The last time the file was modified by anyone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("icon_link").ColumnType(schema.ColumnTypeString).Description("A static, unauthenticated link to the file's icon.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thumbnail_link").ColumnType(schema.ColumnTypeString).Description("A short-lived link to the file's thumbnail, if available.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("export_links").ColumnType(schema.ColumnTypeJSON).Description("Links for exporting Docs Editors files to specific formats.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("explicitly_trashed").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the file has been explicitly trashed, as opposed to recursively trashed from a parent folder.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeInt).Description("A monotonically increasing version number for the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_content_link").ColumnType(schema.ColumnTypeString).Description("A link for downloading the content of the file in a browser.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_view_link").ColumnType(schema.ColumnTypeString).Description("A link for opening the file in a relevant Google editor or viewer in a browser.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("link_share_metadata").ColumnType(schema.ColumnTypeJSON).Description("Contains details about the link URLs that clients are using to refer to this item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("md5_checksum").ColumnType(schema.ColumnTypeString).Description("The MD5 checksum for the content of the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("full_file_extension").ColumnType(schema.ColumnTypeString).Description("The full file extension extracted from the name field.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("head_revision_id").ColumnType(schema.ColumnTypeString).Description("The ID of the file's head revision.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time at which the file was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Description("The full list of permissions for the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sharing_user").ColumnType(schema.ColumnTypeJSON).Description("The user who shared the file with the requesting user, if applicable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shortcut_details").ColumnType(schema.ColumnTypeJSON).Description("Shortcut file details. Only populated for shortcut files, which have the mimeType field set to application/vnd.google-apps.shortcut.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_by_me_time").ColumnType(schema.ColumnTypeTimestamp).Description("The last time the file was modified by the use.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the file has been shared, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("folder_color_rgb").ColumnType(schema.ColumnTypeString).Description("The color for a folder or shortcut to a folder as an RGB hex string.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_thumbnail").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this file has a thumbnail, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("The size of the file's content in bytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trashed").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the file has been trashed, either explicitly or from a trashed parent folder, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_properties").ColumnType(schema.ColumnTypeJSON).Description("A collection of arbitrary key-value pairs which are private to the requesting app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spaces").ColumnType(schema.ColumnTypeJSON).Description("The list of spaces which contain the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("drive_id").ColumnType(schema.ColumnTypeString).Description("ID of the shared drive the file resides in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A short description of the file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("copy_requires_writer_permission").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the options to copy, print, or download this file, should be disabled for readers and commenters, or not.").Build(),
	}
}

func (x *TableGoogleworkspaceDriveMyFileGenerator) GetSubTables() []*schema.Table {
	return nil
}
