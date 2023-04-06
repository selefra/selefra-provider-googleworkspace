# Table: googleworkspace_drive_my_file

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| trashed_time | timestamp | X | √ | The time that the item was trashed. | 
| content_restrictions | json | X | √ | Restrictions for accessing the content of the file. | 
| shared_with_me_time | timestamp | X | √ | The time at which the file was shared with the user. | 
| capabilities | json | X | √ | Describes capabilities the current user has on this file. | 
| properties | json | X | √ | A collection of arbitrary key-value pairs which are visible to all apps. | 
| is_app_authorized | bool | X | √ | Indicates whether the file was created or opened by the requesting app, or not. | 
| mime_type | string | X | √ | The MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploaded content if no value is provided. | 
| viewed_by_me | bool | X | √ | Indicates whether the the file has been viewed by this user, or not. | 
| viewed_by_me_time | timestamp | X | √ | The last time the file was viewed by the user. | 
| content_hints | json | X | √ | Additional information about the content of the file. | 
| permission_ids | json | X | √ | List of permission IDs for users with access to this file. | 
| id | string | X | √ | The ID of the file. | 
| starred | bool | X | √ | Indicates whether the user has starred the file, or not. | 
| writers_can_share | bool | X | √ | Indicates whether users with only writer permission can modify the file's permissions, or not. | 
| video_media_metadata | json | X | √ | Additional metadata about video media. | 
| quota_bytes_used | int | X | √ | The number of storage quota bytes used by the file. | 
| has_augmented_permissions | bool | X | √ | Indicates whether there are permissions directly on this file, or not. | 
| modified_by_me | bool | X | √ | Indicates whether the file has been modified by this user, or not. | 
| last_modifying_user | json | X | √ | The last user to modify the file. | 
| name | string | X | √ | Specifies the name of the file. | 
| file_extension | string | X | √ | The final component of fullFileExtension. | 
| original_file_name | string | X | √ | The original filename of the uploaded content if available, or else the original value of the name field. | 
| resource_key | string | X | √ | A key needed to access the item via a shared link. | 
| thumbnail_version | int | X | √ | The thumbnail version for use in thumbnail cache invalidation. | 
| trashing_user | json | X | √ | Specifies the user who trashed the file explicitly. | 
| owned_by_me | bool | X | √ | Indicates whether the user owns the file, or not. | 
| image_media_metadata | json | X | √ | Additional metadata about image media, if available. | 
| owners | json | X | √ | The owner of this file. Only certain legacy files may have more than one owner. | 
| parents | json | X | √ | The IDs of the parent folders which contain the file. | 
| modified_time | timestamp | X | √ | The last time the file was modified by anyone. | 
| icon_link | string | X | √ | A static, unauthenticated link to the file's icon. | 
| thumbnail_link | string | X | √ | A short-lived link to the file's thumbnail, if available. | 
| export_links | json | X | √ | Links for exporting Docs Editors files to specific formats. | 
| explicitly_trashed | bool | X | √ | Indicates whether the file has been explicitly trashed, as opposed to recursively trashed from a parent folder. | 
| version | int | X | √ | A monotonically increasing version number for the file. | 
| web_content_link | string | X | √ | A link for downloading the content of the file in a browser. | 
| web_view_link | string | X | √ | A link for opening the file in a relevant Google editor or viewer in a browser. | 
| link_share_metadata | json | X | √ | Contains details about the link URLs that clients are using to refer to this item. | 
| md5_checksum | string | X | √ | The MD5 checksum for the content of the file. | 
| full_file_extension | string | X | √ | The full file extension extracted from the name field. | 
| head_revision_id | string | X | √ | The ID of the file's head revision. | 
| created_time | timestamp | X | √ | The time at which the file was created. | 
| permissions | json | X | √ | The full list of permissions for the file. | 
| sharing_user | json | X | √ | The user who shared the file with the requesting user, if applicable. | 
| shortcut_details | json | X | √ | Shortcut file details. Only populated for shortcut files, which have the mimeType field set to application/vnd.google-apps.shortcut. | 
| modified_by_me_time | timestamp | X | √ | The last time the file was modified by the use. | 
| shared | bool | X | √ | Indicates whether the file has been shared, or not. | 
| folder_color_rgb | string | X | √ | The color for a folder or shortcut to a folder as an RGB hex string. | 
| has_thumbnail | bool | X | √ | Indicates whether this file has a thumbnail, or not. | 
| size | int | X | √ | The size of the file's content in bytes. | 
| trashed | bool | X | √ | Indicates whether the file has been trashed, either explicitly or from a trashed parent folder, or not. | 
| app_properties | json | X | √ | A collection of arbitrary key-value pairs which are private to the requesting app. | 
| spaces | json | X | √ | The list of spaces which contain the file. | 
| drive_id | string | X | √ | ID of the shared drive the file resides in. | 
| description | string | X | √ | A short description of the file. | 
| copy_requires_writer_permission | bool | X | √ | Indicates whether the options to copy, print, or download this file, should be disabled for readers and commenters, or not. | 


