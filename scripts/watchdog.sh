#!/usr/bin/env bash

PROJECT_ROOT="/home/alexanderlupatsiy/personal-website"

PORT_A=8080
PORT_B=8081

LOG="$PROJECT_ROOT/logs/website_watchdog.log"
LOCK="$PROJECT_ROOT/.watchdog.lock"

mkdir -p "$PROJECT_ROOT/logs"

# prevent parallel runs
exec 9>"$LOCK" || exit 1
flock -n 9 || exit 0

# check ports
if ss -ltn | grep -Eq ":(8080|8081)\b"; then
    echo "$(date): OK - app listening on 8080 or 8081" >> "$LOG"
    exit 0
fi

echo "$(date): DOWN - no app port listening, running start.sh" >> "$LOG"

cd "$PROJECT_ROOT" || exit 1
/bin/bash "$PROJECT_ROOT/scripts/prod-up.sh" >> "$LOG" 2>&1
