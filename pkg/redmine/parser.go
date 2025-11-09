package redmine

import (
	"math"

	"github.com/mark3labs/mcp-go/mcp"
)

func parsePagination(request *mcp.CallToolRequest) *struct {
	Limit  *int `json:"limit,omitempty" jsonschema:"description=The number of items to be present in the response. If not specified\\, it defaults to 25.,default=25,maximum=100"`
	Nometa *int `json:"nometa,omitempty" jsonschema:"description=If set to 1\\, the response will not include pagination information.,enum=1"`
	Offset *int `json:"offset,omitempty" jsonschema:"description=The offset of the first object to retrieve If not specified\\, it defaults to 0.,default=0"`
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
		Limit  *int `json:"limit,omitempty" jsonschema:"description=The number of items to be present in the response. If not specified\\, it defaults to 25.,default=25,maximum=100"`
		Nometa *int `json:"nometa,omitempty" jsonschema:"description=If set to 1\\, the response will not include pagination information.,enum=1"`
		Offset *int `json:"offset,omitempty" jsonschema:"description=The offset of the first object to retrieve If not specified\\, it defaults to 0.,default=0"`
	}{
		Limit:  pLimit,
		Nometa: pNometa,
		Offset: pOffset,
	}
}

func parseSearchIndexProjectQuery(request *mcp.CallToolRequest) *struct {
	AllWords    *bool   `json:"all_words,omitempty" jsonschema:"description=matched all query strings or not."`
	Attachments *string `json:"attachments,omitempty" jsonschema:"description=Filterd by description and attachment. - \"0\": Seach only in description - \"1\": Search by description and attachment - \"only\": Search only in attachment"`
	Changesets  *bool   `json:"changesets,omitempty" jsonschema:"description=Include changesets or not."`
	Documents   *bool   `json:"documents,omitempty" jsonschema:"description=Include documents or not."`
	Issues      *bool   `json:"issues,omitempty" jsonschema:"description=Include issues or not."`
	Messages    *bool   `json:"messages,omitempty" jsonschema:"description=Include messages or not."`
	News        *bool   `json:"news,omitempty" jsonschema:"description=Include news or not."`
	OpenIssues  *bool   `json:"open_issues,omitempty" jsonschema:"description=Filterd by open issues."`
	Projects    *bool   `json:"projects,omitempty" jsonschema:"description=Include projects or not."`
	Scope       *string `json:"scope,omitempty" jsonschema:"description=Search scope condition. Possible values are: - \"all\": Search all projects - \"my_project\": Search assigned projects - \"bookmarks\": Search bookmarked projects - \"subprojects\": Include subproject when project specified,enum=all,enum=my_project,enum=bookmarks,enum=subprojects"`
	TitlesOnly  *bool   `json:"titles_only,omitempty" jsonschema:"description=matched only title or not."`
	WikiPages   *bool   `json:"wiki_pages,omitempty" jsonschema:"description=Include documents or not."`
} {
	return parseSearchIndexCommonQuery(request)
}

func parseSearchIndexQuery(request *mcp.CallToolRequest) *struct {
	AllWords    *bool   `json:"all_words,omitempty" jsonschema:"description=matched all query strings or not."`
	Attachments *string `json:"attachments,omitempty" jsonschema:"description=Filterd by description and attachment. - \"0\": Seach only in description - \"1\": Search by description and attachment - \"only\": Search only in attachment"`
	Changesets  *bool   `json:"changesets,omitempty" jsonschema:"description=Include changesets or not."`
	Documents   *bool   `json:"documents,omitempty" jsonschema:"description=Include documents or not."`
	Issues      *bool   `json:"issues,omitempty" jsonschema:"description=Include issues or not."`
	Messages    *bool   `json:"messages,omitempty" jsonschema:"description=Include messages or not."`
	News        *bool   `json:"news,omitempty" jsonschema:"description=Include news or not."`
	OpenIssues  *bool   `json:"open_issues,omitempty" jsonschema:"description=Filterd by open issues."`
	Projects    *bool   `json:"projects,omitempty" jsonschema:"description=Include projects or not."`
	Scope       *string `json:"scope,omitempty" jsonschema:"description=Search scope condition. Possible values are: - \"all\": Search all projects - \"my_project\": Search assigned projects - \"bookmarks\": Search bookmarked projects - \"subprojects\": Include subproject when project specified,enum=all,enum=my_project,enum=bookmarks,enum=subprojects"`
	TitlesOnly  *bool   `json:"titles_only,omitempty" jsonschema:"description=matched only title or not."`
	WikiPages   *bool   `json:"wiki_pages,omitempty" jsonschema:"description=Include documents or not."`
} {
	return parseSearchIndexCommonQuery(request)
}

func parseSearchIndexCommonQuery(request *mcp.CallToolRequest) *struct {
	AllWords    *bool   `json:"all_words,omitempty" jsonschema:"description=matched all query strings or not."`
	Attachments *string `json:"attachments,omitempty" jsonschema:"description=Filterd by description and attachment. - \"0\": Seach only in description - \"1\": Search by description and attachment - \"only\": Search only in attachment"`
	Changesets  *bool   `json:"changesets,omitempty" jsonschema:"description=Include changesets or not."`
	Documents   *bool   `json:"documents,omitempty" jsonschema:"description=Include documents or not."`
	Issues      *bool   `json:"issues,omitempty" jsonschema:"description=Include issues or not."`
	Messages    *bool   `json:"messages,omitempty" jsonschema:"description=Include messages or not."`
	News        *bool   `json:"news,omitempty" jsonschema:"description=Include news or not."`
	OpenIssues  *bool   `json:"open_issues,omitempty" jsonschema:"description=Filterd by open issues."`
	Projects    *bool   `json:"projects,omitempty" jsonschema:"description=Include projects or not."`
	Scope       *string `json:"scope,omitempty" jsonschema:"description=Search scope condition. Possible values are: - \"all\": Search all projects - \"my_project\": Search assigned projects - \"bookmarks\": Search bookmarked projects - \"subprojects\": Include subproject when project specified,enum=all,enum=my_project,enum=bookmarks,enum=subprojects"`
	TitlesOnly  *bool   `json:"titles_only,omitempty" jsonschema:"description=matched only title or not."`
	WikiPages   *bool   `json:"wiki_pages,omitempty" jsonschema:"description=Include documents or not."`
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
		AllWords    *bool   `json:"all_words,omitempty" jsonschema:"description=matched all query strings or not."`
		Attachments *string `json:"attachments,omitempty" jsonschema:"description=Filterd by description and attachment. - \"0\": Seach only in description - \"1\": Search by description and attachment - \"only\": Search only in attachment"`
		Changesets  *bool   `json:"changesets,omitempty" jsonschema:"description=Include changesets or not."`
		Documents   *bool   `json:"documents,omitempty" jsonschema:"description=Include documents or not."`
		Issues      *bool   `json:"issues,omitempty" jsonschema:"description=Include issues or not."`
		Messages    *bool   `json:"messages,omitempty" jsonschema:"description=Include messages or not."`
		News        *bool   `json:"news,omitempty" jsonschema:"description=Include news or not."`
		OpenIssues  *bool   `json:"open_issues,omitempty" jsonschema:"description=Filterd by open issues."`
		Projects    *bool   `json:"projects,omitempty" jsonschema:"description=Include projects or not."`
		Scope       *string `json:"scope,omitempty" jsonschema:"description=Search scope condition. Possible values are: - \"all\": Search all projects - \"my_project\": Search assigned projects - \"bookmarks\": Search bookmarked projects - \"subprojects\": Include subproject when project specified,enum=all,enum=my_project,enum=bookmarks,enum=subprojects"`
		TitlesOnly  *bool   `json:"titles_only,omitempty" jsonschema:"description=matched only title or not."`
		WikiPages   *bool   `json:"wiki_pages,omitempty" jsonschema:"description=Include documents or not."`
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
