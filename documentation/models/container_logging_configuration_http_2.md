# ContainerLoggingConfigurationHttp2

Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.

**Properties**

| Name        | Type                                | Required | Description                                                   |
| :---------- | :---------------------------------- | :------- | :------------------------------------------------------------ |
| Host        | string                              | ✅       | The hostname or IP address of the HTTP logging endpoint       |
| Port        | int64                               | ✅       | The port number of the HTTP logging endpoint (1-65535)        |
| Format      | shared.ContainerLoggingHttpFormat   | ✅       | The format in which logs will be delivered                    |
| Compression | any                                 | ✅       |                                                               |
| User        | string                              | ❌       | Optional username for HTTP authentication                     |
| Password    | string                              | ❌       | Optional password for HTTP authentication                     |
| Path        | string                              | ❌       | Optional URL path for the HTTP endpoint                       |
| Headers     | []shared.ContainerLoggingHttpHeader | ❌       | Optional HTTP headers to include in log transmission requests |
