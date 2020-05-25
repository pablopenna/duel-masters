package dm01

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"fmt"
)

// DeadlyFighterBraidClaw ...
func DeadlyFighterBraidClaw(c *match.Card) {

	c.Name = "Deadly Fighter Braid Claw"
	c.Power = 1000
	c.Civ = civ.Fire
	c.Family = family.Dragonoid
	c.ManaCost = 1
	c.ManaRequirement = []string{civ.Fire}

	c.Use(fx.Creature, fx.ForceAttack)

}

// ExplosiveFighterUcarn ...
func ExplosiveFighterUcarn(c *match.Card) {

	c.Name = "Explosive Fighter Ucarn"
	c.Power = 9000
	c.Civ = civ.Fire
	c.Family = family.Dragonoid
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Fire}

	c.Use(fx.Creature, fx.Doublebreaker, func(card *match.Card, ctx *match.Context) {
		if event, ok := ctx.Event.(*match.CardMoved); ok {

			if event.CardID == card.ID && event.To == match.BATTLEZONE {

				cards := match.Search(card.Player, ctx.Match, card.Player, match.MANAZONE, "Explosive Fighter Ucarn: Select 2 cards from your manazone that will be sent to your graveyard", 2, 2, false)

				for _, manaCard := range cards {
					card.Player.MoveCard(manaCard.ID, match.MANAZONE, match.GRAVEYARD)
					ctx.Match.Chat("Server", fmt.Sprintf("%s was sent from %s's manazone to their graveyard", manaCard.ID, card.Name))
				}

			}

		}
	})

}
