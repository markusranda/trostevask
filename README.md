----

# Trostevask

## Prerequisites:

- create file .env with the following keys
    - **input_dir**=<*your_input_dir*>
    - **output_dir**=<*your_output_dir*>
    - **rejected_dir**=<*your_rejected_dir*>


## Example usage

Create service file
```
[Unit]
Description=Service for cleaning up media files

[Install]
WantedBy=multi-user.target

[Service]
Type=simple
ExecStart=/opt/troste/trostevask/trostevask
WorkingDirectory=/opt/troste/trostevask
Restart=always
RestartSec=5
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=%n
```

Setup:

```shell
sudo cp trostevask.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable echo-server.service
sudo systemctl start echo-server.service
sudo systemctl status echo-server.service
```