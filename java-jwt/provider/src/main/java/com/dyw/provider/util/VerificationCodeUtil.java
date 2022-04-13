package com.dyw.provider.util;

import java.util.Random;

/**
 * @author Devil
 * @date 2022-04-13-20:35
 */
public class VerificationCodeUtil {
    /**
     * 生成验证码的位数
     */
    private static final int GENERATE_VERIFICATION_CODE_LENGTH = 6;
    private static final String[] META_CODE ={"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
            "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
            "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"};

    public static String generateVerificationCode(){
        Random random = new Random(System.currentTimeMillis());
        StringBuilder verificationCode = new StringBuilder();
        while(verificationCode.length()<GENERATE_VERIFICATION_CODE_LENGTH){
            int i = random.nextInt(META_CODE.length);
            verificationCode.append(META_CODE[i]);
        }
        return verificationCode.toString();
    }
}
