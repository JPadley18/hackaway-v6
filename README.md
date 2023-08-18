# Cordle Bot
A competitive Wordle discord bot originally developed during Royal Hackaway V6.

## Cordle Legacy
This is the original repo used in Royal Hackaway V6. After the event, we chose to rebuild the bot in Go (see the `cordle2` branch) . The codebase has since been reorganised into two repos, [Cordle Legacy](https://github.com/cordle-bot/cordle-legacy) and [Cordle Bot](https://github.com/cordle-bot/cordle-bot). Legacy preserves the codebase as it was at the end of the hackathon, while Cordle Bot contains the up to date Go code used in Cordle Bot today.

## Bot usage guide
In a server with Cordle Bot present, these following commands can be used:
- `/duel [@player]`: Challenge another player to a duel
- `/accept`: Accept an incoming challenge
- `/decline`: Decline an incoming challenge
- `/guess [guess]`: Submit a guess in an ongoing game
- `/leaderboard`: Display the current top-10 leaderboard
- `/stats [optional player]`: Show your stats by default or those of a specified player

