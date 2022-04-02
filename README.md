# Unofficial Tribun News API

Original site [tribunnews.com](https://www.tribunnews.com)

## Install

```bash
git clone https://github.com/rizki4106/tribun-news-unofficial-api.git
```

```bash
cd tribun-news-unofficial-api && go mod download
```

or run inside docker container

```bash
cd tribun-news-unofficial-api && docker-compose up
```

## Available Topics

| No  | Name          |
| --- | ------------- |
| 1   | makro         |
| 2   | energi        |
| 3   | finansial     |
| 4   | mikro         |
| 5   | investasi     |
| 6   | transportasi  |
| 7   | infrastruktur |
| 8   | insight       |
| 9   | property      |

## Endpoint

```text
{url}/{type}
```

## Example

```text
http://localhost:4000/mikro
```

## Response

```json
{
  "total_item": 90,
  "error": false,
  "error_message": "",
  "items": [
    {
      "title": "Demo di Jakarta",
      "poster": "https://cdn-2.tstatic.net/tribunnews/foto/bank/images2/hariyadi-sukamdani-apindo_20160908_180423.jpg",
      "link": "https://www.tribunnews.com/topic/demo-di-jakarta",
      "date": "",
      "short_description": ""
    },
    {
      "title": "Trump Menang, Indonesia Akan Dibanjiri Barang Impor dari China, Ini Sebabnya",
      "poster": "https://cdn-2.tstatic.net/tribunnews/foto/bank/thumbnails2/donald-trump-as_20161109_230611.jpg",
      "link": "https://www.tribunnews.com/bisnis/2016/11/10/trump-menang-indonesia-akan-dibanjiri-barang-impor-dari-china-ini-sebabnya",
      "date": "Kamis, 10 November 2016 ",
      "short_description": "Trump Menang, Indonesia Akan Dibanjiri Barang Impor dari China, Ini Sebabnya "
    }
  ]
}
```
