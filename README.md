<p align="center">
  <a href="https://www.piloterr.com"><img src="images/banner.png" alt="PR Newswire Scraper API by Piloterr – scrape PR Newswire data as JSON" width="100%"></a>
</p>

# PR Newswire Scraper API

[![Discord](https://img.shields.io/badge/Discord-Join%20the%20community-5865F2?logo=discord&logoColor=white)](https://discord.gg/CFj3ph8aHU) [![Get API key](https://img.shields.io/badge/Piloterr-Get%20a%20free%20API%20key-111827)](https://app.piloterr.com) [![API docs](https://img.shields.io/badge/Docs-OpenAPI-0EA5E9)](https://app.piloterr.com/openapi) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

The **Piloterr PR Newswire Scraper API** lets you extract public data from PR Newswire as **clean, structured JSON** — no proxies, headless browsers or HTML parsing to maintain. Endpoints available in production:

- **PR Newswire Article API** — Read one PR Newswire press release from its URL.
- **PR Newswire Search API** — Search PR Newswire press releases by keyword or by a news-search URL.

Use it for price monitoring, assortment tracking, competitor analysis, product research or catalog enrichment.

| Endpoint | Path | Credits | What you get |
|---|---|---|---|
| [PR Newswire Article API](#pr-newswire-article-api) | `/v2/prnewswire/article` | 1 | Read one PR Newswire press release from its URL. |
| [PR Newswire Search API](#pr-newswire-search-api) | `/v2/prnewswire/search` | 1 | Search PR Newswire press releases by keyword or by a news-search URL. |

## Table of contents

- [How it works](#how-it-works)
- [PR Newswire Article API](#pr-newswire-article-api)
- [PR Newswire Search API](#pr-newswire-search-api)
- [Code examples in other languages](#code-examples-in-other-languages)
- [Practical use cases](#practical-use-cases)
- [Why Piloterr?](#why-piloterr)
- [FAQ](#faq)
- [Related scrapers](#related-scrapers)
- [Support](#support)

## How it works

1. **Create a free Piloterr account** and copy your API key from the [dashboard](https://app.piloterr.com).
2. **Send a `GET` (or `POST`) request** to `https://api.piloterr.com/v2/...` with the header `x-api-key: YOUR_API_KEY`.
3. **Receive structured JSON** — Piloterr handles rendering, proxies, retries and parsing.

```bash
export PILOTERR_API_KEY="your-api-key"
pip install requests        # Python
# TypeScript: Node.js 18+ (native fetch), run with `npx tsx file.ts`
```

> Every endpoint accepts both `GET` (query string) and `POST` (JSON body with the same fields).

## PR Newswire Article API

Read one PR Newswire press release from its URL.

`GET | POST https://api.piloterr.com/v2/prnewswire/article` · **1 credit per request** · [Full reference](https://app.piloterr.com/openapi)

### Request parameters

| Parameter | Type | Required | Description |
|---|---|---|---|
| `query` | string | ✅ | Required. PR Newswire release URL, for example https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html. A keyword is rejected. |

### Python example

```python
import json
import os
import requests

response = requests.get(
    "https://api.piloterr.com/v2/prnewswire/article",
    headers={"x-api-key": os.environ["PILOTERR_API_KEY"]},
    params={
        "query": "https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html",
    },
    timeout=60,
)
response.raise_for_status()
data = response.json()

print(json.dumps(data, indent=2, ensure_ascii=False)[:2000])
```

### TypeScript example

```typescript
type PrNewswireArticle = {
  id?: string;
  url?: string;
  body?: string;
  image?: unknown;
  title?: string;
  company?: string;
  dateline?: string;
  keywords?: string;
  company_url?: string;
  description?: string;
  published_at?: string;
  updated_at_iso?: string;
  published_at_iso?: string;
};

const params = new URLSearchParams({
  "query": "https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html",
});

const response = await fetch(`https://api.piloterr.com/v2/prnewswire/article?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY! },
});
if (!response.ok) throw new Error(`Piloterr error ${response.status}: ${await response.text()}`);

const data = (await response.json()) as PrNewswireArticle;

console.log(JSON.stringify(data, null, 2).slice(0, 2000));
```

<details>
<summary><b>POST variant</b></summary>

```http
POST https://api.piloterr.com/v2/prnewswire/article
x-api-key: YOUR_API_KEY
Content-Type: application/json

{"query": "https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html"}
```
</details>

### Output example

Example response from the API reference, abridged — real values depend on the page you scrape.

```json
{
  "id": "302924942",
  "url": "https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html",
  "body": "NEW YORK , Oct. 30, 2026 /PRNewswire/ -- Pumpkin Pet Insurance is rewarding 11 veterinary clinics with $1,000 to support angel funds that…",
  "image": null,
  "title": "Pumpkin awarded $11,000 in donations to veterinary clinics for cost-of-care education",
  "company": "Pumpkin Pet Insurance",
  "dateline": "NEW YORK",
  "keywords": "Pumpkin Pet Insurance",
  "company_url": "https://www.prnewswire.com/news/pumpkin-pet-insurance/",
  "description": "/PRNewswire/ -- Pumpkin Pet Insurance is rewarding 11 veterinary clinics with $1,000 to support angel funds that help pets access care th…",
  "published_at": "Oct 30, 2026, 09:06 ET",
  "updated_at_iso": "2026-10-30T13:07:53Z",
  "published_at_iso": "2026-10-30T13:06:00Z"
}
```

### Response fields

| Field | Type | Description |
|---|---|---|
| `id` | string |  |
| `url` | string |  |
| `body` | string |  |
| `image` | object |  |
| `title` | string |  |
| `company` | string |  |
| `dateline` | string |  |
| `keywords` | string |  |
| `company_url` | string |  |
| `description` | string |  |
| `published_at` | string |  |
| `updated_at_iso` | string |  |
| `published_at_iso` | string |  |

## PR Newswire Search API

Search PR Newswire press releases by keyword or by a news-search URL.

`GET | POST https://api.piloterr.com/v2/prnewswire/search` · **1 credit per request** · [Full reference](https://app.piloterr.com/openapi)

### Request parameters

| Parameter | Type | Required | Description |
|---|---|---|---|
| `query` | string | ✅ | Required. Keyword, for example apple, or a PR Newswire news-search URL such as https://www.prnewswire.com/search/news/?keyword=apple. A release URL is rejected. |
| `page` | number | — | Optional. Page number starting at 1. Example: 2. Omitted means page 1. A page already present in a search URL is kept when this is omitted or 1. |

### Python example

```python
import json
import os
import requests

response = requests.get(
    "https://api.piloterr.com/v2/prnewswire/search",
    headers={"x-api-key": os.environ["PILOTERR_API_KEY"]},
    params={
        "query": "YOUR_VALUE",
    },
    timeout=60,
)
response.raise_for_status()
data = response.json()

print(json.dumps(data, indent=2, ensure_ascii=False)[:2000])
```

### TypeScript example

```typescript
type PrNewswireSearch = {
  results?: {
    id?: string;
    url?: string;
    image?: unknown;
    title?: string;
    company?: string;
    summary?: string;
    company_url?: string;
    published_at?: string;
  }[];
  pagination?: {
    next?: string;
    page?: number;
    per_page?: number;
    total_pages?: number;
    total_results?: number;
  };
};

const params = new URLSearchParams({
  "query": "YOUR_VALUE",
});

const response = await fetch(`https://api.piloterr.com/v2/prnewswire/search?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY! },
});
if (!response.ok) throw new Error(`Piloterr error ${response.status}: ${await response.text()}`);

const data = (await response.json()) as PrNewswireSearch;

console.log(JSON.stringify(data, null, 2).slice(0, 2000));
```

<details>
<summary><b>POST variant</b></summary>

```http
POST https://api.piloterr.com/v2/prnewswire/search
x-api-key: YOUR_API_KEY
Content-Type: application/json

{"query": "YOUR_VALUE"}
```
</details>

### Output example

Example response from the API reference, abridged — real values depend on the page you scrape.

```json
{
  "results": [
    {
      "id": "302924942",
      "url": "https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html",
      "image": null,
      "title": "Pumpkin awarded $11,000 in donations to veterinary clinics for cost-of-care education",
      "company": "Pumpkin Pet Insurance",
      "summary": "Small Animal Hospital of PA (PA) – donating to Peaceable KingdomBrilliant Veterinary Care (NY) – donating to Frankie's Feline Fund and Bi…",
      "company_url": "https://www.prnewswire.com/news/pumpkin-pet-insurance/",
      "published_at": "Oct 30, 2026, 09:06 ET"
    }
  ],
  "pagination": {
    "next": "https://www.prnewswire.com/search/news/?keyword=apple&pagesize=25&page=2",
    "page": 1,
    "per_page": 25,
    "total_pages": 60,
    "total_results": 1492
  }
}
```

### Response fields

| Field | Type | Description |
|---|---|---|
| `results[]` | object[] |  |
| `results[].id` | string |  |
| `results[].url` | string |  |
| `results[].image` | object |  |
| `results[].title` | string |  |
| `results[].company` | string |  |
| `results[].summary` | string |  |
| `results[].company_url` | string |  |
| `results[].published_at` | string |  |
| `pagination` | object |  |
| `pagination.next` | string |  |
| `pagination.page` | integer |  |
| `pagination.per_page` | integer |  |
| `pagination.total_pages` | integer |  |
| `pagination.total_results` | integer |  |

## Code examples in other languages

Ready-to-run examples live in the [`examples/`](examples) folder. Each one reads your key from the `PILOTERR_API_KEY` environment variable.

| Language | Article | Search |
|---|---|---|
| cURL (Bash) | [`article`](examples/curl/article.sh) | [`search`](examples/curl/search.sh) |
| PHP | [`article`](examples/php/article.php) | [`search`](examples/php/search.php) |
| C# | [`article`](examples/csharp/Article.cs) | [`search`](examples/csharp/Search.cs) |
| Go | [`article`](examples/go/article.go) | [`search`](examples/go/search.go) |
| Java | [`article`](examples/java/ArticleExample.java) | [`search`](examples/java/SearchExample.java) |
| Ruby | [`article`](examples/ruby/article.rb) | [`search`](examples/ruby/search.rb) |
| Node.js | [`article`](examples/nodejs/article.mjs) | [`search`](examples/nodejs/search.mjs) |

## Practical use cases

- **Price monitoring** — Track prices, promotions and availability over time to react to competitors.
- **Assortment & catalog tracking** — Follow new products, delisted items and category changes.
- **Competitor analysis** — Benchmark ranges, brands and positioning across retailers.
- **Product research** — Spot best-sellers and trends from search results.
- **Catalog enrichment** — Pull titles, images and attributes into your PIM or marketplace listings.

## Why Piloterr?

- **No infrastructure** — no proxies, headless browsers or CAPTCHA solving to run yourself.
- **Structured JSON** — stable, documented fields instead of brittle HTML selectors.
- **Pay per request** — transparent credit cost per endpoint, failed requests are retried for you.
- **One API key, hundreds of sources** — PR Newswire plus many other e-commerce, search and company data endpoints in the [Piloterr library](https://www.piloterr.com/library).
- **Any language** — plain HTTPS + JSON, with examples for Python, TypeScript, cURL, PHP, C#, Go, Java, Ruby and Node.js.

## FAQ

### Is it legal to scrape PR Newswire?

Collecting publicly available data is generally allowed in many jurisdictions, but rules depend on your country, the data you collect and how you use it. Avoid personal data unless you have a lawful basis, respect the website's terms where they apply, and get legal advice for your specific use case. This is not legal advice.

### Do I need proxies or a headless browser?

No. Piloterr handles rendering, proxy rotation and retries on its side. You only send an HTTPS request with your `x-api-key` header and receive JSON.

### How much does a request cost?

Credits per call: `/v2/prnewswire/article` **1**, `/v2/prnewswire/search` **1**. See the [Piloterr website](https://www.piloterr.com) for plans.

### What data does the PR Newswire API return?

Each endpoint returns structured JSON: **PR Newswire Article** — `id`, `url`, `body`, `image`, `title`, `company`; **PR Newswire Search** — `results[]`, `results[].id`, `results[].url`, `results[].image`, `results[].title`, `results[].company`.

### How do I paginate PR Newswire results?

Use the `page` parameter and keep requesting while the response indicates more results.

### Which languages are supported?

Any language that can make an HTTP request. This README shows Python and TypeScript; the [`examples/`](examples) folder contains cURL, PHP, C#, Go, Java, Ruby and Node.js.

## Related scrapers

- [About You Scraper API](https://github.com/piloterr/aboutyou-scraper)
- [Action Scraper API](https://github.com/piloterr/action-scraper)
- [Aldi Scraper API](https://github.com/piloterr/aldi-scraper)
- [Alibaba Scraper API](https://github.com/piloterr/alibaba-scraper)
- [AliExpress Scraper API](https://github.com/piloterr/aliexpress-scraper)
- [Allegro Scraper API](https://github.com/piloterr/allegro-scraper)
- [All Piloterr scrapers on GitHub](https://github.com/piloterr)

## Support

- 💬 Join the [Piloterr Discord](https://discord.gg/CFj3ph8aHU) to ask questions and share what you build.
- 📚 Browse the [API reference (OpenAPI)](https://app.piloterr.com/openapi) and the [Piloterr library](https://www.piloterr.com/library).
- 🔑 [Get your API key](https://app.piloterr.com).

---

<p align="center">
  <a href="https://www.piloterr.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="images/logo-dark.svg">
      <img src="images/logo.svg" alt="Piloterr" height="32">
    </picture>
  </a>
  <br>
  <sub>Built by <a href="https://www.piloterr.com">Piloterr</a> — web scraping APIs for e-commerce, search, real estate and more, one API key.</sub>
</p>

<sub>PR Newswire is a trademark of its respective owner. This repository is not affiliated with or endorsed by PR Newswire. Examples are released under the MIT license.</sub>
