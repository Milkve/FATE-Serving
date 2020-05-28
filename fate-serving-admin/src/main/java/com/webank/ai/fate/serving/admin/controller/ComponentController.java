package com.webank.ai.fate.serving.admin.controller;

import com.alibaba.fastjson.JSONObject;
import com.google.common.collect.Maps;
import com.webank.ai.fate.api.networking.common.CommonServiceGrpc;
import com.webank.ai.fate.api.networking.common.CommonServiceProto;
import com.webank.ai.fate.register.zookeeper.ZookeeperRegistry;
import com.webank.ai.fate.serving.admin.services.ComponentService;
import com.webank.ai.fate.serving.core.bean.GrpcConnectionPool;
import com.webank.ai.fate.serving.core.bean.ReturnResult;
import io.grpc.ManagedChannel;
import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

/**
 * @Description Service management
 * @Date: 2020/3/25 11:13
 * @Author: v_dylanxu
 */
@RequestMapping("/api")
@RestController
public class ComponentController {

    private static final Logger logger = LoggerFactory.getLogger(ComponentController.class);

//    @Autowired
//    ZookeeperRegistry  zookeeperRegistry;

    @Autowired
    ComponentService  componentServices;

    @Value("${grpc.timeout:5000}")
    private int timeout;

    GrpcConnectionPool grpcConnectionPool = GrpcConnectionPool.getPool();

    @GetMapping("/component/list")
    public ComponentService.NodeData list() {
      return  componentServices.getCachedNodeData();
    }

    @GetMapping("/component/listProps")
    public ReturnResult listProps(String host, int port, String keyword) {
        ManagedChannel managedChannel = grpcConnectionPool.getManagedChannel(host, port);
        CommonServiceGrpc.CommonServiceBlockingStub blockingStub = CommonServiceGrpc.newBlockingStub(managedChannel);
        blockingStub = blockingStub.withDeadlineAfter(timeout, TimeUnit.MILLISECONDS);

        CommonServiceProto.QueryPropsRequest.Builder builder = CommonServiceProto.QueryPropsRequest.newBuilder();
        if (StringUtils.isNotBlank(keyword)) {
            builder.setKeyword(keyword);
        }

        CommonServiceProto.CommonResponse response = blockingStub.listProps(builder.build());
        Map<String, Object> propMap = JSONObject.parseObject(response.getData().toStringUtf8(), Map.class);

        List<Map> list = propMap.entrySet().stream()
                .sorted(Comparator.comparing(Map.Entry::getKey))
                .map(entry -> {
                    Map map = Maps.newHashMap();
                    map.put("key", entry.getKey());
                    map.put("value", entry.getValue());
                    return map;
                })
                .collect(Collectors.toList());

        Map data = Maps.newHashMap();
        data.put("total", list.size());
        data.put("rows", list);

        return ReturnResult.build(response.getStatusCode(), response.getMessage(), data);
    }








    // 列出集群中各个组件，能查询每个组件的配置信息，能够显示每个组件的健康状况

    // 组件上报后缓存起来，给出ID方便调用

    // 组件启停

}