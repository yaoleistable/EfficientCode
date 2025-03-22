using System;
using System.Windows.Input;
using System.Windows.Interop;
using System.Runtime.InteropServices;
using System.Collections.Generic;

namespace AITranslator.Services
{
    public class HotkeyManager
    {
        private const int WM_HOTKEY = 0x0312;
        private readonly Dictionary<int, Action> _hotkeyActions = new();
        private readonly IntPtr _windowHandle;
        private int _currentId = 0;

        [DllImport("user32.dll")]
        private static extern bool RegisterHotKey(IntPtr hWnd, int id, uint fsModifiers, uint vk);

        [DllImport("user32.dll")]
        private static extern bool UnregisterHotKey(IntPtr hWnd, int id);

        public HotkeyManager(IntPtr windowHandle)
        {
            _windowHandle = windowHandle;
            ComponentDispatcher.ThreadPreprocessMessage += ProcessHotkey;
        }

        public bool RegisterHotkey(string hotkeyString, Action action)
        {
            try
            {
                var (modifiers, key) = ParseHotkeyString(hotkeyString);
                var id = ++_currentId;
                
                if (RegisterHotKey(_windowHandle, id, (uint)modifiers, (uint)key))
                {
                    _hotkeyActions[id] = action;
                    return true;
                }
                return false;
            }
            catch
            {
                return false;
            }
        }

        public void UnregisterAll()
        {
            foreach (var id in _hotkeyActions.Keys)
            {
                UnregisterHotKey(_windowHandle, id);
            }
            _hotkeyActions.Clear();
        }

        private void ProcessHotkey(ref MSG msg, ref bool handled)
        {
            if (msg.message != WM_HOTKEY) return;

            var id = msg.wParam.ToInt32();
            if (_hotkeyActions.TryGetValue(id, out var action))
            {
                action.Invoke();
                handled = true;
            }
        }

        private (ModifierKeys modifiers, Key key) ParseHotkeyString(string hotkeyString)
        {
            var parts = hotkeyString.Split('+');
            ModifierKeys modifiers = ModifierKeys.None;
            Key key = Key.None;

            for (int i = 0; i < parts.Length - 1; i++)
            {
                switch (parts[i].Trim().ToLower())
                {
                    case "ctrl":
                        modifiers |= ModifierKeys.Control;
                        break;
                    case "alt":
                        modifiers |= ModifierKeys.Alt;
                        break;
                    case "shift":
                        modifiers |= ModifierKeys.Shift;
                        break;
                    case "win":
                        modifiers |= ModifierKeys.Windows;
                        break;
                }
            }

            if (Enum.TryParse(parts[^1].Trim(), true, out Key parsedKey))
            {
                key = parsedKey;
            }

            return (modifiers, key);
        }
    }
}