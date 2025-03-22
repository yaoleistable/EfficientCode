using System;
using System.Net.Http;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Newtonsoft.Json;
using AITranslator.Models;

namespace AITranslator.Services
{
    public class TranslationService
    {
        private readonly HttpClient _httpClient;
        private readonly Configuration _config;

        public TranslationService(Configuration config)
        {
            _config = config;
            var baseUrl = _config.ApiSettings.BaseUrl.TrimEnd('/');
            _httpClient = new HttpClient
            {
                BaseAddress = new Uri(baseUrl),
                Timeout = TimeSpan.FromSeconds(_config.ApiSettings.Timeout)
            };
            _httpClient.DefaultRequestHeaders.Add("Authorization", $"Bearer {_config.ApiSettings.ApiKey}");
        }

        public async Task<string> TranslateTextAsync(string text)
        {
            try
            {
                if (!Uri.IsWellFormedUriString(_config.ApiSettings.BaseUrl, UriKind.Absolute))
                {
                    throw new Exception("API基础URL格式不正确，请在设置中检查并更正");
                }

                var targetLanguage = DetectLanguage(text);
                var messages = new[]
                {
                    new { role = "system", content = _config.PromptSettings.SystemPrompt },
                    new
                    {
                        role = "user",
                        content = _config.PromptSettings.UserPromptTemplate
                            .Replace("{source_lang}", targetLanguage == "en" ? "Chinese" : "English")
                            .Replace("{target_lang}", targetLanguage == "en" ? "English" : "Chinese")
                            .Replace("{text}", text)
                    }
                };

                var requestBody = new
                {
                    model = _config.ApiSettings.Model,
                    messages = messages,
                    temperature = _config.PromptSettings.Temperature
                };

                var response = await _httpClient.PostAsync(
                    "chat/completions",
                    new StringContent(JsonConvert.SerializeObject(requestBody), Encoding.UTF8, "application/json")
                );

                if (!response.IsSuccessStatusCode)
                {
                    var errorContent = await response.Content.ReadAsStringAsync();
                    throw new Exception($"API调用失败（状态码：{response.StatusCode}）：{errorContent}");
                }

                var result = await response.Content.ReadAsStringAsync();
                var completion = JsonConvert.DeserializeAnonymousType(result, new
                {
                    choices = new[]
                    {
                        new { message = new { content = "" } }
                    }
                });

                if (completion?.choices == null || completion.choices.Length == 0)
                {
                    throw new Exception("API返回的数据格式不正确");
                }

                return completion.choices[0].message.content ?? string.Empty;
            }
            catch (Exception ex)
            {
                throw new Exception($"翻译失败：{ex.Message}");
            }
        }

        private string DetectLanguage(string text)
        {
            // 使用正则表达式检测文本是否包含中文字符
            var containsChinese = Regex.IsMatch(text, @"[\u4e00-\u9fa5]");
            return containsChinese ? "en" : "zh";
        }

        public async Task<string> TranslateAsync(string text)
        {
            if (string.IsNullOrWhiteSpace(text))
                return string.Empty;

            return await TranslateTextAsync(text);
        }
    }
}