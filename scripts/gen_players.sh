#!/usr/bin/env bash

while read player_id
do
    {
        echo "---"
        echo "player: ${player_id}"
        echo "---"

    } > "content/player/${player_id}.md"
done < /dev/stdin
