# Voice Input for GitHub Copilot with Superwhisper

GitHub Copilot CLI is already one of the most capable agentic coding tools available. Add live voice transcription and it becomes something closer to a hands-free AI pair programmer — the same experience people rave about with Claude Code's voice input, but running on your existing Copilot subscription.

## What you need

- GitHub Copilot subscription (the `copilot` CLI)
- [Superwhisper](https://superwhisper.com) with a **lifetime license** ($249 direct from superwhisper.com — buy direct, not the App Store, to avoid Apple's 20% cut)
- macOS or Windows (Superwhisper launched Windows support in late 2025; some advanced features are still catching up to the macOS version)

## Why the lifetime license

Superwhisper's free tier uses local models only. Live streaming transcription — where words appear on screen as you speak — requires the **Nova (Cloud)** model, which is a Pro feature. The lifetime license at $249 is the best value if you plan to use this daily. Monthly and annual plans are also available (~$9/month or ~$108/year).

## Set up Superwhisper for live transcription

1. Open Superwhisper and navigate to your mode (or create a new one)
2. Open **Advanced Settings** in the sidebar
3. Set **Voice model** to **Nova (Cloud)**
4. A **Realtime** toggle will appear — enable it
5. Set a push-to-talk hotkey (hold to speak, release to insert)
6. Make sure **Output** is set to type into the active app, not just copy to clipboard

If you don't see the Realtime toggle after selecting Nova (Cloud), update Superwhisper to the latest version and restart the app. Confirm your Pro/lifetime license is activated under Settings → Account.

## Using voice input with Copilot

Launch `copilot` in your terminal, click into the prompt, hold your push-to-talk key, and speak. Words stream into the prompt in real time. Release the key and your transcription is ready to submit.

This is exactly the workflow that makes Claude Code feel natural — the difference is you are running it on your Copilot subscription with no additional per-token cost.

## Why this matters

Claude Code's voice input feature is one of the main reasons developers prefer it for long agentic sessions — you can think out loud, describe what you want, and let the agent figure out the implementation without interrupting your flow to type. Superwhisper brings that same capability to `gh copilot`, which already runs Claude Sonnet under the hood. The net result is a voice-driven agentic coding session backed by your existing GitHub Copilot subscription.

If you already have a Claude Pro subscription, you can also use this setup with `claude` (Claude Code CLI) — Claude Pro includes Claude Code access at no extra charge. Run both and use whichever fits the task.

## Alternatives

- **macOS built-in Dictation** (System Settings → Keyboard → Dictation, set to hold `fn`) — free, live transcription, works in any terminal, lower accuracy than Superwhisper Nova
- **Windows Speech Recognition / Voice Access** — built into Windows 11, works system-wide including terminals, free but less accurate than Superwhisper
- **Wispr Flow** — similar push-to-talk product, ~$15/month, no lifetime option

**Reference:** [Superwhisper Realtime Transcription](https://superwhisper.com/docs/common-issues/realtime)
