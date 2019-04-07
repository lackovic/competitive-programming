/*
 * Marco Lackovic, Apr 2019
 * Google Code Jam, Qualification Round
 * Foregone Solution
 */
"use strict";
import * as readline from "readline";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

let i = 0;

rl.on("line", (line: string) => {
  if (i > 0) {
    const n = Number(line);
    const [a, b] = solve(n);
    console.log(`Case #${i}: ${a} ${b}`);
  }
  i++;
});

const solve = (n: number): number[] => {
  const checks = new Array<number>(2);
  const numAsArrayOfDigits = Array.from(n.toString()).map(Number);
  checks[0] = Number(
    numAsArrayOfDigits
      .map((d: Number) => (d == 4 ? 2 : d))
      .map(String)
      .join("")
  );
  checks[1] = Number(
    numAsArrayOfDigits
      .map((d: Number) => (d == 4 ? 2 : 0))
      .map(String)
      .join("")
  );
  return checks;
};
