package redmine

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
	registerAttachmentsDownload(s)
	registerAttachmentsThumbnail(s)
	registerAttachmentsThumbnailSize(s)
	if !readonly {
		registerAttachmentsDestroy(s)
	}
	registerAttachmentsShow(s)
	registerAttachmentsDownloadAll(s)
	registerCustomFieldsIndex(s)
	registerEnumerationsIndexDocumentCategory(s)
	registerEnumerationsIndexIssuePriority(s)
	registerEnumerationsIndexTimeEntryActivity(s)
	registerGroupsIndex(s)
	if !readonly {
		registerGroupsDestroy(s)
	}
	registerGroupsShow(s)
	if !readonly {
		registerGroupsRemoveUser(s)
	}
	registerIssuesIndexCsv(s)
	registerIssuesIndex(s)
	registerIssuesIndexPdf(s)
	registerGanttsShowPdf(s)
	registerGanttsShowPng(s)
	if !readonly {
		registerIssuesDestroy(s)
	}
	registerIssuesShow(s)
	registerIssuesShowPdf(s)
	registerIssueRelationsIndex(s)
	if !readonly {
		registerWatchersDestroyIssue(s)
	}
	if !readonly {
		registerIssueCategoriesDestroy(s)
	}
	registerIssueCategoriesShow(s)
	registerIssueStatusesIndex(s)
	if !readonly {
		registerMembersDestroy(s)
	}
	registerMembersShow(s)
	registerMyAccount(s)
	registerNewsIndex(s)
	if !readonly {
		registerNewsDestroy(s)
	}
	registerNewsShow(s)
	registerProjectsIndexCsv(s)
	registerProjectsIndex(s)
	if !readonly {
		registerProjectsDestroy(s)
	}
	registerProjectsShow(s)
	if !readonly {
		registerProjectsArchivePost(s)
	}
	registerFilesIndex(s)
	if !readonly {
		registerProjectsUnarchivePost(s)
	}
	registerIssuesIndexProjectCsv(s)
	registerIssuesIndexProject(s)
	registerIssuesIndexProjectPdf(s)
	registerIssueCategoriesIndex(s)
	registerGanttsShowProjectPdf(s)
	registerGanttsShowProjectPng(s)
	registerMembersIndex(s)
	registerNewsIndexProject(s)
	if !readonly {
		registerRepositoriesRemoveRelatedIssue(s)
	}
	registerSearchIndexProject(s)
	registerTimelogIndexProjectCsv(s)
	registerTimelogIndexProject(s)
	registerVersionsIndex(s)
	registerWikiShowRoot(s)
	registerWikiIndex(s)
	if !readonly {
		registerWikiDestroy(s)
	}
	registerWikiShow(s)
	registerWikiShowPdf(s)
	registerWikiShowTxt(s)
	registerWikiShowVersion(s)
	registerWikiShowVersionPdf(s)
	registerWikiShowVersionTxt(s)
	registerQueriesIndex(s)
	if !readonly {
		registerIssueRelationsDestroy(s)
	}
	registerIssueRelationsShow(s)
	registerRolesIndex(s)
	registerRolesShow(s)
	registerSearchIndex(s)
	registerTimelogIndexCsv(s)
	registerTimelogIndex(s)
	if !readonly {
		registerTimelogDestroy(s)
	}
	registerTimelogShow(s)
	registerTrackersIndex(s)
	registerUsersIndexCsv(s)
	registerUsersIndex(s)
	if !readonly {
		registerUsersDestroy(s)
	}
	registerUsersShow(s)
	if !readonly {
		registerWatchersDestroy(s)
	}
	registerVersionsShowTxt(s)
	if !readonly {
		registerVersionsDestroy(s)
	}
	registerVersionsShow(s)
}
