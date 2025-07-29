package redmine

import (
	"math"

	"github.com/mark3labs/mcp-go/mcp"
)

func parsePagination(request *mcp.CallToolRequest) *struct {
	Limit  *int `json:"limit,omitempty"`
	Nometa *int `json:"nometa,omitempty"`
	Offset *int `json:"offset,omitempty"`
} {
	var pLimit *int = nil
	var pNometa *int = nil
	var pOffset *int = nil

	limit := request.GetInt("pagination.limit", math.MinInt)
	if limit != math.MinInt {
		pLimit = &limit
	}

	nometa := request.GetInt("pagination.nometa", math.MinInt)
	if nometa != math.MinInt {
		pNometa = &nometa
	}

	offset := request.GetInt("pagination.offset", math.MinInt)
	if offset != math.MinInt {
		pOffset = &offset
	}

	return &struct {
		Limit  *int `json:"limit,omitempty"`
		Nometa *int `json:"nometa,omitempty"`
		Offset *int `json:"offset,omitempty"`
	}{
		Limit:  pLimit,
		Nometa: pNometa,
		Offset: pOffset,
	}
}

func parseSearchIndexProjectQuery(request *mcp.CallToolRequest) *struct {
	AllWords    *bool   `json:"all_words,omitempty"`
	Attachments *string `json:"attachments,omitempty"`
	Changesets  *bool   `json:"changesets,omitempty"`
	Documents   *bool   `json:"documents,omitempty"`
	Issues      *bool   `json:"issues,omitempty"`
	Messages    *bool   `json:"messages,omitempty"`
	News        *bool   `json:"news,omitempty"`
	OpenIssues  *bool   `json:"open_issues,omitempty"`
	Projects    *bool   `json:"projects,omitempty"`
	Scope       *string `json:"scope,omitempty"`
	TitlesOnly  *bool   `json:"titles_only,omitempty"`
	WikiPages   *bool   `json:"wiki_pages,omitempty"`
} {
	return parseSearchIndexCommonQuery(request)
}

func parseSearchIndexQuery(request *mcp.CallToolRequest) *struct {
	AllWords    *bool   `json:"all_words,omitempty"`
	Attachments *string `json:"attachments,omitempty"`
	Changesets  *bool   `json:"changesets,omitempty"`
	Documents   *bool   `json:"documents,omitempty"`
	Issues      *bool   `json:"issues,omitempty"`
	Messages    *bool   `json:"messages,omitempty"`
	News        *bool   `json:"news,omitempty"`
	OpenIssues  *bool   `json:"open_issues,omitempty"`
	Projects    *bool   `json:"projects,omitempty"`
	Scope       *string `json:"scope,omitempty"`
	TitlesOnly  *bool   `json:"titles_only,omitempty"`
	WikiPages   *bool   `json:"wiki_pages,omitempty"`
} {
	return parseSearchIndexCommonQuery(request)
}

func parseSearchIndexCommonQuery(request *mcp.CallToolRequest) *struct {
	AllWords    *bool   `json:"all_words,omitempty"`
	Attachments *string `json:"attachments,omitempty"`
	Changesets  *bool   `json:"changesets,omitempty"`
	Documents   *bool   `json:"documents,omitempty"`
	Issues      *bool   `json:"issues,omitempty"`
	Messages    *bool   `json:"messages,omitempty"`
	News        *bool   `json:"news,omitempty"`
	OpenIssues  *bool   `json:"open_issues,omitempty"`
	Projects    *bool   `json:"projects,omitempty"`
	Scope       *string `json:"scope,omitempty"`
	TitlesOnly  *bool   `json:"titles_only,omitempty"`
	WikiPages   *bool   `json:"wiki_pages,omitempty"`
} {
	var pScope *string = nil
	var pAllWords *bool = nil
	var pTitlesOnly *bool = nil
	var pIssues *bool = nil
	var pNews *bool = nil
	var pDocuments *bool = nil
	var pChangesets *bool = nil
	var pWikiPages *bool = nil
	var pMessages *bool = nil
	var pProjects *bool = nil
	var pOpenIssues *bool = nil
	var pAttachments *string = nil

	scope := request.GetString("scope", "")
	if scope != "" {
		pScope = &scope
	}

	all_words := request.GetBool("all_words", false)
	pAllWords = &all_words

	titles_only := request.GetBool("titles_only", false)
	pTitlesOnly = &titles_only

	issues := request.GetBool("issues", false)
	pIssues = &issues

	news := request.GetBool("news", false)
	pNews = &news

	documents := request.GetBool("documents", false)
	pDocuments = &documents

	changesets := request.GetBool("changesets", false)
	pChangesets = &changesets

	wiki_pages := request.GetBool("wiki_pages", false)
	pWikiPages = &wiki_pages

	messages := request.GetBool("messages", false)
	pMessages = &messages

	projects := request.GetBool("projects", false)
	pProjects = &projects

	open_issues := request.GetBool("open_issues", false)
	pOpenIssues = &open_issues

	attachments := request.GetString("attachments", "")
	if attachments != "" {

		pAttachments = &attachments
	}

	return &struct {
		AllWords    *bool   "json:\"all_words,omitempty\""
		Attachments *string "json:\"attachments,omitempty\""
		Changesets  *bool   "json:\"changesets,omitempty\""
		Documents   *bool   "json:\"documents,omitempty\""
		Issues      *bool   "json:\"issues,omitempty\""
		Messages    *bool   "json:\"messages,omitempty\""
		News        *bool   "json:\"news,omitempty\""
		OpenIssues  *bool   "json:\"open_issues,omitempty\""
		Projects    *bool   "json:\"projects,omitempty\""
		Scope       *string "json:\"scope,omitempty\""
		TitlesOnly  *bool   "json:\"titles_only,omitempty\""
		WikiPages   *bool   "json:\"wiki_pages,omitempty\""
	}{
		AllWords:    pAllWords,
		Attachments: pAttachments,
		Changesets:  pChangesets,
		Documents:   pDocuments,
		Issues:      pIssues,
		Messages:    pMessages,
		News:        pNews,
		OpenIssues:  pOpenIssues,
		Projects:    pProjects,
		Scope:       pScope,
		TitlesOnly:  pTitlesOnly,
		WikiPages:   pWikiPages,
	}
}
