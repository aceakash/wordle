export class Game {
    readonly word: string
    private _isSolved: boolean = false
    
    public get isSolved() : boolean {
        return this._isSolved
    }

    constructor(word: string) {
        this.word = word
    }

    makeGuess(guess: string) {
        if (guess === this.word) {
            this._isSolved = true
        }
    }

    


}