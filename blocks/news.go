package blocks

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	NewsFeedURL        = "https://news.google.com/rss/headlines/section/topic/BUSINESS?hl=zh-CN&gl=CN&ceid=CN:zh-Hans"
	NewsFetchInterval  = 30 * time.Minute
	NewsMaxItems       = 8
	NewsScrollWidth    = 52
	NewsScrollStepRune = 1
)

var newsClient = &http.Client{Timeout: 6 * time.Second}

var (
	newsMu        sync.Mutex
	newsText      string
	newsRunes     []rune
	newsScroll    []rune
	newsOffset    int
	newsLastFetch time.Time
)

type rssDoc struct {
	Channel struct {
		Items []newsTitle `xml:"item"`
	} `xml:"channel"`
}

type atomDoc struct {
	Entries []newsTitle `xml:"entry"`
}

type newsTitle struct {
	Title string `xml:"title"`
}

func BlockNews() string {
	newsMu.Lock()
	defer newsMu.Unlock()

	now := time.Now()
	if newsText == "" || now.Sub(newsLastFetch) > NewsFetchInterval {
		_ = fetchNewsLocked()
	}

	if len(newsRunes) == 0 {
		return "News --"
	}

	if len(newsRunes) <= NewsScrollWidth {
		return string(newsRunes)
	}

	newsOffset = (newsOffset + NewsScrollStepRune) % len(newsScroll)

	var sb strings.Builder
	sb.Grow(NewsScrollWidth)
	for i := 0; i < NewsScrollWidth; i++ {
		idx := (newsOffset + i) % len(newsScroll)
		sb.WriteRune(newsScroll[idx])
	}
	return sb.String()
}

func fetchNewsLocked() error {
	resp, err := newsClient.Get(NewsFeedURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body := &bytes.Buffer{}
	if _, err := body.ReadFrom(resp.Body); err != nil {
		return err
	}

	items := parseNewsTitles(body.Bytes())
	if len(items) == 0 {
		return nil
	}

	newsText = strings.Join(items, " | ")
	newsRunes = []rune(newsText)
	newsScroll = append([]rune{}, newsRunes...)
	newsScroll = append(newsScroll, ' ', ' ', ' ', ' ', ' ')
	newsOffset = 0
	newsLastFetch = time.Now()
	return nil
}

func parseNewsTitles(data []byte) []string {
	var rss rssDoc
	if err := xml.Unmarshal(data, &rss); err == nil && len(rss.Channel.Items) > 0 {
		return limitTitles(rss.Channel.Items)
	}

	var atom atomDoc
	if err := xml.Unmarshal(data, &atom); err == nil && len(atom.Entries) > 0 {
		return limitAtomTitles(atom.Entries)
	}

	return nil
}

func limitTitles(items []newsTitle) []string {
	titles := make([]string, 0, NewsMaxItems)
	for _, item := range items {
		if len(titles) >= NewsMaxItems {
			break
		}
		if title := normalizeTitle(item.Title); title != "" {
			titles = append(titles, title)
		}
	}
	return titles
}

func limitAtomTitles(entries []newsTitle) []string {
	titles := make([]string, 0, NewsMaxItems)
	for _, entry := range entries {
		if len(titles) >= NewsMaxItems {
			break
		}
		if title := normalizeTitle(entry.Title); title != "" {
			titles = append(titles, title)
		}
	}
	return titles
}

func normalizeTitle(title string) string {
	fields := strings.Fields(strings.TrimSpace(title))
	if len(fields) == 0 {
		return ""
	}
	return strings.Join(fields, " ")
}
