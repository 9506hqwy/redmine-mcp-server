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
	if !readonly {
		registerAttachmentsUpdatePatch(s)
	}
	registerAttachmentsShow(s)
	registerAttachmentsDownloadAll(s)
	registerCustomFieldsIndex(s)
	registerEnumerationsIndexDocumentCategory(s)
	registerEnumerationsIndexIssuePriority(s)
	registerEnumerationsIndexTimeEntryActivity(s)
	if !readonly {
		registerGroupsCreate(s)
	}
	registerGroupsIndex(s)
	if !readonly {
		registerGroupsDestroy(s)
	}
	if !readonly {
		registerGroupsUpdatePatch(s)
	}
	registerGroupsShow(s)
	if !readonly {
		registerGroupsAddUsers(s)
	}
	if !readonly {
		registerGroupsRemoveUser(s)
	}
	registerIssuesIndexCsv(s)
	if !readonly {
		registerIssuesCreate(s)
	}
	registerIssuesIndex(s)
	registerIssuesIndexPdf(s)
	registerGanttsShowPdf(s)
	registerGanttsShowPng(s)
	if !readonly {
		registerIssuesDestroy(s)
	}
	if !readonly {
		registerIssuesUpdatePatch(s)
	}
	registerIssuesShow(s)
	registerIssuesShowPdf(s)
	if !readonly {
		registerIssueRelationsCreate(s)
	}
	registerIssueRelationsIndex(s)
	if !readonly {
		registerTimelogCreateIssue(s)
	}
	if !readonly {
		registerWatchersCreateIssue(s)
	}
	if !readonly {
		registerWatchersDestroyIssue(s)
	}
	if !readonly {
		registerIssueCategoriesDestroy(s)
	}
	if !readonly {
		registerIssueCategoriesUpdatePatch(s)
	}
	registerIssueCategoriesShow(s)
	registerIssueStatusesIndex(s)
	if !readonly {
		registerJournalsUpdatePatch(s)
	}
	if !readonly {
		registerMembersDestroy(s)
	}
	if !readonly {
		registerMembersUpdatePatch(s)
	}
	registerMembersShow(s)
	registerMyAccount(s)
	if !readonly {
		registerNewsCreate(s)
	}
	registerNewsIndex(s)
	if !readonly {
		registerNewsDestroy(s)
	}
	if !readonly {
		registerNewsUpdatePatch(s)
	}
	registerNewsShow(s)
	registerProjectsIndexCsv(s)
	if !readonly {
		registerProjectsCreate(s)
	}
	registerProjectsIndex(s)
	if !readonly {
		registerProjectsDestroy(s)
	}
	if !readonly {
		registerProjectsUpdatePatch(s)
	}
	registerProjectsShow(s)
	if !readonly {
		registerProjectsArchivePost(s)
	}
	if !readonly {
		registerFilesCreate(s)
	}
	registerFilesIndex(s)
	if !readonly {
		registerProjectsUnarchivePost(s)
	}
	registerIssuesIndexProjectCsv(s)
	if !readonly {
		registerIssuesCreateProject(s)
	}
	registerIssuesIndexProject(s)
	registerIssuesIndexProjectPdf(s)
	if !readonly {
		registerIssueCategoriesCreate(s)
	}
	registerIssueCategoriesIndex(s)
	registerGanttsShowProjectPdf(s)
	registerGanttsShowProjectPng(s)
	if !readonly {
		registerMembersCreate(s)
	}
	registerMembersIndex(s)
	if !readonly {
		registerNewsCreateProject(s)
	}
	registerNewsIndexProject(s)
	if !readonly {
		registerRepositoriesAddRelatedIssue(s)
	}
	if !readonly {
		registerRepositoriesRemoveRelatedIssue(s)
	}
	registerSearchIndexProject(s)
	registerTimelogIndexProjectCsv(s)
	if !readonly {
		registerTimelogCreateProject(s)
	}
	registerTimelogIndexProject(s)
	if !readonly {
		registerVersionsCreate(s)
	}
	registerVersionsIndex(s)
	registerWikiShowRoot(s)
	registerWikiIndex(s)
	if !readonly {
		registerWikiDestroy(s)
	}
	if !readonly {
		registerWikiUpdatePatch(s)
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
	if !readonly {
		registerTimelogCreate(s)
	}
	registerTimelogIndex(s)
	if !readonly {
		registerTimelogDestroy(s)
	}
	if !readonly {
		registerTimelogUpdatePatch(s)
	}
	registerTimelogShow(s)
	registerTrackersIndex(s)
	// if !readonly { registerAttachmentsUpload(s) }
	registerUsersIndexCsv(s)
	if !readonly {
		registerUsersCreate(s)
	}
	registerUsersIndex(s)
	if !readonly {
		registerUsersDestroy(s)
	}
	if !readonly {
		registerUsersUpdatePatch(s)
	}
	registerUsersShow(s)
	if !readonly {
		registerWatchersDestroy(s)
	}
	if !readonly {
		registerWatchersCreate(s)
	}
	registerVersionsShowTxt(s)
	if !readonly {
		registerVersionsDestroy(s)
	}
	if !readonly {
		registerVersionsUpdatePatch(s)
	}
	registerVersionsShow(s)
}
