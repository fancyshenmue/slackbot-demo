package middleware

import (
	"fmt"

	"github.com/slack-go/slack"
)

func SlackBlockKitDiv() *slack.DividerBlock {
	divSection := slack.NewDividerBlock()
	return divSection
}

func SlackBlockKitHeader(header string) *slack.SectionBlock {
	var section *slack.SectionBlock
	switch header {
	case "TopLevel":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Operator Options*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Terraform":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Terraform Operator*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Ansible":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Ansible Operator*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Kubernetes":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Kubernetes Operator*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Vault":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Vault Operator*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Gitlab":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Gitlab Operator*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Jenkins":
		txt := slack.NewTextBlockObject("mrkdwn", "*Choose Jenkins Operator*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	case "Action":
		txt := slack.NewTextBlockObject("mrkdwn", "*Action*", false, false)
		section = slack.NewSectionBlock(txt, nil, nil)

	}
	return section
}

func Operator() *slack.SectionBlock {
	options := slack.NewOptionsSelectBlockElement(
		"static_select",
		slack.NewTextBlockObject("plain_text", "Operator", true, false),
		"operator",
		slack.NewOptionBlockObject("Terraform", slack.NewTextBlockObject("plain_text", "Terraform", false, false), nil),
		slack.NewOptionBlockObject("Ansible", slack.NewTextBlockObject("plain_text", "Ansible", false, false), nil),
		slack.NewOptionBlockObject("Kubernetes", slack.NewTextBlockObject("plain_text", "Kubernetes", false, false), nil),
		slack.NewOptionBlockObject("Vault", slack.NewTextBlockObject("plain_text", "Vault", false, false), nil),
		slack.NewOptionBlockObject("Gitlab", slack.NewTextBlockObject("plain_text", "Gitlab", false, false), nil),
		slack.NewOptionBlockObject("Jenkins", slack.NewTextBlockObject("plain_text", "Jenkins", false, false), nil),
	)
	txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
	section := slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))
	return section
}

func OperatorOptions(option string) *slack.SectionBlock {
	var section *slack.SectionBlock
	switch option {
	case "Terraform":
		options := slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject("plain_text", "Action", true, false),
			"action",
			slack.NewOptionBlockObject("apply", slack.NewTextBlockObject("plain_text", "apply", false, false), nil),
			slack.NewOptionBlockObject("destroy", slack.NewTextBlockObject("plain_text", "destroy", false, false), nil),
		)
		txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
		section = slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))

	case "Ansible":
		options := slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject("plain_text", "Action", true, false),
			"action",
			slack.NewOptionBlockObject("ad-hoc", slack.NewTextBlockObject("plain_text", "ad-hoc", false, false), nil),
			slack.NewOptionBlockObject("playbook", slack.NewTextBlockObject("plain_text", "playbook", false, false), nil),
		)
		txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
		section = slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))

	case "Kubernetes":
		options := slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject("plain_text", "Action", true, false),
			"action",
			slack.NewOptionBlockObject("apply", slack.NewTextBlockObject("plain_text", "apply", false, false), nil),
			slack.NewOptionBlockObject("delete", slack.NewTextBlockObject("plain_text", "delete", false, false), nil),
		)
		txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
		section = slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))

	case "Vault":
		options := slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject("plain_text", "Action", true, false),
			"action",
			slack.NewOptionBlockObject("apply", slack.NewTextBlockObject("plain_text", "apply", false, false), nil),
			slack.NewOptionBlockObject("destroy", slack.NewTextBlockObject("plain_text", "destroy", false, false), nil),
		)
		txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
		section = slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))

	case "Gitlab":
		options := slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject("plain_text", "Action", true, false),
			"action",
			slack.NewOptionBlockObject("get", slack.NewTextBlockObject("plain_text", "get", false, false), nil),
			slack.NewOptionBlockObject("create", slack.NewTextBlockObject("plain_text", "create", false, false), nil),
			slack.NewOptionBlockObject("update", slack.NewTextBlockObject("plain_text", "update", false, false), nil),
			slack.NewOptionBlockObject("remove", slack.NewTextBlockObject("plain_text", "remove", false, false), nil),
		)
		txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
		section = slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))

	case "Jenkins":
		options := slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject("plain_text", "Action", true, false),
			"action",
			slack.NewOptionBlockObject("get", slack.NewTextBlockObject("plain_text", "get", false, false), nil),
			slack.NewOptionBlockObject("create", slack.NewTextBlockObject("plain_text", "create", false, false), nil),
			slack.NewOptionBlockObject("update", slack.NewTextBlockObject("plain_text", "update", false, false), nil),
			slack.NewOptionBlockObject("remove", slack.NewTextBlockObject("plain_text", "remove", false, false), nil),
			slack.NewOptionBlockObject("trigger", slack.NewTextBlockObject("plain_text", "trigger", false, false), nil),
		)
		txt := slack.NewTextBlockObject("mrkdwn", "*Operator*\nChoose Operator", false, false)
		section = slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))

	}
	return section
}

func DryRun() *slack.SectionBlock {
	options := slack.NewOptionsSelectBlockElement(
		"static_select",
		slack.NewTextBlockObject("plain_text", "Action", true, false),
		"dryRun",
		slack.NewOptionBlockObject("false", slack.NewTextBlockObject("plain_text", "false", false, false), nil),
		slack.NewOptionBlockObject("true", slack.NewTextBlockObject("plain_text", "true", false, false), nil),
	)
	txt := slack.NewTextBlockObject("mrkdwn", "*DryRun*\nChoose DryRun (bool)", false, false)
	section := slack.NewSectionBlock(txt, nil, slack.NewAccessory(options))
	return section
}

func SubmitButton() *slack.ActionBlock {
	button := slack.NewButtonBlockElement(
		"submit",
		"submit_button",
		slack.NewTextBlockObject("plain_text", "Submit", false, false),
	)
	section := slack.NewActionBlock("submit_button", button)
	return section
}

func SlackBlockKitResponse(o OptionsStruct) *slack.ContextBlock {
	var context *slack.ContextBlock

	switch o.Operator {
	case "Terraform":
		context = slack.NewContextBlock(
			"response",
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf(`
*Operator*: *%s*
*Action*: *%s*
*DryRun*: *%s*`,
					o.Operator,
					o.Action,
					o.DryRun,
				),
				false,
				false,
			),
		)

	case "Ansible":
		context = slack.NewContextBlock(
			"response",
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf(`
*Operator*: *%s*
*Action*: *%s*
*DryRun*: *%s*`,
					o.Operator,
					o.Action,
					o.DryRun,
				),
				false,
				false,
			),
		)

	case "Kubernetes":
		context = slack.NewContextBlock(
			"response",
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf(`
*Operator*: *%s*
*Action*: *%s*
*DryRun*: *%s*`,
					o.Operator,
					o.Action,
					o.DryRun,
				),
				false,
				false,
			),
		)

	case "Vault":
		context = slack.NewContextBlock(
			"response",
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf(`
*Operator*: *%s*
*Action*: *%s*
*DryRun*: *%s*`,
					o.Operator,
					o.Action,
					o.DryRun,
				),
				false,
				false,
			),
		)

	case "Gitlab":
		context = slack.NewContextBlock(
			"response",
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf(`
*Operator*: *%s*
*Action*: *%s*
*DryRun*: *%s*`,
					o.Operator,
					o.Action,
					o.DryRun,
				),
				false,
				false,
			),
		)

	case "Jenkins":
		context = slack.NewContextBlock(
			"response",
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf(`
*Operator*: *%s*
*Action*: *%s*
*DryRun*: *%s*`,
					o.Operator,
					o.Action,
					o.DryRun,
				),
				false,
				false,
			),
		)
	}

	return context
}
