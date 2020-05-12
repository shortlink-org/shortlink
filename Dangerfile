require_relative 'danger/lib/gitlab_danger'

danger.import_plugin('danger/plugins/helper.rb')

RELEASE_TEAM = %w[
    batazor
]

MAINTAINERS = %w[
    batazor
]

unless helper.release_automation?
  GitlabDanger.new(helper.gitlab_helper).rule_names.each do |file|
    danger.import_dangerfile(path: File.join('danger', file))
  end
end
