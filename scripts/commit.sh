#!/bin/bash
set -e

if [ -z "$1" ] || [ -z "$2" ]; then
    echo "Usage: ./scripts/commit.sh <date> <message>"
    echo "Date format: YYYY-MM-DD"
    echo "Example: ./scripts/commit.sh 2025-12-30 'feat: initial structure'"
    exit 1
fi

DATE=$1
MESSAGE=$2

if [[ "$OSTYPE" == "darwin"* ]]; then
    TIMESTAMP=$(date -j -f "%Y-%m-%d" "$DATE" "+%Y-%m-%dT10:00:00" 2>/dev/null || echo "${DATE}T10:00:00")
else
    TIMESTAMP=$(date -d "$DATE 10:00:00" "+%Y-%m-%dT10:00:00" 2>/dev/null || echo "${DATE}T10:00:00")
fi

export GIT_AUTHOR_DATE="$TIMESTAMP"
export GIT_COMMITTER_DATE="$TIMESTAMP"

echo "Committing with date: $TIMESTAMP"
echo "Message: $MESSAGE"

git add .
git commit -m "$MESSAGE"

unset GIT_AUTHOR_DATE
unset GIT_COMMITTER_DATE

echo "Commit created successfully"
