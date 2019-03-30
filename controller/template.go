package controller

const index = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="theme-color" content="#000000">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <link rel="apple-touch-icon-precomposed" sizes="144x144" href="/static/favicon.png" />
  <link rel="apple-touch-icon-precomposed" sizes="152x152" href="/static/favicon.png" />
  <link rel="icon" type="image/png" href="/static/favicon.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="/static/favicon.png" sizes="16x16" />
  <title>Star Charts</title>
  <meta name="description" content="StarCharts">
  <meta name="author" content="https://github/caarlos0">
  <link rel="stylesheet" href="/static/styles.css">
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/styles/darcula.min.css">
</head>
<body>
  <div class="title">
    <span class="title"><a href="/">starcharts</a></span>
    <span class="subtitle">Plot your repository stars over time.</span>
  </div>
  {{ if .FullName }}
  <div class="main">
    <p>
      {{ if .StargazersCount }}
        <b>Awesome!</b>
      {{ else }}
        <b>Hang in there!</b>
      {{ end }}
      <a href="https://github.com/{{ .FullName }}">{{ .FullName }}</a>
      was created <time datetime="{{ .CreatedAt }}"></time>
      and now has <b>{{ .StargazersCount }}</b> stars.
    </p>
  </div>
  {{ if .StargazersCount }}
  <div class="chart">
    <p>
      <img src="/{{ .FullName }}.svg"
        alt="Please try again in a few minutes. This might not work for very famous repository.">
    </p>
  </div>
  <div class="code">
    <p>
      You can include the chart on your repository's <code>README.md</code>
      as follows:
      <pre><code class="markdown">
## Stargazers over time

[![Stargazers over time](https://starchart.cc/{{ .FullName }}.svg)](https://starchart.cc/{{ .FullName }})
      </code></pre>
    </p>
  </div>
  {{ end }}
  {{ else }}
  <div class="main">
    Check some repository chart by browsing to user/repo, for example,
    <a href="/caarlos0/starcharts">
      caarlos0/starcharts
    </a>
  </div>
  {{ end }}
  <script src="//cdnjs.cloudflare.com/ajax/libs/timeago.js/3.0.2/timeago.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/highlight.min.js"></script>
  <script>
    hljs.initHighlightingOnLoad();
    timeago().render(document.querySelectorAll('time'));
  </script>
</body>
</html>
`
