package com.dyw.consumer.controller;

import com.dyw.consumer.service.impl.CodeServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author Devil
 * @date 2022-04-12-14:50
 */
@RestController
@RequestMapping("code")
public class CodeController {

    @Autowired
    private CodeServiceImpl codeService;

    @GetMapping
    public void doCode(){
        codeService.getCode();
    }
}
