# Table: googleworkspace_gmail_my_message

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The immutable ID of the message. | 
| raw | string | X | √ | The entire email message in an RFC 2822 formatted and base64url encoded string. | 
| snippet | string | X | √ | A short part of the message text. | 
| label_ids | json | X | √ | A list of IDs of labels applied to this message. | 
| payload | json | X | √ | The parsed email structure in the message parts. | 
| thread_id | string | X | √ | The ID of the thread the message belongs to. | 
| history_id | string | X | √ | The ID of the last history record that modified this message. | 
| sender_email | string | X | √ | Specifies the email address of the sender. | 
| internal_date | timestamp | X | √ | The internal message creation timestamp which determines ordering in the inbox. | 
| size_estimate | int | X | √ | Estimated size in bytes of the message. | 


