# GitLab configuration

```json:table
{
  "caption" : "GitLab configuration",
  "fields" : [
    {"key": "folder", "label": "folder"},
    {"key": "decsription", "label": "decsription"},
    {"key": "type", "label": "type", "sortable": true}
  ],
  "items" : [
    {
      "folder": "agents", 
      "decsription": "List k8s agent for GitLab",
      "type": "folder"
    },
    {
      "folder": "issue_templates", 
      "decsription": "Templates for issues",
      "type": "folder"
    },
    {
      "folder": "merge_request_templates", 
      "decsription": "Templates for merge requests",
      "type": "folder"
    },
    {
      "folder": "CODEOWNERS", 
      "decsription": "The list owners",
      "type": "file"
    }
  ],
  "filter" : true
}
```
