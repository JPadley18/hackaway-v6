var fs = require("fs")
var path = require("path")

// For testing user input
const prompt = require("prompt-sync")({sigint: true});

// TODO:
// KEEP TRACK OF USED CHARACTERS and create getters for them
// CREATE GETTER FOR PAST ATTEMPTS
// INPUT FILTERING = DONE

const Result = {
    CORRECT_CHARACTER: "correct_character",
    INCORRECT_POSITION: "incorrect_position",
    INCORRECT_CHARACTER: "incorrect_character",
};

const wordList = importWords();

var incorrectCharacters = new Set();

class WordleGame{
    constructor(){
        // Choose a random word
        this.word = wordList[Math.floor(Math.random() * wordList.length)];
        this.guesses = [];
    }

    submitGuess(guess){
        guess = guess.toLowerCase();
        // Input filtering
        if(!this.checkInput(guess)){
            return null; // NULL RETURN WHEN INPUT INVALID
        }
        
        let toAdd = [];
        // The following stores each character of the guess under the key 'char' with the correctness of the character under "STATE"
        // The state is any value from char_values
        // the char is the individual characters from the guess
        // Structure example: {CHAR: 'KEY', STATE: 'STATE FROM char_values' }
        // NOTE: CHAR: null WHEN STATE IS EQUAL TO correct_guess. THIS IS WHEN THE GUESS WAS CORRECT
        let charUsed = [];
        for(let i = 0; i < guess.length; i++){
            let key = guess.charAt(i);
            let result;
            
            if(guess.charAt(i) == this.word.charAt(i))
                result = Result.CORRECT_CHARACTER;
            else if(this.word.includes(guess.charAt(i))) // needs to check for cases where the character only occurs once but the guess has two instances of it
            { // Check in for loop against word and if equal to word[index] then check that gues[index] doesnt equal word[index]
                result = Result.INCORRECT_POSITION;

                charUsed.push(key); // push key

                let idk = false;
                for(let i = 0; i < guess.length; i++)
                {
                    // If key is equal to character in guess
                    if(Object.is(key, guess.charAt(i)))
                    {
                        if(Object.is(key, this.word.charAt(i)))
                        {
                            result = Result.INCORRECT_CHARACTER;
                            idk = true;
                        }
                    }
                }

                if(!idk)
                {
                    let charInWord = 0;
                    for(let i = 0; i < this.word.length; i++)
                    {
                        if(Object.is(key, this.word.charAt(i)))
                        {
                            charInWord++;
                        }
                    }
    
                    let charSeen = 0;
                    for(let i = 0; i < charUsed.length; i++)
                    {
                        if(Object.is(key, charUsed[i]))
                        {
                            charSeen++;
                        }
                    }
    
                    if(charSeen > charInWord)
                        result = Result.INCORRECT_CHARACTER;
                }
            }
            else{
                result = Result.INCORRECT_CHARACTER;
                // Add incorrect characters to used incorrect set
                incorrectCharacters.add(guess.charAt(i));
            }
            
            toAdd.push({char: key, result: result});
        }
        
        let result = ({
            correct: toAdd.every(r => r.result == Result.CORRECT_CHARACTER), 
            guess: toAdd,
        });

        this.guesses.push(result);

        return result;
    }

    // checks for input being 5 characters long and no numbers only alpha characters
    checkInput(input){
        if(input.length != 5)
            return false;
        return /^[a-zA-Z()]*$/.test(input);
    }

    hasRemainingAttempts()
    {
        return this.guesses.length < 6;
    }
}

function importWords(){
    let data = fs.readFileSync("./words.json", {encoding:"utf8", flag:"r"});
    return JSON.parse(data);
}

function isValidWord(word){
    return wordList.includes(word);
}

module.exports = {
    WordleGame,
    Result,
    isValidWord,
};