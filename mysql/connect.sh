#!/bin/bash
# Connect to the mysql container from the host.
# Hat tip to https://stackoverflow.com/a/34489879/2382333
mysql -h localhost -P 3306 --protocol=tcp -u root -p
