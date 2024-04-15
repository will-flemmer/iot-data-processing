# IoT Data Processing
A project which simulates IoT data being sent to a server. This data is then processed calculated metrics are displayed via a simple UI.

### IoT Simulator
Sends data to the central server every 5 seconds. The data looks like this:

#### Temperature Datafile
```
{
  sensor_id: 1,
  data: [
    { datetime: 2024-04-14T10:38:30.622Z, temp_c: 20.0 },
    { datetime: 2024-04-14T10:39:30.622Z, temp_c: 20.5 },
    { datetime: 2024-04-14T10:40:30.622Z, temp_c: 20.5 },
    { datetime: 2024-04-14T10:41:30.622Z, temp_c: 20.5 },
    { datetime: 2024-04-14T10:42:30.622Z, temp_c: 20.5 },
  ]
}
```

### Central Server
Once data is received, the server saves the record to the DB and starts several jobs to analyse the file. The jobs are:
- AlertingJob (checks if the temp is over 25 degrees C, if so it send an alert and increments alert counter)
- DataAnalysisJob (Updates top level metrics and stores them in the DB. E.g. average_temp, max_temp, min_temp)

### User Interface
**Temperature Metrics**: max_temp, min_temp, avg_temp
**System Metrics**: num_alerts
