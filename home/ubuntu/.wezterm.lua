local wezterm = require("wezterm")
local mux = wezterm.mux

----------------------------- detect_os ----------------------------

local function detect_os()
  if wezterm.target_triple == "x86_64-apple-darwin" or wezterm.target_triple == "aarch64-apple-darwin" then
    return "macos"
  elseif wezterm.target_triple == "x86_64-pc-windows-msvc" then
    return "windows"
  elseif wezterm.target_triple == "x86_64-unknown-linux-gnu" then
    return "linux"
  else
    return "unknown"
  end
end

local myos = detect_os()

--------------------------- detect_shell ---------------------------

local function detect_shell()
  if myos == "windows" then
    return { "wsl.exe" }
  elseif myos == "mac" then
    return { "/opt/homebrew/bin/bash", "--login", }
  else
    return {"/bin/bash", "--login"}
  end
end

---------------------------- detect_font ---------------------------

local function detect_font()
  if myos == "windows" then
    return wezterm.font("Ubuntu Mono")
  else
    return wezterm.font("UbuntuMono Nerd Font")
  end
end

------------------------------- main -------------------------------

return {
  window_close_confirmation = 'NeverPrompt',
  default_prog = detect_shell(),

  color_scheme = 'Gruvbox Material (Gogh)',
  font = detect_font(),
  font_size = 26,

  default_domain =  myos:match('windows') and "WSL:Ubuntu" or "",

  colors = {
    cursor_bg = '#928374',
    cursor_border = '#928374',
  },

  window_padding = {
   left = 40,
   right = 0,
   top = 0,
   bottom = 2,
  },

  term = "xterm-256color",
  animation_fps = 1,
  max_fps = 1,

  ----------- rwxrob streaming stuff, fyi --------------

  --font_size = 41,
  --initial_cols = 74,
  --initial_rows = 22,

  --[[
  -- exactly position
  wezterm.on('gui-startup', function(cmd)
    local tab, pane, window = mux.spawn_window(cmd or {})
    window:gui_window():set_position(0, 130)
end),
  ]]--

  --window_decorations = "RESIZE",
  --enable_tab_bar = false,
  --hide_tab_bar_if_only_one_tab = true,

}





