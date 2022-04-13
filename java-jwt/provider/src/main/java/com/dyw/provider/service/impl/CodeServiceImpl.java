package com.dyw.provider.service.impl;

import com.dyw.jwt.proto.code.Code;
import com.dyw.jwt.proto.code.CodeRequest;
import com.dyw.jwt.proto.code.CodeResponse;
import com.dyw.jwt.proto.code.getCodeGrpc;
import com.dyw.provider.util.VerificationCodeUtil;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;
import org.apache.rocketmq.client.producer.SendCallback;
import org.apache.rocketmq.client.producer.SendResult;
import org.apache.rocketmq.spring.core.RocketMQTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.mail.javamail.JavaMailSenderImpl;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.thymeleaf.TemplateEngine;
import org.thymeleaf.context.Context;

import javax.mail.MessagingException;
import javax.mail.internet.MimeMessage;
import java.util.Arrays;
import java.util.Date;
import java.util.concurrent.TimeUnit;

/**
 * @author Devil
 * @date 2022-04-12-14:24
 */
@Slf4j
@GrpcService
public class CodeServiceImpl extends getCodeGrpc.getCodeImplBase {
    @Autowired
    private StringRedisTemplate stringRedisTemplate;
    @Autowired
    private RocketMQTemplate rocketMQTemplate;
    @Value("${spring.mail.username}")
    private String from;
    @Autowired
    private JavaMailSenderImpl mailSender;
    @Autowired
    private TemplateEngine templateEngine;
    @Override
    public void getCode(CodeRequest request, StreamObserver<CodeResponse> responseObserver){
        String email = request.getEmail();
        System.out.println(email);
        String code = VerificationCodeUtil.generateVerificationCode();
        stringRedisTemplate.opsForValue().set("code"+email,code,15, TimeUnit.MINUTES);
        try {
            log.info("消费者接收到发送验证码的通知");
            mailSender.setUsername(from);
            MimeMessage mimeMessage = mailSender.createMimeMessage();
            MimeMessageHelper mimeMessageHelper = new MimeMessageHelper(mimeMessage);
            mimeMessageHelper.setTo(email);
            mimeMessageHelper.setSubject("验证码");
            mimeMessageHelper.setFrom(from);
            mimeMessageHelper.setSentDate(new Date());
            Context context = new Context();

            context.setVariable("verifyCode", Arrays.asList(code.split("")));
            String process = templateEngine.process("code.html", context);
            mimeMessageHelper.setText(process,true);
            mailSender.send(mimeMessage);
        } catch (MessagingException e) {
            log.error("邮件发送失败");
            e.printStackTrace();
        }
        CodeResponse response = CodeResponse.newBuilder().setMessage("成功").build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
