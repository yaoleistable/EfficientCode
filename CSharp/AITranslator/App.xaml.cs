using System;
using System.Windows;

namespace AITranslator
{
    public partial class App : Application
    {
        private void Application_Startup(object sender, StartupEventArgs e)
        {
            // TODO: 初始化配置
            // TODO: 初始化服务
        }

        private void Application_Exit(object sender, ExitEventArgs e)
        {
            // TODO: 保存配置
            // TODO: 清理资源
        }

        private void Application_DispatcherUnhandledException(object sender, System.Windows.Threading.DispatcherUnhandledExceptionEventArgs e)
        {
            MessageBox.Show($"发生未处理的异常：{e.Exception.Message}", "错误", MessageBoxButton.OK, MessageBoxImage.Error);
            e.Handled = true;
        }
    }
}