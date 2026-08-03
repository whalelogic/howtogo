package main

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	baseURL  = "https://howtogo.dev"
	pagesDir = "views/pages"
	outFile  = "public/sitemap.xml"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []URL    `xml:"url"`
}

type URL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

// routeOverrides maps a .templ filename (minus extension, underscores intact)
// to its actual registered route in main.go, for the cases where the two diverge.
var routeOverrides = map[string]string{
	"workerpools": "worker-pools",
}

// excludedRoutes are routes that exist but shouldn't be published in the sitemap
// (internal dashboards, not content pages).
var excludedRoutes = map[string]bool{
	"analytics": true,
}

func main() {
	now := time.Now().Format("2006-01-02")
	var urls []URL

	err := filepath.WalkDir(pagesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		name := d.Name()

		if !strings.HasSuffix(name, ".templ") {
			return nil
		}
		if strings.Contains(name, "_templ") {
			return nil
		}
		if strings.HasSuffix(name, ".bak") || strings.HasSuffix(name, ".backup") {
			return nil
		}

		route := strings.TrimSuffix(name, ".templ")
		if route == "index" {
			route = ""
		} else if override, ok := routeOverrides[route]; ok {
			route = override
		} else {
			route = strings.ReplaceAll(route, "_", "-")
		}

		if excludedRoutes[route] {
			return nil
		}

		urls = append(urls, URL{
			Loc:     baseURL + "/" + route,
			LastMod: now,
		})

		return nil
	})

	if err != nil {
		panic(err)
	}

	sitemap := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}

	if err := os.MkdirAll(filepath.Dir(outFile), 0755); err != nil {
		panic(err)
	}

	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := xml.NewEncoder(f)
	enc.Indent("", "  ")
	if err := enc.Encode(sitemap); err != nil {
		panic(err)
	}
}
