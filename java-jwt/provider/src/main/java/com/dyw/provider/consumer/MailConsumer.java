package com.dyw.provider.consumer;

import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.spring.annotation.ConsumeMode;
import org.apache.rocketmq.spring.annotation.MessageModel;
import org.apache.rocketmq.spring.annotation.RocketMQMessageListener;
import org.apache.rocketmq.spring.annotation.SelectorType;
import org.apache.rocketmq.spring.core.RocketMQListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.javamail.JavaMailSenderImpl;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.thymeleaf.TemplateEngine;
import org.thymeleaf.context.Context;

import javax.mail.MessagingException;
import javax.mail.internet.MimeMessage;
import java.util.Date;

/**
 * @author Devil
 * @date 2022-04-12-23:11
 */
@Slf4j
@RocketMQMessageListener(consumerGroup = "jwt-group",consumeMode = ConsumeMode.ORDERLY,topic = "jwt-iris-spring:code",
selectorType = SelectorType.TAG,selectorExpression = "code",messageModel = MessageModel.CLUSTERING,
consumeTimeout = 1,maxReconsumeTimes = 1,nameServer = "${rocketmq.name-server}")
public class MailConsumer implements RocketMQListener<String> {
    @Value("${spring.mail.username}")
    private String from;
    @Autowired
    private JavaMailSenderImpl mailSender;
    @Autowired
    private TemplateEngine templateEngine;
    @Override
    public void onMessage(String message) {
        try {
            log.info("消费者接收到发送验证码的通知");
            mailSender.setUsername(from);
            MimeMessage mimeMessage = mailSender.createMimeMessage();
            MimeMessageHelper mimeMessageHelper = new MimeMessageHelper(mimeMessage);
            mimeMessageHelper.setTo("1842501760");
            mimeMessageHelper.setSubject("验证码");
            mimeMessageHelper.setFrom(from);
            mimeMessageHelper.setSentDate(new Date());
            Context context = new Context();
            context.setVariable("verifyCode","123456");
            String process = templateEngine.process("code.html", context);
            mimeMessageHelper.setText(process,true);
            mailSender.send(mimeMessage);
        } catch (MessagingException e) {
            log.error("邮件发送失败");
            e.printStackTrace();
        }
    }
}
