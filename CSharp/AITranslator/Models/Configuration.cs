using System;
using System.IO;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace AITranslator.Models
{
    public class Configuration
    {
        public ApiSettings ApiSettings { get; set; } = new();
        public HotkeySettings HotkeySettings { get; set; } = new();
        public PromptSettings PromptSettings { get; set; } = new();

        private static readonly string ConfigPath = Path.Combine(
            Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData),
            "AITranslator",
            "config.json"
        );

        public static Configuration Load()
        {
            try
            {
                if (File.Exists(ConfigPath))
                {
                    var json = File.ReadAllText(ConfigPath);
                    return JsonConvert.DeserializeObject<Configuration>(json) ?? new Configuration();
                }
            }
            catch (Exception ex)
            {
                System.Windows.MessageBox.Show($"加载配置文件失败：{ex.Message}", "错误");
            }
            return new Configuration();
        }

        public void Save()
        {
            try
            {
                var directory = Path.GetDirectoryName(ConfigPath);
                if (!string.IsNullOrEmpty(directory) && !Directory.Exists(directory))
                {
                    Directory.CreateDirectory(directory);
                }

                var json = JsonConvert.SerializeObject(this, Formatting.Indented);
                File.WriteAllText(ConfigPath, json);
            }
            catch (Exception ex)
            {
                System.Windows.MessageBox.Show($"保存配置文件失败：{ex.Message}", "错误");
            }
        }
    }

    public class ApiSettings
    {
        public string BaseUrl { get; set; } = "https://api.openai.com/v1";
        public string ApiKey { get; set; } = string.Empty;
        public string Model { get; set; } = "gpt-3.5-turbo";
        public int Timeout { get; set; } = 30;
    }

    public class HotkeySettings
    {
        public string TranslateHotkey { get; set; } = "Alt+T";
    }

    public class PromptSettings
    {
        public string SystemPrompt { get; set; } = "You are a professional translator.";
        public string UserPromptTemplate { get; set; } = "Translate the following {source_lang} text to {target_lang}:\n{text}";
        public double Temperature { get; set; } = 0.3;
    }
}