# Table: googleworkspace_calendar_my_event

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ | Status of the event. | 
| end_time | timestamp | X | √ | Specifies the event end time. | 
| color_id | string | X | √ | The color of the event. | 
| description | string | X | √ | A short user-defined description of the event. | 
| private_copy | bool | X | √ | Indicates whether event propagation is disabled, or not. | 
| id | string | X | √ | Opaque identifier of the event. | 
| creator | json | X | √ | Specifies the creator details of the event. | 
| extended_properties | json | X | √ | A list of extended properties of the event. | 
| etag | string | X | √ | ETag of the resource. | 
| attachments | json | X | √ | A list of file attachments for the event. | 
| updated_at | timestamp | X | √ | Last modification time of the event. | 
| calendar_id | string | X | √ | Identifier of the calendar. | 
| hangout_link | string | X | √ | An absolute link to the Google Hangout associated with this event. | 
| event_type | string | X | √ | Specifies the type of the event. | 
| html_link | string | X | √ | An absolute link to this event in the Google Calendar Web UI. | 
| created_at | timestamp | X | √ | Creation time of the event. | 
| guests_can_modify | bool | X | √ | Indicates whether attendees other than the organizer can modify the event, or not. | 
| ical_uid | string | X | √ | Specifies the event unique identifier as defined in RFC5545. It is used to uniquely identify events accross calendaring systems and must be supplied when importing events via the import method. | 
| summary | string | X | √ | Specifies the title of the event. | 
| attendees | json | X | √ | A list of attendees of the event. | 
| conference_data | json | X | √ | The conference-related information, such as details of a Google Meet conference. | 
| locked | bool | X | √ | Indicates whether this is a locked event copy where no changes can be made to the main event fields "summary", "description", "location", "start", "end" or "recurrence". | 
| visibility | string | X | √ | Visibility of the event. | 
| organizer | json | X | √ | Specifies the organizer details of the event. | 
| location | string | X | √ | Geographic location of the event as free-form text. | 
| attendees_omitted | bool | X | √ | Indicates whether attendees may have been omitted from the event's representation, or not. | 
| end_time_unspecified | bool | X | √ | Indicates whether the end time is actually unspecified, or not. | 
| guests_can_invite_others | bool | X | √ | Indicates whether attendees other than the organizer can invite others to the event, or not. | 
| timezone | string | X | √ | The time zone of the calendar. | 
| recurrence | json | X | √ | A list of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. | 
| day | string | X | √ | Specifies the day of a week. | 
| transparency | string | X | √ | Indicates whether the event blocks time on the calendar. | 
| reminders | json | X | √ | Information about the event's reminders for the authenticated user. | 
| recurring_event_id | string | X | √ | For an instance of a recurring event, this is the id of the recurring event to which this instance belongs. | 
| guests_can_see_other_guests | bool | X | √ | Indicates whether attendees other than the organizer can modify the event, or not. | 
| sequence | int | X | √ | Sequence number as per iCalendar. | 
| original_start_time | json | X | √ | For an instance of a recurring event, this is the time at which this event would start according to the recurrence data in the recurring event identified by recurringEventId. | 
| source | json | X | √ | Source from which the event was created. | 
| start_time | timestamp | X | √ | Specifies the event start time. | 


