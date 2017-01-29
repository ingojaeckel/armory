#!/bin/bash
/usr/local/bin/supervisorctl shutdown
rm -rf /tmp/app /home/ec2-user/appspec.yml /home/ec2-home/app /home/ec2-home/*.go /home/ec2-home/supervisord.conf
