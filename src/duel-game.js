const { WordleGame, isValidWord } = require("./wordle");

const Conditions = {
    OUT_OF_GUESSES: "OUT_OF_GUESSES",
    INVALID_ID: "INVALID_ID",
    WIN: "WIN",
    INVALID_INPUT: "INVALID_INPUT",
    INVALID_WORD: "INVALID_WORD",
    BOTH_PLAYERS_OUT: "BOTH_PLAYERS_OUT",
}

// Two players with individual guesses
// supply user ID to decide whos game to choose

class DuelGame extends WordleGame{
    constructor(playerId){
        super();
        this.playerId = playerId;
        this.guessing = true;
        //setTimeout(playerGuessingOutOfTime() ,30000);
    }

} 

// Time -> if player runs out of time they loose
// 

class DuelWordle{
    constructor(player1, player2){
        
        this.player1= new DuelGame(player1);
        this.player2 = new DuelGame(player2);
        this.player2.word = this.player1.word;
    }

    submitGuess(playerId, guess)
    {
        guess = guess.toLowerCase();
        let player = null;
        if(Object.is(playerId, this.player1.playerId))
            player = this.player1;
        else if(Object.is(playerId, this.player2.playerId))
            player = this.player2;
        else
            return {condition: Conditions.INVALID_ID, result: null};

        if(!player.checkInput(guess))
            return {condition: Conditions.INVALID_INPUT, result: null};
        if(!isValidWord(guess))
            return {condition: Conditions.INVALID_WORD, result: null};

        let result = null;
        if(player.hasRemainingAttempts()){
            result = player.submitGuess(guess);

            if(result.correct)
                return {condition: Conditions.WIN, result: result, attempts: player.guesses.length};
        }
        if(!player.hasRemainingAttempts()){
            if(!this.player1.hasRemainingAttempts() && !this.player2.hasRemainingAttempts())
                return {condition: Conditions.BOTH_PLAYERS_OUT, result: result};
            return {condition: Conditions.OUT_OF_GUESSES, result: result};
        }
            
        return {condition: null, result: result};
    }

}

module.exports = {
    DuelWordle,
    Conditions,
}