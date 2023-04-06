# Table: googleworkspace_people_contact_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| formatted_name | string | X | √ | The name translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale for system groups names. | 
| member_count | int | X | √ | The total number of contacts in the group irrespective of max members in specified in the request. | 
| deleted | bool | X | √ | Indicates whether the contact group resource has been deleted, or not. | 
| updated_time | timestamp | X | √ | The time the group was last updated. | 
| client_data | json | X | √ | The group's client data. | 
| member_resource_names | json | X | √ | A list of contact person resource names that are members of the contact group. | 
| resource_name | string | X | √ | The resource name for the contact group, assigned by the server. | 
| name | string | X | √ | The contact group name set by the group owner or a system provided name for system groups. | 
| group_type | string | X | √ | The contact group type. | 


