----

# Trostevask

## Prerequisites:

- create file .env with the following keys
    - **input_dir**=<*your_input_dir*>
    - **output_dir**=<*your_output_dir*>
    - **rejected_dir**=<*your_rejected_dir*>


## Example usage

### Arguments
Run the program with different arguments.

- test | The program will create files in input directory to test with.
- dispose | Will cleanup input, ouput and rejected directory before proceeding with anything.
- debug | Will print extra debug messages.

example:
```
./trostevask debug
```

It's also possible to run with all arguments at the same time, 
just add more and separate them with spaces.

### Create service file named <service-name>.service
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
sudo systemctl enable <service-name>.service
sudo systemctl start <service-name>.service
sudo systemctl status <service-name>.service
```
