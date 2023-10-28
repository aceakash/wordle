import { expect, test } from "bun:test";
import { Game } from "../domain/game";

test("Getting the word correct in the first guess marks the game as solved", () => {
  const game = new Game("POINT");
  game.makeGuess("POINT");

  expect(game.isSolved).toBe(true);
});

test("Getting the word wrong in the first guess does not mark the game as solved", () => {
  const game = new Game("POINT");
  game.makeGuess("BLANK");

  expect(game.isSolved).toBe(false);
});
