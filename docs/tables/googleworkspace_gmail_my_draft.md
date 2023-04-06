# Table: googleworkspace_gmail_my_draft

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| message_thread_id | string | X | √ | The ID of the thread the message belongs to. | 
| message_raw | string | X | √ | The entire email message in an RFC 2822 formatted and base64url encoded string. | 
| message_snippet | string | X | √ | A short part of the message text. | 
| message_internal_date | timestamp | X | √ | The internal message creation timestamp which determines ordering in the inbox. | 
| message_size_estimate | int | X | √ | Estimated size in bytes of the message. | 
| message_label_ids | json | X | √ | A list of IDs of labels applied to this message. | 
| message_payload | json | X | √ | The parsed email structure in the message parts. | 
| draft_id | string | X | √ | The immutable ID of the draft. | 
| message_id | string | X | √ | The immutable ID of the message. | 
| message_history_id | string | X | √ | The ID of the last history record that modified this message. | 


