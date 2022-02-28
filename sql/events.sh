#!/bin/bash

set -euo pipefail

{
    sqlite3 $1 <<EOF
.mode json
select
  name,
  date,
  city,
  location,
  format,
  website as event_url,
  tabletop_url as signup_url
from Events
order by 2 desc;
EOF
} | python -m json.tool
