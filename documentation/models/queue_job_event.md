# QueueJobEvent

Represents an event for queue job

**Properties**

| Name   | Type          | Required | Description |
| :----- | :------------ | :------- | :---------- |
| Action | queues.Action | ✅       |             |
| Time   | string        | ✅       |             |

# Action

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| created   | string | ✅       | "created"   |
| started   | string | ✅       | "started"   |
| succeeded | string | ✅       | "succeeded" |
| cancelled | string | ✅       | "cancelled" |
| failed    | string | ✅       | "failed"    |
