type Clue = "CORRECT" | "MISPLACED" | "ABSENT";

export class Game {
  readonly word: string;
  readonly maxGuessesAllowed: number;
  private _isSolved: boolean = false;

  private guessesMade: number = 0;

  public get isSolved(): boolean {
    return this._isSolved;
  }

  constructor(word: string, maxGuessesAllowed: number) {
    this.word = word;
    this.maxGuessesAllowed = maxGuessesAllowed;
  }

  makeGuess(guess: string): Clue[] {
    if (this.guessesMade == this.maxGuessesAllowed) {
      throw new Error("Max guesses exceeded");
    }
    if (guess === this.word) {
      this._isSolved = true;
    }
    this.guessesMade++;

    let clues: Clue[] = [];

    for (let i = 0; i < guess.length; i++) {
      const letter = guess[i];
      if (letter === this.word[i]) {
        clues.push("CORRECT");
        continue;
      }
      if (this.word.includes(letter)) {
        clues.push("MISPLACED");
        continue;
      }
      clues.push("ABSENT");
    }

    return clues;
  }
}
