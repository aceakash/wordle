import { expect, test } from "bun:test";
import { Game } from "../domain/game";

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
