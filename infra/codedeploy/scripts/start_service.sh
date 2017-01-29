#!/bin/sh
mv /home/ec2-user/supervisord.conf /etc/supervisord.conf
mv /home/ec2-user/app /tmp/app
chmod +x /tmp/app

/usr/local/bin/supervisord -c /etc/supervisord.conf
/usr/local/bin/supervisorctl start app
