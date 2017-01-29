#!/bin/sh
# This is not executed when launching the app.
# Instead the frontend will trigger this after it copied the latest app file.
echo "Starting application.."
chmod +x /tmp/app
nohup /tmp/app >/tmp/application.log 2>&1 &
