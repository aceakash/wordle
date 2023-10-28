import { expect, test } from "bun:test";
import { Game } from "../domain/game";
import { fail } from "assert";

test("Getting the word correct in the first guess marks the game as solved", () => {
  const game = new Game("POINT", 6);
  game.makeGuess("POINT");

  expect(game.isSolved).toBe(true);
});

test("Getting the word wrong in the first guess does not mark the game as solved", () => {
  const game = new Game("POINT", 6);
  game.makeGuess("BLANK");

  expect(game.isSolved).toBe(false);
});

test("When max allowed guesses are exhausted, making further guesses is not allowed", () => {
  const game = new Game("POINT", 1);

  game.makeGuess("BLANK");

  expect(() => {
    game.makeGuess("PRISM");
  }).toThrow();

  // todo: look for a specific error; maybe use an enum of error codes
});

// todo: don't expose the word until the game is over
// todo: maybe have a list of states of the game

test("[clues] Guessing the first letter correctly gives the CORRECT clue for that letter", () => {
  const game = new Game("POINT", 6);
  const clues = game.makeGuess("PRAMS");

  expect(clues[0]).toBe("CORRECT");
});

test("[clues] If the first letter of a guess is present in the word but in the wrong place, we get the MISPLACED clue for that letter in the guess", () => {
  const game = new Game("POINT", 6);
  const clues = game.makeGuess("TRUCK");

  expect(clues[0]).toBe("MISPLACED");
});

test("[clues] If the first letter of a guess is not in the word at all, we get the ABSENT clue for that letter in the guess", () => {
  const game = new Game("POINT", 6);
  const clues = game.makeGuess("SPARE");

  expect(clues[0]).toBe("ABSENT");
});

test("[clues] With word POINT, and guess POUND, we get the correct clues", () => {
  const game = new Game("POINT", 6);
  const clue = game.makeGuess("POUND");
  expect(clue).toStrictEqual([
    "CORRECT",
    "CORRECT",
    "ABSENT",
    "CORRECT",
    "ABSENT",
  ]);
});

test("[clues] With word POINT, and guess MASON, we get the correct clues", () => {
  const game = new Game("POINT", 6);
  const clue = game.makeGuess("MASON");
  expect(clue).toStrictEqual([
    "ABSENT",
    "ABSENT",
    "ABSENT",
    "MISPLACED",
    "MISPLACED",
  ]);
});

test("[clues] With various words and guesses, we get the correct clues", () => {
  const clueColorLookup = {
    ABSENT: "🔴",
    MISPLACED: "🟡",
    CORRECT: "🟢",
  };

  const game = new Game("POINT", Infinity);
  const testCases = [
    ["POINT", "MASON", "🔴🔴🔴🟡🟡"],
    ["POINT", "SLURP", "🔴🔴🔴🔴🟡"],
    ["POINT", "DEALS", "🔴🔴🔴🔴🔴"],
    ["POINT", "STORY", "🔴🟡🟡🔴🔴"],
    ["POINT", "STORY", "🔴🟡🟡🔴🔴"],
  ];

  testCases.forEach((testCase) => {
    const [word, guess, expectedClues] = testCase;

    const actualClues = game.makeGuess(guess);
    const actualClueColours = actualClues
      .map((c) => clueColorLookup[c])
      .join("");

    if (actualClueColours !== expectedClues) {
      fail(
        `Want ${testCase[2]}, got ${actualClueColours} (for word ${word} & guess ${guess})`
      );
    }
  });
});
