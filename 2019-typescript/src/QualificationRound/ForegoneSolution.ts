/*
 * Marco Lackovic, Apr 2019
 * Google Code Jam, Qualification Round
 * Foregone Solution
 */

declare var require: any;
declare var process: any;
const readline = require("readline");
const rl = readline.createInterface({ input: process.stdin });

let i = 0;
rl.on("line", (num: string) => {
  if (i > 0) {
    const [a, b] = solve(num);
    console.log(`Case #${i}: ${a} ${b}`);
  }
  i++;
});

const solve = (num: string): number[] => {
  const checks = new Array<number>(2);
  const numAsArrayOfDigits = num.split("").map(Number);
  checks[1] = Number(
    numAsArrayOfDigits
      .map((d: Number) => (d == 4 ? 2 : 0))
      .map(String)
      .join("")
  );
  checks[0] = Number(num) - checks[1];
  return checks;
};
