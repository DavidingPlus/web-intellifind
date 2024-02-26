package captcha

import (
	"backend/utils"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象，并配置存储 Key 的前缀
		store := RedisStore{
			RedisClient: utils.Redis,
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			viper.GetInt("Capatcha.height"),      // 宽
			viper.GetInt("Capatcha.width"),       // 高
			viper.GetInt("Capatcha.length"),      // 长度
			viper.GetFloat64("Capatcha.maxskew"), // 数字的最大倾斜角度
			viper.GetInt("Capatcha.dotcount"),    // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)

	})

	return internalCaptcha
}

// 生成验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, answer string, err error) {

	return c.Base64Captcha.Generate()

}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	match = c.Base64Captcha.Verify(id, answer, false)
	//fmt.Println(utils.Redis.Get(id))
	//utils.Redis.Del(id)
	return match

}
