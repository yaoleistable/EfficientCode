using System;
using System.Windows;
using System.Windows.Controls;

namespace AITranslator.Views
{
    public partial class TranslationWindow : Window
    {
        public TranslationWindow(string sourceText, string translatedText)
        {
            InitializeComponent();
            SourceTextBox.Text = sourceText;
            TranslatedTextBox.Text = translatedText;
            System.Windows.Clipboard.SetText(translatedText); // 自动复制翻译结果到剪贴板
        }

        private void CopySource_Click(object sender, RoutedEventArgs e)
        {
            if (!string.IsNullOrEmpty(SourceTextBox.Text))
            {
                System.Windows.Clipboard.SetText(SourceTextBox.Text);
                MessageBox.Show("原文已复制到剪贴板", "提示", MessageBoxButton.OK, MessageBoxImage.Information);
            }
        }

        private void CopyTranslated_Click(object sender, RoutedEventArgs e)
        {
            if (!string.IsNullOrEmpty(TranslatedTextBox.Text))
            {
                System.Windows.Clipboard.SetText(TranslatedTextBox.Text);
                MessageBox.Show("译文已复制到剪贴板", "提示", MessageBoxButton.OK, MessageBoxImage.Information);
            }
        }

        private void Close_Click(object sender, RoutedEventArgs e)
        {
            Close();
        }
    }
}