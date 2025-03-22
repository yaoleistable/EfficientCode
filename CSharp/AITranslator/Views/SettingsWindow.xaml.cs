using System;
using System.Windows;
using System.Windows.Input;
using System.Windows.Controls;
using System.Windows.Media;
using System.Windows.Shapes;
using AITranslator.Models;

namespace AITranslator.Views
{
    public partial class SettingsWindow : Window
    {
        private Configuration _config;
        private string _currentHotkey = string.Empty;
        private int _selectedTabIndex = -1;

        public SettingsWindow(Configuration config, int selectedTabIndex = -1)
        {
            InitializeComponent();
            _config = config;
            _selectedTabIndex = selectedTabIndex;
            LoadSettings();
        }



        private void LoadSettings()
        {
            // 加载热键设置
            HotkeyTextBox.Text = _config.HotkeySettings.TranslateHotkey;
            _currentHotkey = HotkeyTextBox.Text;

            // 加载API设置
            BaseUrlTextBox.Text = _config.ApiSettings.BaseUrl;
            ApiKeyPasswordBox.Password = _config.ApiSettings.ApiKey;
            ModelTextBox.Text = _config.ApiSettings.Model;
            TimeoutTextBox.Text = _config.ApiSettings.Timeout.ToString();

            // 初始化密码显示状态
            EyeIcon.Data = Geometry.Parse("M12 4.5C7 4.5 2.73 7.61 1 12c1.73 4.39 6 7.5 11 7.5s9.27-3.11 11-7.5c-1.73-4.39-6-7.5-11-7.5zM12 17c-2.76 0-5-2.24-5-5s2.24-5 5-5 5 2.24 5 5-2.24 5-5 5zm0-8c-1.66 0-3 1.34-3 3s1.34 3 3 3 3-1.34 3-3-1.34-3-3-3z");
            EyeIcon.Fill = new SolidColorBrush(Colors.Gray);

            // 加载提示词设置
            SystemPromptTextBox.Text = _config.PromptSettings.SystemPrompt;
            UserPromptTemplateTextBox.Text = _config.PromptSettings.UserPromptTemplate;
            TemperatureTextBox.Text = _config.PromptSettings.Temperature.ToString("0.0");

            // 设置选中的选项卡
            if (_selectedTabIndex >= 0 && _selectedTabIndex < 3)
            {
                ((TabControl)this.FindName("SettingsTabControl")).SelectedIndex = _selectedTabIndex;
            }
        }

        private void HotkeyTextBox_PreviewKeyDown(object sender, KeyEventArgs e)
        {
            e.Handled = true;

            // 获取修饰键
            var modifiers = Keyboard.Modifiers;
            var key = e.Key;

            // 忽略单独的修饰键
            if (key == Key.LeftCtrl || key == Key.RightCtrl ||
                key == Key.LeftAlt || key == Key.RightAlt ||
                key == Key.LeftShift || key == Key.RightShift ||
                key == Key.LWin || key == Key.RWin)
            {
                return;
            }

            var hotkeyText = "";

            if ((modifiers & ModifierKeys.Control) == ModifierKeys.Control)
                hotkeyText += "Ctrl+";
            if ((modifiers & ModifierKeys.Alt) == ModifierKeys.Alt)
                hotkeyText += "Alt+";
            if ((modifiers & ModifierKeys.Shift) == ModifierKeys.Shift)
                hotkeyText += "Shift+";
            if ((modifiers & ModifierKeys.Windows) == ModifierKeys.Windows)
                hotkeyText += "Win+";

            hotkeyText += key.ToString();
            HotkeyTextBox.Text = hotkeyText;
        }

        private void ClearHotkey_Click(object sender, RoutedEventArgs e)
        {
            HotkeyTextBox.Text = string.Empty;
        }

        private void CheckHotkey_Click(object sender, RoutedEventArgs e)
        {
            if (string.IsNullOrWhiteSpace(HotkeyTextBox.Text))
            {
                MessageBox.Show("请先设置热键", "提示", MessageBoxButton.OK, MessageBoxImage.Information);
                return;
            }

            var windowHandle = new System.Windows.Interop.WindowInteropHelper(this).Handle;
            var hotkeyManager = new Services.HotkeyManager(windowHandle);
            var success = hotkeyManager.RegisterHotkey(HotkeyTextBox.Text, () => { });
            hotkeyManager.UnregisterAll();

            if (success)
            {
                MessageBox.Show("热键可以成功注册", "提示", MessageBoxButton.OK, MessageBoxImage.Information);
            }
            else
            {
                MessageBox.Show("热键注册失败，请尝试其他组合键", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
            }
        }

        private void TogglePasswordButton_Click(object sender, RoutedEventArgs e)
        {
            var parent = (Grid)ApiKeyPasswordBox.Parent;
            if (ApiKeyPasswordBox.Visibility == Visibility.Visible)
            {
                // 显示密码
                var textBox = new TextBox
                {
                    Text = ApiKeyPasswordBox.Password,
                    VerticalContentAlignment = VerticalAlignment.Center
                };
                Grid.SetColumn(textBox, 0);
                parent.Children.Remove(ApiKeyPasswordBox);
                parent.Children.Add(textBox);
                ApiKeyPasswordBox.Visibility = Visibility.Collapsed;
                EyeIcon.Data = Geometry.Parse("M12 7c2.76 0 5 2.24 5 5 0 .65-.13 1.26-.36 1.83l2.92 2.92c1.51-1.26 2.7-2.89 3.43-4.75-1.73-4.39-6-7.5-11-7.5-1.4 0-2.74.25-3.98.7l2.16 2.16C10.74 7.13 11.35 7 12 7zM2 4.27l2.28 2.28.46.46C3.08 8.3 1.78 10.02 1 12c1.73 4.39 6 7.5 11 7.5 1.55 0 3.03-.3 4.38-.84l.42.42L19.73 22 21 20.73 3.27 3 2 4.27zM7.53 9.8l1.55 1.55c-.05.21-.08.43-.08.65 0 1.66 1.34 3 3 3 .22 0 .44-.03.65-.08l1.55 1.55c-.67.33-1.41.53-2.2.53-2.76 0-5-2.24-5-5 0-.79.2-1.53.53-2.2zm4.31-.78l3.15 3.15.02-.16c0-1.66-1.34-3-3-3l-.17.01z");
            }
            else
            {
                // 隐藏密码
                var textBox = parent.Children.OfType<TextBox>().FirstOrDefault();
                if (textBox != null)
                {
                    ApiKeyPasswordBox.Password = textBox.Text;
                    parent.Children.Remove(textBox);
                    parent.Children.Add(ApiKeyPasswordBox);
                    ApiKeyPasswordBox.Visibility = Visibility.Visible;
                    EyeIcon.Data = Geometry.Parse("M12 4.5C7 4.5 2.73 7.61 1 12c1.73 4.39 6 7.5 11 7.5s9.27-3.11 11-7.5c-1.73-4.39-6-7.5-11-7.5zM12 17c-2.76 0-5-2.24-5-5s2.24-5 5-5 5 2.24 5 5-2.24 5-5 5zm0-8c-1.66 0-3 1.34-3 3s1.34 3 3 3 3-1.34 3-3-1.34-3-3-3z");
                }
            }
        }

        private bool ValidateSettings()
        {
            // 验证热键设置
            if (string.IsNullOrWhiteSpace(HotkeyTextBox.Text))
            {
                MessageBox.Show("请设置翻译热键", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            // 验证API设置
            if (string.IsNullOrWhiteSpace(BaseUrlTextBox.Text))
            {
                MessageBox.Show("请输入API基础URL", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            if (!Uri.IsWellFormedUriString(BaseUrlTextBox.Text, UriKind.Absolute))
            {
                MessageBox.Show("API基础URL格式不正确，请输入完整的URL地址（例如：https://api.openai.com/v1）", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            if (string.IsNullOrWhiteSpace(ApiKeyPasswordBox.Password))
            {
                MessageBox.Show("请输入API密钥", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            if (string.IsNullOrWhiteSpace(ModelTextBox.Text))
            {
                MessageBox.Show("请输入模型名称", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            if (!int.TryParse(TimeoutTextBox.Text, out int timeout) || timeout <= 0)
            {
                MessageBox.Show("超时时间必须是大于0的整数", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            // 验证提示词设置
            if (string.IsNullOrWhiteSpace(SystemPromptTextBox.Text))
            {
                MessageBox.Show("请输入系统提示词", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            if (string.IsNullOrWhiteSpace(UserPromptTemplateTextBox.Text))
            {
                MessageBox.Show("请输入用户提示词模板", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            if (!double.TryParse(TemperatureTextBox.Text, out double temperature) || temperature < 0 || temperature > 1)
            {
                MessageBox.Show("温度值必须在0.0到1.0之间", "验证错误", MessageBoxButton.OK, MessageBoxImage.Warning);
                return false;
            }

            return true;
        }

        private void SaveButton_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                if (!ValidateSettings())
                    return;

                // 保存热键设置
                _config.HotkeySettings.TranslateHotkey = HotkeyTextBox.Text;

                // 保存API设置
                _config.ApiSettings.BaseUrl = BaseUrlTextBox.Text.TrimEnd('/');
                _config.ApiSettings.ApiKey = ApiKeyPasswordBox.Password;
                _config.ApiSettings.Model = ModelTextBox.Text.Trim();
                _config.ApiSettings.Timeout = int.Parse(TimeoutTextBox.Text);

                // 保存提示词设置
                _config.PromptSettings.SystemPrompt = SystemPromptTextBox.Text.Trim();
                _config.PromptSettings.UserPromptTemplate = UserPromptTemplateTextBox.Text.Trim();
                _config.PromptSettings.Temperature = double.Parse(TemperatureTextBox.Text);

                _config.Save();
                DialogResult = true;
                Close();
            }
            catch (Exception ex)
            {
                MessageBox.Show($"保存设置失败：{ex.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
            }
        }

        private void CancelButton_Click(object sender, RoutedEventArgs e)
        {
            DialogResult = false;
            Close();
        }

        private void NumberValidationTextBox(object sender, TextCompositionEventArgs e)
        {
            if (sender is not TextBox textBox)
                return;
                
            var fullText = textBox.Text + e.Text;

            // 检查是否为温度输入框
            if (textBox == TemperatureTextBox)
            {
                // 允许输入0-1之间的小数
                e.Handled = !double.TryParse(fullText, out double value) || value < 0 || value > 1;
            }
            else if (textBox == TimeoutTextBox)
            {
                // 只允许输入正整数
                e.Handled = !int.TryParse(fullText, out int value) || value <= 0;
            }
        }

        private async void CheckApiButton_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                // 禁用检查按钮
                CheckApiButton.IsEnabled = false;
                CheckApiButton.Content = "正在检查...";

                // 创建临时配置
                var tempConfig = new Configuration
                {
                    ApiSettings = new ApiSettings
                    {
                        BaseUrl = BaseUrlTextBox.Text.TrimEnd('/'),
                        ApiKey = ApiKeyPasswordBox.Password,
                        Model = ModelTextBox.Text.Trim(),
                        Timeout = int.Parse(TimeoutTextBox.Text)
                    },
                    PromptSettings = _config.PromptSettings,
                    HotkeySettings = _config.HotkeySettings
                };

                // 创建临时翻译服务
                var translationService = new Services.TranslationService(tempConfig);

                // 发送测试请求
                await translationService.TranslateAsync("测试");

                MessageBox.Show("API连接测试成功！", "成功", MessageBoxButton.OK, MessageBoxImage.Information);
            }
            catch (Exception ex)
            {
                MessageBox.Show($"API连接测试失败：{ex.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
            }
            finally
            {
                // 恢复检查按钮
                CheckApiButton.IsEnabled = true;
                CheckApiButton.Content = "检查API连接";
            }
        }
    }
}