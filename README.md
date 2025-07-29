# Redmine MCP Server

This repository provides model context protocol server for Redmine server.

## Build

Build binary.

```sh
go build -o bin/redmine-mcp-server ./cmd/redmine-mcp-server/main.go
```

Or build container image.

```sh
docker build -t redmine-mcp-server .
```

Add `Z` option at bind mount operation in *Dockerfile* if using podman with SELinux.

## Usage

Run application.

Specify API key or username and password.

```sh
$ ./bin/redmine-mcp-server -h
Redmine MCP Server

Usage:
  redmine-mcp-server [flags]

Flags:
      --apikey string     Redmine server API key.
  -h, --help              help for redmine-mcp-server
      --password string   Redmine server password.
      --readonly          HTTP GET method only. (default true)
      --url string        Redmine server URL. (default "http://127.0.0.1:3000")
      --user string       Redmine server username.
  -v, --version           version for redmine-mcp-server
```

Set environment variable instead of arguments.

| Argument   | Environment Variable |
| :--------- | :------------------- |
| --url      | REDMINE_URL          |
| --user     | REDMINE_USER         |
| --password | REDMINE_PASSWORD     |
| --apikey   | REDMINE_APIKEY       |
| --readonly | REDMINE_READONLY     |

Or run container.

```sh
docker run --rm -i -e REDMINE_URL=<URL> -e REDMINE_APIKEY=<API Key> redmine-mcp-server
```

### Usage with VS code

Add `redmine-mcp-server` binary to `PATH` environment variable and configure VS code.

```json
{
    "servers": {
        "redmine": {
            "type": "stdio",
            "command": "redmine-mcp-server",
            "env": {
                "REDMINE_URL": "${input:redmine_url}",
                "REDMINE_APIKEY": "${input:redmine_apikey}",
            }
        }
    },
    "inputs": [
        {
            "type": "promptString",
            "id": "redmine_url",
            "description": "Redmine Server URL",
            "password": false
        },
        {
            "type": "promptString",
            "id": "redmine_apikey",
            "description": "Redmine Server API key",
            "password": true
        }
    ]
}
```

## Tools

| Tool                                   | Description                                                                                                                                                                  |
| :------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| attachments_download                   | Download the attachment with the specified ID.                                                                                                                               |
| attachments_thumbnail                  | Download the thumbnail of the attachment with the specified ID.                                                                                                              |
| attachments_thumbnail_size             | Downloads the attachment thumbnail with the specified ID.                                                                                                                    |
| attachments_destroy                    | Deletes the attachment with the specified ID.                                                                                                                                |
| attachments_show                       | Returns the attachment with the specified ID.                                                                                                                                |
| attachments_download_all               | Downloads the attachment with the specified ID.                                                                                                                              |
| custom_fields_index                    | Returns all custom field definitions.                                                                                                                                        |
| enumerations_index_document_category   | Returns a list of document categories.                                                                                                                                       |
| enumerations_index_issue_priority      | Returns a list of issue priorities.                                                                                                                                          |
| enumerations_index_time_entry_activity | Returns a list of time entry activities.                                                                                                                                     |
| groups_index                           | Returns a list of all groups.                                                                                                                                                |
| groups_destroy                         | Deletes the group with the specified ID.                                                                                                                                     |
| groups_show                            | Returns the group with the specified ID.                                                                                                                                     |
| groups_remove_user                     | Removes a user from a group.                                                                                                                                                 |
| issues_index_csv                       | Returns a paginated list of issues. By default, it returns open issues only.                                                                                                 |
| issues_index                           | Returns a paginated list of issues. By default, it returns open issues only.                                                                                                 |
| issues_index_pdf                       | Returns a paginated list of issues. By default, it returns open issues only.                                                                                                 |
| gantts_show_pdf                        | Download the Gantt chart.                                                                                                                                                    |
| gantts_show_png                        | Download the Gantt chart.                                                                                                                                                    |
| issues_destroy                         | Deletes the issue with the specified ID.                                                                                                                                     |
| issues_show                            | Returns the issue with the specified ID.                                                                                                                                     |
| issues_show_pdf                        | Returns the issue with the specified ID.                                                                                                                                     |
| issue_relations_index                  | Returns the relations for the specified issue ID.                                                                                                                            |
| watchers_destroy_issue                 | Deletes the watcher with the specified ID from the issue.                                                                                                                    |
| issue_categories_destroy               | Deletes the issue category with the specified ID.                                                                                                                            |
| issue_categories_show                  | Returns the issue category with the specified ID.                                                                                                                            |
| issue_statuses_index                   | Returns a list of all issue statuses.                                                                                                                                        |
| members_destroy                        | Deletes the membership with the specified ID.                                                                                                                                |
| members_show                           | Returns the membership with the specified ID.                                                                                                                                |
| my_account                             | Returns the current user's account information.                                                                                                                              |
| news_index                             | Returns all news items across all projects with pagination.                                                                                                                  |
| news_destroy                           | Deletes the news with the specified ID.                                                                                                                                      |
| news_show                              | Returns the news item with the specified ID.                                                                                                                                 |
| projects_index_csv                     | Returns all projects (including all public projects and private projects to which the user has access).                                                                      |
| projects_index                         | Returns all projects (including all public projects and private projects to which the user has access).                                                                      |
| projects_destroy                       | Deletes the project with the specified ID or identifier.                                                                                                                     |
| projects_show                          | Returns the project with the specified ID or identifier.                                                                                                                     |
| projects_archive_post                  | Archives the project with the specified ID or identifier.                                                                                                                    |
| files_index                            | Returns a list of all files.                                                                                                                                                 |
| projects_unarchive_post                | Unarchives the project with the specified ID or identifier.                                                                                                                  |
| issues_index_project_csv               | Returns a paginated list of issues. By default, it returns open issues only.                                                                                                 |
| issues_index_project                   | Returns a paginated list of issues. By default, it returns open issues only.                                                                                                 |
| issues_index_project_pdf               | Returns a paginated list of issues. By default, it returns open issues only.                                                                                                 |
| issue_categories_index                 | Returns the issue categories available for the specified project by ID or identifier.                                                                                        |
| gantts_show_project_pdf                | Download the Gantt chart for the specified project.                                                                                                                          |
| gantts_show_project_png                | Download the Gantt chart for the specified project.                                                                                                                          |
| members_index                          | Returns a paginated list of project memberships.                                                                                                                             |
| news_index_project                     | Returns all news items across all projects with pagination.                                                                                                                  |
| repositories_remove_related_issue      | Remove a related issue from the specified revision.                                                                                                                          |
| search_index_project                   | Returns search results based on the specified query parameters.                                                                                                              |
| timelog_index_project_csv              | Returns a list of time entries for the specified project in CSV format.                                                                                                      |
| timelog_index_project                  | Returns a list of time entries for the specified project.                                                                                                                    |
| versions_index                         | Returns the versions available for the project with the specified ID or identifier (:project_id). The response may include shared versions from other projects.              |
| wiki_show_root                         | Returns the details of the root wiki page.                                                                                                                                   |
| wiki_index                             | Returns a list of all pages in the project wiki.                                                                                                                             |
| wiki_destroy                           | Deletes a wiki page, its attachments and its history with the specified ID. If the deleted page is a parent page, its child pages are not deleted but changed as root pages. |
| wiki_show                              | Returns the details of a wiki page with the specified ID.                                                                                                                    |
| wiki_show_pdf                          | Returns the details of a wiki page with the specified ID.                                                                                                                    |
| wiki_show_txt                          | Returns the details of a wiki page with the specified ID.                                                                                                                    |
| wiki_show_version                      | Returns the details of an old version of a wiki page with the specified ID.                                                                                                  |
| wiki_show_version_pdf                  | Returns the details of an old version of a wiki page with the specified ID.                                                                                                  |
| wiki_show_version_txt                  | Returns the details of an old version of a wiki page with the specified ID.                                                                                                  |
| queries_index                          | Returns a list of all queries.                                                                                                                                               |
| issue_relations_destroy                | Deletes the relation with the specified ID.                                                                                                                                  |
| issue_relations_show                   | Returns the relation with the specified ID.                                                                                                                                  |
| roles_index                            | Returns a list of all roles.                                                                                                                                                 |
| roles_show                             | Returns the role with the specified ID.                                                                                                                                      |
| search_index                           | Returns search results based on the specified query parameters.                                                                                                              |
| timelog_index_csv                      | Returns a list of time entries.                                                                                                                                              |
| timelog_index                          | Returns a list of time entries.                                                                                                                                              |
| timelog_destroy                        | Deletes the time entry with the specified ID.                                                                                                                                |
| timelog_show                           | Returns the time entry with the specified ID.                                                                                                                                |
| trackers_index                         | Returns a list of all trackers.                                                                                                                                              |
| users_index_csv                        | Returns a list of all users in CSV format.                                                                                                                                   |
| users_index                            | Returns a list of all users.                                                                                                                                                 |
| users_destroy                          | Deletes the user with the specified ID.                                                                                                                                      |
| users_show                             | Returns the user with the specified ID. Use /users/current.json to retrieve the user whose credentials is used to access the API.                                            |
| watchers_destroy                       | Deletes the watcher with the specified ID.                                                                                                                                   |
| versions_show_txt                      | Returns the version with the specified ID.                                                                                                                                   |
| versions_destroy                       | Deletes the version with the specified ID.                                                                                                                                   |
| versions_show                          | Returns the version with the specified ID.                                                                                                                                   |

## Testing

TODO

## Notes

* Check off unused tools at [tool icon] in GitHub Copilot Chat panel bacause vscode limits max 128 tools.
