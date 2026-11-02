#!/usr/bin/env bash
# PR Newswire Article API – cURL example
# Usage: PILOTERR_API_KEY=your-key ./article.sh
set -euo pipefail

curl -sS --fail-with-body -G "https://api.piloterr.com/v2/prnewswire/article" \
  -H "x-api-key: ${PILOTERR_API_KEY}" \
  --data-urlencode query=https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html
