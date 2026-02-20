# Designing Enemies

My philosophy is that, with peripherals being "enemies that are never required to kill in order to win", we could go more into the direction of peripherals giving more buffs, while the brunt of the damage comes from the main enemy.

All enemies also have a passive.

#### Do's and Don'ts of Peripherals

* They **DO** buff the main enemy, giving more inevitability to the fight.
* They **DON'T** do things that are hard to evaluate the value of.
* They **DO** debuff the player in specific ways (meaning not general like "minus damage/block"), taxing specific decks moreso than others. Think "cards now cost 2 life instead of resources".
* They **DO** have "biome specificity," meaning peripherals from the same biome are sometimes shared between many enemies in that biome.
* They **DO** have predictable patterns, often doing the same thing every turn.
* They **DON'T** necessarily produce good outcomes if dealt with. It is often the case, but not a rule.

With this frame of peripherals being -- literally -- extensions responsible for one specific thing that can be reused throughout multiple encounters we start entering a design space that is based on designing fun building blocks and then just playing tetris and combining them together. "Get the arm from that guy, thrusters from that other guy, make them come from Industry instead of Tech and you've built a new enemy out of parts".

## Cluster Biomes

The pools of enemies are based on a "cluster biome," with each biome having an overarching theme it tackles.

Each biome has specific nodes, and ideally we build it so you **have to visit a bit of each biome,** even if not necessarily going towards the end of it.

The reason for introducing the biomes system is to help the exploration problems I've outlined in **The graph of ITG** and **Encounter balance** sections of the **Overview** document. In summary, it is to make your deck get taxed more heavily through the quest for the rewards being more guarded behind a heavy commitment to harder encounters that are -- themselves -- more specialized (meaning not a damage race all the time).

Through these clusters we are basically coupling a few rooms together in a theme and giving the player the option of a few quests, namely:

* Go to the end of the cluster, fighting more and more specialized enemies to get a big reward, **if** their deck is particularly good against this biome;
* Use their tools as a means to completing the quest through "currency" instead of combat, getting the big reward (this is essentially a delayed shop purchase);
* Not engage with the cluster because their deck is not particularly well positioned against this biome.

Either way we are making exploration more engaging and thought-inducing because of the natural escalation of enemies inside a cluster. **A cluster is now a mini-act.**


#### Do's and Don'ts of Biomes

* They **DO** have very different layouts from one another in terms of density of enemy-to-node ratios and unlocking needed.
* They **DON'T** all have to be balanced between each other. It's fine to have an easier and a harder biome as long as the quest rewards are balanced for the difficulty in completing them.
* They **DO** have specific enemies tied to them.
* They **DON'T** ask the player to navigate to the end of them every time to complete a quest.
* They **DO** sometimes tie progression to their completion, acting basically as nodes (in the graph terminology) that need to be traversed in order to advance the run.

For this last point, something like a biome having a key needed to unlock the boss' room would work.

The exploration graph this system produces is **hopefully more structured than what we currently have,** which is a change I will now attempt to defend to be positive.

### How does this work with the alarm system?

I'm not too sure yet, but that's basically because I don't like the alarm system in itself.

I believe the alarm system has the _intent_ to be a balance mechanism. I think it doesn't accomplish it in a fun way, and it feels like a system _appended_ to the game, rather than a system that elevates and interacts with the game in a thought-provoking way.

## Brainstorming Biomes

### Generic Biomes

Nodes for general goodness (card rewards, healing, shops).

Enemies here could be interesting if they simple and scaled with their peripherals instead, creating "random enemies" more or less. That is, an enemy you fight is a combination of

    [Alarm Level] * [Random Peripherals] + [Base Simple Enemy]

for instance. This way we make the alarm system less boring. Instead of a raw % increase to stats we get to make more decisions on what to kill.

#### Generic Peripherals

* Heat-Viz Goggles - 50 Health
  * [100%] Give -10 Barrier to the player

* Hand-gun - 50 Health
  * [75%] Deal 10 Damage
  * [25%] Deal 20 Damage

* Claws - 20 Health
  * [100%] Deal 5 Damage. If it hits, inflict **Shred 5**. _(If the player would take damage, instead they take an extra 5 damage. Lasts forever.)_

* Poison Gas Spitter - 20 Health
  * [100%] Adds **2 Confused cards** to the player's Discard pile. _(Confused can't be played. Decompiles at the end of the round.)_

* Shield Melter
  * [100%] Inflicts **2 Slow.** _(If the player would gain barrier, instead they gain that much minus 2. Lasts forever.)_

### Industry Biome

Nodes for riches (credits, shops) and scaling.

### Recon Biome

Nodes for manipulating exploration (sentinel breakers, gate openers, revealing more biomes, teleports).

### Mod Biome

Nodes for deck and character manipulation (drivers, unlocking new commands, health, mana, specific card nodes like a protocol-only node)

Enemies are scaling brutes that tax health and the deck (adding Curses on turn 5 for example).

Enemies have lots of peripherals focused on **scaling their power** if the fight lasts long enough moreso than flat out damaging.



### Creating balanced Biomes

We have a formula for creating a balanced biome that goes like so:

Each biome has an overall weight for "how good it generally is" called `base_weight`.

Each room layout has

<!-- 
**Layer 1** enemies are taxing decks in the following ways:

1. They take away your health by a bit every turn but scale slowly, forcing you to be "in a hurry" to end them in ~5 turns.
1. They tackle your ability to get out of the gates fast.

## Layer 2 Enemies

**Layer 2** enemies are taxing decks in the following ways:

1. They have varying mechanics that ask your deck to be well rounded, and the patterns of attack take a while and then go for an ultimate, forcing you to be a in hurry if you're not particularly good at handling the side-effects of their ultimate.
1. They create
-->