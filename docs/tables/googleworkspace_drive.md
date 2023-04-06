# Table: googleworkspace_drive

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| background_image_link | string | X | √ | A short-lived link to this shared drive's background image. | 
| copy_requires_writer_permission | bool | X | √ | Indicates whether the options to copy, print, or download files inside this shared drive, should be disabled for readers and commenters, or not. | 
| domain_users_only | bool | X | √ | Indicates whether access to this shared drive and items inside this shared drive is restricted to users of the domain to which this shared drive belongs. | 
| id | string | X | √ | The ID of this shared drive which is also the ID of the top level folder of this shared drive. | 
| name | string | X | √ | The name of this shared drive. | 
| created_time | timestamp | X | √ | The time at which the shared drive was created. | 
| hidden | bool | X | √ | Indicates whether the shared drive is hidden from default view, or not. | 
| admin_managed_restrictions | bool | X | √ | Indicates whether administrative privileges on this shared drive are required to modify restrictions, or not. | 
| theme_id | string | X | √ | The ID of the theme from which the background image and color will be set. | 
| background_image_file | json | X | √ | An image file and cropping parameters from which a background image for this shared drive is set. | 
| color_rgb | string | X | √ | The color of this shared drive as an RGB hex string. | 
| drive_members_only | bool | X | √ | Indicates whether access to items inside this shared drive is restricted to its members, or not. | 
| capabilities | json | X | √ | Describes the capabilities the current user has on this shared drive. | 


