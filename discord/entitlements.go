package discord

import "time"

type Entitlement struct {
	UserID         *Snowflake      `json:"user_id,omitempty"`
	GiftCodeFlags  *GiftCodeFlags  `json:"gift_code_flags,omitempty"`
	StartsAt       *time.Time      `json:"starts_at,omitempty"`
	EndsAt         *time.Time      `json:"ends_at,omitempty"`
	GuildID        *Snowflake      `json:"guild_id,omitempty"`
	SubscriptionID *Snowflake      `json:"subscription_id,omitempty"`
	ID             Snowflake       `json:"id"`
	SkuID          Snowflake       `json:"sku_id"`
	ApplicationID  Snowflake       `json:"application_id"`
	Type           EntitlementType `json:"type"`
	Deleted        bool            `json:"deleted"`
	Consumed       bool            `json:"consumed"`
}

// EntitlementParams represents the payload sent to discord.
type EntitlementParams struct {
	SkuID     Snowflake `json:"sku_id"`
	OwnerID   Snowflake `json:"owner_id"`
	OwnerType OwnerType `json:"owner_type"`
}

// EntitlementType represents the type of an entitlement.
type EntitlementType uint16

const (
	// EntitlementTypePurchase is an entltmenet that was purchased by user.
	EntitlementTypePurchase EntitlementType = 1
	// EntitlementTypePremiumGift is an entitlement for a Discord Nitro subscription.
	EntitlementTypePremiumGift EntitlementType = 2
	// EntitlementTypeDeveloperGift is an entitlement gifted by developer.
	EntitlementTypeDeveloperGift EntitlementType = 3
	// EntitlementTypeTestModePurchase is an entitlement that was purchased by a dev in application test mode.
	EntitlementTypeTestModePurchase EntitlementType = 4
	// EntitlementTypeFreePurchase is an entitlement that was granted when the SKU was free.
	EntitlementTypeFreePurchase EntitlementType = 5
	// EntitlementTypeUserGift is an entitlement that was gifted by another user.
	EntitlementTypeUserGift EntitlementType = 6
	// EntitlementTypePremiumPurchase is an entitlement that was claimed by the user for free as a Nitro Subscriber.
	EntitlementTypePremiumPurchase EntitlementType = 7
	// EntitlementTypeApplicationSubscription is an entitlement that was purchased as an app subscription.
	EntitlementTypeApplicationSubscription EntitlementType = 8
)

// GiftCodeFlags is undocumented, but present in the API.
type GiftCodeFlags uint16

// OwnerType represents who owns the entitlement.
type OwnerType uint16

const (
	OwnerTypeGuild OwnerType = 1
	OwnerTypeUser  OwnerType = 2
)
