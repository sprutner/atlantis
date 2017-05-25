package main

type HelpExecutor struct {
	BaseExecutor
}

var helpComment = "```cmake\n" +
`atlantis - Terraform collaboration tool that enables you to collaborate on infrastructure
safely and securely. (v` + version + `)

Usage: atlantis <command> [environment] [--verbose]

Commands:
plan           Runs 'terraform plan' on the files changed in the pull request
apply          Runs 'terraform apply' using the plans generated by 'atlantis plan'
help           Get help

Examples:

# Generates a plan for staging environment
atlantis plan staging

# Generates a plan for a standalone terraform project
atlantis plan

# Applies a plan for staging environment
atlantis apply staging

# Applies a plan for a standalone terraform project
atlantis apply
`

func (h *HelpExecutor) execute(ctx *ExecutionContext, github *GithubClient) {
	prCtx := &PullRequestContext{
		owner:                 ctx.repoOwner,
		repoName:              ctx.repoName,
		head:                  ctx.head,
		base:                  ctx.base,
		number:                ctx.pullNum,
		pullRequestLink:       ctx.pullLink,
		terraformApplier:      ctx.requesterUsername,
		terraformApplierEmail: ctx.requesterEmail,
	}
	github.CreateComment(prCtx, helpComment)

	ctx.log.Info("generating help comment....")
	return
}
