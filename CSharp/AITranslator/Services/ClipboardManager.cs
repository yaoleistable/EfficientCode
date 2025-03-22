using System;
using System.Windows;
using System.Runtime.InteropServices;
using System.Text;

namespace AITranslator.Services
{
    public class ClipboardManager
    {
        [DllImport("user32.dll")]
        public static extern IntPtr GetForegroundWindow();

        [DllImport("user32.dll")]
        private static extern void keybd_event(byte bVk, byte bScan, uint dwFlags, UIntPtr dwExtraInfo);

        private const byte VK_CONTROL = 0x11;
        private const byte VK_C = 0x43;
        private const uint KEYEVENTF_KEYUP = 0x0002;

        public string GetSelectedText()
        {
            string originalText = string.Empty;
            try
            {
                // 保存原始剪贴板内容
                if (Clipboard.ContainsText())
                    originalText = Clipboard.GetText();

                // 模拟Ctrl+C复制选中文本
                var foregroundWindow = GetForegroundWindow();
                if (foregroundWindow != IntPtr.Zero)
                {
                    // 按下Ctrl+C
                    keybd_event(VK_CONTROL, 0, 0, UIntPtr.Zero);
                    keybd_event(VK_C, 0, 0, UIntPtr.Zero);

                    // 释放Ctrl+C
                    keybd_event(VK_C, 0, KEYEVENTF_KEYUP, UIntPtr.Zero);
                    keybd_event(VK_CONTROL, 0, KEYEVENTF_KEYUP, UIntPtr.Zero);

                    // 等待剪贴板更新
                    System.Threading.Thread.Sleep(100);

                    // 获取选中的文本
                    if (Clipboard.ContainsText())
                    {
                        var selectedText = Clipboard.GetText();
                        // 恢复原始剪贴板内容
                        if (!string.IsNullOrEmpty(originalText))
                            Clipboard.SetText(originalText);
                        return selectedText;
                    }
                }

                // 恢复原始剪贴板内容
                if (!string.IsNullOrEmpty(originalText))
                    Clipboard.SetText(originalText);

                return string.Empty;
            }
            catch (Exception ex)
            {
                MessageBox.Show($"获取选中文本失败：{ex.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
                return string.Empty;
            }
        }

        public void SetText(string text)
        {
            try
            {
                Clipboard.SetText(text);
            }
            catch (Exception ex)
            {
                MessageBox.Show($"设置剪贴板文本失败：{ex.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
            }
        }
    }
}