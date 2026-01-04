package discord

// Endpoints.
var (
	EndpointDiscord = "https://discord.com"
	EndpointCDN     = "https://cdn.discordapp.com"

	EndpointGuilds     = "/guilds/"
	EndpointChannels   = "/channels/"
	EndpointUsers      = "/users/"
	EndpointGateway    = "/gateway"
	EndpointGatewayBot = EndpointGateway + "/bot"
	EndpointWebhooks   = "/webhooks/"

	EndpointCDNAttachments  = "/attachments/"
	EndpointCDNAvatars      = "/avatars/"
	EndpointCDNIcons        = "/icons/"
	EndpointCDNSplashes     = "/splashes/"
	EndpointCDNChannelIcons = "/channel-icons/"
	EndpointCDNBanners      = "/banners/"

	EndpointUser = func(userID string) string {
		return EndpointUsers + userID
	}
	EndpointUserAvatar = func(userID, avatarID string) string {
		return EndpointCDNAvatars + userID + "/" + avatarID + ".png"
	}
	EndpointUserAvatarAnimated = func(userID, avatarID string) string {
		return EndpointCDNAvatars + userID + "/" + avatarID + ".gif"
	}
	EndpointDefaultUserAvatar = func(index string) string {
		return "/embed/avatars/" + index + ".png"
	}
	EndpointUserGuilds = func(userID string) string {
		return EndpointUsers + userID + "/guilds"
	}
	EndpointUserGuild = func(userID, guildID string) string {
		return EndpointUsers + userID + "/guilds/" + guildID
	}
	EndpointUserGuildMember = func(userID, guildID string) string {
		return EndpointUserGuild(userID, guildID) + "/member"
	}
	EndpointUserChannels = func(userID string) string {
		return EndpointUsers + userID + "/channels"
	}

	EndpointGuild = func(guildID string) string {
		return EndpointGuilds + guildID
	}
	EndpointGuildVanityURL = func(guildID string) string {
		return EndpointGuilds + guildID + "/vanity-url"
	}
	EndpointGuildPreview = func(guildID string) string {
		return EndpointGuilds + guildID + "/preview"
	}
	EndpointGuildChannels = func(guildID string) string {
		return EndpointGuilds + guildID + "/channels"
	}
	EndpointGuildMembers = func(guildID string) string {
		return EndpointGuilds + guildID + "/members"
	}
	EndpointGuildMembersSearch = func(guildID string) string {
		return EndpointGuildMembers(guildID) + "/search"
	}
	EndpointGuildMember = func(guildID, userID string) string {
		return EndpointGuilds + guildID + "/members/" + userID
	}
	EndpointGuildMemberRole = func(guildID, userID, roleID string) string {
		return EndpointGuilds + guildID + "/members/" + userID + "/roles/" + roleID
	}
	EndpointGuildBans = func(guildID string) string {
		return EndpointGuilds + guildID + "/bans"
	}
	EndpointGuildBan = func(guildID, userID string) string {
		return EndpointGuilds + guildID + "/bans/" + userID
	}
	EndpointGuildIntegrations = func(guildID string) string {
		return EndpointGuilds + guildID + "/integrations"
	}
	EndpointGuildIntegration = func(guildID, integrationID string) string {
		return EndpointGuilds + guildID + "/integrations/" + integrationID
	}
	EndpointGuildIntegrationSync = func(guildID, integrationID string) string {
		return EndpointGuilds + guildID + "/integrations/" + integrationID + "/sync"
	}
	EndpointGuildRoles = func(guildID string) string {
		return EndpointGuilds + guildID + "/roles"
	}
	EndpointGuildRole = func(guildID, roleID string) string {
		return EndpointGuilds + guildID + "/roles/" + roleID
	}
	EndpointGuildInvites = func(guildID string) string {
		return EndpointGuilds + guildID + "/invites"
	}
	EndpointGuildWidget = func(guildID string) string {
		return EndpointGuilds + guildID + "/widget"
	}
	EndpointGuildWidgetImage = func(guildID string) string {
		return EndpointGuilds + guildID + "/widget.png"
	}
	EndpointGuildWidgetJSON = func(guildID string) string {
		return EndpointGuilds + guildID + "/widget.json"
	}
	EndpointGuildWelcomeScreen = func(guildID string) string {
		return EndpointGuilds + guildID + "/welcome-screen"
	}
	EndpointGuildStickers = func(guildID string) string {
		return EndpointGuilds + guildID + "/stickers"
	}
	EndpointGuildSticker = func(guildID, stickerID string) string {
		return EndpointGuilds + guildID + "/stickers/" + stickerID
	}
	EndpointGuildTemplates = func(guildID string) string {
		return EndpointGuilds + guildID + "/templates"
	}
	EndpointGuildTemplate = func(guildID, templateCode string) string {
		return EndpointGuilds + guildID + "/templates/" + templateCode
	}
	EndpointGuildVoiceRegions = func(guildID string) string {
		return EndpointGuilds + guildID + "/regions"
	}
	EndpointGuildVoiceStates = func(guildID string) string {
		return EndpointGuilds + guildID + "/voice-states"
	}
	EndpointGuildVoiceState = func(guildID, userID string) string {
		return EndpointGuilds + guildID + "/voice-states/" + userID
	}
	EndpointGuildVoiceStateSelf = func(guildID string) string {
		return EndpointGuilds + guildID + "/voice-states/@me"
	}

	EndpointGuildEmbed = EndpointGuildWidget

	EndpointGuildPrune = func(guildID string) string {
		return EndpointGuilds + guildID + "/prune"
	}
	EndpointGuildIcon = func(guildID, hash string) string {
		return EndpointCDNIcons + guildID + "/" + hash + ".png"
	}
	EndpointGuildIconAnimated = func(guildID, hash string) string {
		return EndpointCDNIcons + guildID + "/" + hash + ".gif"
	}
	EndpointGuildSplash = func(guildID, hash string) string {
		return EndpointCDNSplashes + guildID + "/" + hash + ".png"
	}
	EndpointGuildWebhooks = func(guildID string) string {
		return EndpointGuilds + guildID + "/webhooks"
	}
	EndpointGuildAuditLogs = func(guildID string) string {
		return EndpointGuilds + guildID + "/audit-logs"
	}
	EndpointGuildEmojis = func(guildID string) string {
		return EndpointGuilds + guildID + "/emojis"
	}
	EndpointGuildEmoji = func(guildID, emojiID string) string {
		return EndpointGuilds + guildID + "/emojis/" + emojiID
	}
	EndpointGuildBanner = func(guildID, hash string) string {
		return EndpointCDNBanners + guildID + "/" + hash + ".png"
	}

	EndpointGuildScheduledEvents = func(guildID string) string {
		return EndpointGuilds + guildID + "/scheduled-events"
	}
	EndpointGuildScheduledEvent = func(guildID, eventID string) string {
		return EndpointGuilds + guildID + "/scheduled-events/" + eventID
	}
	EndpointGuildScheduledEventUsers = func(guildID, eventID string) string {
		return EndpointGuildScheduledEvent(guildID, eventID) + "/users"
	}

	EndpointChannel = func(channelID string) string {
		return EndpointChannels + channelID
	}
	EndpointChannelPermissions = func(channelID string) string {
		return EndpointChannels + channelID + "/permissions"
	}
	EndpointChannelPermission = func(channelID, overwriteID string) string {
		return EndpointChannels + channelID + "/permissions/" + overwriteID
	}
	EndpointChannelInvites = func(channelID string) string {
		return EndpointChannels + channelID + "/invites"
	}
	EndpointChannelTyping = func(channelID string) string {
		return EndpointChannels + channelID + "/typing"
	}
	EndpointChannelMessages = func(channelID string) string {
		return EndpointChannels + channelID + "/messages"
	}
	EndpointChannelMessage = func(channelID, messageID string) string {
		return EndpointChannels + channelID + "/messages/" + messageID
	}
	EndpointChannelMessageAck = func(channelID, messageID string) string {
		return EndpointChannels + channelID + "/messages/" + messageID + "/ack"
	}
	EndpointChannelMessagesBulkDelete = func(channelID string) string {
		return EndpointChannel(channelID) + "/messages/bulk-delete"
	}
	EndpointChannelMessagesPins = func(channelID string) string {
		return EndpointChannel(channelID) + "/pins"
	}
	EndpointChannelMessagePin = func(channelID, messageID string) string {
		return EndpointChannel(channelID) + "/pins/" + messageID
	}
	EndpointChannelMessageCrosspost = func(channelID, messageID string) string {
		return EndpointChannel(channelID) + "/messages/" + messageID + "/crosspost"
	}
	EndpointChannelFollow = func(channelID string) string {
		return EndpointChannel(channelID) + "/followers"
	}

	EndpointGroupIcon = func(channelID, hash string) string {
		return EndpointCDNChannelIcons + channelID + "/" + hash + ".png"
	}

	EndpointChannelWebhooks = func(channelID string) string {
		return EndpointChannel(channelID) + "/webhooks"
	}
	EndpointWebhook = func(webhookID string) string {
		return EndpointWebhooks + webhookID
	}
	EndpointWebhookToken = func(webhookID, token string) string {
		return EndpointWebhooks + webhookID + "/" + token
	}
	EndpointWebhookMessage = func(webhookID, token, messageID string) string {
		return EndpointWebhookToken(webhookID, token) + "/messages/" + messageID
	}

	EndpointMessageReactionsAll = func(channelID, messageID string) string {
		return EndpointChannelMessage(channelID, messageID) + "/reactions"
	}
	EndpointMessageReactions = func(channelID, messageID, emojiID string) string {
		return EndpointChannelMessage(channelID, messageID) + "/reactions/" + emojiID
	}
	EndpointMessageReaction = func(channelID, messageID, emojiID, targetID string) string {
		return EndpointMessageReactions(channelID, messageID, emojiID) + "/" + targetID
	}

	EndpointApplicationGlobalCommands = func(applicationID string) string {
		return EndpointApplication(applicationID) + "/commands"
	}
	EndpointApplicationGlobalCommand = func(applicationID, channelID string) string {
		return EndpointApplicationGlobalCommands(applicationID) + "/" + channelID
	}

	EndpointApplicationGuildCommands = func(applicationID, guildID string) string {
		return EndpointApplication(applicationID) + "/guilds/" + guildID + "/commands"
	}
	EndpointApplicationGuildCommandsPermissions = func(applicationID, guildID string) string {
		return EndpointApplication(applicationID) + "/guilds/" + guildID + "/commands/permissions"
	}
	EndpointApplicationGuildCommand = func(applicationID, guildID, commandID string) string {
		return EndpointApplicationGuildCommands(applicationID, guildID) + "/" + commandID
	}
	EndpointApplicationGuildCommandPermissions = func(applicationID, guildID, commandID string) string {
		return EndpointApplicationGuildCommands(applicationID, guildID) + "/" + commandID + "/permissions"
	}
	EndpointInteraction = func(applicationID, interactionToken string) string {
		return "/interactions/" + applicationID + "/" + interactionToken
	}
	EndpointInteractionResponse = func(interactionID, interactionToken string) string {
		return EndpointInteraction(interactionID, interactionToken) + "/callback"
	}
	EndpointInteractionResponseActions = func(applicationID, interactionToken string) string {
		return EndpointWebhookMessage(applicationID, interactionToken, "@original")
	}
	EndpointFollowupMessage = func(applicationID, interactionToken string) string {
		return EndpointWebhookToken(applicationID, interactionToken)
	}
	EndpointFollowupMessageActions = func(applicationID, interactionToken, messageID string) string {
		return EndpointWebhookMessage(applicationID, interactionToken, messageID)
	}

	EndpointGuildCreate = "/guilds"

	EndpointInvite = func(inviteID string) string {
		return "/invites/" + inviteID
	}

	EndpointIntegrationsJoin = func(interactionID string) string {
		return "/integrations/" + interactionID + "/join"
	}

	EndpointEmoji = func(emojiID string) string {
		return "/emojis/" + emojiID + ".png"
	}
	EndpointEmojiAnimated = func(emojiID string) string {
		return "/emojis/" + emojiID + ".gif"
	}

	EndpointChannelThreads = func(channelID string) string {
		return EndpointChannels + channelID + "/threads"
	}

	EndpointApplications = "/applications"
	EndpointApplication  = func(applicationID string) string {
		return EndpointApplications + "/" + applicationID
	}

	EndpointEntitlements = func(applicationID string) string {
		return EndpointApplication(applicationID) + "/entitlements"
	}

	EndpointEntitlement = func(applicationID, entitlementID string) string {
		return EndpointEntitlements(applicationID) + "/" + entitlementID
	}

	EndpointOAuth2             = "/oauth2"
	EndpointOAuth2Me           = EndpointOAuth2 + "/@me"
	EndpointOAuth2Applications = EndpointOAuth2 + "/applications"
	EndpointOAuth2Authorize    = EndpointOAuth2 + "/authorize"
	EndpointOAuth2Token        = EndpointOAuth2 + "/token"
	EndpointOAuth2TokenRevoke  = EndpointOAuth2Token + "/revoke"

	EndpointOAuth2Application = func(applicationID string) string {
		return EndpointOAuth2Applications + "/" + applicationID
	}
	EndpointOAuth2ApplicationsBot = func(applicationID string) string {
		return EndpointOAuth2Applications + "/" + applicationID + "/bot"
	}
	EndpointOAuth2ApplicationAssets = func(applicationID string) string {
		return EndpointOAuth2Applications + "/" + applicationID + "/assets"
	}
)
