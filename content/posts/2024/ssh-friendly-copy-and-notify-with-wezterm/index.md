---
author: lawrencegripper
category:
  - programming
  - wezterm
  - terminal
date: "2025-01-08T07:02:00+00:00"
title: "WezTerm: Easily copy text or send notification to local machine (even when connected via SSH)"
url: /2025/01/08/wezterm-easily-copy-text-or-send-notifications-to-local-machine-even-when-connected-via-ssh/
draft: false
---

Ever had a long-running command in terminal but forget to check back on it?

I've used `do_long_think; notify-send "thing finished"` in the past to help. I can do other stuff then be interrupted when it finishes.

Annoyingly that doesn't work if you're over an SSH connection, in a [DevContainer](https://code.visualstudio.com/docs/devcontainers/containers) or using [Codespaces](https://github.com/features/codespaces).

To fix that up I'm now using [WezTerm](https://wezfurlong.org/wezterm/)'s [`user-var-changed`](https://wezfurlong.org/wezterm/config/lua/window-events/user-var-changed.html) event, and it's sooo good!

Any window in WezTerm which writes out the `user var` escape sequence triggers a function in the
WezTerm lua config. For example 👇 will send `name: foo value: bar`

`printf "\033]1337;SetUserVar=%s=%s\007" foo `echo -n bar | base64`

This doesn't care if it's in TMUX over nested SSH. It **just works**.

To receive these events do stuff with them you add `wezterm.on('user-var-changed..` to [WezTerm config](https://wezfurlong.org/wezterm/config/files.html).

Here I wire up `weznot` and `weznot`. `weznot` triggers a notification with the `value` passed in and `wezcopy` copies the `value` to the clipboard.

```lua
wezterm.on('user-var-changed', function(window, pane, name, value)
  wezterm.log_info('var', name, value)
  if name == 'wez_not' then
    window:toast_notification('wezterm', 'msg: ' .. value, nil, 1000)
  end

  if name == 'wez_copy' then
    window:copy_to_clipboard(value, 'Clipboard')
  end
end)
```

To use these easily I create functions in dotfiles which output the escape sequences (these work over ssh as load into my SSH connections with [sshrc](https://github.com/cdown/sshrc)).

```bash
# Send a notification with wezterm use like `do think && weznot "think is done"`
function weznot() {
    title=$1
    printf "\033]1337;SetUserVar=%s=%s\007" wez_not $(echo -n "$title" | base64)
}

# Pipeline content to the clipboard `echo "hello" | wezcopy`
function wezcopy() {
    read clip_stuff
    printf "\033]1337;SetUserVar=%s=%s\007" wez_copy $(echo -n "$clip_stuff" | base64)
}

# Run a command and notify that the command has failed or succeeded
function wezmon() {
    command=$*
    
    eval $command
    
    last_exit_code=$?
    if [ $last_exit_code -eq 0 ]; then
        weznot "✅ '$command' completed successfully"
    else
        weznot "❌ '$command' failed"
    fi
}
```

Now you can use them like 👇

```bash
weznot "Hey stuff finished!"

# or 

cat /some/log/file.txt | wezcopy
```

Even better, to save a few keystrokes with `dothing; weznot "thing finished"` you can do

```bash
wezmon do_thing
```

Notification will be: `✅ 'do_thing' completed successfully` or `❌ 'do_thing' failed` automatically 

🎉 Happy terminal-ing