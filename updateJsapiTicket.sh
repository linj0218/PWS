updateAccessTokenUrl="http://127.0.0.1:8080/pws/pws006"
curl -G $updateAccessTokenUrl
# 添加定时器
# crontab -e
# 每2小时运行一次脚本
# 0 */2 * * * /bin/sh /Users/ljpc/goworkspace/src/PWS/updateJsapiTicket.sh