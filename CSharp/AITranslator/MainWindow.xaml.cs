using System;
using System.Windows;

namespace AITranslator
{
    public partial class MainWindow : Window
    {
        private Services.HotkeyManager _hotkeyManager;
        private readonly Services.ClipboardManager _clipboardManager;
        private Services.TranslationService _translationService;

        public MainWindow()
        {
            InitializeComponent();
            var config = Models.Configuration.Load();
            _clipboardManager = new Services.ClipboardManager();
            _translationService = new Services.TranslationService(config);
            var windowHandle = new System.Windows.Interop.WindowInteropHelper(this).Handle;
            _hotkeyManager = new Services.HotkeyManager(windowHandle);
            InitializeHotkeys();
        }

        private void InitializeHotkeys()
        {
            var config = Models.Configuration.Load();
            var success = _hotkeyManager.RegisterHotkey(config.HotkeySettings.TranslateHotkey, async () =>
            {
                try
                {
                    // 保存当前活动窗口
                    var foregroundWindow = System.Windows.Interop.HwndSource.FromHwnd(
                        Services.ClipboardManager.GetForegroundWindow());

                    // 获取选中文本并等待剪贴板更新
                    var selectedText = _clipboardManager.GetSelectedText();
                    if (string.IsNullOrWhiteSpace(selectedText))
                    {
                        MessageBox.Show("请先选择要翻译的文本\n1. 已选中要翻译的文本\n2. 当前窗口允许复制操作\n3. 没有其他程序占用剪贴板\n4. 选中的文本不为空", 
                            "提示", MessageBoxButton.OK, MessageBoxImage.Information);
                        return;
                    }

                    // 显示主窗口并激活
                    this.Show();
                    this.Activate();
                    this.WindowState = WindowState.Normal;
                    this.Topmost = true;

                    // 设置文本并开始翻译
                    SourceTextBox.Text = selectedText;
                    await TranslateText();
                    this.Topmost = false;

                    // 如果原窗口还存在，重新激活它
                    if (foregroundWindow?.Handle != IntPtr.Zero)
                    {
                        var window = System.Windows.Window.GetWindow(foregroundWindow.RootVisual);
                        window?.Activate();
                    }
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"翻译过程中发生错误：{ex.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
                }
            });
            // 格式化热键显示
            var hotkeyDisplay = config.HotkeySettings.TranslateHotkey
                .Replace("Alt", "Alt+")
                .Replace("Ctrl", "Ctrl+")
                .Replace("Shift", "Shift+")
                .Replace("Win", "Win+");

            StatusText.Text = success 
                ? $"当前状态：已启动，热键：{hotkeyDisplay}" 
                : $"当前状态：热键 {hotkeyDisplay} 注册失败，请在设置中更换其他组合键";
        }

        private void OpenHotkeySettings_Click(object sender, RoutedEventArgs e)
        {
            OpenSettings(0);
        }

        private void OpenApiSettings_Click(object sender, RoutedEventArgs e)
        {
            OpenSettings(1);
        }

        private void OpenPromptSettings_Click(object sender, RoutedEventArgs e)
        {
            OpenSettings(2);
        }

        private void OpenSettings(int selectedTabIndex)
        {
            var config = Models.Configuration.Load();
            var settingsWindow = new Views.SettingsWindow(config, selectedTabIndex)
            {
                Owner = this
            };
            if (settingsWindow.ShowDialog() == true)
            {
                // 重新加载配置并更新热键
                config = Models.Configuration.Load();
                _translationService = new Services.TranslationService(config);
                _hotkeyManager.UnregisterAll();
                InitializeHotkeys();
            }
        }

        protected override void OnSourceInitialized(EventArgs e)
        {
            base.OnSourceInitialized(e);
            var windowHandle = new System.Windows.Interop.WindowInteropHelper(this).Handle;
            _hotkeyManager = new Services.HotkeyManager(windowHandle);
            InitializeHotkeys();
        }

        protected override void OnClosed(EventArgs e)
        {
            base.OnClosed(e);
            _hotkeyManager?.UnregisterAll();
        }



        private async Task TranslateText()
        {
            try
            {
                var sourceText = SourceTextBox.Text;
                if (string.IsNullOrWhiteSpace(sourceText))
                {
                    MessageBox.Show("请输入要翻译的文本", "提示", MessageBoxButton.OK, MessageBoxImage.Information);
                    return;
                }

                StatusBarText.Text = "正在翻译...";
                var translatedText = await _translationService.TranslateAsync(sourceText);
                TranslatedTextBox.Text = translatedText;
                StatusBarText.Text = "就绪";
            }
            catch (Exception ex)
            {
                StatusBarText.Text = "翻译失败";
                MessageBox.Show($"翻译失败：{ex.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
            }
        }

        private async void TranslateButton_Click(object sender, RoutedEventArgs e)
        {
            await TranslateText();
        }

        private void ClearButton_Click(object sender, RoutedEventArgs e)
        {
            SourceTextBox.Clear();
            TranslatedTextBox.Clear();
            StatusBarText.Text = "就绪";
        }
    }
}