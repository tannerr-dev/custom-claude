import { load } from "https://deno.land/std/dotenv/mod.ts";
const env = await load();
console.log(env)


import Anthropic from 'npm:@anthropic-ai/sdk';

const anthropic = new Anthropic({
  apiKey: env["CLAUDEKEY"]// defaults to process.env["ANTHROPIC_API_KEY"]
});



let prompt = "how can i check to see if my codebase is virus free?"



const msg = await anthropic.messages.create({
  model: "claude-3-7-sonnet-20250219",
  max_tokens: 1024,
  messages: [{ role: "user", content: prompt }],
});
console.log(msg);



// To run this code, you need to provide the --allow-write flag
// e.g., deno run --allow-write your-script.ts

const file = await Deno.open("log.txt", { append: true, create: true });
await file.write(new TextEncoder()
  .encode(`=====\nPrompt: ${prompt}\n\n======\nResponse: ${msg.content[0].text}\n`)
);
file.close();

console.log("Files have been appended successfully!");
